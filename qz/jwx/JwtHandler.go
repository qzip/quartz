/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

http://www.apache.org/licenses/LICENSE-2.0

Author: Ashish Banerjee, <ashish@qzip.in>

*/

//Package jwx handles JWT http(s) authentication
package jwx

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/pascaldekloe/jwt"
)

// JwtAuthorization checks authorization from config
type JwtAuthorization interface {
	ValidAudience(audiences []string) bool
	ValidIssuerRoles(issuer string, roles []string) bool
	JwtPubKeys() *jwt.KeyRegister
}

// NewJwtWrapper JWT auth & role wrapper
func NewJwtWrapper(jwtAuth JwtAuthorization, targetHandler http.Handler) http.Handler {

	roleFunc := &JwtRolesHandler{
		jwtAuthorization: jwtAuth,
	}
	comp := &JwtHandler{
		Target: targetHandler,
		Keys:   jwtAuth.JwtPubKeys(),
		// map two claims to HTTP headers
		HeaderBinding: map[string]string{
			"sub": "X-Verified-User",   // registered [standard] claim
			"iss": "X-Verified-Issuer", // private [custom] claim
		},
		Func: roleFunc.CheckRoleAuthorization,
	}
	// TODO: decide upon Func func(http.ResponseWriter, *http.Request, *Claims) (pass bool)
	// see: https://github.com/pascaldekloe/jwt/blob/master/web.go
	// see: https://play.golang.org/p/k_TF_54oJPX
	// the func checks if the caller's role allows
	return comp

}

// JwtHandler protects an http.Handler with security enforcements.
// Requests are only passed to Target if the JWT checks out.
// cloned from jwt/web.go
type JwtHandler struct {
	// Target is the secured service.
	Target http.Handler

	// Keys defines the trusted credentials.
	Keys *jwt.KeyRegister

	// HeaderBinding maps JWT claim names to HTTP header names.
	// All requests passed to Target have these headers set. In
	// case of failure the request is rejected with status code
	// 401 (Unauthorized) and a description.
	HeaderBinding map[string]string

	// ContextKey places the validated Claims in the context of
	// each respective request passed to Target when set. See
	// http.Request.Context and context.Context.Value.
	ContextKey interface{}

	// When not nil, then Func is called after the JWT validation
	// succeeds and before any header bindings. Target is skipped
	// [request drop] when the return is false.
	// This feature may be used to further customise requests or
	// as a filter or as an extended http.HandlerFunc.
	Func func(http.ResponseWriter, *http.Request, *jwt.Claims) (pass bool)
}

// ServeHTTP honors the http.Handler interface.
func (h *JwtHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// verify claims
	claims, err := h.Keys.CheckHeader(r)
	//fmt.Println("JwtHandler: ServeHTTP")
	if err == jwt.ErrNoHeader {
		r.Header["X-Verified-Roles"] = []string{"public"}
		//w.Header().Set("WWW-Authenticate", "Bearer")
		h.Target.ServeHTTP(w, r)
		return
	}

	if err != nil {
		w.Header().Set("WWW-Authenticate", `Bearer error="invalid_token", error_description=`+strconv.QuoteToASCII(err.Error()))
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// verify time constraints
	if !claims.Valid(time.Now()) {
		w.Header().Set("WWW-Authenticate", `Bearer error="invalid_token", error_description="jwt: time constraints exceeded"`)
		http.Error(w, "jwt: time constraints exceeded", http.StatusUnauthorized)
		return
	}

	// apply the custom function when set
	if h.Func != nil && !h.Func(w, r, claims) {
		return
	}

	// claim propagation
	for claimName, headerName := range h.HeaderBinding {
		s, ok := claims.String(claimName)
		if !ok {
			msg := "jwt: want string for claim " + claimName
			w.Header().Set("WWW-Authenticate", `Bearer error="invalid_token", error_description=`+strconv.QuoteToASCII(msg))
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}

		r.Header.Set(headerName, s)
	}

	// place claims in request context
	if h.ContextKey != nil {
		r = r.WithContext(context.WithValue(r.Context(), h.ContextKey, claims))
	}

	h.Target.ServeHTTP(w, r)
}

// JwtRolesHandler wraps around a HTTP handler
type JwtRolesHandler struct {
	jwtAuthorization JwtAuthorization
}

// CheckRoleAuthorization for a given jwt id can it authorize a set of roles?
//
func (jf *JwtRolesHandler) CheckRoleAuthorization(w http.ResponseWriter, req *http.Request, claims *jwt.Claims) (pass bool) {
	//log.Printf("got a valid JWT %q for %q", claims.ID, claims.Audiences)

	ok := jf.jwtAuthorization.ValidAudience(claims.Audiences)
	iroles, ok := claims.Set["roles"]
	if !ok {
		http.Error(w, "jwt: roles not set", http.StatusForbidden)
		return false
	}
	var roles []string
	rolOne, ok := iroles.(string)
	if !ok {
		rolx, ok := iroles.([]interface{})
		if !ok {
			http.Error(w, "jwt: roles must be string or string array", http.StatusForbidden)
			return false
		}
		roles = make([]string, len(rolx))
		for i, sval := range rolx {
			roles[i] = fmt.Sprint(sval)
		}

	} else {
		roles = make([]string, 1)
		roles[0] = rolOne
	}
	ok = jf.jwtAuthorization.ValidIssuerRoles(claims.Issuer, roles)
	if !ok {
		http.Error(w, "jwt: Issuer not allowed for specified roles", http.StatusForbidden)
		return false
	}
	req.Header["X-Verified-Roles"] = roles

	return true
	// https://www.iana.org/assignments/jwt/jwt.xhtml
	// aud (Audience) must include the web site its serving
	// iss (Issuer) must have the rigts to issue the roles to sub (Subject)
	// https://tools.ietf.org/html/rfc7519
	// https://github.com/pascaldekloe/jwt/blob/master/jwt.go
	//
}

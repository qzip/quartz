/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

http://www.apache.org/licenses/LICENSE-2.0

Author: Ashish Banerjee, <ashish@qzip.in>
*/

//Package jwx handles JWT authentication
package jwx

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"fmt"
	"qz/util"
	"strings"

	"github.com/pascaldekloe/jwt"
)

/*
// GenKey generates ECDSA P256 curve Private Key Pair
func GenKey() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	prvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	pubKey := prvKey.Public()
	return prvKey, &pubKey, nil
}
*/

// JwkAuthConfig helper class to Jwt Http(s) Auth Handler
type JwkAuthConfig struct {
	JwkAuthPubFiles []string                    `json:"jwkAuthPubFiles"`
	Audience        string                      `json:"audience"`
	IsserRoles      map[string][]string         `json:"issuerRoles"`
	RoleAccess      map[string]RoleObjectAccess `json:"roleObjectAccess"`
}

// RoleObjectAccess maps role to object fields access
type RoleObjectAccess struct {
	IsSU      bool                `json:"superuser,omitempty"`
	ClassPath map[string][]string `json:"classPath,omitempty"`
}

// JwkAuth authenticates the JWT issuer against roles
type JwkAuth struct {
	keys *jwt.KeyRegister
	cfg  *JwkAuthConfig
}

// NewJwkAuth loads the keys from files & sets up config.
func NewJwkAuth(authCfg *JwkAuthConfig) (*JwkAuth, error) {
	auth := &JwkAuth{keys: &jwt.KeyRegister{}, cfg: authCfg}
	for _, pubKFl := range authCfg.JwkAuthPubFiles {
		pubKey, err := util.GetJSONWebKey(pubKFl)
		if err != nil {
			return nil, err
		}
		if err = auth.addKey(pubKey.Key); err != nil {
			return nil, err
		}
	}

	return auth, nil
}

// ValidAudience checks audience validity
func (ja *JwkAuth) ValidAudience(audiences []string) bool {
	for _, aud := range audiences {
		if strings.EqualFold(aud, ja.cfg.Audience) {
			return true
		}
	}
	return false
}

// ValidIssuerRoles checks if issuer is entitled to allow roles (to subject)
func (ja *JwkAuth) ValidIssuerRoles(issuer string, roles []string) bool {
	irs, ok := ja.cfg.IsserRoles[issuer]
	if !ok {
		return false
	}
	x := make(map[string]bool)
	for _, ir := range irs {
		x[ir] = true
	}
	for _, rol := range roles {
		_, ok = x[rol]
		if !ok {
			return false
		}
	}

	return true
}

// JwtPubKeys returns the JWK key Register
func (ja *JwkAuth) JwtPubKeys() *jwt.KeyRegister {
	return ja.keys
}

// MaskFields masks out fields, if none allowed then return (nil, true)
func (ja *JwkAuth) MaskFields(roles []string, class string,
	data map[string]interface{}) (map[string]interface{}, bool) {
	fieldPaths := make(map[string]bool)
	// check if su or allows all in any role
	for _, ro := range roles {
		acc, ok := ja.cfg.RoleAccess[ro]
		//fmt.Println("JwkAuth.MaskFields", "role:", ro, ok)
		if !ok {
			continue
		}
		if acc.IsSU {
			return data, false
		}
		cpArr, ok := acc.ClassPath[class]
		//fmt.Println("JwkAuth.MaskFields", "class", class, "cpArr len", len(cpArr), ok)
		if ok {
			if len(cpArr) == 0 {
				return data, false // allows all
			}
			for _, cp := range cpArr {
				fieldPaths[cp] = true
			}
		}
	} // for roles
	if len(fieldPaths) > 0 {
		spaths := make([]string, 0)
		for k := range fieldPaths {
			spaths = append(spaths, k)
		}
		out := util.StripMap(data, spaths)
		stripped := false
		if len(out) < len(data) {
			stripped = true
		}
		return out, stripped
	}
	return nil, true // none allowed
}
func (ja *JwkAuth) addKey(key interface{}) error {
	keys := ja.keys
	switch t := key.(type) {
	case *ecdsa.PublicKey:
		keys.ECDSAs = append(keys.ECDSAs, t)
	case *ecdsa.PrivateKey:
		keys.ECDSAs = append(keys.ECDSAs, &t.PublicKey)
	case *rsa.PublicKey:
		keys.RSAs = append(keys.RSAs, t)
	case *rsa.PrivateKey:
		keys.RSAs = append(keys.RSAs, &t.PublicKey)
	default:
		return fmt.Errorf("jwt: unsupported key type %T", t)
	}
	return nil
}

// DefaultJwkAuth provides a default tempate for Auth
func DefaultJwkAuth() *JwkAuthConfig {

	//
	// in.qzip.bc.v01
	pubFlNames := []string{"jwk_sig_EdDSA_com.aloagri.auth.v01.pub", "jwk_sig_EdDSA_com.khetose.auth.v01.pub"}
	irols := map[string][]string{
		"in.qzip.blockchain.01may2019": {"su"},
		"com.khetose.auth.v01":         {"public", "farmtrade", "khetose"},
		"com.aloagri.auth.v01":         {"public", "farmtrade", "aloagri"},
	}

	ja := &JwkAuthConfig{
		JwkAuthPubFiles: pubFlNames,
		Audience:        "in.qzip.bc.v01",
		IsserRoles:      irols,
		RoleAccess: map[string]RoleObjectAccess{
			"public": {
				IsSU:      false,
				ClassPath: map[string][]string{"in.qzip.gpay.block": {""}},
			},
			"su": {IsSU: true},
		},
	}

	ja.RoleAccess["aloagri"] = RoleObjectAccess{ClassPath: map[string][]string{
		"com.aloagri":     {""},
		"in.gpay.sim.txn": {"eventID", "amount"},
	}}
	ja.RoleAccess["khetose"] = RoleObjectAccess{ClassPath: map[string][]string{
		"com.khetose":     {""},
		"in.gpay.sim.txn": {"tmStamp", "amount"},
	}}
	ja.RoleAccess["farmtrade"] = RoleObjectAccess{ClassPath: map[string][]string{
		"biz.farmtrade":   {""},
		"in.gpay.sim.txn": {"eventID", "fromAccount", "toAccount"},
	}}
	return ja
}

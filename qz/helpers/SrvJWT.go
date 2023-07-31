/*
Copyright (c) 2021, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"qz/commands"
	"qz/jwx"
	"reflect"
)

// SrvJWT implements qz/commands.HelperFactory
type SrvJWT struct {
}

//SrvJWTConfig configuration
type SrvJWTConfig struct {
	ContentHandlerCtxName string            `json:"contentHandlerCtxName"`
	JwkAuth               jwx.JwkAuthConfig `json:"jwkAuth"`
}

//Name implements qz/commands.HelperFactory
func (fs *SrvJWT) Name() string {
	return reflect.TypeOf(fs).Elem().String()
}

//Help implements qz/commands.HelperFactory
func (fs *SrvJWT) Help() string {
	return `
	"component": "qz/helpers.SrvJWT",
	"param": {
		//TODO: fill in a sample
	}
	`
}

//ComponentType implements qz/commands.HelperFactory
func (fs *SrvJWT) ComponentType() reflect.Type {
	return reflect.TypeOf(fs)
}

//CreateHelper returns http.Handler implements qz/commands.HelperFactory.
func (fs *SrvJWT) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (mux interface{}, err error) {

	jwCfg, err := fs.getParams(param)
	if err != nil {
		return
	}
	ctxName := jwCfg.ContentHandlerCtxName

	if ctxName == "" {
		err = commands.NewFatalError(fmt.Sprintf("helper.SrvJWT.CreateHelper: Context Handler missing from param"))
		return
	}
	handlerx := ctx.Value(ctxName)
	if handlerx == nil {
		err = commands.NewFatalError(fmt.Sprintf("context  %v handler not in context", ctxName))
		return
	}
	handler, ok := handlerx.(http.Handler)
	if !ok {
		err = commands.NewFatalError(fmt.Sprintf("context  %v handler not of http.Handler type", ctxName))
		return
	}

	jwkAuth, err := jwx.NewJwkAuth(&jwCfg.JwkAuth)
	mux = jwx.NewJwtWrapper(jwkAuth, handler)

	return
}

func (fs *SrvJWT) getParams(param interface{}) (cfg *SrvJWTConfig, err error) {
	cfg = &SrvJWTConfig{}

	by, err := json.Marshal(param)
	if err != nil {
		err = commands.NewFatalError(err.Error())
		return
	}
	err = json.Unmarshal(by, cfg)

	return
}

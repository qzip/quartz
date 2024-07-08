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
	handler *jwx.JwtHandler
}

// SrvJWTConfig configuration
type SrvJWTConfig struct {
	ContentHandlerCtxName string            `json:"contentHandlerCtxName"`
	JwkAuth               jwx.JwkAuthConfig `json:"jwkAuth"`
}

// Name implements qz/commands.HelperFactory
func (fs *SrvJWT) Name() string {
	return reflect.TypeOf(fs).Elem().String()
}

// Help implements qz/commands.HelperFactory
func (fs *SrvJWT) Help() string {
	return `
	"component": "qz/helpers.SrvJWT",
	"param": {
		//TODO: fill in a sample
	}
	`
}

// ComponentType implements qz/commands.HelperFactory
func (fs *SrvJWT) ComponentType() reflect.Type {
	return reflect.TypeOf(fs)
}

// CreateHelper returns http.Handler implements qz/commands.HelperFactory.
func (fs *SrvJWT) CreateHelper(_ context.Context, param interface{}, _ map[string]interface{}) (mux interface{}, err error) {
	fs.handler = &jwx.JwtHandler{}
	mux = fs.handler
	return
}

func (fs *SrvJWT) InstallChildren(ctx context.Context, param interface{}) error {

	jwCfg, err := fs.getParams(param)
	if err != nil {
		return err
	}
	ctxName := jwCfg.ContentHandlerCtxName

	if ctxName == "" {
		err = commands.NewFatalError("helper.SrvJWT.CreateHelper: Context Handler name missing from param")
		return err
	}
	handlerx := ctx.Value(ctxName)
	if handlerx == nil {
		err = commands.NewFatalError(fmt.Sprintf("context  %v handler not in context", ctxName))
		return err
	}
	handler, ok := handlerx.(http.Handler)
	if !ok {
		err = commands.NewFatalError(fmt.Sprintf("context  %v handler not of http.Handler type", ctxName))
		return err
	}
	fs.handler.Target = handler
	jwAuth, err := jwx.NewJwkAuth(&jwCfg.JwkAuth)
	if err != nil {
		jwx.SetJwtWrapper(fs.handler, jwAuth)
	}

	return err
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

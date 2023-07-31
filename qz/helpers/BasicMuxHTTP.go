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
	"reflect"
)

// Http Mux Handler factory implements qz/commands.HelperFactory

// BasicMuxHTTP implements qz/commands.HelperFactory
type BasicMuxHTTP struct {
}

// Name implements qz/commands.HelperFactory
func (mh *BasicMuxHTTP) Name() string {
	return reflect.TypeOf(mh).Elem().String() //"qz/helpers.BasicMuxHTTP"
}

// Help implements qz/commands.HelperFactory
func (mh *BasicMuxHTTP) Help() string {
	return `
	"component": "qz/helpers.BasicMuxHTTP",
	"param": {
		"path": "HandlerHelperFactCtxName"
	}
	`
}

// ComponentType implements qz/commands.HelperFactory
func (mh *BasicMuxHTTP) ComponentType() reflect.Type {
	return reflect.TypeOf(mh)
}

// CreateHelper returns http.Handler implements qz/commands.HelperFactory.
// returns instance of http.Handler
func (mh *BasicMuxHTTP) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (mux interface{}, err error) {
	m := http.NewServeMux()
	err = mh.installMux(ctx, m, param)
	if err != nil {
		return
	}
	mux = m
	return
}

func (mh *BasicMuxHTTP) installMux(ctx context.Context, mux *http.ServeMux, param interface{}) error {
	if param == nil {
		return commands.NewFatalError("handler.BasicMuxHTTP: parameter expected")

	}
	pathCtx, err := mh.getParams(param)
	if err != nil {
		return err
	}
	// iterate the list
	for path, ctxName := range pathCtx {
		handlerx := ctx.Value(ctxName)
		if handlerx == nil {
			return commands.NewFatalError(fmt.Sprintf("context %v for path %v handler not in context", ctxName, path))
		}
		handler, ok := handlerx.(http.Handler)
		if !ok {
			return commands.NewFatalError(fmt.Sprintf("context %v for path %v handler not of http.Handler type", ctxName, path))

		}
		mux.Handle(path, handler)
	}
	return nil
}

func (mh *BasicMuxHTTP) getParams(param interface{}) (pathCtx map[string]string, err error) {
	pathCtx = make(map[string]string)
	by, err := json.Marshal(param)
	if err != nil {
		err = commands.NewFatalError(err.Error())
		return
	}
	err = json.Unmarshal(by, &pathCtx)

	return
}

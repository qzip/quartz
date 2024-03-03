package compfact

import (
	"context"
	"encoding/json"
	"jshttp/comp"
	"qz/commands"
	"reflect"
)

type JavaScriptRunFactory struct {
}

// ComponentType implements qz/commands.HelperFactory
func (jsf *JavaScriptRunFactory) Name() string {
	return "json"
}

// ComponentType implements qz/commands.HelperFactory
func (jsf *JavaScriptRunFactory) Help() string {
	help := ``
	return help
}

// ComponentType implements qz/commands.HelperFactory
func (jsf *JavaScriptRunFactory) ComponentType() reflect.Type {
	return reflect.TypeOf(jsf)
}

// CreateHelper returns http.Handler implements qz/commands.HelperFactory.
func (jsf *JavaScriptRunFactory) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (jsr interface{}, err error) {
	if param == nil {
		return nil, commands.NewFatalError(jsf.Name() + ": parameter expected")

	}
	params, err := jsf.getParams(param)
	if err != nil {
		return
	}
	jsr = &comp.JsHandler{
		Ctx:    ctx,
		Cfg:    cfg,
		Params: *params,
	}
	return
}

func (jsf *JavaScriptRunFactory) getParams(param interface{}) (scriptParams *comp.JsHandlerParam, err error) {
	scriptParams = &comp.JsHandlerParam{}
	by, err := json.Marshal(param)
	if err != nil {
		err = commands.NewFatalError(err.Error())
		return
	}
	err = json.Unmarshal(by, scriptParams)

	return
}

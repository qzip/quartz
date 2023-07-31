package comp

import (
	"context"
	"net/http"
)

type JsHandler struct {
	Ctx      context.Context        // for accessing Helpers from JS VM
	Cfg      map[string]interface{} // full config is available to JS VM
	Params   map[string]string
	Function string // <Function Name>(Request , Response)
	Script   string
}

func (jsh *JsHandler) ServeHTTP(http.ResponseWriter, *http.Request) {

}

func (jsh *JsHandler) ExtractJs() error {

	return nil
}

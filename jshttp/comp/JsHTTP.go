package comp

import (
	"context"
	"net/http"
)

/*
JsHandler is a http.Handler that can be used to serve HTTP requests using a JS function.
The JS function is called with a Request and Response object.

Each URL path can be associated with a different JS function.
*/
type JsHandler struct {
	Ctx    context.Context        // for accessing Helpers from JS VM
	Cfg    map[string]interface{} // full config is available to JS VM
	Params JsHandlerParam
}

type JsHandlerParam struct {
	Function string `json:"function"` // <Function Name>(Request , Response)
	Script   string `json:"script"`
}

// ServeHTTP implements http.Handler
func (jsh *JsHandler) ServeHTTP(http.ResponseWriter, *http.Request) {

}

func (jsh *JsHandler) ExtractJs() error {

	return nil
}

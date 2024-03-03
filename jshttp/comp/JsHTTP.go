package comp

import (
	"context"
	"net/http"
	"os"

	//"github.com/Workiva/go-datastructures/threadsafe/err"
	"github.com/dop251/goja"
)

/*
JsHandler is a http.Handler that can be used to serve HTTP requests using a JS function.
The JS function is called with a Request and Response object.

Each URL path can be associated with a different JS function.
*/
type JsHandler struct {
	Ctx    context.Context        // for accessing Helpers from JS VM
	Cfg    map[string]interface{} // full config is available to JS VM
	prog   *goja.Program
	Params JsHandlerParam
}

type JsHandlerParam struct {
	Function string `json:"function"` // <Function Name>(Request , Response)
	Script   string `json:"script"`
}

// ServeHTTP implements http.Handler
func (jsh *JsHandler) ServeHTTP(req http.ResponseWriter, res *http.Request) {
	//call the JS function
	vm := goja.New()
	reqVal := vm.ToValue(&req)
	resVal := vm.ToValue(res)
	_, err := vm.RunProgram(jsh.prog)
	if err != nil {
		http.Error(req, err.Error(), http.StatusInternalServerError)
		return
	}
	serv, ok := goja.AssertFunction(vm.Get(jsh.Params.Function))
	if !ok {
		err := (jsh.Params.Function + " not found in" + jsh.Params.Script)
		http.Error(req, err, http.StatusInternalServerError)
		return
	}

	if _, err := serv(nil, reqVal, resVal); err != nil {
		http.Error(req, err.Error(), http.StatusInternalServerError)
		return
	}
	//res = resVal.Export().(*http.ResponseWriter)

}

func (jsh *JsHandler) ExtractJs() error {
	prog, err := os.ReadFile(jsh.Params.Script)
	if err != nil {
		return err
	}
	jsh.prog, err = goja.Compile(jsh.Params.Script, string(prog), true)
	if err != nil {
		return err
	}
	return nil
}

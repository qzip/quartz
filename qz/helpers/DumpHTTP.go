package helpers

import (
	"context"
	"net/http"
	u "net/http/httputil"
	"reflect"
)

// DumpHTTP implements qz/commands.HelperFactory
type DumpHTTP struct {
}

//Name implements qz/commands.HelperFactory
func (dh *DumpHTTP) Name() string {
	return reflect.TypeOf(dh).Elem().String()
}

//Help implements qz/commands.HelperFactory
func (dh *DumpHTTP) Help() string {
	return `
	"component": "qz/helpers.DumpHTTP",
	"param": ""
	`
}

//ComponentType implements qz/commands.HelperFactory
func (dh *DumpHTTP) ComponentType() reflect.Type {
	return reflect.TypeOf(dh)
}

//CreateHelper returns http.Handler implements qz/commands.HelperFactory.
func (dh *DumpHTTP) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (hnd interface{}, err error) {
	hnd = dh
	return
}

// ServeHTTP implements https://golang.org/pkg/net/http/#Handler
func (dh *DumpHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := u.DumpRequest(r, true)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		w.Write([]byte("\nRemote Address:[" + r.RemoteAddr + "]\n"))
	}

}

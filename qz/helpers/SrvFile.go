package helpers

import (
	"context"
	"fmt"
	"net/http"
	"qz/commands"
	"reflect"
)

// SrvFile implements qz/commands.HelperFactory
type SrvFile struct {
}

//Name implements qz/commands.HelperFactory
func (fs *SrvFile) Name() string {
	return reflect.TypeOf(fs).Elem().String()
}

//Help implements qz/commands.HelperFactory
func (fs *SrvFile) Help() string {
	return `
	"component": "qz/helpers.SrvFile",
	"param": "/path/to/root"
	`
}

//ComponentType implements qz/commands.HelperFactory
func (fs *SrvFile) ComponentType() reflect.Type {
	return reflect.TypeOf(fs)
}

//CreateHelper returns http.Handler implements qz/commands.HelperFactory.
func (fs *SrvFile) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (hnd interface{}, err error) {
	if param == nil {
		err = commands.NewFatalError("parameter expected")
		return
	}
	root, ok := param.(string)
	if !ok {
		err = commands.NewFatalError(fmt.Errorf("helper.SrvFile.CreateHelper: %v not a string", param).Error())
		return
	}
	if root == "" {
		err = commands.NewFatalError(fmt.Errorf("helper.SrvFile.CreateHelper: empty root param").Error())
		return
	}
	hnd = http.FileServer(http.Dir(root))

	return
}

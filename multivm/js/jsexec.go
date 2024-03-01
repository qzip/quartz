package js

import (
	"context"
	"os"
	"qz/util"

	"github.com/dop251/goja"
)

func JsExec(ctx context.Context, filename string) error {

	buf, err := os.ReadFile(filename)
	if err != nil {
		util.DebugInfo(ctx, err.Error())
		return err
	}
	vm := goja.New() // goja.Runtime
	var randsrc goja.RandSource
	randsrc, err = NewRandSource()
	if err != nil {
		util.DebugInfo(ctx, err.Error())
		return err
	}

	vm.SetRandSource(randsrc)
	// install the functions from helpers
	// vm.Set(name, interface{})
	jval, err := vm.RunScript(filename, string(buf))
	if err != nil {
		util.DebugInfo(ctx, err.Error())
		return err
	}
	util.DebugInfo(ctx, jval.String())

	return nil
}

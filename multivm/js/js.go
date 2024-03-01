package js

import (
	"context"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"qz/commands"
	"qz/util"

	crand "crypto/rand"

	"github.com/dop251/goja"
)

type VMjs struct {
}

// Name implements command interface method
func (vjs *VMjs) Name() string {
	return "vm_js"
}

// Help implements command interface method
func (vjs *VMjs) Help() string {
	return `
	  version: 31jul23 
      JS VM, file extension .js,  "Goja is an implementation of ECMAScript 5.1 in pure Go"
	  see: https://github.com/dop251/goja
	`
}

func (vjs *VMjs) Exec(ctx context.Context, cfg map[string]interface{}, errChan chan error) {

	src := cfg[commands.CmdVmSource].(string)
	//filename := cfg[commands.CmdFileName].(string)
	buf, err := os.ReadFile(src)
	if err != nil {
		util.DebugInfo(ctx, err.Error())
		errChan <- err
		return
	}
	// FIXME:
	fmt.Print(buf)   // remove it
	vm := goja.New() // goja.Runtime
	var randsrc goja.RandSource
	randsrc, err = newRandSource()
	if err != nil {
		util.DebugInfo(ctx, err.Error())
		errChan <- err
		return
	}

	vm.SetRandSource(randsrc)
	// install the functions from helpers

	jval, err := vm.RunString(string(buf))
	if err != nil {
		util.DebugInfo(ctx, err.Error())
		errChan <- err
	} else {
		util.DebugInfo(ctx, jval.String())
	}
}

// TODO: this is a source of non determinism, should not be used for smart contracts
// set it to throw exception or return zero to make it deterministic?
func newRandSource() (goja.RandSource, error) {
	var seed int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed); err != nil {
		return nil, fmt.Errorf("Could not read random bytes: %v", err)
	}
	return rand.New(rand.NewSource(seed)).Float64, nil
}

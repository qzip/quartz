package compfact

import (
	"context"
	"encoding/json"
	"l2/algorand/comp"
	"qz/commands"
	"qz/seq"
	"reflect"
)

type Md5HashFileFact struct {
}

// implements qz/command/ComponentFactory
func (anf *Md5HashFileFact) Name() string {
	return "comp.Md5HashFileFactory"
}

// implements qz/command/ComponentFactory
func (anf *Md5HashFileFact) Help() string {
	help := ``
	return help
}

// implements qz/command/ComponentFactory
func (anf *Md5HashFileFact) ComponentType() reflect.Type {
	return reflect.TypeOf(anf)
}

// Create implements seq.Runner interface method
func (anf *Md5HashFileFact) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	an := comp.Md5HashFile{}
	by, err := json.Marshal(param)
	if err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errCh <- err
		return nil
	}
	err = json.Unmarshal(by, &an)
	if err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errCh <- err
		return nil
	}
	an.SetCtxHelper(helper)
	an.SetChanErr(errCh)

	return &an
}

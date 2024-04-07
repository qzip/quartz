package compfact

import (
	"context"
	"encoding/json"
	"l2/algorand/comp"
	"qz/commands"
	"qz/seq"
	"reflect"
)

type AlgoNotarizeFact struct {
}

// implements qz/command/ComponentFactory
func (anf *AlgoNotarizeFact) Name() string {
	return "comp.DoltInsertFactory"
}

// implements qz/command/ComponentFactory
func (anf *AlgoNotarizeFact) Help() string {
	help := ``
	return help
}

// implements qz/command/ComponentFactory
func (anf *AlgoNotarizeFact) ComponentType() reflect.Type {
	return reflect.TypeOf(anf)
}

// Create implements seq.Runner interface method
func (anf *AlgoNotarizeFact) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	an := comp.AlgoNotarize{}
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

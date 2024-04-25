package compfact

import (
	"context"
	"encoding/json"
	"l2/algorand/comp"
	"qz/commands"
	"qz/seq"
	"reflect"
)

type AlgoTxn2FileFact struct {
}

// implements qz/command/ComponentFactory
func (anf *AlgoTxn2FileFact) Name() string {
	return "comp.AlgoTxn2FileFact"
}

// implements qz/command/ComponentFactory
func (anf *AlgoTxn2FileFact) Help() string {
	help := `

	`
	return help
}

// implements qz/command/ComponentFactory
func (anf *AlgoTxn2FileFact) ComponentType() reflect.Type {
	return reflect.TypeOf(anf)
}

// Create implements seq.Runner interface method
func (anf *AlgoTxn2FileFact) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	an := comp.AlgoTxn2File{}
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

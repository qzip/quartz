package compfact

import (
	"context"
	"encoding/json"
	"qz/commands"
	"qz/seq"
	"reflect"
)

type component interface {
	commands.Pipeline
	SetCtxHelper(helper seq.CtxHelper)
	SetChanErr(errCh chan error)
}

func NewCompFact(name, help string, comp component) *GenFact {
	return &GenFact{help: help, name: name, comp: comp}
}

// GenFact implements qz/command/ComponentFactory
type GenFact struct {
	help string
	name string
	comp component
}

func (gf *GenFact) ComponentType() reflect.Type {
	return reflect.TypeOf(gf.comp)
}

func (gf *GenFact) Help() string {
	return gf.help
}

func (gf *GenFact) Name() string {
	return gf.name
}

func (gf *GenFact) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	by, err := json.Marshal(param)
	if err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errCh <- err
		return nil
	}
	err = json.Unmarshal(by, &gf.comp)
	if err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errCh <- err
		return nil
	}
	gf.comp.SetCtxHelper(helper)
	gf.comp.SetChanErr(errCh)

	return gf.comp

}

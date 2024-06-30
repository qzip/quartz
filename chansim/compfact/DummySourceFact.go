package compfact

import (
	"chansim/comp"
	"context"
	"fmt"
	"qz/util"
	"reflect"
)

type DummySourceFact struct {
}

// Name implements qz/command/ComponentFactory
func (acf *DummySourceFact) Name() string {
	return reflect.TypeOf(*acf).String()
}

// ComponentType implements qz/command/ComponentFactory
func (acf *DummySourceFact) ComponentType() reflect.Type {
	return reflect.TypeOf(*acf)
}

// CreateHelper implements qz/commands/InstHelpers.go#HelperFactory
func (acf *DummySourceFact) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (ret interface{}, err error) {
	c := &comp.DummySource{}
	ret = c
	if param != nil {
		p, ok := param.(int)
		if !ok {
			util.DebugInfo(ctx, fmt.Sprintf("%v: expected param as int got [%v]", acf.Name(), reflect.TypeOf(param).String()))
		}
		c.MaxCount = p
	}
	return
}

// implements qz/command/ComponentFactory
func (acf *DummySourceFact) Help() string {
	return `
= DummyTransformerFact

creates the DummyTransformer helper component for ChanExec channel pipeline.

	
	`
}

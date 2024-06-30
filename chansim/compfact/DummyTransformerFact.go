package compfact

import (
	"chansim/comp"
	"context"
	"reflect"
)

type DummyTransformerFact struct {
}

// Name implements qz/command/ComponentFactory
func (acf *DummyTransformerFact) Name() string {
	return reflect.TypeOf(*acf).String()
}

// ComponentType implements qz/command/ComponentFactory
func (acf *DummyTransformerFact) ComponentType() reflect.Type {
	return reflect.TypeOf(*acf)
}

// CreateHelper implements qz/commands/InstHelpers.go#HelperFactory
func (acf *DummyTransformerFact) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (ret interface{}, err error) {
	ret = &comp.DummyTransformer{}
	return
}

// implements qz/command/ComponentFactory
func (acf *DummyTransformerFact) Help() string {
	return `
= DummyTransformerFact

creates the DummyTransformer helper component for ChanExec channel pipeline.

	
	`
}

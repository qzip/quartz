package compfact

import (
	"chansim/comp"
	"context"
	"reflect"
)

type DummySinkFact struct {
}

// Name implements qz/command/ComponentFactory
func (acf *DummySinkFact) Name() string {
	return reflect.TypeOf(*acf).String()
}

// ComponentType implements qz/command/ComponentFactory
func (acf *DummySinkFact) ComponentType() reflect.Type {
	return reflect.TypeOf(*acf)
}

// CreateHelper implements qz/commands/InstHelpers.go#HelperFactory
func (acf *DummySinkFact) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (ret interface{}, err error) {
	ret = &comp.DummySink{}

	return
}

// implements qz/command/ComponentFactory
func (acf *DummySinkFact) Help() string {
	return `
= DummyTransformerFact

creates the DummyTransformer helper component for ChanExec channel pipeline.

	
	`
}

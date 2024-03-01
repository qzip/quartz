package js

import (
	"context"
	"reflect"

	"github.com/dop251/goja"
)

type JsInstaller interface {
	Install(ctx context.Context, vm *goja.Runtime) error
}

// install event infrastructure
// implements command.ComponentFactory
type EventInfra struct {
}

func (ei *EventInfra) Name() string {
	return "event"
}

func (ei *EventInfra) Help() string {
	return `
	   installs event handling infrastructure (pub/sub)
	`
}

func (ei *EventInfra) ComponentType() reflect.Type {
	return reflect.TypeOf(ei)
}

func (ei *EventInfra) Install(ctx context.Context, vm *goja.Runtime) error {

	return nil
}

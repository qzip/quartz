package starlark

import (
	"reflect"

	"go.starlark.net/lib/json"
	"go.starlark.net/lib/math"
	"go.starlark.net/lib/time"
	"go.starlark.net/starlarkstruct"
)

/**
Define the components forbuiltin modules.
Implements json, time and math buildin modules factories
*/

type JsonModFactory struct {
}

// ComponentType implements qz/commands.HelperFactory
// see: go.starlark.net/cmd/starlark/starlark.go
func (jmf *JsonModFactory) Name() string {
	return "json"
}

// ComponentType implements qz/commands.HelperFactory
func (jmf *JsonModFactory) Help() string {
	help := `creates the built in json module
	   see: https://github.com/google/starlark-go/  
	   starlark-go/lib/json/
	`
	return help
}

// ComponentType implements qz/commands.HelperFactory
func (jmf *JsonModFactory) ComponentType() reflect.Type {
	return reflect.TypeOf(jmf)
}

// CreateModule implements StarlarkModule
// lib/json/json.go#L75
func (jmf *JsonModFactory) CreateModule() *starlarkstruct.Module {
	return json.Module
}

// TimeModFactory time module factory
type TimeModFactory struct {
}

// ComponentType implements qz/commands.HelperFactory
// see: go.starlark.net/cmd/starlark/starlark.go
func (tmf *TimeModFactory) Name() string {
	return "time"
}

// ComponentType implements qz/commands.HelperFactory
func (tmf *TimeModFactory) Help() string {
	help := `creates the built in json module
	   see: https://github.com/google/starlark-go/  
	   starlark-go/lib/time/
	`
	return help
}

// ComponentType implements qz/commands.HelperFactory
func (tmf *TimeModFactory) ComponentType() reflect.Type {
	return reflect.TypeOf(tmf)
}

// CreateModule implements StarlarkModule
// lib/json/json.go#L75
func (tmf *TimeModFactory) CreateModule() *starlarkstruct.Module {
	return time.Module
}

// MathModFactory time module factory
type MathModFactory struct {
}

// ComponentType implements qz/commands.HelperFactory
// see: go.starlark.net/cmd/starlark/starlark.go
func (mmf *MathModFactory) Name() string {
	return "math"
}

// ComponentType implements qz/commands.HelperFactory
func (mmf *MathModFactory) Help() string {
	help := `creates the built in json module
	   see: https://github.com/google/starlark-go/  
	   starlark-go/lib/math/
	`
	return help
}

// ComponentType implements qz/commands.HelperFactory
func (mmf *MathModFactory) ComponentType() reflect.Type {
	return reflect.TypeOf(mmf)
}

// CreateModule implements StarlarkModule
// lib/json/json.go#L75
func (mmf *MathModFactory) CreateModule() *starlarkstruct.Module {
	return math.Module
}

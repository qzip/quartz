package compfact

import (
	"Dolt/comp"
	"context"
	"encoding/json"
	"qz/commands"
	"reflect"
)

// see: qz/commands/InstHelpers.go

type DoltClientFactory struct {
}

// implements qz/command/ComponentFactory
func (acf *DoltClientFactory) Name() string {
	return "comp.DoltClientFactory"
}

// implements qz/command/ComponentFactory
func (acf *DoltClientFactory) Help() string {
	help := `
	"helpers": [
        {
            "component": "comp.DoltClientFactory",
            "name": "DoltClientLocal",
            "param": {
                "dbURL": "http://localhost:8529",
                "user": "root",
                "pass" : "$Dolt_PASS_LOCAL"
            }
        }
    ]	
	`
	return help
}

// implements qz/command/ComponentFactory
func (acf *DoltClientFactory) ComponentType() reflect.Type {
	return reflect.TypeOf(acf)
}

// implements qz/commands/InstHelpers.go#HelperFactory
func (acf *DoltClientFactory) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (client interface{}, err error) {
	by, err := json.Marshal(param)
	if err != nil {
		err = commands.NewFatalError(err.Error())
		return
	}
	client = &comp.DoltClient{}
	err = json.Unmarshal(by, &client)
	if err != nil {
		return
	}

	return
}

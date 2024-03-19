package compfact

import (
	"arango/comp"
	"context"
	"encoding/json"
	"qz/commands"
	"reflect"
)

// see: qz/commands/InstHelpers.go

type ArangoClientFactory struct {
}

// implements qz/command/ComponentFactory
func (acf *ArangoClientFactory) Name() string {
	return "comp.ArangoClientFactory"
}

// implements qz/command/ComponentFactory
func (acf *ArangoClientFactory) Help() string {
	help := `
	"helpers": [
        {
            "component": "comp.ArangoClientFactory",
            "name": "arangoClientLocal",
            "param": {
                "dbURL": "http://localhost:8529",
                "user": "root",
                "pass" : "$ARANGO_PASS_LOCAL"
            }
        }
    ]	
	`
	return help
}

// implements qz/command/ComponentFactory
func (acf *ArangoClientFactory) ComponentType() reflect.Type {
	return reflect.TypeOf(acf)
}

// implements qz/commands/InstHelpers.go#HelperFactory
func (acf *ArangoClientFactory) CreateHelper(ctx context.Context, param interface{}, cfg map[string]interface{}) (client interface{}, err error) {
	by, err := json.Marshal(param)
	if err != nil {
		err = commands.NewFatalError(err.Error())
		return
	}
	client = &comp.ArangoClient{}
	err = json.Unmarshal(by, &client)
	if err != nil {
		return
	}

	return
}

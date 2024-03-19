package compfact

import (
	"arango/comp"
	"context"
	"encoding/json"
	"qz/commands"
	"qz/seq"
	"reflect"
)

type ArangoInsertFactory struct {
}

// implements qz/command/ComponentFactory
func (acf *ArangoInsertFactory) Name() string {
	return "comp.ArangoInsertFactory"
}

// implements qz/command/ComponentFactory
func (acf *ArangoInsertFactory) Help() string {
	help := `
	{
		"component": "comp.ArangoInsertFactory",
		"param": {
			"database": "kheti-badi",
			"collection": "categories",
			"dbHelper": "arangoClientLocal",
			"DataInCtxName": "data",
			"createDbIfNotExists": true,
			"createCollectionIfNotExists": true,
			"updateDocumentIfExists": true
		}
	}
	`
	return help
}

// implements qz/command/ComponentFactory
func (acf *ArangoInsertFactory) ComponentType() reflect.Type {
	return reflect.TypeOf(acf)
}

// Create implements seq.Runner interface method
func (acf *ArangoInsertFactory) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	ai := &comp.ArangoInsert{}
	by, err := json.Marshal(param)
	if err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errCh <- err
		return nil
	}
	err = json.Unmarshal(by, &ai)
	if err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errCh <- err
		return nil
	}
	ai.SetCtxHelper(helper)
	ai.SetChanErr(errCh)

	return ai
}

package compfact

import (
	"Dolt/comp"
	"context"
	"encoding/json"
	"qz/commands"
	"qz/seq"
	"reflect"
)

type DoltInsertFactory struct {
}

// implements qz/command/ComponentFactory
func (acf *DoltInsertFactory) Name() string {
	return "comp.DoltInsertFactory"
}

// implements qz/command/ComponentFactory
func (acf *DoltInsertFactory) Help() string {
	help := `
	{
		"component": "comp.DoltInsertFactory",
		"param": {
			"database": "kheti-badi",
			"collection": "categories",
			"dbHelper": "DoltClientLocal",
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
func (acf *DoltInsertFactory) ComponentType() reflect.Type {
	return reflect.TypeOf(acf)
}

// Create implements seq.Runner interface method
func (acf *DoltInsertFactory) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	ai := &comp.DoltInsert{}
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

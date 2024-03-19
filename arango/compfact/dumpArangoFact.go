package compfact

import (
	"arango/comp"
	"context"
	"encoding/json"
	"qz/commands"
	"qz/seq"
	"reflect"
)

type DumpArangoFactory struct {
}

// implements qz/command/ComponentFactory
func (acf *DumpArangoFactory) Name() string {
	return "comp.DumpArangoFactory"
}

// implements qz/command/ComponentFactory
func (acf *DumpArangoFactory) Help() string {
	help := `
	{
		"component": "comp.DumpArangoFactory",
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
func (acf *DumpArangoFactory) ComponentType() reflect.Type {
	return reflect.TypeOf(acf)
}

// Create implements seq.Runner interface method
func (acf *DumpArangoFactory) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	ai := &comp.DumpArango{}
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

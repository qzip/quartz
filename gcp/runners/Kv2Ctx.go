package runners

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/commands"
	"qz/seq"
	"qz/util"
	"reflect"
)

//Kv2Ctx  inserts param to help context
type Kv2Ctx struct {
	helper seq.CtxHelper
}

//Name implements commands.ComponentFactory methods
func (t *Kv2Ctx) Name() string {
	return "runners.Kv2Ctx"
}

//Help implements commands.ComponentFactory methods
func (t *Kv2Ctx) Help() string {

	return `
	{
		"component" : "firebase.Kv2Ctx",
		"param" : {
			"key": "val",
			"k1", "v2"
		}
	}	   
	`
}

//ComponentType implements commands.ComponentFactory methods
func (t *Kv2Ctx) ComponentType() reflect.Type {
	return reflect.TypeOf(t)
}

//Create implements seq.Runner interface
func (t *Kv2Ctx) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errChan chan error) commands.Pipeline {
	t.helper = helper
	kv := make(map[string]interface{}, 0)
	by, err := json.Marshal(param)
	if err != nil {
		errChan <- commands.NewFatalError("Kv2Ctx.Create: param marshalling:" + err.Error())
		return nil
	}
	err = json.Unmarshal(by, &kv)
	if err != nil {
		errChan <- commands.NewFatalError("Kv2Ctx.Create: param unmarshalling to map[string]string:" + err.Error())
		return nil
	}
	for k, v := range kv {
		t.helper.SetKeyValue(k, v)
		util.DebugInfo(ctx, fmt.Sprintf("Kv2Ctx.Create: Setting ctx DocRecID [%v] = %v \n", k, v))

	}

	helper.SetExecStatus(seq.ExSinit)
	return t
}

//Process implements commands.Pipeline method
func (t *Kv2Ctx) Process(ctx context.Context) {

	t.helper.SetExecStatus(seq.ExSok)
}

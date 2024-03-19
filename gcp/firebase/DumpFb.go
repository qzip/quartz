package firebase

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/commands"
	"qz/seq"
	"qz/util"
	"reflect"

	firebase "firebase.google.com/go/v4/db"
)

// DumpFb dump the Firebase DB to a Zip or PostgreSQL
// format: collection name->key->record
// assumes that the database can fit into memory (db.Size << mem.Size)
type DumpFb struct {
	fsClient *firebase.Client
	helper   seq.CtxHelper
	errCh    chan error
	param    *DumpFbParam
}

// DumpFbParam parameters
type DumpFbParam struct {
	FirebaseClientCtxName string `json:"firebaseClientCtxName"`
	Path                  string `json:"path,omitempty"`
	DataOutCtxName        string `json:"dataOutCtxName"`
}

//***** component seq.Runner interface methods *****

// Name implements seq.Runner interface method
func (ks *DumpFb) Name() string {
	return reflect.TypeOf(ks).Elem().String() // "firebase.DumpFb"
}

// Help implements seq.Runner interface method
func (ks *DumpFb) Help() string {
	return `
	  # cmd.RunSeq: stage 
	  {
		  "component" : "firebase.UniDocSubCols",
		  "param" : {
			#compoent specific parameters
			}
	  }
	`
}

// ComponentType implements seq.Runner interface method
func (ks *DumpFb) ComponentType() reflect.Type {
	return reflect.TypeOf(ks)
}

// Create implements seq.Runner interface method
func (ks *DumpFb) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	ks.helper = helper
	ks.errCh = errCh

	if err := ks.getParams(ctx, param, cfg); err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errCh <- err
		return nil
	}
	if err := ks.getHelper(helper); err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errCh <- err
		return nil
	}
	helper.SetExecStatus(seq.ExSinit)
	return ks
}

// Process implements Pipeline (required by seq.RunSeq)
func (ks *DumpFb) Process(ctx context.Context) {
	ks.helper.SetExecStatus(seq.ExSrunning)
	// assumes format: collectioname->key->record
	//           [collection][doc][key]Value
	//           address->1fNRRkDCX9iadxUTnEhZ -> phone -> 98123456
	//https://pkg.go.dev/firebase.google.com/go/v4@v4.2.0/db#Ref
	ref := ks.fsClient.NewRef(ks.param.Path)
	util.DebugInfo(ctx, fmt.Sprintf("DumpFb.Process: key[%v] path[%v]\n", ref.Key, ref.Path))
	data := make(map[string]interface{})
	if err := ref.Get(ctx, &data); err != nil {
		util.DebugInfo(ctx, err.Error())
		ks.errCh <- err
		ks.helper.SetExecStatus(seq.ExSerror)
		return
	}
	ks.helper.SetKeyValue(ks.param.DataOutCtxName, data)
	ks.helper.SetExecStatus(seq.ExSok)
	return
}

func (ks *DumpFb) getHelper(helper seq.CtxHelper) (err error) {
	h := helper.Value(ks.param.FirebaseClientCtxName)
	if fsc, ok := h.(*firebase.Client); ok {
		ks.fsClient = fsc
	} else {
		err = fmt.Errorf("DumpFb.getHelper: %v context helper not of type *firebase.Client", ks.param.FirebaseClientCtxName)
	}
	return
}
func (ks *DumpFb) getParams(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	if param == nil {
		return fmt.Errorf("DumpFb.getParams: nil param")
	}
	ks.param = &DumpFbParam{}
	by, err := json.Marshal(param)
	if err != nil {
		return err
	}
	err = json.Unmarshal(by, ks.param)
	if err != nil {
		return err
	}

	return nil
}

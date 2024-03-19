package firestore

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/commands"
	"qz/seq"
	"reflect"

	"cloud.google.com/go/firestore"
)

// DumpFs dump the Firebase DB to a Zip or PostgreSQL
// format: collection name->key->record
// assumes that the database can fit into memory (db.Size << mem.Size)
type DumpFs struct {
	fsClient *firestore.Client
	helper   seq.CtxHelper
	errCh    chan error
	param    *DumpFsParam
}

// DumpFsParam parameters
type DumpFsParam struct {
	FirebaseClientCtxName string `json:"firebaseClientCtxName"`
	DataOutCtxName        string `json:"dataOutCtxName"`
}

//***** component seq.Runner interface methods *****

// Name implements seq.Runner interface method
func (ks *DumpFs) Name() string {
	return reflect.TypeOf(ks).Elem().String() // "khetose.DumpFs"
}

// Help implements seq.Runner interface method
func (ks *DumpFs) Help() string {
	return `
	  # cmd.RunSeq: stage 
	  {
		  "component" : "firestore.UniDocSubCols",
		  "param" : {
			#component specific parameters
			}
	  }
	`
}

// ComponentType implements seq.Runner interface method
func (ks *DumpFs) ComponentType() reflect.Type {
	return reflect.TypeOf(ks)
}

// Create implements seq.Runner interface method
func (ks *DumpFs) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
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
func (ks *DumpFs) Process(ctx context.Context) {
	ks.helper.SetExecStatus(seq.ExSrunning)
	// assumes format: collection name->key->record
	//           [collection][doc][key]Value
	//           address->1fNRRkDCX9iadxUTnEhZ -> phone -> 98123456
	data := make(map[string]map[string]map[string]interface{})
	// get all the collections
	// https://pkg.go.dev/cloud.google.com/go/firestore#Client.Collections

	cols, err := ks.fsClient.Collections(ctx).GetAll()
	if err != nil {
		ks.helper.SetExecStatus(seq.ExSerror)
		ks.errCh <- err
		return
	}
	// top level collections, like, "users", "address"
	for _, colRef := range cols {
		/*	path := colRef.Path

			n := strings.LastIndex(path, "/")
			key := ""
			if n < 0 {
				key = path // no '/' found
			}
			if n+1 < len(path) {
				key = path[n+1:]
			}
		*/
		key := colRef.ID
		data[key], err = ks.getDocuments(ctx, colRef) //map[string]map[string]interface{}
		if err != nil {
			ks.helper.SetExecStatus(seq.ExSerror)
			ks.errCh <- err
			return
		}
	}
	ks.helper.SetKeyValue(ks.param.DataOutCtxName, data)
	ks.helper.SetExecStatus(seq.ExSok)
	return
}
func (ks *DumpFs) getDocuments(ctx context.Context, colRef *firestore.CollectionRef) (docs map[string]map[string]interface{}, err error) {
	docs = make(map[string]map[string]interface{})
	if docRefs, er := colRef.DocumentRefs(ctx).GetAll(); er == nil {
		for _, docRef := range docRefs {
			docs[docRef.ID] = make(map[string]interface{})
			snap, er := docRef.Get(ctx)
			if er != nil {
				err = er
				return
			}
			docs[docRef.ID] = snap.Data()
		}
	} else {
		err = er
	}

	return
}
func (ks *DumpFs) getHelper(helper seq.CtxHelper) (err error) {
	h := helper.Value(ks.param.FirebaseClientCtxName)
	if fsc, ok := h.(*firestore.Client); ok {
		ks.fsClient = fsc
	} else {
		err = fmt.Errorf("DumpFs.getHelper: %v context helper not of type *firestore.Client", ks.param.FirebaseClientCtxName)
	}
	return
}
func (ks *DumpFs) getParams(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	if param == nil {
		return fmt.Errorf("DumpFs.getParams: nil param")
	}
	ks.param = &DumpFsParam{}
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

/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package firebase

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/commands"
	"qz/seq"
	"reflect"

	firebase "firebase.google.com/go/v4/db"
)

// DefaFsClient firebase base component registration name
const DefaFsClient = "firebase.Client"

// FbCollectionLoader implements seq.RunSeq
// This component may be called with different arguments
// by difference factory classes, so should not store
// instance info.
type FbCollectionLoader struct {
}

// CfgFbCollectionLoader contructs init params
type CfgFbCollectionLoader struct {
	FbClientCtxKey string `json:"FbClientCtxKey,omitempty"`
	CollectionName string `json:"collectionName"`
	DataOutCtxName string `json:"dataOutCtxName"`
}

// Create implements seq.Runner interface method
func (fs *FbCollectionLoader) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errChan chan error) commands.Pipeline {
	if param == nil {
		helper.SetExecStatus(seq.ExSerror)
		errChan <- fmt.Errorf("FbCollectionLoader.Load: no param specified ")
		return nil
	}
	ret, err := fs.getParams(ctx, param, cfg)
	if err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errChan <- err
		return nil
	}
	ret.helper = helper
	helper.SetExecStatus(seq.ExSinit)
	ret.errChan = errChan
	return ret
}

func (fs *FbCollectionLoader) getParams(ctx context.Context, param interface{}, cfg map[string]interface{}) (rl *fbRecLoader, err error) {
	rl = &fbRecLoader{}
	cfgRl := &CfgFbCollectionLoader{}
	by, err := json.Marshal(param)
	if err != nil {
		return
	}
	err = json.Unmarshal(by, cfgRl)
	if err != nil {
		return
	}

	if len(cfgRl.FbClientCtxKey) == 0 {
		cfgRl.FbClientCtxKey = DefaFsClient
	}
	ok := false
	if rl.fsClient, ok = ctx.Value(cfgRl.FbClientCtxKey).(*firebase.Client); !ok {
		err = commands.NewFatalError(fmt.Sprintf("FbCollectionLoader.getParams: firebase context key %v not instantiated", cfgRl.FbClientCtxKey))
		return
	}
	rl.collectionName = cfgRl.CollectionName
	rl.dataCtx = cfgRl.DataOutCtxName
	return
}

type fbRecLoader struct {
	fsClient       *firebase.Client
	collectionName string
	errChan        chan error
	helper         seq.CtxHelper
	dataCtx        string
}

func (fs *fbRecLoader) Process(ctx context.Context) {
	recs := make(map[string]map[string]interface{}, 0)
	// https://pkg.go.dev/firebase.google.com/go/v4@v4.1.0/db#Client.NewRef
	ref := fs.fsClient.NewRef(fs.collectionName)
	err := ref.Get(ctx, &recs)
	if err != nil {
		fs.helper.SetExecStatus(seq.ExSerror)
		fs.errChan <- err
		return
	}
	fs.helper.SetKeyValue(fs.dataCtx, recs)
	fs.helper.SetExecStatus(seq.ExSok)
	return
}

//**** component methods *****

// Name to register in the component registry
func (fs *FbCollectionLoader) Name() string {
	return "firestore.FbCollectionLoader"
}

// Help return help string
func (fs *FbCollectionLoader) Help() string {
	return `
	  
	  
	  For printing to a file, specify the file path.
	  "param": {

	  }  
	`
}

// ComponentType returns the component type
func (fs *FbCollectionLoader) ComponentType() reflect.Type {
	return reflect.TypeOf(fs)
}

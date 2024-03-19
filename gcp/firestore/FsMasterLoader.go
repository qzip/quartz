/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package firestore

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/commands"
	"qz/util"
	"reflect"

	"cloud.google.com/go/firestore"
)

// FbMasterLoader implements runners.RecsMasterLoader interface &
// loads  Master from firestore.
// This component may be called with different arguments
// by difference factory classes, so should not store
// instance info.
type FbMasterLoader struct {
}

// CfgFbMasterLoader contructs init params
type CfgFbMasterLoader struct {
	FsClientCtxKey string `json:"fsClientCtxKey,omitempty"`
	CollectionName string `json:"collectionName"`
	DocKey         string `json:"docKey,omitempty"`
}

// Load the master collection from firestore (assumed to be a small collection)
func (fs *FbMasterLoader) Load(ctx context.Context, param interface{}, cfg map[string]interface{}, errChan chan error) interface{} {
	if param == nil {
		errChan <- fmt.Errorf("FbMasterLoader.Load: no param specified ")
		return nil
	}
	ret, err := fs.getParams(ctx, param, cfg)
	if err != nil {
		errChan <- err
		return nil
	}
	return ret.load(ctx)
}

// DefaFsClient firestore base component registration name
const DefaFsClient = "firestore.Client"

func (fs *FbMasterLoader) getParams(ctx context.Context, param interface{}, cfg map[string]interface{}) (rl *fbRecLoader, err error) {
	rl = &fbRecLoader{}
	cfgRl := &CfgFbMasterLoader{}
	by, err := json.Marshal(param)
	if err != nil {
		return
	}
	err = json.Unmarshal(by, cfgRl)
	if err != nil {
		return
	}

	if len(cfgRl.FsClientCtxKey) == 0 {
		cfgRl.FsClientCtxKey = DefaFsClient
	}
	ok := false
	if rl.fsClient, ok = ctx.Value(cfgRl.FsClientCtxKey).(*firestore.Client); !ok {
		err = commands.NewFatalError(fmt.Sprintf("FbMasterLoader.getParams: firestore context key %v not instantiated", cfgRl.FsClientCtxKey))
		return
	}
	rl.collectionName = cfgRl.CollectionName
	rl.docKey = cfgRl.DocKey
	return
}

type fbRecLoader struct {
	fsClient       *firestore.Client
	collectionName string
	docKey         string // "id" by default
	errChan        chan error
}

// DefaultDocKey firestore doc key default is "id"
//const DefaultDocKey = "id"

func (fs *fbRecLoader) load(ctx context.Context) map[string]interface{} {
	if fs.docKey == "" {
		fs.docKey = DefaultDocKey
	}
	recs := make(map[string]interface{}, 0)
	// https://firebase.google.com/docs/firestore/quickstart
	snaps, err := fs.fsClient.Collection(fs.collectionName).Documents(ctx).GetAll()
	if err != nil {
		fs.errChan <- err
		return nil
	}
	for j := 0; j < len(snaps); j++ {
		dmap := snaps[j].Data()
		id, ok := (dmap[fs.docKey]).(string)
		if id == "" || !ok {
			id = util.RandRecID()
		}
		recs[id] = dmap

	} // for j
	return recs
}

//**** component methods *****

// Name to register in the component registry
func (fs *FbMasterLoader) Name() string {
	return "firestore.FbMasterLoader"
}

// Help return help string
func (fs *FbMasterLoader) Help() string {
	return `
	  
	  
	  For printing to a file, specify the file path.
	  "param": {

	  }  
	`
}

// ComponentType returns the component type
func (fs *FbMasterLoader) ComponentType() reflect.Type {
	return reflect.TypeOf(fs)
}

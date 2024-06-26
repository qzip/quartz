package comp

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/seq"
	"qz/util"
	"strings"

	driver "github.com/arangodb/go-driver"
	"github.com/google/uuid"
)

type DumpArango struct {
	Database        string `json:"database"`
	CollectionName  string `json:"collection"`
	DbHelperCtxName string `json:"dbHelper"`
	// Assumes map[string]interface{}
	DataOutCtxName string `json:"data,omitempty"`

	//private fields
	helper  seq.CtxHelper
	errChan chan error

	clientHelper *ArangoClient
	client       driver.Client
	database     driver.Database
	collection   driver.Collection
}

func (ad *DumpArango) SetCtxHelper(hlp seq.CtxHelper) {
	ad.helper = hlp
}

func (ad *DumpArango) SetChanErr(ce chan error) {
	ad.errChan = ce
}

// Process implements Pipeline (required by seq.RunSeq)
func (ad *DumpArango) Process(ctx context.Context) {
	ad.helper.SetExecStatus(seq.ExSrunning)
	if ad.DataOutCtxName == "" {
		ad.DataOutCtxName = DefaultDataCtxName
	}

	// create cleint
	clhlp := ad.helper.Value(ad.DbHelperCtxName)
	if clhlp == nil {
		ad.helper.SetExecStatus(seq.ExSerror)
		ad.errChan <- fmt.Errorf("ArangoInsert.Process:#4 %v context db helper not found", ad.DbHelperCtxName)
		return
	}
	var ok bool
	ad.clientHelper, ok = clhlp.(*ArangoClient)
	if !ok {
		ad.helper.SetExecStatus(seq.ExSerror)
		ad.errChan <- fmt.Errorf("ArangoInsert.Process:#5 %v context db helper not of type comp.ArangoClient", ad.DbHelperCtxName)
		return
	}
	util.DebugInfo(ctx, "ArangoInsert.Process: Creating NewClient")
	var err error
	ad.client, err = ad.clientHelper.NewClient()
	if err != nil {
		ad.helper.SetExecStatus(seq.ExSerror)
		ad.errChan <- err
		return
	}
	ad.database, err = ad.client.Database(ctx, ad.Database)
	if err != nil {
		ad.helper.SetExecStatus(seq.ExSerror)
		ad.errChan <- err
		return
	}

	/*FOR doc in @col return doc */
	qry := strings.Replace("FOR doc in @col return doc", "@col", ad.CollectionName, -1)
	crsr, err := ad.database.Query(ctx, qry, nil)
	//ad.collection, err = ad.database.Collection(ctx, ad.CollectionName)
	if err != nil {
		ad.helper.SetExecStatus(seq.ExSerror)
		ad.errChan <- err
		return
	}
	defer crsr.Close()

	collMap := make(map[string]map[string]*json.RawMessage)
	for crsr.HasMore() {
		bodyMap := make(map[string]*json.RawMessage)
		meta, err := crsr.ReadDocument(ctx, &bodyMap)
		if err != nil {
			ad.helper.SetExecStatus(seq.ExSerror)
			ad.errChan <- err
			return
		}
		docKey := uuid.NewString()
		if len(meta.Key) > 0 {
			docKey = meta.Key
		} else if len(meta.ID) > 0 {
			docKey = meta.ID.String()
		}
		collMap[docKey] = bodyMap

	}

	util.DebugInfo(ctx, "ArangoInsert.Process: get documents")
	ad.helper.SetKeyValue(ad.DataOutCtxName, collMap)
	ad.helper.SetExecStatus(seq.ExSok)

}

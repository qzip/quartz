package comp

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/seq"
	"qz/util"

	driver "github.com/arangodb/go-driver"
)

type DumpArango struct {
	Database        string `json:"database"`
	CollectionName  string `json:"collection"`
	DbHelperCtxName string `json:"dbHelper"`
	// Assumes map[string]interface{}
	DataInCtxName               string `json:"data,omitempty"`
	CreateDbIfNotExists         bool   `json:"createDbIfNotExists,omitempty"`
	CreateCollectionIfNotExists bool   `json:"createCollectionIfNotExists,omitempty"`
	UpdateDocumentIfExists      bool   `json:"updateDocumentIfExists,omitempty"`

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
	if ad.DataInCtxName == "" {
		ad.DataInCtxName = DefaultDataCtxName
	}
	// get data value from previous pipeline element
	data := ad.helper.Value(ad.DataInCtxName)
	if data == nil {
		ad.helper.SetExecStatus(seq.ExSerror)
		ad.errChan <- fmt.Errorf("ArangoInsert.Process:#1 %v context helper not found", ad.DataInCtxName)
		return
	}
	// transform
	col := make(map[string]map[string]interface{})
	if by, err := json.Marshal(data); err == nil {
		if err = json.Unmarshal(by, &col); err != nil {
			ad.helper.SetExecStatus(seq.ExSerror)
			ad.errChan <- fmt.Errorf("ArangoInsert.Process:#2 %v JSON unmarshall error", ad.DataInCtxName)
			return
		}
	} else {
		ad.helper.SetExecStatus(seq.ExSerror)
		ad.errChan <- fmt.Errorf("ArangoInsert.Process:#3 %v JSON marshall error", ad.DataInCtxName)
		return
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
	client, err := ad.clientHelper.NewClient()
	if err != nil {
		ad.helper.SetExecStatus(seq.ExSerror)
		ad.errChan <- err
		return
	}
	util.DebugInfo(ctx, "ArangoInsert.Process: Insert/Update documents")

	// col is the key,value collection
	if err = ad.insertDocuments(ctx, col); err != nil {
		ad.helper.SetExecStatus(seq.ExSerror)
		ad.errChan <- err
		return
	}

	ad.helper.SetExecStatus(seq.ExSok)

}

func (ad *DumpArango) getRecord() {
	data := make(map[string]interface{})
	if err := ref.Get(ctx, &data); err != nil {
		util.DebugInfo(ctx, err.Error())
		ad.errChan <- err
		ad.helper.SetExecStatus(seq.ExSerror)
		return
	}

}

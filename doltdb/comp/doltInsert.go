package comp

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/seq"
	"qz/util"

	driver "github.com/Doltdb/go-driver"
)

const (
	DefaultDataCtxName = "data"
	DoltKeyName        = "_key"
)

type ClientHelper interface {
	NewClient() (driver.Client, error)
}
type DoltInsert struct {
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

	clientHelper *DoltClient
	client       driver.Client
	database     driver.Database
	collection   driver.Collection
}

func (ai *DoltInsert) SetCtxHelper(hlp seq.CtxHelper) {
	ai.helper = hlp
}

func (ai *DoltInsert) SetChanErr(ce chan error) {
	ai.errChan = ce
}

// Process implements Pipeline (required by seq.RunSeq)
func (ai *DoltInsert) Process(ctx context.Context) {
	ai.helper.SetExecStatus(seq.ExSrunning)
	if ai.DataInCtxName == "" {
		ai.DataInCtxName = DefaultDataCtxName
	}
	// get data value from previous pipeline element
	data := ai.helper.Value(ai.DataInCtxName)
	if data == nil {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- fmt.Errorf("DoltInsert.Process:#1 %v context helper not found", ai.DataInCtxName)
		return
	}
	// transform
	col := make(map[string]map[string]interface{})
	if by, err := json.Marshal(data); err == nil {
		if err = json.Unmarshal(by, &col); err != nil {
			ai.helper.SetExecStatus(seq.ExSerror)
			ai.errChan <- fmt.Errorf("DoltInsert.Process:#2 %v JSON unmarshall error", ai.DataInCtxName)
			return
		}
	} else {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- fmt.Errorf("DoltInsert.Process:#3 %v JSON marshall error", ai.DataInCtxName)
		return
	}
	// create cleint
	clhlp := ai.helper.Value(ai.DbHelperCtxName)
	if clhlp == nil {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- fmt.Errorf("DoltInsert.Process:#4 %v context db helper not found", ai.DbHelperCtxName)
		return
	}
	var ok bool
	ai.clientHelper, ok = clhlp.(*DoltClient)
	if !ok {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- fmt.Errorf("DoltInsert.Process:#5 %v context db helper not of type comp.DoltClient", ai.DbHelperCtxName)
		return
	}
	util.DebugInfo(ctx, "DoltInsert.Process: Creating NewClient")
	client, err := ai.clientHelper.NewClient()
	if err != nil {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- err
		return
	}
	util.DebugInfo(ctx, "DoltInsert.Process: Creating Database")

	ai.client = client
	if err = ai.createDatabase(ctx); err != nil {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- err
		return
	}
	util.DebugInfo(ctx, "DoltInsert.Process: Creating Collection")

	if err = ai.createCollection(ctx); err != nil {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- err
		return
	}
	util.DebugInfo(ctx, "DoltInsert.Process: Insert/Update documents")

	// col is the key,value collection
	if err = ai.insertDocuments(ctx, col); err != nil {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- err
		return
	}

	ai.helper.SetExecStatus(seq.ExSok)

}

func (ai *DoltInsert) insertDocuments(ctx context.Context, col map[string]map[string]interface{}) error {
	for key, doc := range col {
		doc[DoltKeyName] = key
		if err := ai.insertDocument(ctx, doc); err != nil {
			return err
		}
	}
	return nil
}
func (ai *DoltInsert) insertDocument(ctx context.Context, doc map[string]interface{}) error {
	ok, err := ai.collection.DocumentExists(ctx, doc[DoltKeyName].(string))
	if err != nil {
		return err
	}
	if ok {
		if !ai.UpdateDocumentIfExists {
			return fmt.Errorf("DoltInsert.insertDocument: documrnt [%v] exists and UpdateDocumentIfExists is false", doc[DoltKeyName].(string))
		}
		// update document
		if _, err := ai.collection.UpdateDocument(ctx, doc[DoltKeyName].(string), doc); err != nil {
			return nil
		}
	} else {
		// document does not exist
		if _, err := ai.collection.CreateDocument(ctx, doc); err != nil {
			return nil
		}
	}
	return nil
}

func (ai *DoltInsert) createCollection(ctx context.Context) error {
	ok, err := ai.database.CollectionExists(ctx, ai.CollectionName)
	if err != nil {
		return err
	}
	if ok {
		if ai.collection, err = ai.database.Collection(ctx, ai.CollectionName); err != nil {
			return err
		}
	} else {
		if !ok && ai.CreateCollectionIfNotExists {
			if ai.collection, err = ai.database.CreateCollection(ctx, ai.CollectionName, nil); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("DoltInsert.createCollection: collection [%v] does not exists ", ai.CollectionName)
		}
	}

	return nil
}

func (ai *DoltInsert) createDatabase(ctx context.Context) error {
	// create DB
	ok, err := ai.client.DatabaseExists(ctx, ai.Database)
	if err != nil {
		return err
	}
	util.DebugInfo(ctx, fmt.Sprintf("DoltInsert.createDatabase#1: db [%v] exits = %v ", ai.Database, ok))
	if ok {
		if ai.database, err = ai.client.Database(ctx, ai.Database); err != nil {
			return err
		}
	} else {
		if !ok && ai.CreateDbIfNotExists {
			var dbOpt *driver.CreateDatabaseOptions = nil
			if ai.clientHelper.User != "" {
				util.DebugInfo(ctx, fmt.Sprintf("DoltInsert.createDatabase#2: user %v ", ai.clientHelper.User))
				dbOpt = &driver.CreateDatabaseOptions{
					Users: []driver.CreateDatabaseUserOptions{
						{
							UserName: ai.clientHelper.User,
							Password: ai.clientHelper.Pass,
						},
					},
				}
			}
			if ai.database, err = ai.client.CreateDatabase(ctx, ai.Database, dbOpt); err != nil {
				return err
			}
			util.DebugInfo(ctx, fmt.Sprintf("DoltInsert.createDatabase#3: db created %v ", ai.Database))

		} else {
			return fmt.Errorf("DoltInsert.createDatabase: database [%v] does not exists ", ai.Database)
		}
	}

	return nil
}

package comp

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/seq"
	"qz/util"

	driver "github.com/arangodb/go-driver"
)

const (
	DefaultDataCtxName = "data"
	ArangoKeyName      = "_key"
)

type ClientHelper interface {
	NewClient() (driver.Client, error)
}
type ArangoInsert struct {
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

func (ai *ArangoInsert) SetCtxHelper(hlp seq.CtxHelper) {
	ai.helper = hlp
}

func (ai *ArangoInsert) SetChanErr(ce chan error) {
	ai.errChan = ce
}

// Process implements Pipeline (required by seq.RunSeq)
func (ai *ArangoInsert) Process(ctx context.Context) {
	ai.helper.SetExecStatus(seq.ExSrunning)
	if ai.DataInCtxName == "" {
		ai.DataInCtxName = DefaultDataCtxName
	}
	// get data value from previous pipeline element
	data := ai.helper.Value(ai.DataInCtxName)
	if data == nil {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- fmt.Errorf("ArangoInsert.Process:#1 %v context helper not found", ai.DataInCtxName)
		return
	}
	// transform
	col := make(map[string]map[string]interface{})
	if by, err := json.Marshal(data); err == nil {
		if err = json.Unmarshal(by, &col); err != nil {
			ai.helper.SetExecStatus(seq.ExSerror)
			ai.errChan <- fmt.Errorf("ArangoInsert.Process:#2 %v JSON unmarshall error", ai.DataInCtxName)
			return
		}
	} else {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- fmt.Errorf("ArangoInsert.Process:#3 %v JSON marshall error", ai.DataInCtxName)
		return
	}
	// create cleint
	clhlp := ai.helper.Value(ai.DbHelperCtxName)
	if clhlp == nil {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- fmt.Errorf("ArangoInsert.Process:#4 %v context db helper not found", ai.DbHelperCtxName)
		return
	}
	var ok bool
	ai.clientHelper, ok = clhlp.(*ArangoClient)
	if !ok {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- fmt.Errorf("ArangoInsert.Process:#5 %v context db helper not of type comp.ArangoClient", ai.DbHelperCtxName)
		return
	}
	util.DebugInfo(ctx, "ArangoInsert.Process: Creating NewClient")
	client, err := ai.clientHelper.NewClient()
	if err != nil {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- err
		return
	}
	util.DebugInfo(ctx, "ArangoInsert.Process: Creating Database")

	ai.client = client
	if err = ai.createDatabase(ctx); err != nil {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- err
		return
	}
	util.DebugInfo(ctx, "ArangoInsert.Process: Creating Collection")

	if err = ai.createCollection(ctx); err != nil {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- err
		return
	}
	util.DebugInfo(ctx, "ArangoInsert.Process: Insert/Update documents")

	// col is the key,value collection
	if err = ai.insertDocuments(ctx, col); err != nil {
		ai.helper.SetExecStatus(seq.ExSerror)
		ai.errChan <- err
		return
	}

	ai.helper.SetExecStatus(seq.ExSok)

}

func (ai *ArangoInsert) insertDocuments(ctx context.Context, col map[string]map[string]interface{}) error {
	for key, doc := range col {
		doc[ArangoKeyName] = key
		if err := ai.insertDocument(ctx, doc); err != nil {
			return err
		}
	}
	return nil
}
func (ai *ArangoInsert) insertDocument(ctx context.Context, doc map[string]interface{}) error {
	ok, err := ai.collection.DocumentExists(ctx, doc[ArangoKeyName].(string))
	if err != nil {
		return err
	}
	if ok {
		if !ai.UpdateDocumentIfExists {
			return fmt.Errorf("ArangoInsert.insertDocument: documrnt [%v] exists and UpdateDocumentIfExists is false", doc[ArangoKeyName].(string))
		}
		// update document
		if _, err := ai.collection.UpdateDocument(ctx, doc[ArangoKeyName].(string), doc); err != nil {
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

func (ai *ArangoInsert) createCollection(ctx context.Context) error {
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
			return fmt.Errorf("ArangoInsert.createCollection: collection [%v] does not exists ", ai.CollectionName)
		}
	}

	return nil
}

func (ai *ArangoInsert) createDatabase(ctx context.Context) error {
	// create DB
	ok, err := ai.client.DatabaseExists(ctx, ai.Database)
	if err != nil {
		return err
	}
	util.DebugInfo(ctx, fmt.Sprintf("ArangoInsert.createDatabase#1: db [%v] exits = %v ", ai.Database, ok))
	if ok {
		if ai.database, err = ai.client.Database(ctx, ai.Database); err != nil {
			return err
		}
	} else {
		if !ok && ai.CreateDbIfNotExists {
			var dbOpt *driver.CreateDatabaseOptions = nil
			if ai.clientHelper.User != "" {
				util.DebugInfo(ctx, fmt.Sprintf("ArangoInsert.createDatabase#2: user %v ", ai.clientHelper.User))
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
			util.DebugInfo(ctx, fmt.Sprintf("ArangoInsert.createDatabase#3: db created %v ", ai.Database))

		} else {
			return fmt.Errorf("ArangoInsert.createDatabase: database [%v] does not exists ", ai.Database)
		}
	}

	return nil
}

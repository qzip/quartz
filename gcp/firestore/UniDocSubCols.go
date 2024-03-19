package firestore

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"qz/commands"
	"qz/defs"
	"qz/seq"
	"qz/util"
	"reflect"
	"strings"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

//UniDocSubCols implements seq.Runner interface.
// fetches a single record and passes it to the next stage
type UniDocSubCols struct {
}

//***** component seq.Runner interface methods *****

// Name implements seq.Runner interface method
func (bi *UniDocSubCols) Name() string {
	return "firestore.UniDocSubCols"
}

//Help implements seq.Runner interface method
func (bi *UniDocSubCols) Help() string {
	return `
	  # cmd.RunSeq: stage 
	  {
		  "component" : "firestore.UniDocSubCols",
		  "param" : {
			#compoent specific parameters
			}
	  }
	`
}

//ComponentType implements seq.Runner interface method
func (bi *UniDocSubCols) ComponentType() reflect.Type {
	return reflect.TypeOf(bi)
}

// Create implements seq.Runner interface method
func (bi *UniDocSubCols) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	fetcher := &fetchOneRec{
		helper: helper,
		errCh:  errCh,
	}
	//TODO: delegate the to firestore component
	if err := fetcher.open(ctx, param, cfg); err != nil {
		errCh <- commands.NewFatalError(fmt.Sprintf("UniDocSubCols.Create: %v", err.Error()))
		helper.SetExecStatus(seq.ExSerror)
		return nil

	}
	helper.SetExecStatus(seq.ExSinit)
	return fetcher
}

//DocSubColParam configuration parameters
type DocSubColParam struct {
	ProjectID string `json:"projectID,omitempty"`
	// CredPath is optional, or can be "default" or <path to cred file>
	CredPath       string   `json:"credentialsPath,omitempty"`
	CollectionName string   `json:"collectionName"`
	SubColArr      []string `json:"subCollectionNames,omitempty"`
	MasterRecs     []string `json:"masterRecs,omitempty"`
}

//fetchOneRec fetches a records
type fetchOneRec struct {
	fsClient *firestore.Client
	params   *DocSubColParam
	helper   seq.CtxHelper
	errCh    chan error
}

//Process implements Pipeline method
// assumes defs.DocRecID in ctx contains record ID
func (fdc *fetchOneRec) Process(ctx context.Context) {

	defer fdc.close()
	val := fmt.Sprintf("%v", ctx.Value(defs.DocRecID))

	util.DebugInfo(ctx, fmt.Sprintf("fetchOneRec.Process: ctx DocRecID [%v] = %v \n", defs.DocRecID, val))

	fdc.helper.SetExecStatus(seq.ExSrunning)
	doc, err := fdc.fsClient.Collection(fdc.params.CollectionName).Doc(val).Get(ctx)
	if err != nil {
		fdc.errCh <- commands.NewFatalError(fmt.Sprintf("UniDocSubCols.Create: %v", err.Error()))
		fdc.helper.SetExecStatus(seq.ExSerror)
		return
	}
	dsc := &defs.DocSubCols{
		SubCols: make(map[string]map[string]interface{}),
	}
	dsc.Doc = doc.Data()
	dsc.Doc[defs.DocPATH], dsc.Doc[defs.DocID] = doc.Ref.Path, doc.Ref.ID

	if fdc.params.SubColArr == nil {
		fdc.params.SubColArr = make([]string, 0)
	}

	for i := 0; i < len(fdc.params.SubColArr); i++ {
		//fmt.Println("inside for i", i)
		colref := doc.Ref.Collection(fdc.params.SubColArr[i]) // *CollectionRef
		dsc.SubCols[fdc.params.SubColArr[i]] = make(map[string]interface{})
		dsc.SubCols[fdc.params.SubColArr[i]][defs.DocPATH] = colref.Path
		dsc.SubCols[fdc.params.SubColArr[i]][defs.DocID] = colref.ID
		snaps, err := colref.Documents(ctx).GetAll()
		if err != nil {
			fdc.errCh <- err
			fdc.helper.SetExecStatus(seq.ExSerror)
			return
		}
		for j := 0; j < len(snaps); j++ {
			dmap := snaps[j].Data()
			dmap[defs.DocID] = snaps[j].Ref.ID
			dmap[defs.DocPATH] = snaps[j].Ref.Path
			// collection(<name>)map[<sub col ID>]map[string]interface{}
			dsc.SubCols[fdc.params.SubColArr[i]][snaps[j].Ref.ID] = dmap
		} // for j
	}

	// load mater recs

	for _, mstr := range fdc.params.MasterRecs {
		recLoader := &firestoreRecLoader{collectionName: mstr,
			fsClient: fdc.fsClient, errChan: fdc.errCh,
		}

		master := recLoader.load(ctx)
		if master == nil {
			util.DebugInfo(ctx, fmt.Sprintf("fetchOneRec.Process: load master %v failed\n", mstr))
			fdc.helper.SetExecStatus(seq.ExSerror)
			return
		}
		dsc.SubCols[mstr] = master
	}

	// Add the rec to context
	fdc.helper.SetKeyValue(defs.DocRec, dsc)
	//util.DebugInfo(ctx, fmt.Sprintf("fetchOneRec.Process: Setting ctx DocRec [%v] = %v \n", defs.DocRec, dsc))
	fdc.helper.SetExecStatus(seq.ExSok)
}

//Open the collection
func (fdc *fetchOneRec) open(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	err := fdc.getParams(ctx, param, cfg)
	if err != nil {
		return err
	}
	if fdc.params.CredPath == "" || strings.EqualFold(fdc.params.CredPath, "default") {
		if fdc.fsClient, err = firestore.NewClient(ctx, fdc.params.ProjectID); err != nil {
			return err
		}
	} else {
		opt := option.WithCredentialsFile(fdc.params.CredPath)
		if fdc.fsClient, err = firestore.NewClient(ctx, fdc.params.ProjectID, opt); err != nil {
			return err
		}
	}
	return nil
}
func (fdc *fetchOneRec) getParams(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	if param == nil {
		return fmt.Errorf("fetchOneRec.open: nil param")
	}
	fdc.params = &DocSubColParam{}
	// get project ID
	by, err := json.Marshal(param)
	if err != nil {
		return err
	}
	err = json.Unmarshal(by, fdc.params)
	if err != nil {
		return err
	}
	if len(fdc.params.CollectionName) == 0 {
		return fmt.Errorf("fetchOneRec.open: collection name param not specified")
	}

	if fdc.params.ProjectID == "" {
		pid, ok := cfg[defs.CfgParamFirebasePfix+"."+defs.CfgParamProjectID]
		if !ok {
			return fmt.Errorf("fetchOneRec.open: firebase projectID param not specified")
		}
		fdc.params.ProjectID, ok = pid.(string)
		if !ok {
			return fmt.Errorf("fetchOneRec.open: firebase projectID param not a string")
		}
	}
	// check if credential path present
	if fdc.params.CredPath == "" {
		pid, ok := cfg[defs.CfgParamFirebasePfix+"."+defs.CfgParamCredPath]
		if !ok {
			// assumed default
			fdc.params.CredPath = ""
		}
		if p, ok := pid.(string); ok {
			fdc.params.CredPath = p
		} else {
			return fmt.Errorf("fetchOneRec.open: firebase credential param not specified")
		}
	}
	return nil
}

func (fdc *fetchOneRec) close() {
	if fdc.fsClient != nil {
		fdc.fsClient.Close()
		fdc.fsClient = nil
	}

}

//*************** master loader *********
type firestoreRecLoader struct {
	fsClient       *firestore.Client
	collectionName string
	docKey         string // "id" by default
	errChan        chan error
}

// DefaultDocKey firebase doc key default is "id"
const DefaultDocKey = "id"

func (fs *firestoreRecLoader) load(ctx context.Context) map[string]interface{} {
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

func (fs *firestoreRecLoader) loadx(ctx context.Context) map[string]map[string]interface{} {
	if fs.docKey == "" {
		fs.docKey = DefaultDocKey
	}
	recs := make(map[string]map[string]interface{}, 0)
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

//*************** unit test components

//TestParam2Ctx unit test inserts param to help context
type TestParam2Ctx struct {
	helper seq.CtxHelper
}

//Name implements commands.ComponentFactory methods
func (t *TestParam2Ctx) Name() string {
	return "firestore.TestParam2Ctx"
}

//Help implements commands.ComponentFactory methods
func (t *TestParam2Ctx) Help() string {

	return `
	{
		"component" : "firestore.TestParam2Ctx",
		"param" : "record-id-of-Doc"
	}	   
	`
}

//ComponentType implements commands.ComponentFactory methods
func (t *TestParam2Ctx) ComponentType() reflect.Type {
	return reflect.TypeOf(t)
}

//Create implements seq.Runner interface
func (t *TestParam2Ctx) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errChan chan error) commands.Pipeline {
	t.helper = helper
	key, ok := param.(string)
	if key == "" || !ok {
		errChan <- commands.NewFatalError("TestParam2Ctx.Create: param is not a string or is empty")
		helper.SetExecStatus(seq.ExSerror)
		return nil
	}
	t.helper.SetKeyValue(defs.DocRecID, key)
	util.DebugInfo(ctx, fmt.Sprintf("fetchOneRec.Process: Setting ctx DocRecID [%v] = %v \n", defs.DocRecID, key))

	helper.SetExecStatus(seq.ExSinit)
	return t
}

//Process implements commands.Pipeline method
func (t *TestParam2Ctx) Process(ctx context.Context) {

	t.helper.SetExecStatus(seq.ExSok)
}

//********* unit test component Doc 2 File

//TestDoc2File unit test inserts param to help context
type TestDoc2File struct {
	helper  seq.CtxHelper
	errChan chan error
	w       io.Writer
}

//Name implements commands.ComponentFactory methods
func (t *TestDoc2File) Name() string {
	return "firestore.TestDoc2File"
}

//Help implements commands.ComponentFactory methods
func (t *TestDoc2File) Help() string {

	return `
	{
		"component" : "firestore.TestDoc2File",
		"param" : "output-file-name.json"
	}	   
	`
}

//ComponentType implements commands.ComponentFactory methods
func (t *TestDoc2File) ComponentType() reflect.Type {
	return reflect.TypeOf(t)
}

//Create implements seq.Runner interface
func (t *TestDoc2File) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	t.helper = helper
	if param == nil {
		t.w = os.Stdout
	} else {
		fl, ok := param.(string)
		if !ok {
			t.errChan <- commands.NewFatalError(fmt.Sprintf("TestDoc2File.Create: %v not a string, expecting filename", param))
			t.helper.SetExecStatus(seq.ExSerror)
			return nil
		} else if strings.EqualFold(fl, "stdio") {
			t.w = os.Stdout
		} else {
			var err error
			if t.w, err = os.OpenFile(fl, os.O_RDWR|os.O_CREATE, 0755); err != nil {
				t.errChan <- fmt.Errorf("TestDoc2File.Create: error [ %v ]opening %v, reverting to Stdout", err.Error(), fl)
				t.w = os.Stdout
			}
		}

	}

	helper.SetExecStatus(seq.ExSinit)
	return t
}

//Process implements commands.Pipeline method
func (t *TestDoc2File) Process(ctx context.Context) {
	t.helper.SetExecStatus(seq.ExSrunning)
	var dsc *defs.DocSubCols
	v := ctx.Value(defs.DocRec)
	if v == nil {
		err := commands.NewFatalError("TestDoc2File.Process: NOT FOUND: context value for key reflect.TypeOf(*defs.DocSubCols)")
		util.DebugInfo(ctx, err.Error())
		t.errChan <- err
		t.helper.SetExecStatus(seq.ExSerror)
		return
	}
	dsc, ok := v.(*defs.DocSubCols)
	if !ok {
		err := commands.NewFatalError("TestDoc2File.Process: NOT of type. context value for key reflect.TypeOf(*defs.DocSubCols)")
		util.DebugInfo(ctx, err.Error())
		t.errChan <- err
		t.helper.SetExecStatus(seq.ExSerror)
		return
	}
	if t.w != os.Stdout {
		f, ok := (t.w).(*os.File)
		if ok {
			defer f.Close()
		}

	}

	if err := FPrintUniDocSubCols(ctx, t.w, dsc); err != nil {
		t.helper.SetExecStatus(seq.ExSok)
		t.errChan <- commands.NewFatalError(err.Error())
	} else {
		t.helper.SetExecStatus(seq.ExSok)
	}

}

//******************** unit test component  File 2 Doc

//TestFile2Doc unit test inserts param to help context
type TestFile2Doc struct {
	helper  seq.CtxHelper
	errChan chan error
	recs    defs.DocSubCols
}

//Name implements commands.ComponentFactory methods
func (t *TestFile2Doc) Name() string {
	return "firestore.TestFile2Doc"
}

//Help implements commands.ComponentFactory methods
func (t *TestFile2Doc) Help() string {

	return `
	{
		"component" : "firestore.TestFile2Doc",
		"param" : "input-file-name.json"
	}	   
	`
}

//ComponentType implements commands.ComponentFactory methods
func (t *TestFile2Doc) ComponentType() reflect.Type {
	return reflect.TypeOf(t)
}

//Create implements seq.Runner interface
func (t *TestFile2Doc) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errChan chan error) commands.Pipeline {
	t.helper = helper
	util.DebugInfo(ctx, fmt.Sprintf("TestFile2Doc.Create: param [%v]", param))

	fl, ok := param.(string)
	if !ok {
		err := commands.NewFatalError(fmt.Sprintf("TestFile2Doc.Create: %v not a string, expecting filename", param))
		util.DebugInfo(ctx, err.Error())
		t.errChan <- err
		t.helper.SetExecStatus(seq.ExSerror)
		return nil
	}
	jsonb, err := ioutil.ReadFile(fl)
	if err != nil {
		err := commands.NewFatalError(fmt.Sprintf("TestFile2Doc.Create.read: %v", err.Error()))
		util.DebugInfo(ctx, err.Error())
		t.helper.SetExecStatus(seq.ExSerror)
		errChan <- err
		return nil
	}
	err = json.Unmarshal(jsonb, &t.recs)
	if err != nil {
		err := commands.NewFatalError(fmt.Sprintf("TestFile2Doc.Create.load.json: %v", err.Error()))
		util.DebugInfo(ctx, err.Error())
		t.helper.SetExecStatus(seq.ExSerror)
		errChan <- err
		return nil
	}
	// Add the rec to context
	t.helper.SetKeyValue(defs.DocRec, &t.recs)
	util.DebugInfo(ctx, fmt.Sprintf("TestFile2Doc.Create set key %v", defs.DocRec))
	helper.SetExecStatus(seq.ExSinit)
	return t
}

//Process loads the file
func (t *TestFile2Doc) Process(ctx context.Context) {
	t.helper.SetExecStatus(seq.ExSok)
}

//*******************

// FPrintUniDocSubCols prints in json format to io.Writer
// TODO: add a file handle as aparam for redirection to a file
func FPrintUniDocSubCols(ctx context.Context, w io.Writer, dat *defs.DocSubCols) error {

	p, err := json.MarshalIndent(dat, "\n", " ")
	if err != nil {
		util.DebugInfo(ctx, "firestore.FPrintDocSubCol: ret err")
		return err
	}
	fmt.Fprintln(w, string(p))
	util.DebugInfo(ctx, "firestore.FPrintDocSubCol: ret ok\n")
	return nil
}

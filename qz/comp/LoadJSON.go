/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package comp

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"qz/commands"
	"qz/seq"
	"reflect"
)

// LoadJSON implement seq.RunSeq Factory
type LoadJSON struct {
}

// Load the master collection from file (assumed to be a small collection)
func (fs *LoadJSON) Load(ctx context.Context, param interface{}, cfg map[string]interface{}, errChan chan error) interface{} {
	if param == nil {
		errChan <- fmt.Errorf("LoadJSON.Load: no param specified ")
		return nil
	}
	fl, ok := param.(string)
	if fl == "" || !ok {
		errChan <- commands.NewFatalError("LoadJSON.Load: param is not a string or is empty")
		return nil
	}
	jsonb, err := ioutil.ReadFile(fl)
	if err != nil {
		errChan <- commands.NewFatalError(fmt.Sprintf("LoadJSON.Load.read: %v", err.Error()))
		return nil
	}
	recs := make(map[string]interface{})
	err = json.Unmarshal(jsonb, &recs)
	if err != nil {
		errChan <- commands.NewFatalError(fmt.Sprintf("LoadJSON.Load.json: %v", err.Error()))
		return nil
	}
	return recs
}

//Fload loads a JSON file of type map[string]interface{}
type Fload struct {
	param   FloadParam
	errChan chan error
	helper  seq.CtxHelper
}

//FloadParam params for Fload
type FloadParam struct {
	InFileName     string `json:"inFileName"`
	DataOutCtxName string `json:"dataOutCtxName"`
}

func (fl *Fload) load() interface{} {

	jsonb, err := ioutil.ReadFile(fl.param.InFileName)
	if err != nil {
		fl.helper.SetExecStatus(seq.ExSerror)
		fl.errChan <- commands.NewFatalError(fmt.Sprintf("Fload.Load.read: %v", err.Error()))
		return nil
	}
	recs := make(map[string]interface{})
	err = json.Unmarshal(jsonb, &recs)
	if err != nil {
		fl.helper.SetExecStatus(seq.ExSerror)
		fl.errChan <- commands.NewFatalError(fmt.Sprintf("Fload.Load.json: %v", err.Error()))
		return nil
	}
	fl.helper.SetExecStatus(seq.ExSrunning)
	return recs
}

// Process implements seq.RunSeq method
func (fl *Fload) Process(ctx context.Context) {
	if v := fl.load(); v != nil {
		fl.helper.SetKeyValue(fl.param.DataOutCtxName, v)
		fl.helper.SetExecStatus(seq.ExSok)
	}

}

//Create implements seq.RunSeq method
func (fs *LoadJSON) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline {
	fl := &Fload{helper: helper, errChan: errCh}
	var err error
	fl.param, err = fs.getParams(param)
	if err != nil {
		fl.helper.SetExecStatus(seq.ExSerror)
		errCh <- err
		return nil
	}
	helper.SetExecStatus(seq.ExSinit)
	return fl
}

func (fs *LoadJSON) getParams(param interface{}) (fp FloadParam, err error) {
	if param == nil {
		err = fmt.Errorf("LoadJSON.getParams: nil param")
		return
	}
	by, err := json.Marshal(param)
	if err != nil {
		return
	}
	err = json.Unmarshal(by, &fp)

	return
}

//**** component methods *****

// Name to register in the component registry
func (fs *LoadJSON) Name() string {
	return "comp.LoadJSON"
}

// Help return help string
func (fs *LoadJSON) Help() string {
	return `
	  
	  
	  For printing to a file, specify the file path.
	  "param": {

	  }  
	`
}

// ComponentType returns the component type
func (fs *LoadJSON) ComponentType() reflect.Type {
	return reflect.TypeOf(fs)
}

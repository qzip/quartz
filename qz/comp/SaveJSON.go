package comp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"qz/commands"
	"qz/seq"
	"reflect"
	"strings"
)

// SaveJSON loads record, optionally transforms & saves record
type SaveJSON struct {
	helper  seq.CtxHelper
	errChan chan error
	param   *SaveJSONParam
	w       io.Writer
}

//SaveJSONParam parameters
type SaveJSONParam struct {
	DataInCtxName string `json:"dataInCtxName"`
	OutFileName   string `json:"outFileName"`
}

// Name implements component interface
func (sav *SaveJSON) Name() string {
	return "comp.SaveJSON"
}

// Help implements component interface
func (sav *SaveJSON) Help() string {
	return `
	{
		"component" :  "comp.SaveJSON",
		"param": {
		  "dataInCtxName": "data",
		  "outFileName": "data.json"
		}
	}`
}

// ComponentType implements component interface
func (sav *SaveJSON) ComponentType() reflect.Type {
	return reflect.TypeOf(sav)
}

//Create implements seq.Runner interface
func (sav *SaveJSON) Create(helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errChan chan error) commands.Pipeline {
	sav.helper = helper
	sav.errChan = errChan

	if err := sav.getParams(helper, param, cfg); err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errChan <- err
		return nil
	}
	helper.SetExecStatus(seq.ExSinit)
	return sav
}

func (sav *SaveJSON) getParams(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	if param == nil {
		return fmt.Errorf("SaveJSON.getParams: nil param")
	}
	sav.param = &SaveJSONParam{}
	by, err := json.Marshal(param)
	if err != nil {
		return err
	}
	err = json.Unmarshal(by, sav.param)
	if err != nil {
		return err
	}

	return nil
}

//Process implements commands.Pipeline method
func (sav *SaveJSON) Process(ctx context.Context) {
	sav.helper.SetExecStatus(seq.ExSrunning)
	// get data from context helper
	data := sav.helper.Value(sav.param.DataInCtxName)
	if data == nil {
		sav.helper.SetExecStatus(seq.ExSerror)
		sav.errChan <- fmt.Errorf("SaveJSON.Process: context data not present in %v", sav.param.DataInCtxName)
		return
	}
	p, er := json.MarshalIndent(data, "\n", " ")
	if er != nil {
		sav.errChan <- er
		return
	}
	sav.open()
	f, ok := (sav.w).(*os.File)
	if ok {
		defer f.Close()
	}

	if _, er := sav.w.Write(p); er != nil {
		sav.helper.SetExecStatus(seq.ExSerror)
		sav.errChan <- er
		return
	}
	sav.helper.SetExecStatus(seq.ExSok)
}

func (sav *SaveJSON) open() {
	if sav.param.OutFileName == "" {
		sav.w = os.Stdout
	} else if strings.EqualFold(sav.param.OutFileName, "stdio") {
		sav.w = os.Stdout
	} else {
		if w, err := os.OpenFile(sav.param.OutFileName, os.O_RDWR|os.O_CREATE, 0755); err != nil {
			sav.w = os.Stdout
		} else {
			sav.w = w
		}
	}

}

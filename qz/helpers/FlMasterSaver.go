/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"qz/util"
	"reflect"
	"strings"
)

// FlMasterSaver saves to file
type FlMasterSaver struct {
	w       io.Writer
	errChan chan error
}

//	**** component methods *****

// Name to register in the component registry
func (fs *FlMasterSaver) Name() string {
	return "helpers.FlMasterSaver"
}

// Save the record, implements runners.RecJSONSaver
func (fs *FlMasterSaver) Save(ctx context.Context, param interface{}, cfg map[string]interface{}, data interface{}, errChan chan error) {
	util.DebugInfo(ctx, "FlMasterSaver.Save: enter")
	fs.errChan = errChan
	if er := fs.open(ctx, param, cfg); er != nil {
		errChan <- er
		return
	}
	if fs.w != os.Stdout {
		f, ok := (fs.w).(*os.File)
		if ok {
			defer f.Close()
		}

	}
	p, er := json.MarshalIndent(data, "\n", " ")
	if er != nil {
		errChan <- er
		return
	}
	if i, er := fs.w.Write(p); er != nil {
		util.DebugInfo(ctx, fmt.Sprintf("Error: bytes written %v , %v", i, er.Error()))
		errChan <- er
	}
}

func (fs *FlMasterSaver) open(ctx context.Context, param interface{}, cfg map[string]interface{}) error {
	if param == nil {
		fs.w = os.Stdout
	} else {
		fl, ok := param.(string)
		if !ok {
			return fmt.Errorf("FlMasterSaver.Save.open: %v not a string, expecting filename", param)
		} else if strings.EqualFold(fl, "stdio") {
			fs.w = os.Stdout
		} else {
			var err error
			if fs.w, err = os.OpenFile(fl, os.O_RDWR|os.O_CREATE, 0755); err != nil {
				util.DebugInfo(ctx, fmt.Sprintf("FactFlDocSink.Pipeline: error [ %v ]opening %v, reverting to Stdout", err.Error(), fl))
				fs.w = os.Stdout
			}
		}
	}
	return nil
}

// Help return help string
func (fs *FlMasterSaver) Help() string {
	return `
	  
	  For printing to a file, specify the file path.
	  "param": "path/to/file/name.json
	`
}

// ComponentType returns the component type
func (fs *FlMasterSaver) ComponentType() reflect.Type {
	return reflect.TypeOf(fs)
}

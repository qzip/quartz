package comp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"qz/commands"
	"qz/event"
	"qz/seq"
	"qz/util"
	"reflect"
	"strings"
)

type DebugComp struct {
}

type DebugCompParam struct {
	OutFileName string `json:"outFileName,omitempty"`
}

type dbgPrint struct {
	helper seq.CtxHelper
	w      io.Writer
}

// Name implements component interface
func (dc *dbgPrint) Name() string {
	return "comp.DebugPrint"
}

// Help implements component interface
func (dc *dbgPrint) Help() string {
	return `
	{
		"component" :  "comp.DebugComp",
		"param": {
		  "outFileName": "log.txt"
		}
	}`
}

// ComponentType implements component interface
func (dc *dbgPrint) ComponentType() reflect.Type {
	return reflect.TypeOf(dc)
}

// implements DebugInfoHandler
func (dc *dbgPrint) Print(ctx context.Context, msg string) {
	if _, err := fmt.Fprintln(dc.w, msg); err != nil {
		log.Println(err)
		ex := &event.ExitEvent{Err: err}
		event.EvtBusFromContext(ctx).Publish(ctx, ex)
	}
}

// Process implements commands.Pipeline method
func (dc *dbgPrint) Process(ctx context.Context) {
	dc.helper.SetExecStatus(seq.ExSok)
	util.DebugInfo(ctx, "dbgPrint.Process: ok")
}

// Name implements component interface
func (dc *DebugComp) Name() string {
	return "comp.DebugComp"
}

// Help implements component interface
func (dc *DebugComp) Help() string {
	return `
	{
		"component" :  "comp.DebugComp",
		"param": {
		  "outFileName": "log.txt"
		}
	}`
}

// ComponentType implements component interface
func (dc *DebugComp) ComponentType() reflect.Type {
	return reflect.TypeOf(dc)
}

// Create implements seq.Runner interface
func (dc *DebugComp) Create(ctx context.Context, helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errChan chan error) commands.Pipeline {
	w, err := dc.getParams(param)
	if err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errChan <- err
		return nil
	}
	dp := &dbgPrint{w: w, helper: helper}
	helper.SetKeyValue(commands.CfgDebugKey, dp)
	helper.SetExecStatus(seq.ExSinit)
	util.DebugInfo(ctx, "DebugComp.Create: ok")
	return dp
}

func (dc *DebugComp) getParams(param interface{}) (io.Writer, error) {
	var w io.Writer
	if param == nil {
		return nil, fmt.Errorf("DebugComp.getParams: nil param")
	}
	dp := &DebugCompParam{}
	by, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(by, dp)
	if err != nil {
		return nil, err
	}
	if dp.OutFileName == "" {
		w = os.Stdout
	} else if strings.EqualFold(dp.OutFileName, "stdio") {
		w = os.Stdout
	} else {
		if w, err = os.OpenFile(dp.OutFileName, os.O_RDWR|os.O_CREATE, 0755); err != nil {
			w = os.Stdout
		}
	}
	return w, nil
}

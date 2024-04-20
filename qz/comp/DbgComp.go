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
	"reflect"
	"strings"
)

type DebugComp struct {
	w io.Writer
}

type DebugCompParam struct {
	OutFileName string `json:"outFileName"`
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
func (dc *DebugComp) Create(helper seq.CtxHelper, param interface{}, cfg map[string]interface{}, errChan chan error) commands.Pipeline {

	if err := dc.getParams(param); err != nil {
		helper.SetExecStatus(seq.ExSerror)
		errChan <- err
		return nil
	}
	helper.SetKeyValue(commands.CfgDebugKey, dc)
	helper.SetExecStatus(seq.ExSinit)
	return dc
}

// implements DebugInfoHandler
func (dc *DebugComp) Print(ctx context.Context, msg string) {
	if _, err := fmt.Fprintln(dc.w, msg); err != nil {
		log.Println(err)
		ex := &event.ExitEvent{Err: err}
		event.EvtBusFromContext(ctx).Publish(ctx, ex)
	}
}

func (dc *DebugComp) getParams(param interface{}) error {
	if param == nil {
		return fmt.Errorf("DebugComp.getParams: nil param")
	}
	dp := &DebugCompParam{}
	by, err := json.Marshal(param)
	if err != nil {
		return err
	}
	err = json.Unmarshal(by, dp)
	if err != nil {
		return err
	}
	if dp.OutFileName == "" {
		dc.w = os.Stdout
	} else if strings.EqualFold(dp.OutFileName, "stdio") {
		dc.w = os.Stdout
	} else {
		if w, err := os.OpenFile(dp.OutFileName, os.O_RDWR|os.O_CREATE, 0755); err != nil {
			dc.w = os.Stdout
		} else {
			dc.w = w
		}
	}
	return nil
}

// Process implements commands.Pipeline method
func (dc *DebugComp) Process(_ context.Context) {

}

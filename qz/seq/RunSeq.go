/*
Copyright (c) 2020, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package seq

import (
	"context"
	"fmt"
	"qz/commands"
	"qz/util"
	"reflect"
)

//ExecStatus Execution Status
type ExecStatus string

//Execution Status
const (
	ExSundef   ExecStatus = "undef"
	ExSinit    ExecStatus = "init"
	ExSrunning ExecStatus = "running"
	ExSerror   ExecStatus = "error"
	ExSok      ExecStatus = "ok"
)

//Configuration parameter constants
const (
	RunSeqParam     = "runners"
	ComponentSubKey = "component"
	ParamSubKey     = "param"
)

//CtxHelper parent context setter helper
type CtxHelper interface {
	context.Context
	SetExecStatus(status ExecStatus)
	GetExecStatus() ExecStatus
	StatusOk() bool
	SetKeyValue(key interface{}, val interface{})
}

//CtxHelperTypeName get the interface path
func CtxHelperTypeName() string {
	return reflect.TypeOf((*CtxHelper)(nil)).Elem().String()
}

//KeyName  context key name
func (xs ExecStatus) KeyName() string {
	return CtxHelperTypeName() + ".ExecStatus"
}

//ContextHelper extends commands.ContextHelper
type ContextHelper struct {
	*commands.ContextHelper
}

// SetExecStatus sets the execution status
func (ch *ContextHelper) SetExecStatus(status ExecStatus) {
	ch.KeyValMap()[status.KeyName()] = status

}

//GetExecStatus gets the execution status
func (ch *ContextHelper) GetExecStatus() (status ExecStatus) {
	status = ExSundef
	sts, ok := ch.KeyValMap()[status.KeyName()]
	if ok {
		if status, ok = sts.(ExecStatus); !ok {
			status = ExSerror
		}
	}
	return
}

//StatusOk true if exec status is set to ok
func (ch *ContextHelper) StatusOk() bool {
	return ch.GetExecStatus() == ExSok
}

//SetKeyValue sets the context
func (ch *ContextHelper) SetKeyValue(key interface{}, val interface{}) {
	ch.KeyValMap()[key] = val
}

//NewContextHelper creates an instance of Context helper
func NewContextHelper(ctx context.Context) *ContextHelper {
	ch := &ContextHelper{ContextHelper: commands.NewContextHelper(ctx)}
	ch.SetExecStatus(ExSinit)
	return ch
}

// Runner component factory to create a sequence pipeline runner
type Runner interface {
	commands.ComponentFactory
	Create(ctx context.Context, helper CtxHelper, param interface{}, cfg map[string]interface{}, errCh chan error) commands.Pipeline
}

//***** command interface methods *****

//RunSeq sequence of pipelines, that are executed if the previous
// stage is successful
type RunSeq struct {
	pipes   []commands.Pipeline
	errChan chan error
	helper  CtxHelper
}

// Name implements command interface method
func (rs *RunSeq) Name() string {
	return "seq.RunSeq"
}

//Help implements command interface method
func (rs *RunSeq) Help() string {
	return `
	  # cmd.RunSeq: executes the runners in sequence,
	  # the next stage is executed only if previous
	  # stage is successful
	  {
		  "command" : "seq.RunSeq",
		  "stages": [
			  {
			  	"component": "<component name>",
			  	"param" : {
				  #compoent specific parameters
			  	}
			}, #...
		  ]
	  }
	`
}

//Exec executes the stages, if the prev. stage is successful
func (rs *RunSeq) Exec(ctx context.Context, cfg map[string]interface{}, errChan chan error) {
	rs.helper = NewContextHelper(ctx)
	rs.errChan = errChan
	if err := rs.setupRunners(rs.helper, cfg); err != nil {
		util.DebugInfo(ctx, err.Error())
		errChan <- err
		return
	}
	util.DebugInfo(ctx, fmt.Sprintf("RunSeq.Exec: total %v sequence of pipes\n", len(rs.pipes)))
	for i, v := range rs.pipes {
		v.Process(rs.helper) // will set context value with ExecStatus
		if !rs.helper.StatusOk() {
			erx := fmt.Errorf("RunSeq.Exec: ERROR breaking off at %v of %v sequence", i+1, len(rs.pipes))
			util.DebugInfo(rs.helper, erx.Error())
			errChan <- erx
			break
		}
	}
	return
}

func (rs *RunSeq) setupRunners(ctx context.Context, cfg map[string]interface{}) error {
	rs.pipes = make([]commands.Pipeline, 0)
	rnv, ok := cfg[RunSeqParam]
	if !ok {
		return commands.NewFatalError(fmt.Sprintf("RunSeq.setupRunners:#0 %v  component name not found", RunSeqParam))
	}
	runarr, ok := rnv.([]interface{})
	if !ok {
		return commands.NewFatalError(fmt.Sprintf("RunSeq.setupRunners:#1 %v  component not array type", RunSeqParam))
	}
	for i, r := range runarr {
		me, ok := r.(map[string]interface{})
		if !ok {
			return commands.NewFatalError(fmt.Sprintf("RunSeq.setupRunners:#2 %v  component array index %v not an object", RunSeqParam, i))
		}
		ci, ok := me[ComponentSubKey]
		if !ok {
			return commands.NewFatalError(fmt.Sprintf("RunSeq.setupRunners:#3 %v  component array index %v , element %v not found", RunSeqParam, i, ComponentSubKey))
		}
		cname := fmt.Sprintf("%v", ci)
		comp, ok := commands.LookUpComponentFactory(cname)
		if !ok {
			return commands.NewFatalError(fmt.Sprintf("RunSeq.setupRunners:#4 %v  component array index %v , component %v not registered", RunSeqParam, i, cname))
		}
		runner, ok := comp.(Runner)
		if !ok {
			return commands.NewFatalError(fmt.Sprintf("RunSeq.setupRunners:#5 %v  component array index %v , component %v does not implement Runner interface", RunSeqParam, i, cname))
		}
		param, ok := me[ParamSubKey]
		if !ok {
			param = nil
		}
		pipe := runner.Create(ctx, rs.helper, param, cfg, rs.errChan)
		if pipe == nil {
			erx := fmt.Errorf("RunSeq.setupRunners:#6 %v runners component array index %v , component %v  Runner failed to create a Pipeline element", RunSeqParam, i, cname)
			rs.errChan <- erx
			return commands.NewFatalError(erx.Error())
		}
		rs.pipes = append(rs.pipes, pipe)
	}
	return nil
}

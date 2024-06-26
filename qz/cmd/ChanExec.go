/*
Copyright (c) 2024, Ashish Banerjee

Author: Ashish Banerjee, <tech@AshishBanerjee.com>

SPDX-License-Identifier: Apache-2.0

*/

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"qz/commands"
	"qz/util"
	"reflect"
	"sync"
)

// ChanExec implements commands.Commander
type ChanExec struct {
	source       commands.Source
	sink         commands.Sink
	transformers []commands.Transformer
}

const (
	sourceKey       = "source"
	sinkKey         = "sink"
	transformersKey = "transformers"
)

// Exec implements qz/commands.Commander Exec method. Creates channel driven pipeline
func (e *ChanExec) Exec(ctx context.Context, cfg map[string]interface{}, errCh chan error) {
	var ok bool
	util.DebugInfo(ctx, "cmd.ExecChan.Exec: entered")

	if hlp, err := e.getHelper(ctx, cfg, sourceKey); err != nil {
		util.DebugInfo(ctx, "cmd.ExecChan.Exec: "+err.Error())
		errCh <- commands.NewFatalError(err.Error())
		return
	} else if e.source, ok = hlp.(commands.Source); !ok {
		emsg := "cmd.ChanExec.Exec: helper not of type commands.Source"
		util.DebugInfo(ctx, "cmd.ExecChan: "+emsg)
		errCh <- commands.NewFatalError(emsg)
		return
	}

	if hlp, err := e.getHelper(ctx, cfg, sinkKey); err != nil {
		util.DebugInfo(ctx, "cmd.ExecChan.Exec: "+err.Error())
		errCh <- commands.NewFatalError(err.Error())
		return
	} else if e.sink, ok = hlp.(commands.Sink); !ok {
		emsg := "cmd.ChanExec.Exec: helper not of type commands.Sink"
		util.DebugInfo(ctx, emsg)
		errCh <- commands.NewFatalError(emsg)
		return
	}

	// deserialize the transformer array, transformers are optional
	tarrKey, ok := cfg[transformersKey]
	var tarr []string
	if ok {
		if tarrKey != nil {
			if by, err := json.Marshal(tarrKey); err != nil {
				util.DebugInfo(ctx, err.Error())
				errCh <- err
				return
			} else {
				if err = json.Unmarshal(by, &tarr); err != nil {
					util.DebugInfo(ctx, err.Error())
					errCh <- err
					return
				}
			}
		}
	} else if !ok {
		erf := commands.NewFatalError("cmd.ChanExec.Exec: transformer key [" + transformersKey + "] string array expected")
		util.DebugInfo(ctx, erf.Error())
		errCh <- erf
		return
	}

	e.transformers = make([]commands.Transformer, len(tarr))
	for i, t := range tarr {

		if handler := ctx.Value(t); handler == nil {
			erf := commands.NewFatalError("cmd.ChanExec.Exec: helper [" + t + "  not found")
			util.DebugInfo(ctx, erf.Error())
			errCh <- erf
			return
		} else if e.transformers[i], ok = handler.(commands.Transformer); !ok {
			erf := commands.NewFatalError("cmd.ChanExec.Exec: helper not of type commands.Transformer")
			util.DebugInfo(ctx, erf.Error())
			errCh <- erf
			return
		}

	}

	util.DebugInfo(ctx, "cmd.ExecChan.Exec: before spinoff")
	e.spinOff(ctx, errCh)
	util.DebugInfo(ctx, "cmd.ExecChan.Exec: after spinoff")
}

func (e *ChanExec) spinOff(ctx context.Context, errCh chan error) {
	cctx, cancel := context.WithCancel(ctx)
	schan := e.source.Chan(cctx, errCh, cancel)

	for _, t := range e.transformers {
		schan = t.Chan(cctx, schan, errCh, cancel)
	}
	e.sink.Chan(cctx, schan, errCh, cancel)

	var wg sync.WaitGroup
	f := func(p commands.Pipeline) {
		p.Process(cctx)
		wg.Done()
	}
	wg.Add(2)
	go f(e.source)
	go f(e.sink)
	for _, t := range e.transformers {
		wg.Add(1)
		go f(t)
	}
	util.DebugInfo(ctx, "cmd.ExecChan.spinoff: before wait")
	wg.Wait()
	util.DebugInfo(ctx, "cmd.ExecChan.spinoff: after wait")

}

func (e *ChanExec) getHelper(ctx context.Context, cfg map[string]interface{}, key string) (helper interface{}, err error) {
	handlerFact, ok := cfg[key]
	if !ok {
		return nil, fmt.Errorf("cmd.ChanExec.getHelper:  helper param %v not specified", key)
	}
	helper, ok = handlerFact.(string)
	if !ok {
		return nil, fmt.Errorf("cmd.ChanExec.getHelper:  param %v not a string", key)
	}
	helper = ctx.Value(helper)
	if helper == nil {
		return nil, fmt.Errorf("cmd.ChanExec.getHelper:  param %v helper not found", key)
	}
	return
}

func (e *ChanExec) Name() string {
	return reflect.TypeOf(*e).String()
}

func (e *ChanExec) Help() string {
	return `
	  {
		"desription" : "
		 Creates a channel pipeline. 
		",
		"command": "cmd.ChanExec",
		"debug": true,
		"source": "<commands.Source helper>",
		"sink":  "<commands.Sink helper>"
		"transformers": [
			"<commands.Transformer helper>"
		],
		"helpers": [
			{
				"component": "<commands.Source helper Factory>"
				"param": {}
			},
			{
				"component": "<commands.Sink helper Factory>"
				"param": {}
			},
			{
				"component": "<commands.Transformer helper Factory>"
				"param": {}
			}
		]
	  }
	
	`
}

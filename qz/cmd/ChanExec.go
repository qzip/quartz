/*
Copyright (c) 2024, Ashish Banerjee

Author: Ashish Banerjee, <tech@AshishBanerjee.com>

SPDX-License-Identifier: Apache-2.0

*/

package cmd

import (
	"context"
	"fmt"
	"qz/commands"
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

func (e *ChanExec) Exec(ctx context.Context, cfg map[string]interface{}, errCh chan error) {
	var ok bool

	if hlp, err := e.getHelper(ctx, cfg, sourceKey); err != nil {
		errCh <- commands.NewFatalError(err.Error())
		return
	} else if e.source, ok = hlp.(commands.Source); !ok {
		errCh <- commands.NewFatalError("cmd.ChanExec.Exec: helper not of type commands.Source")
		return
	}

	if hlp, err := e.getHelper(ctx, cfg, sinkKey); err != nil {
		errCh <- commands.NewFatalError(err.Error())
		return
	} else if e.sink, ok = hlp.(commands.Sink); !ok {
		errCh <- commands.NewFatalError("cmd.ChanExec.Exec: helper not of type commands.Sink")
		return
	}

	// deserialize the transformer array, transformers are optional
	tarrKey, ok := cfg[transformersKey]
	if ok {
		tarr, ok := tarrKey.([]string)
		if !ok {
			errCh <- commands.NewFatalError("cmd.ChanExec.Exec: transformer key [" + transformersKey + "] string array expected")
			return
		}
		e.transformers = make([]commands.Transformer, len(tarr))
		for i, t := range tarr {

			if handler, err := e.getHelper(ctx, cfg, t); err != nil {
				errCh <- commands.NewFatalError("cmd.ChanExec.Exec: helper not of type commands.Transformer array")
				return
			} else if e.transformers[i], ok = handler.(commands.Transformer); !ok {
				errCh <- commands.NewFatalError("cmd.ChanExec.Exec: helper not of type commands.Transformer array")
				return
			}

		}
	}
	e.spinOff(ctx, errCh)
}

func (e *ChanExec) spinOff(ctx context.Context, errCh chan error) {

	schan, err := e.source.Chan(ctx)
	if err != nil {
		errCh <- commands.NewFatalError(err.Error())
		return
	}

	for _, t := range e.transformers {
		schan, err = t.Chan(ctx, schan)
		if err != nil {
			errCh <- commands.NewFatalError(err.Error())
			return
		}
	}
	if err = e.sink.Chan(ctx, schan); err != nil {
		errCh <- commands.NewFatalError(err.Error())
		return
	}

	var wg sync.WaitGroup
	f := func(p commands.Pipeline) {
		p.Process(ctx)
		wg.Done()
	}
	wg.Add(2)
	go f(e.source)
	go f(e.sink)
	for _, t := range e.transformers {
		wg.Add(1)
		go f(t)
	}
	wg.Wait()

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

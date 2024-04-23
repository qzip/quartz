/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

/*
Package commands provides a configurable framework for commands
while avoiding multiple main packages.
*/
package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"qz/util"
	"reflect"
	"strings"
	"time"

	//"sync"
	"syscall"
)

// FatalError signals to the error handler to close the channel
type FatalError struct {
	Fatal string
}

func (fe FatalError) Error() string {
	return fe.Fatal
}

// NewFatalError on fatal error the handler will close the chan
func NewFatalError(fe string) (fatal error) {
	if !strings.HasSuffix(fe, "\n") {
		fe = fe + "\n"
	}
	fatal = FatalError{Fatal: fe}

	return
}

// Commander interface that all commands must implement
type Commander interface {
	Name() string
	Help() string
	Exec(ctx context.Context, cfg map[string]interface{}, errCh chan error)
}

var registry = make(map[string]Commander)

// RegisterCommand commands must register to be visible to the commander
func RegisterCommand(cmd Commander) {
	registry[cmd.Name()] = cmd
}

// LookUpCommand gets command from registry
func LookUpCommand(cmd string) Commander {
	return registry[cmd]
}
func CommandExists(cmd string) (exists bool) {
	_, exists = registry[cmd]
	return
}

// IterateCommands for each registered command call back
func IterateCommands(callBack func(Commander)) {
	for _, v := range registry {
		callBack(v)
	}
}

// ComponentFactory helpers for commands for dynamically switching components
// Componet factory can define it's own methods, by extending
// the ComponentFactory
type ComponentFactory interface {
	Name() string // registry
	Help() string
	ComponentType() reflect.Type
	//	Create(ctx context.Context, cfg map[string]interface{}, errChan chan error) (component interface{})
}

var regComponentFactory = make(map[string]ComponentFactory)

// RegisterComponentFactory register a component factory
func RegisterComponentFactory(cfact ComponentFactory) {
	regComponentFactory[cfact.Name()] = cfact
}

// LookUpComponentFactory gets command from registry
func LookUpComponentFactory(comp string) (cf ComponentFactory, ok bool) {
	cf, ok = regComponentFactory[comp]
	return
}

// IterateComponentFactory for each component call back
func IterateComponentFactory(callBack func(ComponentFactory)) {
	for _, v := range regComponentFactory {
		callBack(v)
	}
}

// CommandRunner command wrapper & CTL-C os Signal handler
type CommandRunner struct {
	Cmd Commander
}

func (run *CommandRunner) setdbg(cfg map[string]interface{}) {
	dbg, ok := (cfg[CfgDebugKey]).(bool)
	if ok {
		util.Debug = dbg
	} else {
		util.Debug = true
	}
	log.Println("debug=", util.Debug)
}

// Exec executes the wrapped command & also handles CTL-C os signal
func (run *CommandRunner) Exec(pctx context.Context, cfg map[string]interface{}, errCh chan error) {
	run.setdbg(cfg)
	startTm := time.Now()
	util.DebugInfo(pctx, "CommandRunner.Exec: Exec start")
	ctx, err := BuildCtxHandlers(pctx, cfg)
	if err != nil {
		errCh <- NewFatalError("CommandRunner.Exec(handlers): " + err.Error())
		return
	}
	util.DebugInfo(pctx, "CommandRunner.Exec: BuildCtxHandlers done")
	if err = InstallGlobalListners(ctx, cfg); err != nil {
		errCh <- NewFatalError("CommandRunner.Exec(listeners): " + err.Error())
		return
	}
	util.DebugInfo(pctx, "CommandRunner.Exec: InstallGlobalListners done")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-ctx.Done():
			return
		case sig := <-sigs:
			errCh <- NewFatalError(sig.String())
		}
	}() // do not wait on this
	run.Cmd.Exec(ctx, cfg, errCh)
	elapsed := time.Since(startTm)
	util.DebugInfo(pctx, fmt.Sprintf("CommandRunner.Exec: ret, execution duration %v", elapsed))
}

// Name returns the warpped command name
func (run *CommandRunner) Name() string { return run.Cmd.Name() }

// Help returns the wrapped command help
func (run *CommandRunner) Help() string { return run.Cmd.Help() }

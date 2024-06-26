package starlark

/*

"go.starlark.net/internal/compile"

	"go.starlark.net/lib/math"
	"go.starlark.net/lib/time"

	"go.starlark.net/resolve"
	"go.starlark.net/starlark"
	"golang.org/x/term"
	"qz/util"

*/
import (
	"context"
	"fmt"
	"os"
	"qz/commands"

	"qz/util"

	"go.starlark.net/repl"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
	"golang.org/x/term"
)

func init() {
	commands.RegisterComponentFactory(&JsonModFactory{})
	commands.RegisterComponentFactory(&TimeModFactory{})
	commands.RegisterComponentFactory(&MathModFactory{})
}

type VmStarlark struct {
	thread  *starlark.Thread
	globals starlark.StringDict
}

// Name implements command interface method
func (vs *VmStarlark) Name() string {
	return "vm_star"
}

// Help implements command interface method
func (vs *VmStarlark) Help() string {
	return `
	  version: 05jun23 7pm
      Starlark VM, file extension .star 
	  see: https://github.com/google/starlark-go
	`
}

// Exec executes the stages, if the prev. stage is successful
func (vs *VmStarlark) Exec(ctx context.Context, cfg map[string]interface{}, errChan chan error) {
	util.Debug = true

	vs.thread = &starlark.Thread{Load: repl.MakeLoad()}
	vs.globals = make(starlark.StringDict)
	// set context
	vs.thread.SetLocal("context", ctx)
	vs.setEnv()

	// execute
	var err error
	replx := false
	src, ok := cfg[commands.CmdVmSource]
	if !ok {
		replx = true
	}
	cmdf := cfg[commands.CmdFileName]
	filename := ""
	if cmdf != nil {
		filename = cmdf.(string)
	}

	if replx {
		filename = "cmdline"
		stdinIsTerminal := term.IsTerminal(int(os.Stdin.Fd()))
		if stdinIsTerminal {
			fmt.Println("Welcome to Starlark (go.starlark.net)")
		}
		vs.thread.Name = "REPL"
		repl.REPL(vs.thread, vs.globals)
		if stdinIsTerminal {
			fmt.Println()
		}
	} else {
		vs.globals, err = starlark.ExecFile(vs.thread, filename, src, nil)
		if err != nil {
			util.DebugInfo(ctx, err.Error())
			errChan <- err
		}
	}

}

type StarlarkModule interface {
	commands.ComponentFactory
	CreateModule() *starlarkstruct.Module
}

func (vs *VmStarlark) setEnv() {

	iter := func(cfact commands.ComponentFactory) {
		mod, ok := cfact.(StarlarkModule)
		if ok {
			starlark.Universe[cfact.Name()] = mod.CreateModule()
		}
	}
	commands.IterateComponentFactory(iter)

}

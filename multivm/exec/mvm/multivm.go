package main

import (
	"multivm/starlark"
	"os"
	"qz/commands"
	"qz/event"
	"qz/run"
)

// Register the commands & components.
// this is done explicitly here to provide control over what commands & components
// are included with a command runner. No Magic :)
func init() {
	//	commands.RegisterCommand(&cmdfact.NormDocSubColsCmd{})
	//	commands.RegisterCommand(&seq.RunSeq{})
	commands.RegisterCommand(&starlark.VmStarlark{})
	event.GlobalEventBus.Subscribe(&event.ExitEventListner{})

}

func main() {

	if len(os.Args) == 1 {
		run.PrintDefaultConfig()
		return
	}
	run.Runner(os.Args[1])
}

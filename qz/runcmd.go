package main

import (
	"os"
	"qz/cmd"
	"qz/commands"
	"qz/event"
	"qz/run"
	"qz/seq"
)

// Register the commands & components.
// this is done explicitly here to provide control over what commands & components
// are included with a command runner. No Magic :)
func init() {
	//	commands.RegisterCommand(&cmdfact.NormDocSubColsCmd{})
	//	commands.RegisterCommand(&seq.RunSeq{})
	commands.RegisterCommand(&seq.RunSeq{})
	commands.RegisterCommand(&cmd.SrvHTTP{})

	event.GlobalEventBus.Subscribe(&event.ExitEventListner{})

}

func main() {

	if len(os.Args) == 1 {
		run.PrintDefaultConfig()
		return
	}
	run.Runner(os.Args[1])
}

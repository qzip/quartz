/*
Copyright (c) 2022, Ashish Banerjee <tech@AshishBanerjee.com>

All Rights Reserved.

Apache Licese 2.0

*/
package main

import (
	"os"
	"qz/run"
)

// Register the commands & components.
// this is done explicitly here to provide control over what commands & components
// are included with a command runner. No Magic :)
func init() {
	//	commands.RegisterCommand(&cmdfact.NormDocSubColsCmd{})

}

func main() {

	if len(os.Args) == 1 {
		run.PrintDefaultConfig()
		return
	}
	run.Runner(os.Args[1])
}

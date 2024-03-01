// JS runner

package main

import (
	"context"
	"fmt"
	"multivm/js"
	"os"
	"qz/commands"
	"qz/event"
)

func init() {
	event.GlobalEventBus.Subscribe(&event.ExitEventListner{})
	commands.RegisterComponentFactory(&js.EventInfra{})
}
func main() {

	if len(os.Args) == 1 {
		fmt.Println(os.Args[0], "<js file name>")
		return
	}
	ctx := context.Background()
	js.JsExec(ctx, os.Args[1])
}

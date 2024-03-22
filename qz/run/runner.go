package run

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"qz/commands"
	"qz/event"
	"qz/util"
	"strings"
	"time"
)

func Runner(cfgFile string) {

	cfg, err := readCfg(os.Args[1])
	if err != nil {
		log.Fatal(fmt.Errorf("main.readCfg: %s", err.Error()))
		return
	}
	cmd := fmt.Sprintf("%v", cfg[commands.CmdKey])
	if cmd == "<nil>" {
		log.Fatal(fmt.Errorf("[%s] not found in config %s", commands.CmdKey, os.Args[1]))
		return
	}
	util.SetDebugFlag(cfg)
	commander := commands.LookUpCommand(cmd)
	if commander == nil {
		log.Fatal(fmt.Errorf("[%s] not found a registered command ", commands.CmdKey))
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	run := &commands.CommandRunner{Cmd: commander}
	errChan := make(chan error, 1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				util.DebugInfo(ctx, "main: ctx done")
				return // for
			case err := <-errChan:
				util.DebugInfo(ctx, err.Error())
				if _, ok := err.(commands.FatalError); ok {
					cancel()
					fmt.Println("main: after calling cancel, sleep 2 sec")
					time.Sleep(2 * time.Second)
					event.GlobalEventBus.Publish(ctx, &event.ExitEvent{Err: err})
					return
				}
			}
		}

	}()

	run.Exec(ctx, cfg, errChan)
	errChan <- commands.NewFatalError("OK: Exit runcmd")
	util.DebugInfo(ctx, "qz: main ret")
}

func readCfg(fname string) (cfg map[string]interface{}, err error) {
	cfg = make(map[string]interface{}, 0)
	jsonb, erx := os.ReadFile(fname)
	if erx != nil {
		err = erx
		return
	}

	// check if extension

	// extract the file extension
	ext := getExtension(fname)
	cmdName := commands.CmdVmPrefix + strings.ToLower(ext)

	if commands.CommandExists(cmdName) {
		cfg[commands.CmdKey] = cmdName
		cfg[commands.CmdFileName] = fname
		cfg[commands.CmdVmSource] = jsonb
	} else {
		err = json.Unmarshal(jsonb, &cfg)
	}

	return
}

func getExtension(fname string) (ext string) {
	f := func(c rune) bool {
		return c == '.'
	}
	r := strings.FieldsFunc(fname, f)
	ext = ""
	if len(r) > 1 {
		ext = r[len(r)-1]
	}
	return
}

// PrintDefaultConfig prints the default configs for commands & components
func PrintDefaultConfig() {
	fmt.Println("{")
	fmt.Println("\"commands\": [")
	commands.IterateCommands(func(cf commands.Commander) {
		fmt.Printf("{\n\"name\": \"%v\", \n", cf.Name())
		fmt.Printf("\"description\": \"%v\" \n}\n", cf.Help())
	})
	fmt.Println("],")

	fmt.Println("\"components\": [")
	commands.IterateComponentFactory(func(cf commands.ComponentFactory) {
		fmt.Printf("{\n\"name\": \"%v\", \n", cf.Name())
		fmt.Printf("\"description\": \"%v\", \n", cf.Help())
		fmt.Printf("\"type\": \"%v\" \n}\n,", cf.ComponentType())
	})
	fmt.Println("]")

	fmt.Println("}")
}

package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"qz/commands"
	"qz/dsl"
	"qz/util"
	"reflect"
	"strings"
)

//DslFile2Json check syntax of a DSL File
type DslFile2Json struct {
}

//Help implements Commander interface method
func (bi *DslFile2Json) Help() string {
	return `
	  # test the DSL parser
	  {
		  "command" : "comp.DslFile2Json",
		  "dsl" : "<dsl file>",
		  "json": "<optional json output>",
	  }
	`
}

// Name implements Commander interface method
func (bi *DslFile2Json) Name() string {
	return "cmd.DslFile2Json"
}

//Exec parses the DSL file & converts to JSON
func (bi *DslFile2Json) Exec(ctx context.Context, cfg map[string]interface{}, errChan chan error) {
	util.DebugInfo(ctx, "DslFile2Json.Exec: enter")
	dslFname, ok := cfg["dsl"]
	if !ok {
		errChan <- commands.NewFatalError("param dsl not specified\n")
		return
	}
	fl := fmt.Sprintf("%v", dslFname)
	util.DebugInfo(ctx, fmt.Sprintf("DslFile2Json.Exec: dsl file : %v\n", fl))
	fbuf, err := ioutil.ReadFile(fl)
	if err != nil {
		util.DebugInfo(ctx, err.Error())
		errChan <- commands.NewFatalError(err.Error() + "\n")
		return
	}
	util.DebugInfo(ctx, fmt.Sprintf("DslFile2Json.Exec: dsl file : %v opened\n", fl))
	prs := dsl.NewParser(bytes.NewReader(fbuf))
	blk, err := prs.Parse(ctx)
	if err != nil {
		errChan <- commands.NewFatalError(err.Error())
		return
	}
	util.DebugInfo(ctx, "DslFile2Json.Exec: blocks parsed")
	jsonFname, ok := cfg["json"]
	if !ok {
		jsonFname = "stdio"
	}
	fl = fmt.Sprintf("%v", jsonFname)
	var jw io.Writer
	if strings.EqualFold(fl, "stdio") {
		jw = os.Stdout
	} else {
		var err error
		if jw, err = os.OpenFile(fl, os.O_RDWR|os.O_CREATE, 0755); err != nil {
			errChan <- fmt.Errorf("DslFile2Json.Exec: error [ %v ]opening %v, reverting to Stdout", err.Error(), fl)
			jw = os.Stdout
		} else {
			defer ((jw).(*os.File)).Close()
		}
	}
	if err = FPrintBlocks(ctx, jw, blk); err != nil {
		errChan <- err
	}
}

//ComponentType implements Component interface method
func (bi *DslFile2Json) ComponentType() reflect.Type {
	return reflect.TypeOf(bi)
}

//FPrintBlocks prints the Block array in json format
func FPrintBlocks(ctx context.Context, w io.Writer, dat []dsl.Block) error {
	p, err := json.MarshalIndent(dat, "\n", " ")
	if err != nil {
		return err
	}
	fmt.Fprintln(w, string(p))
	/*
		//fmt.Fprintf(w, "%v\n", dat)
		// debug
		if util.Debug {
			util.DebugInfo(ctx, fmt.Sprintf("FPrintBlocks START\n blocks %v\n", len(dat)))
			for _, blk := range dat {
				util.DebugInfo(ctx, customPrint(&blk))
			}
			util.DebugInfo(ctx, "FPrintBlocks END\n")
		}
	*/
	return nil
}

func customPrint(blk *dsl.Block) string {
	var b bytes.Buffer
	b.WriteString("\n{") // start
	name := fmt.Sprintf("\"name\": \"%v\",\n", blk.Key)
	b.WriteString(name)

	parent := fmt.Sprintf("\"parent\": \"%v\",\n", blk.ParentNdx)
	b.WriteString(parent)
	b.WriteString("\"nodes\":[\n") // start nodes
	for _, n := range blk.Nodes {
		b.WriteString("{")
		b.WriteString(fmt.Sprintf("\t\"name\": \"%v\",\n", n.Key))
		b.WriteString(fmt.Sprintf("\t\"lnum\": %v,\n", n.Number))

		b.WriteString("\t\"lines\": [\n")
		for n, lines := range n.Lines {
			if n > 0 {
				b.WriteString(",\n")
			}
			b.WriteString(fmt.Sprintf("\t\"%v\"", lines))
		}

		b.WriteString("\t]}\n")
	}
	b.WriteString("\n]")   // end nodes
	b.WriteString("\n}\n") // end block
	return b.String()
}

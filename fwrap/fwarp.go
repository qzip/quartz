// https://github.com/antlr/antlr4/blob/master/doc/go-target.md

package main

import (
	"fmt"
	"os"

	"fwrap/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseClarityListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (tsl *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	lexer := parser.NewFuncWrapLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewClarityParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Sexpr()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

package dsl

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"qz/util"
	"strings"
)

// Parser parses the DSL
type Parser struct {
	in        *bufio.Reader
	blocks    []Block
	stack     stack
	multiLine *KeyLines
	curLine   int
}

// DSL Keywords
const (
	Comment        = "#"
	BlockStart     = "BEGIN"
	BlockEnd       = "END"
	MultiLineStart = "[["
	MultiLineEnd   = "]]"
)

type stack struct {
	ele []int
}

// NewParser Instantiates a new Parser.
// Block 0 is the root block. "__root__"
func NewParser(in io.Reader) *Parser {
	p := &Parser{
		in:     bufio.NewReader(in),
		blocks: make([]Block, 0),
	}
	p.blocks = append(p.blocks, *NewBlock(RootBlock))
	p.stack.push(0)

	return p
}

// Parse the DSL
func (p *Parser) Parse(ctx context.Context) ([]Block, error) {
	p.multiLine = nil
	for p.curLine = 1; ; p.curLine++ {
		lnr, err := p.in.ReadString('\n')
		if lnr != "" {
			ln := strings.TrimSpace(lnr)
			if strings.HasPrefix(ln, Comment) {
				continue
			}

			switch strings.ToUpper(strings.Split(ln, " ")[0]) {
			case "#":
				if p.multiLine != nil {
					p.appendLine(ctx, ln)
				}
				continue
			case "[[":
				util.DebugInfo(ctx, fmt.Sprintf("Parse:line %v  multiline start", p.curLine))
				if p.multiLine != nil {
					return nil, fmt.Errorf("error: line:%v invalid token '[[', already in multiline", p.curLine)
				}
				p.startMultiLine()
			case "]]":
				util.DebugInfo(ctx, fmt.Sprintf("Parse:line %v   multiline end", p.curLine))
				if p.multiLine == nil {
					return nil, fmt.Errorf("error: line:%v invalid token ']]', NOT in multiline", p.curLine)
				}
				p.endMultiLine()
			case "BEGIN":
				util.DebugInfo(ctx, fmt.Sprintf("Parse:line %v BEGIN", p.curLine))
				if p.multiLine != nil {
					p.appendLine(ctx, ln)
				} else {
					p.startBlock(ln)
				}
			case "END":
				util.DebugInfo(ctx, fmt.Sprintf("Parse:line %v  multiline END", p.curLine))
				if p.multiLine != nil {
					p.appendLine(ctx, ln)
				} else {
					p.endBlock()
				}
			default:
				//util.DebugInfo(ctx, "Parse:  DEFAULT")
				p.appendLine(ctx, lnr)
			}

		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

	}

	return p.blocks, nil
}

func (p *Parser) startBlock(ln string) {
	if p.multiLine != nil {
		return
	}
	toks := strings.Split(ln, " ")
	bname := ""
	if len(toks) > 1 {
		bname = toks[1]
	}
	blk := NewBlock(bname)
	if pndx, ok := p.stack.peek(); ok {
		blk.ParentNdx = pndx
	}
	p.blocks = append(p.blocks, *blk)
	p.stack.push(len(p.blocks) - 1)

}
func (p *Parser) endBlock() {
	if p.multiLine != nil {
		return
	}
	p.stack.pop()
}

func (p *Parser) startMultiLine() {
	if p.multiLine != nil {
		return
	}
	p.multiLine = NewKeyLines("", p.curLine)
}

func (p *Parser) endMultiLine() {
	top, _ := p.stack.peek()
	curBlock := &p.blocks[top]
	if len(curBlock.Nodes) == 0 {
		curBlock.Nodes = append(curBlock.Nodes, *p.multiLine)
	} else {
		keyln := &curBlock.Nodes[len(curBlock.Nodes)-1]

		keyln.Lines = append(keyln.Lines, p.multiLine.Lines...)
		/*
			for _, kl := range p.multiLine.Lines {
				keyln.Lines = append(keyln.Lines, kl)
			}
		*/

	}

	p.multiLine = nil
}

func (p *Parser) appendLine(ctx context.Context, ln string) {
	if p.multiLine != nil {
		p.multiLine.Lines = append(p.multiLine.Lines, ln)
	} else {
		ln = strings.TrimSpace(ln)
		if len(ln) == 0 {
			return
		}
		toks := strings.Split(ln, " ")
		kname := ""
		if len(toks) >= 1 {
			kname = toks[0]
		}
		top, _ := p.stack.peek()
		curBlock := &p.blocks[top]
		l := NewKeyLines(kname, p.curLine)

		ndx := strings.Index(ln, " ")
		if ndx > 1 && ndx < len(ln)-1 {
			val := ln[ndx:]
			l.Lines = append(l.Lines, val)
			//			util.DebugInfo(ctx, fmt.Sprintf("Parse.appendLine: appeded %v\n", val))
		}
		curBlock.Nodes = append(curBlock.Nodes, *l)
		util.DebugInfo(ctx, fmt.Sprintf("Parse.appendLine: single line appended.\n Block name %v \n, line %v ", curBlock.Key, l))

	}
}

//**************************** STACK *************

func (s *stack) push(blockNdx int) {
	s.ele = append(s.ele, blockNdx)
}
func (s *stack) pop() (top int) {
	n := len(s.ele) - 1 // Top element
	if n >= 0 {
		top = s.ele[n]
		s.ele = s.ele[:n] // Pop
	}
	return
}
func (s *stack) peek() (top int, ok bool) {
	n := len(s.ele) - 1 // Top element
	if n >= 0 {
		top = s.ele[n]
		ok = true
	}
	return
}

/*
func (s *stack) depth() int {
	return len(s.ele)
}
*/

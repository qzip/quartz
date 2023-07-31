//Package dsl Domain Specific Language
package dsl

import (
	"strings"
)

//Node can be a Block or a KeyLinesArr
type Node interface {
	Name() string
}

//TODO: make Block part of Array (Blocks)

//Block  BEGIN <name> \n ..key lines. END
// # is a comment line
// Keywords must be at the begining of a new line
type Block struct {
	Key       string     `json:"name"`
	ParentNdx int        `json:"parent"`
	Nodes     []KeyLines `json:"nodes"`
}

const (
	//RootBlock any keyword without a block is part of root
	RootBlock = "__root__"
)

//KeyLines keyword at start of line followed by 1 or multi lines
// Multi lines start & end with [[ ]] at begining of a new line
type KeyLines struct {
	Key    string   `json:"name"`
	Number int      `json:"lnum"`
	Lines  []string `json:"lines"`
}

func NewBlock(name string) *Block {
	return &Block{
		Key:   name,
		Nodes: make([]KeyLines, 0),
	}
}
func (n *Block) Name() string {
	return n.Key
}

//Blocks block array for operating on array of blocks
type Blocks []Block

//ChilderenOf collect the children
func (bx Blocks) ChilderenOf(rootNdx int) (childeren []int) {
	childeren = make([]int, 0)
	if rootNdx >= len(bx) {
		return nil
	}
	barr := []Block(bx)
	for n, blk := range barr {
		if n == 0 {
			continue // zero is always root
		}

		if blk.ParentNdx == rootNdx {
			if n == rootNdx {
				continue // self referential
			}
			childeren = append(childeren, n)
		}
	}
	return
}

//TreeOfBlocks block array to tree
type TreeOfBlocks struct {
	Node      *Block
	Childeren []*Block
}

//BuildTreeOfBlocks  block array to tree
func BuildTreeOfBlocks(barr []Block) *TreeOfBlocks {
	tob := &TreeOfBlocks{}
	return tob
}

/*
//MarshalJSON custom marshal
func (n *Block) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("\n{") // start
	name := fmt.Sprintf("\"name\": \"%v\",\n", n.Key)
	b.WriteString(name)

	parent := fmt.Sprintf("\"parent\": \"%v\",\n", n.Parent)
	b.WriteString(parent)
	b.WriteString("\"nodes\":[\n") // start nodes

	b.WriteString("\n]") // end nodes
	b.WriteString("\n}") // end block
}
*/

func NewKeyLines(key string, lineNum int) *KeyLines {
	return &KeyLines{
		Key: key, Number: lineNum, Lines: make([]string, 0),
	}
}

func (n *KeyLines) Name() string {
	return n.Key
}

//Lines2String lines to string
func (n *KeyLines) Lines2String() string {
	var b strings.Builder
	for i := 1; i < len(n.Lines); i++ {
		b.WriteString(n.Lines[i])
	}
	return b.String()
}

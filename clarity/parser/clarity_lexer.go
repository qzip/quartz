// Code generated from Clarity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ClarityLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var claritylexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func claritylexerLexerInit() {
	staticData := &claritylexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "':'", "'{'", "','", "'}'", "", "", "", "", "", "", "", "", "'('",
		"')'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "STRING", "UTF8", "WHITESPACE", "NUMBER", "HEX",
		"SYMBOL", "BUILTIN", "PRINCIPAL", "LPAREN", "RPAREN", "DOT", "BlockComment",
		"LineComment",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "STRING", "UTF8", "WHITESPACE", "NUMBER",
		"HEX", "SYMBOL", "BUILTIN", "PRINCIPAL", "LPAREN", "RPAREN", "DOT",
		"SYMBOL_START", "DIGIT", "BlockComment", "LineComment",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 143, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 52, 8, 4, 10, 4,
		12, 4, 55, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 4, 6, 63, 8, 6, 11,
		6, 12, 6, 64, 1, 6, 1, 6, 1, 7, 3, 7, 70, 8, 7, 1, 7, 4, 7, 73, 8, 7, 11,
		7, 12, 7, 74, 1, 7, 1, 7, 4, 7, 79, 8, 7, 11, 7, 12, 7, 80, 3, 7, 83, 8,
		7, 1, 7, 3, 7, 86, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 92, 8, 8, 1, 9,
		1, 9, 1, 9, 5, 9, 97, 8, 9, 10, 9, 12, 9, 100, 9, 9, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 3,
		15, 115, 8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 123, 8,
		17, 10, 17, 12, 17, 126, 9, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 5, 18, 137, 8, 18, 10, 18, 12, 18, 140, 9, 18, 1,
		18, 1, 18, 1, 124, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 0, 33, 0,
		35, 16, 37, 17, 1, 0, 7, 2, 0, 34, 34, 92, 92, 3, 0, 9, 10, 13, 13, 32,
		32, 2, 0, 43, 43, 45, 45, 2, 0, 65, 70, 97, 102, 2, 0, 33, 33, 63, 63,
		5, 0, 42, 43, 45, 47, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 154,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 35, 1,
		0, 0, 0, 0, 37, 1, 0, 0, 0, 1, 39, 1, 0, 0, 0, 3, 41, 1, 0, 0, 0, 5, 43,
		1, 0, 0, 0, 7, 45, 1, 0, 0, 0, 9, 47, 1, 0, 0, 0, 11, 58, 1, 0, 0, 0, 13,
		62, 1, 0, 0, 0, 15, 85, 1, 0, 0, 0, 17, 87, 1, 0, 0, 0, 19, 93, 1, 0, 0,
		0, 21, 101, 1, 0, 0, 0, 23, 104, 1, 0, 0, 0, 25, 107, 1, 0, 0, 0, 27, 109,
		1, 0, 0, 0, 29, 111, 1, 0, 0, 0, 31, 114, 1, 0, 0, 0, 33, 116, 1, 0, 0,
		0, 35, 118, 1, 0, 0, 0, 37, 132, 1, 0, 0, 0, 39, 40, 5, 58, 0, 0, 40, 2,
		1, 0, 0, 0, 41, 42, 5, 123, 0, 0, 42, 4, 1, 0, 0, 0, 43, 44, 5, 44, 0,
		0, 44, 6, 1, 0, 0, 0, 45, 46, 5, 125, 0, 0, 46, 8, 1, 0, 0, 0, 47, 53,
		5, 34, 0, 0, 48, 49, 5, 92, 0, 0, 49, 52, 9, 0, 0, 0, 50, 52, 8, 0, 0,
		0, 51, 48, 1, 0, 0, 0, 51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51,
		1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 56, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0,
		56, 57, 5, 34, 0, 0, 57, 10, 1, 0, 0, 0, 58, 59, 5, 117, 0, 0, 59, 60,
		3, 9, 4, 0, 60, 12, 1, 0, 0, 0, 61, 63, 7, 1, 0, 0, 62, 61, 1, 0, 0, 0,
		63, 64, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 1,
		0, 0, 0, 66, 67, 6, 6, 0, 0, 67, 14, 1, 0, 0, 0, 68, 70, 7, 2, 0, 0, 69,
		68, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 72, 1, 0, 0, 0, 71, 73, 3, 33,
		16, 0, 72, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74,
		75, 1, 0, 0, 0, 75, 82, 1, 0, 0, 0, 76, 78, 5, 46, 0, 0, 77, 79, 3, 33,
		16, 0, 78, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80,
		81, 1, 0, 0, 0, 81, 83, 1, 0, 0, 0, 82, 76, 1, 0, 0, 0, 82, 83, 1, 0, 0,
		0, 83, 86, 1, 0, 0, 0, 84, 86, 3, 17, 8, 0, 85, 69, 1, 0, 0, 0, 85, 84,
		1, 0, 0, 0, 86, 16, 1, 0, 0, 0, 87, 88, 5, 48, 0, 0, 88, 91, 5, 120, 0,
		0, 89, 92, 3, 33, 16, 0, 90, 92, 7, 3, 0, 0, 91, 89, 1, 0, 0, 0, 91, 90,
		1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 18, 1, 0, 0, 0, 93, 98, 3, 31, 15,
		0, 94, 97, 3, 31, 15, 0, 95, 97, 3, 33, 16, 0, 96, 94, 1, 0, 0, 0, 96,
		95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0,
		0, 0, 99, 20, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 102, 3, 19, 9, 0, 102,
		103, 7, 4, 0, 0, 103, 22, 1, 0, 0, 0, 104, 105, 5, 39, 0, 0, 105, 106,
		3, 19, 9, 0, 106, 24, 1, 0, 0, 0, 107, 108, 5, 40, 0, 0, 108, 26, 1, 0,
		0, 0, 109, 110, 5, 41, 0, 0, 110, 28, 1, 0, 0, 0, 111, 112, 5, 46, 0, 0,
		112, 30, 1, 0, 0, 0, 113, 115, 7, 5, 0, 0, 114, 113, 1, 0, 0, 0, 115, 32,
		1, 0, 0, 0, 116, 117, 2, 48, 57, 0, 117, 34, 1, 0, 0, 0, 118, 119, 5, 47,
		0, 0, 119, 120, 5, 42, 0, 0, 120, 124, 1, 0, 0, 0, 121, 123, 9, 0, 0, 0,
		122, 121, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 124,
		122, 1, 0, 0, 0, 125, 127, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 128,
		5, 42, 0, 0, 128, 129, 5, 47, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 6,
		17, 0, 0, 131, 36, 1, 0, 0, 0, 132, 133, 5, 59, 0, 0, 133, 134, 5, 59,
		0, 0, 134, 138, 1, 0, 0, 0, 135, 137, 8, 6, 0, 0, 136, 135, 1, 0, 0, 0,
		137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		141, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 6, 18, 0, 0, 142, 38,
		1, 0, 0, 0, 15, 0, 51, 53, 64, 69, 74, 80, 82, 85, 91, 96, 98, 114, 124,
		138, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ClarityLexerInit initializes any static state used to implement ClarityLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewClarityLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClarityLexerInit() {
	staticData := &claritylexerLexerStaticData
	staticData.once.Do(claritylexerLexerInit)
}

// NewClarityLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewClarityLexer(input antlr.CharStream) *ClarityLexer {
	ClarityLexerInit()
	l := new(ClarityLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &claritylexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Clarity.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ClarityLexer tokens.
const (
	ClarityLexerT__0         = 1
	ClarityLexerT__1         = 2
	ClarityLexerT__2         = 3
	ClarityLexerT__3         = 4
	ClarityLexerSTRING       = 5
	ClarityLexerUTF8         = 6
	ClarityLexerWHITESPACE   = 7
	ClarityLexerNUMBER       = 8
	ClarityLexerHEX          = 9
	ClarityLexerSYMBOL       = 10
	ClarityLexerBUILTIN      = 11
	ClarityLexerPRINCIPAL    = 12
	ClarityLexerLPAREN       = 13
	ClarityLexerRPAREN       = 14
	ClarityLexerDOT          = 15
	ClarityLexerBlockComment = 16
	ClarityLexerLineComment  = 17
)

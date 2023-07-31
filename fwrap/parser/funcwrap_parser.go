// Code generated from FuncWrapParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // FuncWrapParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 220,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 6, 2, 56,
	10, 2, 13, 2, 14, 2, 57, 3, 3, 3, 3, 5, 3, 62, 10, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 70, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	5, 6, 78, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 83, 10, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 5, 8, 89, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 97,
	10, 10, 3, 10, 3, 10, 5, 10, 101, 10, 10, 3, 10, 3, 10, 7, 10, 105, 10,
	10, 12, 10, 14, 10, 108, 11, 10, 3, 10, 5, 10, 111, 10, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 5, 11, 117, 10, 11, 3, 12, 3, 12, 5, 12, 121, 10, 12, 3,
	13, 3, 13, 5, 13, 125, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	132, 10, 14, 3, 15, 3, 15, 3, 15, 7, 15, 137, 10, 15, 12, 15, 14, 15, 140,
	11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 148, 10, 16, 12,
	16, 14, 16, 151, 11, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 157, 10, 16,
	12, 16, 14, 16, 160, 11, 16, 3, 16, 5, 16, 163, 10, 16, 3, 17, 3, 17, 3,
	17, 3, 17, 7, 17, 169, 10, 17, 12, 17, 14, 17, 172, 11, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 5, 17, 178, 10, 17, 3, 18, 3, 18, 3, 18, 3, 19, 5, 19, 184,
	10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 5, 20, 192, 10, 20, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 200, 10, 22, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 6, 25, 209, 10, 25, 13, 25, 14, 25,
	210, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 2, 2, 28,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 42, 44, 46, 48, 50, 52, 2, 4, 3, 2, 21, 28, 3, 2, 30, 33, 2, 226, 2,
	55, 3, 2, 2, 2, 4, 61, 3, 2, 2, 2, 6, 69, 3, 2, 2, 2, 8, 71, 3, 2, 2, 2,
	10, 77, 3, 2, 2, 2, 12, 79, 3, 2, 2, 2, 14, 88, 3, 2, 2, 2, 16, 90, 3,
	2, 2, 2, 18, 92, 3, 2, 2, 2, 20, 114, 3, 2, 2, 2, 22, 120, 3, 2, 2, 2,
	24, 124, 3, 2, 2, 2, 26, 131, 3, 2, 2, 2, 28, 133, 3, 2, 2, 2, 30, 141,
	3, 2, 2, 2, 32, 177, 3, 2, 2, 2, 34, 179, 3, 2, 2, 2, 36, 183, 3, 2, 2,
	2, 38, 191, 3, 2, 2, 2, 40, 193, 3, 2, 2, 2, 42, 199, 3, 2, 2, 2, 44, 201,
	3, 2, 2, 2, 46, 203, 3, 2, 2, 2, 48, 206, 3, 2, 2, 2, 50, 212, 3, 2, 2,
	2, 52, 215, 3, 2, 2, 2, 54, 56, 5, 4, 3, 2, 55, 54, 3, 2, 2, 2, 56, 57,
	3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 3, 3, 2, 2, 2,
	59, 62, 5, 18, 10, 2, 60, 62, 5, 6, 4, 2, 61, 59, 3, 2, 2, 2, 61, 60, 3,
	2, 2, 2, 62, 5, 3, 2, 2, 2, 63, 70, 5, 34, 18, 2, 64, 70, 5, 8, 5, 2, 65,
	70, 5, 12, 7, 2, 66, 70, 5, 30, 16, 2, 67, 70, 5, 46, 24, 2, 68, 70, 5,
	52, 27, 2, 69, 63, 3, 2, 2, 2, 69, 64, 3, 2, 2, 2, 69, 65, 3, 2, 2, 2,
	69, 66, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 7, 3, 2,
	2, 2, 71, 72, 5, 34, 18, 2, 72, 73, 7, 19, 2, 2, 73, 74, 5, 10, 6, 2, 74,
	9, 3, 2, 2, 2, 75, 78, 5, 26, 14, 2, 76, 78, 5, 16, 9, 2, 77, 75, 3, 2,
	2, 2, 77, 76, 3, 2, 2, 2, 78, 11, 3, 2, 2, 2, 79, 80, 5, 14, 8, 2, 80,
	82, 7, 13, 2, 2, 81, 83, 5, 28, 15, 2, 82, 81, 3, 2, 2, 2, 82, 83, 3, 2,
	2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 7, 14, 2, 2, 85, 13, 3, 2, 2, 2, 86,
	89, 5, 16, 9, 2, 87, 89, 5, 34, 18, 2, 88, 86, 3, 2, 2, 2, 88, 87, 3, 2,
	2, 2, 89, 15, 3, 2, 2, 2, 90, 91, 7, 12, 2, 2, 91, 17, 3, 2, 2, 2, 92,
	93, 7, 8, 2, 2, 93, 94, 5, 16, 9, 2, 94, 96, 7, 13, 2, 2, 95, 97, 5, 28,
	15, 2, 96, 95, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98,
	100, 7, 14, 2, 2, 99, 101, 5, 22, 12, 2, 100, 99, 3, 2, 2, 2, 100, 101,
	3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 106, 7, 15, 2, 2, 103, 105, 5, 6,
	4, 2, 104, 103, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2,
	106, 107, 3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109,
	111, 5, 20, 11, 2, 110, 109, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 112,
	3, 2, 2, 2, 112, 113, 7, 16, 2, 2, 113, 19, 3, 2, 2, 2, 114, 116, 7, 9,
	2, 2, 115, 117, 5, 24, 13, 2, 116, 115, 3, 2, 2, 2, 116, 117, 3, 2, 2,
	2, 117, 21, 3, 2, 2, 2, 118, 121, 5, 42, 22, 2, 119, 121, 5, 16, 9, 2,
	120, 118, 3, 2, 2, 2, 120, 119, 3, 2, 2, 2, 121, 23, 3, 2, 2, 2, 122, 125,
	5, 26, 14, 2, 123, 125, 5, 16, 9, 2, 124, 122, 3, 2, 2, 2, 124, 123, 3,
	2, 2, 2, 125, 25, 3, 2, 2, 2, 126, 132, 5, 34, 18, 2, 127, 132, 5, 36,
	19, 2, 128, 132, 5, 12, 7, 2, 129, 132, 5, 42, 22, 2, 130, 132, 5, 48,
	25, 2, 131, 126, 3, 2, 2, 2, 131, 127, 3, 2, 2, 2, 131, 128, 3, 2, 2, 2,
	131, 129, 3, 2, 2, 2, 131, 130, 3, 2, 2, 2, 132, 27, 3, 2, 2, 2, 133, 138,
	5, 26, 14, 2, 134, 135, 7, 20, 2, 2, 135, 137, 5, 28, 15, 2, 136, 134,
	3, 2, 2, 2, 137, 140, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138, 139, 3, 2,
	2, 2, 139, 29, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 141, 142, 7, 3, 2, 2,
	142, 143, 7, 13, 2, 2, 143, 144, 5, 32, 17, 2, 144, 145, 7, 14, 2, 2, 145,
	149, 7, 15, 2, 2, 146, 148, 5, 6, 4, 2, 147, 146, 3, 2, 2, 2, 148, 151,
	3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 152, 3, 2,
	2, 2, 151, 149, 3, 2, 2, 2, 152, 162, 7, 16, 2, 2, 153, 154, 7, 4, 2, 2,
	154, 158, 7, 15, 2, 2, 155, 157, 5, 6, 4, 2, 156, 155, 3, 2, 2, 2, 157,
	160, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 161,
	3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 161, 163, 7, 16, 2, 2, 162, 153, 3, 2,
	2, 2, 162, 163, 3, 2, 2, 2, 163, 31, 3, 2, 2, 2, 164, 170, 5, 26, 14, 2,
	165, 166, 5, 40, 21, 2, 166, 167, 5, 32, 17, 2, 167, 169, 3, 2, 2, 2, 168,
	165, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170, 171,
	3, 2, 2, 2, 171, 178, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 173, 174, 7, 13,
	2, 2, 174, 175, 5, 32, 17, 2, 175, 176, 7, 14, 2, 2, 176, 178, 3, 2, 2,
	2, 177, 164, 3, 2, 2, 2, 177, 173, 3, 2, 2, 2, 178, 33, 3, 2, 2, 2, 179,
	180, 7, 6, 2, 2, 180, 181, 7, 12, 2, 2, 181, 35, 3, 2, 2, 2, 182, 184,
	5, 38, 20, 2, 183, 182, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 185, 3,
	2, 2, 2, 185, 186, 7, 17, 2, 2, 186, 187, 5, 28, 15, 2, 187, 188, 7, 18,
	2, 2, 188, 37, 3, 2, 2, 2, 189, 192, 5, 34, 18, 2, 190, 192, 5, 48, 25,
	2, 191, 189, 3, 2, 2, 2, 191, 190, 3, 2, 2, 2, 192, 39, 3, 2, 2, 2, 193,
	194, 9, 2, 2, 2, 194, 41, 3, 2, 2, 2, 195, 200, 7, 7, 2, 2, 196, 200, 7,
	34, 2, 2, 197, 200, 7, 35, 2, 2, 198, 200, 5, 44, 23, 2, 199, 195, 3, 2,
	2, 2, 199, 196, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 198, 3, 2, 2, 2,
	200, 43, 3, 2, 2, 2, 201, 202, 9, 3, 2, 2, 202, 45, 3, 2, 2, 2, 203, 204,
	7, 10, 2, 2, 204, 205, 5, 34, 18, 2, 205, 47, 3, 2, 2, 2, 206, 208, 5,
	34, 18, 2, 207, 209, 5, 50, 26, 2, 208, 207, 3, 2, 2, 2, 209, 210, 3, 2,
	2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 49, 3, 2, 2, 2,
	212, 213, 7, 11, 2, 2, 213, 214, 7, 12, 2, 2, 214, 51, 3, 2, 2, 2, 215,
	216, 5, 48, 25, 2, 216, 217, 7, 19, 2, 2, 217, 218, 5, 10, 6, 2, 218, 53,
	3, 2, 2, 2, 26, 57, 61, 69, 77, 82, 88, 96, 100, 106, 110, 116, 120, 124,
	131, 138, 149, 158, 162, 170, 177, 183, 191, 199, 210,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'if'", "'else'", "'nil'", "'$'", "", "'func'", "'return'", "'object'",
	"'.'", "", "'('", "')'", "'{'", "'}'", "'['", "']'", "'='", "','", "'||'",
	"'&&'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'!'",
}
var symbolicNames = []string{
	"", "IF", "ELSE", "NIL_LIT", "VAR", "BOOLEAN", "FUNC", "RETURN", "OBJ",
	"DOT", "IDENTIFIER", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET",
	"R_BRACKET", "ASSIGN", "COMMA", "LOGICAL_OR", "LOGICAL_AND", "EQUALS",
	"NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
	"EXCLAMATION", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "RAW_STRING_LIT",
	"INTERPRETED_STRING_LIT", "WS", "COMMENT", "TERMINATOR", "LINE_COMMENT",
}

var ruleNames = []string{
	"prog", "pstatement", "statement", "varAssign", "asnVal", "funCall", "funValName",
	"funcName", "funDef", "retStmt", "fVal", "retVal", "rVal", "rVals", "cond",
	"eval", "variable", "array", "arrRef", "op", "literals", "numbers", "objectDef",
	"objMemberRef", "objElement", "objAssign",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FuncWrapParser struct {
	*antlr.BaseParser
}

func NewFuncWrapParser(input antlr.TokenStream) *FuncWrapParser {
	this := new(FuncWrapParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FuncWrapParser.g4"

	return this
}

// FuncWrapParser tokens.
const (
	FuncWrapParserEOF                    = antlr.TokenEOF
	FuncWrapParserIF                     = 1
	FuncWrapParserELSE                   = 2
	FuncWrapParserNIL_LIT                = 3
	FuncWrapParserVAR                    = 4
	FuncWrapParserBOOLEAN                = 5
	FuncWrapParserFUNC                   = 6
	FuncWrapParserRETURN                 = 7
	FuncWrapParserOBJ                    = 8
	FuncWrapParserDOT                    = 9
	FuncWrapParserIDENTIFIER             = 10
	FuncWrapParserL_PAREN                = 11
	FuncWrapParserR_PAREN                = 12
	FuncWrapParserL_CURLY                = 13
	FuncWrapParserR_CURLY                = 14
	FuncWrapParserL_BRACKET              = 15
	FuncWrapParserR_BRACKET              = 16
	FuncWrapParserASSIGN                 = 17
	FuncWrapParserCOMMA                  = 18
	FuncWrapParserLOGICAL_OR             = 19
	FuncWrapParserLOGICAL_AND            = 20
	FuncWrapParserEQUALS                 = 21
	FuncWrapParserNOT_EQUALS             = 22
	FuncWrapParserLESS                   = 23
	FuncWrapParserLESS_OR_EQUALS         = 24
	FuncWrapParserGREATER                = 25
	FuncWrapParserGREATER_OR_EQUALS      = 26
	FuncWrapParserEXCLAMATION            = 27
	FuncWrapParserDECIMAL_LIT            = 28
	FuncWrapParserOCTAL_LIT              = 29
	FuncWrapParserHEX_LIT                = 30
	FuncWrapParserFLOAT_LIT              = 31
	FuncWrapParserRAW_STRING_LIT         = 32
	FuncWrapParserINTERPRETED_STRING_LIT = 33
	FuncWrapParserWS                     = 34
	FuncWrapParserCOMMENT                = 35
	FuncWrapParserTERMINATOR             = 36
	FuncWrapParserLINE_COMMENT           = 37
)

// FuncWrapParser rules.
const (
	FuncWrapParserRULE_prog         = 0
	FuncWrapParserRULE_pstatement   = 1
	FuncWrapParserRULE_statement    = 2
	FuncWrapParserRULE_varAssign    = 3
	FuncWrapParserRULE_asnVal       = 4
	FuncWrapParserRULE_funCall      = 5
	FuncWrapParserRULE_funValName   = 6
	FuncWrapParserRULE_funcName     = 7
	FuncWrapParserRULE_funDef       = 8
	FuncWrapParserRULE_retStmt      = 9
	FuncWrapParserRULE_fVal         = 10
	FuncWrapParserRULE_retVal       = 11
	FuncWrapParserRULE_rVal         = 12
	FuncWrapParserRULE_rVals        = 13
	FuncWrapParserRULE_cond         = 14
	FuncWrapParserRULE_eval         = 15
	FuncWrapParserRULE_variable     = 16
	FuncWrapParserRULE_array        = 17
	FuncWrapParserRULE_arrRef       = 18
	FuncWrapParserRULE_op           = 19
	FuncWrapParserRULE_literals     = 20
	FuncWrapParserRULE_numbers      = 21
	FuncWrapParserRULE_objectDef    = 22
	FuncWrapParserRULE_objMemberRef = 23
	FuncWrapParserRULE_objElement   = 24
	FuncWrapParserRULE_objAssign    = 25
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllPstatement() []IPstatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPstatementContext)(nil)).Elem())
	var tst = make([]IPstatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPstatementContext)
		}
	}

	return tst
}

func (s *ProgContext) Pstatement(i int) IPstatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPstatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPstatementContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *FuncWrapParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FuncWrapParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FuncWrapParserIF)|(1<<FuncWrapParserVAR)|(1<<FuncWrapParserFUNC)|(1<<FuncWrapParserOBJ)|(1<<FuncWrapParserIDENTIFIER))) != 0) {
		{
			p.SetState(52)
			p.Pstatement()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPstatementContext is an interface to support dynamic dispatch.
type IPstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPstatementContext differentiates from other interfaces.
	IsPstatementContext()
}

type PstatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPstatementContext() *PstatementContext {
	var p = new(PstatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_pstatement
	return p
}

func (*PstatementContext) IsPstatementContext() {}

func NewPstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PstatementContext {
	var p = new(PstatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_pstatement

	return p
}

func (s *PstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PstatementContext) FunDef() IFunDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunDefContext)
}

func (s *PstatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *PstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterPstatement(s)
	}
}

func (s *PstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitPstatement(s)
	}
}

func (p *FuncWrapParser) Pstatement() (localctx IPstatementContext) {
	localctx = NewPstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FuncWrapParserRULE_pstatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FuncWrapParserFUNC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.FunDef()
		}

	case FuncWrapParserIF, FuncWrapParserVAR, FuncWrapParserOBJ, FuncWrapParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Statement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *StatementContext) VarAssign() IVarAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarAssignContext)
}

func (s *StatementContext) FunCall() IFunCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunCallContext)
}

func (s *StatementContext) Cond() ICondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondContext)
}

func (s *StatementContext) ObjectDef() IObjectDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectDefContext)
}

func (s *StatementContext) ObjAssign() IObjAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjAssignContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *FuncWrapParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FuncWrapParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.VarAssign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.FunCall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(64)
			p.Cond()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(65)
			p.ObjectDef()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(66)
			p.ObjAssign()
		}

	}

	return localctx
}

// IVarAssignContext is an interface to support dynamic dispatch.
type IVarAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarAssignContext differentiates from other interfaces.
	IsVarAssignContext()
}

type VarAssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarAssignContext() *VarAssignContext {
	var p = new(VarAssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_varAssign
	return p
}

func (*VarAssignContext) IsVarAssignContext() {}

func NewVarAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarAssignContext {
	var p = new(VarAssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_varAssign

	return p
}

func (s *VarAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *VarAssignContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VarAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserASSIGN, 0)
}

func (s *VarAssignContext) AsnVal() IAsnValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsnValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsnValContext)
}

func (s *VarAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterVarAssign(s)
	}
}

func (s *VarAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitVarAssign(s)
	}
}

func (p *FuncWrapParser) VarAssign() (localctx IVarAssignContext) {
	localctx = NewVarAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FuncWrapParserRULE_varAssign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Variable()
	}
	{
		p.SetState(70)
		p.Match(FuncWrapParserASSIGN)
	}
	{
		p.SetState(71)
		p.AsnVal()
	}

	return localctx
}

// IAsnValContext is an interface to support dynamic dispatch.
type IAsnValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAsnValContext differentiates from other interfaces.
	IsAsnValContext()
}

type AsnValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsnValContext() *AsnValContext {
	var p = new(AsnValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_asnVal
	return p
}

func (*AsnValContext) IsAsnValContext() {}

func NewAsnValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsnValContext {
	var p = new(AsnValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_asnVal

	return p
}

func (s *AsnValContext) GetParser() antlr.Parser { return s.parser }

func (s *AsnValContext) RVal() IRValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRValContext)
}

func (s *AsnValContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *AsnValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsnValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsnValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterAsnVal(s)
	}
}

func (s *AsnValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitAsnVal(s)
	}
}

func (p *FuncWrapParser) AsnVal() (localctx IAsnValContext) {
	localctx = NewAsnValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FuncWrapParserRULE_asnVal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.RVal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.FuncName()
		}

	}

	return localctx
}

// IFunCallContext is an interface to support dynamic dispatch.
type IFunCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunCallContext differentiates from other interfaces.
	IsFunCallContext()
}

type FunCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunCallContext() *FunCallContext {
	var p = new(FunCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_funCall
	return p
}

func (*FunCallContext) IsFunCallContext() {}

func NewFunCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunCallContext {
	var p = new(FunCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_funCall

	return p
}

func (s *FunCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunCallContext) FunValName() IFunValNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunValNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunValNameContext)
}

func (s *FunCallContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserL_PAREN, 0)
}

func (s *FunCallContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserR_PAREN, 0)
}

func (s *FunCallContext) RVals() IRValsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRValsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRValsContext)
}

func (s *FunCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterFunCall(s)
	}
}

func (s *FunCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitFunCall(s)
	}
}

func (p *FuncWrapParser) FunCall() (localctx IFunCallContext) {
	localctx = NewFunCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FuncWrapParserRULE_funCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.FunValName()
	}
	{
		p.SetState(78)
		p.Match(FuncWrapParserL_PAREN)
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(FuncWrapParserVAR-4))|(1<<(FuncWrapParserBOOLEAN-4))|(1<<(FuncWrapParserIDENTIFIER-4))|(1<<(FuncWrapParserL_BRACKET-4))|(1<<(FuncWrapParserDECIMAL_LIT-4))|(1<<(FuncWrapParserOCTAL_LIT-4))|(1<<(FuncWrapParserHEX_LIT-4))|(1<<(FuncWrapParserFLOAT_LIT-4))|(1<<(FuncWrapParserRAW_STRING_LIT-4))|(1<<(FuncWrapParserINTERPRETED_STRING_LIT-4)))) != 0 {
		{
			p.SetState(79)
			p.RVals()
		}

	}
	{
		p.SetState(82)
		p.Match(FuncWrapParserR_PAREN)
	}

	return localctx
}

// IFunValNameContext is an interface to support dynamic dispatch.
type IFunValNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunValNameContext differentiates from other interfaces.
	IsFunValNameContext()
}

type FunValNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunValNameContext() *FunValNameContext {
	var p = new(FunValNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_funValName
	return p
}

func (*FunValNameContext) IsFunValNameContext() {}

func NewFunValNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunValNameContext {
	var p = new(FunValNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_funValName

	return p
}

func (s *FunValNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FunValNameContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *FunValNameContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FunValNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunValNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunValNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterFunValName(s)
	}
}

func (s *FunValNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitFunValName(s)
	}
}

func (p *FuncWrapParser) FunValName() (localctx IFunValNameContext) {
	localctx = NewFunValNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FuncWrapParserRULE_funValName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FuncWrapParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.FuncName()
		}

	case FuncWrapParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Variable()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFuncNameContext is an interface to support dynamic dispatch.
type IFuncNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncNameContext differentiates from other interfaces.
	IsFuncNameContext()
}

type FuncNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncNameContext() *FuncNameContext {
	var p = new(FuncNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_funcName
	return p
}

func (*FuncNameContext) IsFuncNameContext() {}

func NewFuncNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncNameContext {
	var p = new(FuncNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_funcName

	return p
}

func (s *FuncNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserIDENTIFIER, 0)
}

func (s *FuncNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterFuncName(s)
	}
}

func (s *FuncNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitFuncName(s)
	}
}

func (p *FuncWrapParser) FuncName() (localctx IFuncNameContext) {
	localctx = NewFuncNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FuncWrapParserRULE_funcName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(FuncWrapParserIDENTIFIER)
	}

	return localctx
}

// IFunDefContext is an interface to support dynamic dispatch.
type IFunDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunDefContext differentiates from other interfaces.
	IsFunDefContext()
}

type FunDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunDefContext() *FunDefContext {
	var p = new(FunDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_funDef
	return p
}

func (*FunDefContext) IsFunDefContext() {}

func NewFunDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunDefContext {
	var p = new(FunDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_funDef

	return p
}

func (s *FunDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunDefContext) FUNC() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserFUNC, 0)
}

func (s *FunDefContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *FunDefContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserL_PAREN, 0)
}

func (s *FunDefContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserR_PAREN, 0)
}

func (s *FunDefContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserL_CURLY, 0)
}

func (s *FunDefContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserR_CURLY, 0)
}

func (s *FunDefContext) RVals() IRValsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRValsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRValsContext)
}

func (s *FunDefContext) FVal() IFValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFValContext)
}

func (s *FunDefContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *FunDefContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FunDefContext) RetStmt() IRetStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetStmtContext)
}

func (s *FunDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterFunDef(s)
	}
}

func (s *FunDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitFunDef(s)
	}
}

func (p *FuncWrapParser) FunDef() (localctx IFunDefContext) {
	localctx = NewFunDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FuncWrapParserRULE_funDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(FuncWrapParserFUNC)
	}
	{
		p.SetState(91)
		p.FuncName()
	}
	{
		p.SetState(92)
		p.Match(FuncWrapParserL_PAREN)
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(FuncWrapParserVAR-4))|(1<<(FuncWrapParserBOOLEAN-4))|(1<<(FuncWrapParserIDENTIFIER-4))|(1<<(FuncWrapParserL_BRACKET-4))|(1<<(FuncWrapParserDECIMAL_LIT-4))|(1<<(FuncWrapParserOCTAL_LIT-4))|(1<<(FuncWrapParserHEX_LIT-4))|(1<<(FuncWrapParserFLOAT_LIT-4))|(1<<(FuncWrapParserRAW_STRING_LIT-4))|(1<<(FuncWrapParserINTERPRETED_STRING_LIT-4)))) != 0 {
		{
			p.SetState(93)
			p.RVals()
		}

	}
	{
		p.SetState(96)
		p.Match(FuncWrapParserR_PAREN)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-5)&-(0x1f+1)) == 0 && ((1<<uint((_la-5)))&((1<<(FuncWrapParserBOOLEAN-5))|(1<<(FuncWrapParserIDENTIFIER-5))|(1<<(FuncWrapParserDECIMAL_LIT-5))|(1<<(FuncWrapParserOCTAL_LIT-5))|(1<<(FuncWrapParserHEX_LIT-5))|(1<<(FuncWrapParserFLOAT_LIT-5))|(1<<(FuncWrapParserRAW_STRING_LIT-5))|(1<<(FuncWrapParserINTERPRETED_STRING_LIT-5)))) != 0 {
		{
			p.SetState(97)
			p.FVal()
		}

	}
	{
		p.SetState(100)
		p.Match(FuncWrapParserL_CURLY)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FuncWrapParserIF)|(1<<FuncWrapParserVAR)|(1<<FuncWrapParserOBJ)|(1<<FuncWrapParserIDENTIFIER))) != 0 {
		{
			p.SetState(101)
			p.Statement()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FuncWrapParserRETURN {
		{
			p.SetState(107)
			p.RetStmt()
		}

	}
	{
		p.SetState(110)
		p.Match(FuncWrapParserR_CURLY)
	}

	return localctx
}

// IRetStmtContext is an interface to support dynamic dispatch.
type IRetStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetStmtContext differentiates from other interfaces.
	IsRetStmtContext()
}

type RetStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetStmtContext() *RetStmtContext {
	var p = new(RetStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_retStmt
	return p
}

func (*RetStmtContext) IsRetStmtContext() {}

func NewRetStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetStmtContext {
	var p = new(RetStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_retStmt

	return p
}

func (s *RetStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RetStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserRETURN, 0)
}

func (s *RetStmtContext) RetVal() IRetValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetValContext)
}

func (s *RetStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterRetStmt(s)
	}
}

func (s *RetStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitRetStmt(s)
	}
}

func (p *FuncWrapParser) RetStmt() (localctx IRetStmtContext) {
	localctx = NewRetStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FuncWrapParserRULE_retStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(FuncWrapParserRETURN)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(FuncWrapParserVAR-4))|(1<<(FuncWrapParserBOOLEAN-4))|(1<<(FuncWrapParserIDENTIFIER-4))|(1<<(FuncWrapParserL_BRACKET-4))|(1<<(FuncWrapParserDECIMAL_LIT-4))|(1<<(FuncWrapParserOCTAL_LIT-4))|(1<<(FuncWrapParserHEX_LIT-4))|(1<<(FuncWrapParserFLOAT_LIT-4))|(1<<(FuncWrapParserRAW_STRING_LIT-4))|(1<<(FuncWrapParserINTERPRETED_STRING_LIT-4)))) != 0 {
		{
			p.SetState(113)
			p.RetVal()
		}

	}

	return localctx
}

// IFValContext is an interface to support dynamic dispatch.
type IFValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFValContext differentiates from other interfaces.
	IsFValContext()
}

type FValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFValContext() *FValContext {
	var p = new(FValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_fVal
	return p
}

func (*FValContext) IsFValContext() {}

func NewFValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FValContext {
	var p = new(FValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_fVal

	return p
}

func (s *FValContext) GetParser() antlr.Parser { return s.parser }

func (s *FValContext) Literals() ILiteralsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralsContext)
}

func (s *FValContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *FValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterFVal(s)
	}
}

func (s *FValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitFVal(s)
	}
}

func (p *FuncWrapParser) FVal() (localctx IFValContext) {
	localctx = NewFValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FuncWrapParserRULE_fVal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FuncWrapParserBOOLEAN, FuncWrapParserDECIMAL_LIT, FuncWrapParserOCTAL_LIT, FuncWrapParserHEX_LIT, FuncWrapParserFLOAT_LIT, FuncWrapParserRAW_STRING_LIT, FuncWrapParserINTERPRETED_STRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Literals()
		}

	case FuncWrapParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.FuncName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRetValContext is an interface to support dynamic dispatch.
type IRetValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetValContext differentiates from other interfaces.
	IsRetValContext()
}

type RetValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetValContext() *RetValContext {
	var p = new(RetValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_retVal
	return p
}

func (*RetValContext) IsRetValContext() {}

func NewRetValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetValContext {
	var p = new(RetValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_retVal

	return p
}

func (s *RetValContext) GetParser() antlr.Parser { return s.parser }

func (s *RetValContext) RVal() IRValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRValContext)
}

func (s *RetValContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
}

func (s *RetValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterRetVal(s)
	}
}

func (s *RetValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitRetVal(s)
	}
}

func (p *FuncWrapParser) RetVal() (localctx IRetValContext) {
	localctx = NewRetValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FuncWrapParserRULE_retVal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.RVal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.FuncName()
		}

	}

	return localctx
}

// IRValContext is an interface to support dynamic dispatch.
type IRValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRValContext differentiates from other interfaces.
	IsRValContext()
}

type RValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRValContext() *RValContext {
	var p = new(RValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_rVal
	return p
}

func (*RValContext) IsRValContext() {}

func NewRValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RValContext {
	var p = new(RValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_rVal

	return p
}

func (s *RValContext) GetParser() antlr.Parser { return s.parser }

func (s *RValContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *RValContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *RValContext) FunCall() IFunCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunCallContext)
}

func (s *RValContext) Literals() ILiteralsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralsContext)
}

func (s *RValContext) ObjMemberRef() IObjMemberRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjMemberRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjMemberRefContext)
}

func (s *RValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterRVal(s)
	}
}

func (s *RValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitRVal(s)
	}
}

func (p *FuncWrapParser) RVal() (localctx IRValContext) {
	localctx = NewRValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FuncWrapParserRULE_rVal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Array()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			p.FunCall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(127)
			p.Literals()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)
			p.ObjMemberRef()
		}

	}

	return localctx
}

// IRValsContext is an interface to support dynamic dispatch.
type IRValsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRValsContext differentiates from other interfaces.
	IsRValsContext()
}

type RValsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRValsContext() *RValsContext {
	var p = new(RValsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_rVals
	return p
}

func (*RValsContext) IsRValsContext() {}

func NewRValsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RValsContext {
	var p = new(RValsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_rVals

	return p
}

func (s *RValsContext) GetParser() antlr.Parser { return s.parser }

func (s *RValsContext) RVal() IRValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRValContext)
}

func (s *RValsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FuncWrapParserCOMMA)
}

func (s *RValsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FuncWrapParserCOMMA, i)
}

func (s *RValsContext) AllRVals() []IRValsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRValsContext)(nil)).Elem())
	var tst = make([]IRValsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRValsContext)
		}
	}

	return tst
}

func (s *RValsContext) RVals(i int) IRValsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRValsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRValsContext)
}

func (s *RValsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RValsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RValsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterRVals(s)
	}
}

func (s *RValsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitRVals(s)
	}
}

func (p *FuncWrapParser) RVals() (localctx IRValsContext) {
	localctx = NewRValsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FuncWrapParserRULE_rVals)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.RVal()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(132)
				p.Match(FuncWrapParserCOMMA)
			}
			{
				p.SetState(133)
				p.RVals()
			}

		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// ICondContext is an interface to support dynamic dispatch.
type ICondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCondContext differentiates from other interfaces.
	IsCondContext()
}

type CondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondContext() *CondContext {
	var p = new(CondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_cond
	return p
}

func (*CondContext) IsCondContext() {}

func NewCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondContext {
	var p = new(CondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_cond

	return p
}

func (s *CondContext) GetParser() antlr.Parser { return s.parser }

func (s *CondContext) IF() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserIF, 0)
}

func (s *CondContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserL_PAREN, 0)
}

func (s *CondContext) Eval() IEvalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEvalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEvalContext)
}

func (s *CondContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserR_PAREN, 0)
}

func (s *CondContext) AllL_CURLY() []antlr.TerminalNode {
	return s.GetTokens(FuncWrapParserL_CURLY)
}

func (s *CondContext) L_CURLY(i int) antlr.TerminalNode {
	return s.GetToken(FuncWrapParserL_CURLY, i)
}

func (s *CondContext) AllR_CURLY() []antlr.TerminalNode {
	return s.GetTokens(FuncWrapParserR_CURLY)
}

func (s *CondContext) R_CURLY(i int) antlr.TerminalNode {
	return s.GetToken(FuncWrapParserR_CURLY, i)
}

func (s *CondContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *CondContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CondContext) ELSE() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserELSE, 0)
}

func (s *CondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterCond(s)
	}
}

func (s *CondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitCond(s)
	}
}

func (p *FuncWrapParser) Cond() (localctx ICondContext) {
	localctx = NewCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FuncWrapParserRULE_cond)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(FuncWrapParserIF)
	}
	{
		p.SetState(140)
		p.Match(FuncWrapParserL_PAREN)
	}
	{
		p.SetState(141)
		p.Eval()
	}
	{
		p.SetState(142)
		p.Match(FuncWrapParserR_PAREN)
	}
	{
		p.SetState(143)
		p.Match(FuncWrapParserL_CURLY)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FuncWrapParserIF)|(1<<FuncWrapParserVAR)|(1<<FuncWrapParserOBJ)|(1<<FuncWrapParserIDENTIFIER))) != 0 {
		{
			p.SetState(144)
			p.Statement()
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(150)
		p.Match(FuncWrapParserR_CURLY)
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FuncWrapParserELSE {
		{
			p.SetState(151)
			p.Match(FuncWrapParserELSE)
		}
		{
			p.SetState(152)
			p.Match(FuncWrapParserL_CURLY)
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FuncWrapParserIF)|(1<<FuncWrapParserVAR)|(1<<FuncWrapParserOBJ)|(1<<FuncWrapParserIDENTIFIER))) != 0 {
			{
				p.SetState(153)
				p.Statement()
			}

			p.SetState(158)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(159)
			p.Match(FuncWrapParserR_CURLY)
		}

	}

	return localctx
}

// IEvalContext is an interface to support dynamic dispatch.
type IEvalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEvalContext differentiates from other interfaces.
	IsEvalContext()
}

type EvalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEvalContext() *EvalContext {
	var p = new(EvalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_eval
	return p
}

func (*EvalContext) IsEvalContext() {}

func NewEvalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EvalContext {
	var p = new(EvalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_eval

	return p
}

func (s *EvalContext) GetParser() antlr.Parser { return s.parser }

func (s *EvalContext) RVal() IRValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRValContext)
}

func (s *EvalContext) AllOp() []IOpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOpContext)(nil)).Elem())
	var tst = make([]IOpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOpContext)
		}
	}

	return tst
}

func (s *EvalContext) Op(i int) IOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *EvalContext) AllEval() []IEvalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEvalContext)(nil)).Elem())
	var tst = make([]IEvalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEvalContext)
		}
	}

	return tst
}

func (s *EvalContext) Eval(i int) IEvalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEvalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEvalContext)
}

func (s *EvalContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserL_PAREN, 0)
}

func (s *EvalContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserR_PAREN, 0)
}

func (s *EvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EvalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EvalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterEval(s)
	}
}

func (s *EvalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitEval(s)
	}
}

func (p *FuncWrapParser) Eval() (localctx IEvalContext) {
	localctx = NewEvalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FuncWrapParserRULE_eval)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FuncWrapParserVAR, FuncWrapParserBOOLEAN, FuncWrapParserIDENTIFIER, FuncWrapParserL_BRACKET, FuncWrapParserDECIMAL_LIT, FuncWrapParserOCTAL_LIT, FuncWrapParserHEX_LIT, FuncWrapParserFLOAT_LIT, FuncWrapParserRAW_STRING_LIT, FuncWrapParserINTERPRETED_STRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.RVal()
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(163)
					p.Op()
				}
				{
					p.SetState(164)
					p.Eval()
				}

			}
			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
		}

	case FuncWrapParserL_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.Match(FuncWrapParserL_PAREN)
		}
		{
			p.SetState(172)
			p.Eval()
		}
		{
			p.SetState(173)
			p.Match(FuncWrapParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VAR() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserVAR, 0)
}

func (s *VariableContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserIDENTIFIER, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *FuncWrapParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FuncWrapParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(FuncWrapParserVAR)
	}
	{
		p.SetState(178)
		p.Match(FuncWrapParserIDENTIFIER)
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserL_BRACKET, 0)
}

func (s *ArrayContext) RVals() IRValsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRValsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRValsContext)
}

func (s *ArrayContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserR_BRACKET, 0)
}

func (s *ArrayContext) ArrRef() IArrRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrRefContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *FuncWrapParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FuncWrapParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FuncWrapParserVAR {
		{
			p.SetState(180)
			p.ArrRef()
		}

	}
	{
		p.SetState(183)
		p.Match(FuncWrapParserL_BRACKET)
	}
	{
		p.SetState(184)
		p.RVals()
	}
	{
		p.SetState(185)
		p.Match(FuncWrapParserR_BRACKET)
	}

	return localctx
}

// IArrRefContext is an interface to support dynamic dispatch.
type IArrRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrRefContext differentiates from other interfaces.
	IsArrRefContext()
}

type ArrRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrRefContext() *ArrRefContext {
	var p = new(ArrRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_arrRef
	return p
}

func (*ArrRefContext) IsArrRefContext() {}

func NewArrRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrRefContext {
	var p = new(ArrRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_arrRef

	return p
}

func (s *ArrRefContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrRefContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ArrRefContext) ObjMemberRef() IObjMemberRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjMemberRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjMemberRefContext)
}

func (s *ArrRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterArrRef(s)
	}
}

func (s *ArrRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitArrRef(s)
	}
}

func (p *FuncWrapParser) ArrRef() (localctx IArrRefContext) {
	localctx = NewArrRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FuncWrapParserRULE_arrRef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.ObjMemberRef()
		}

	}

	return localctx
}

// IOpContext is an interface to support dynamic dispatch.
type IOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpContext differentiates from other interfaces.
	IsOpContext()
}

type OpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpContext() *OpContext {
	var p = new(OpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_op
	return p
}

func (*OpContext) IsOpContext() {}

func NewOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpContext {
	var p = new(OpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_op

	return p
}

func (s *OpContext) GetParser() antlr.Parser { return s.parser }

func (s *OpContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserLOGICAL_OR, 0)
}

func (s *OpContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserLOGICAL_AND, 0)
}

func (s *OpContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserEQUALS, 0)
}

func (s *OpContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserNOT_EQUALS, 0)
}

func (s *OpContext) LESS() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserLESS, 0)
}

func (s *OpContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserLESS_OR_EQUALS, 0)
}

func (s *OpContext) GREATER() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserGREATER, 0)
}

func (s *OpContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserGREATER_OR_EQUALS, 0)
}

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterOp(s)
	}
}

func (s *OpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitOp(s)
	}
}

func (p *FuncWrapParser) Op() (localctx IOpContext) {
	localctx = NewOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FuncWrapParserRULE_op)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FuncWrapParserLOGICAL_OR)|(1<<FuncWrapParserLOGICAL_AND)|(1<<FuncWrapParserEQUALS)|(1<<FuncWrapParserNOT_EQUALS)|(1<<FuncWrapParserLESS)|(1<<FuncWrapParserLESS_OR_EQUALS)|(1<<FuncWrapParserGREATER)|(1<<FuncWrapParserGREATER_OR_EQUALS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralsContext is an interface to support dynamic dispatch.
type ILiteralsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralsContext differentiates from other interfaces.
	IsLiteralsContext()
}

type LiteralsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralsContext() *LiteralsContext {
	var p = new(LiteralsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_literals
	return p
}

func (*LiteralsContext) IsLiteralsContext() {}

func NewLiteralsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralsContext {
	var p = new(LiteralsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_literals

	return p
}

func (s *LiteralsContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralsContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserBOOLEAN, 0)
}

func (s *LiteralsContext) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserRAW_STRING_LIT, 0)
}

func (s *LiteralsContext) INTERPRETED_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserINTERPRETED_STRING_LIT, 0)
}

func (s *LiteralsContext) Numbers() INumbersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumbersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumbersContext)
}

func (s *LiteralsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterLiterals(s)
	}
}

func (s *LiteralsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitLiterals(s)
	}
}

func (p *FuncWrapParser) Literals() (localctx ILiteralsContext) {
	localctx = NewLiteralsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FuncWrapParserRULE_literals)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(197)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FuncWrapParserBOOLEAN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Match(FuncWrapParserBOOLEAN)
		}

	case FuncWrapParserRAW_STRING_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.Match(FuncWrapParserRAW_STRING_LIT)
		}

	case FuncWrapParserINTERPRETED_STRING_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(195)
			p.Match(FuncWrapParserINTERPRETED_STRING_LIT)
		}

	case FuncWrapParserDECIMAL_LIT, FuncWrapParserOCTAL_LIT, FuncWrapParserHEX_LIT, FuncWrapParserFLOAT_LIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(196)
			p.Numbers()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumbersContext is an interface to support dynamic dispatch.
type INumbersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumbersContext differentiates from other interfaces.
	IsNumbersContext()
}

type NumbersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumbersContext() *NumbersContext {
	var p = new(NumbersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_numbers
	return p
}

func (*NumbersContext) IsNumbersContext() {}

func NewNumbersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumbersContext {
	var p = new(NumbersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_numbers

	return p
}

func (s *NumbersContext) GetParser() antlr.Parser { return s.parser }

func (s *NumbersContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserDECIMAL_LIT, 0)
}

func (s *NumbersContext) OCTAL_LIT() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserOCTAL_LIT, 0)
}

func (s *NumbersContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserHEX_LIT, 0)
}

func (s *NumbersContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserFLOAT_LIT, 0)
}

func (s *NumbersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumbersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumbersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterNumbers(s)
	}
}

func (s *NumbersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitNumbers(s)
	}
}

func (p *FuncWrapParser) Numbers() (localctx INumbersContext) {
	localctx = NewNumbersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FuncWrapParserRULE_numbers)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FuncWrapParserDECIMAL_LIT)|(1<<FuncWrapParserOCTAL_LIT)|(1<<FuncWrapParserHEX_LIT)|(1<<FuncWrapParserFLOAT_LIT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IObjectDefContext is an interface to support dynamic dispatch.
type IObjectDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectDefContext differentiates from other interfaces.
	IsObjectDefContext()
}

type ObjectDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectDefContext() *ObjectDefContext {
	var p = new(ObjectDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_objectDef
	return p
}

func (*ObjectDefContext) IsObjectDefContext() {}

func NewObjectDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectDefContext {
	var p = new(ObjectDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_objectDef

	return p
}

func (s *ObjectDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectDefContext) OBJ() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserOBJ, 0)
}

func (s *ObjectDefContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ObjectDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterObjectDef(s)
	}
}

func (s *ObjectDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitObjectDef(s)
	}
}

func (p *FuncWrapParser) ObjectDef() (localctx IObjectDefContext) {
	localctx = NewObjectDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FuncWrapParserRULE_objectDef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(FuncWrapParserOBJ)
	}
	{
		p.SetState(202)
		p.Variable()
	}

	return localctx
}

// IObjMemberRefContext is an interface to support dynamic dispatch.
type IObjMemberRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjMemberRefContext differentiates from other interfaces.
	IsObjMemberRefContext()
}

type ObjMemberRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjMemberRefContext() *ObjMemberRefContext {
	var p = new(ObjMemberRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_objMemberRef
	return p
}

func (*ObjMemberRefContext) IsObjMemberRefContext() {}

func NewObjMemberRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjMemberRefContext {
	var p = new(ObjMemberRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_objMemberRef

	return p
}

func (s *ObjMemberRefContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjMemberRefContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ObjMemberRefContext) AllObjElement() []IObjElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjElementContext)(nil)).Elem())
	var tst = make([]IObjElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjElementContext)
		}
	}

	return tst
}

func (s *ObjMemberRefContext) ObjElement(i int) IObjElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjElementContext)
}

func (s *ObjMemberRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjMemberRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjMemberRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterObjMemberRef(s)
	}
}

func (s *ObjMemberRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitObjMemberRef(s)
	}
}

func (p *FuncWrapParser) ObjMemberRef() (localctx IObjMemberRefContext) {
	localctx = NewObjMemberRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FuncWrapParserRULE_objMemberRef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Variable()
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FuncWrapParserDOT {
		{
			p.SetState(205)
			p.ObjElement()
		}

		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IObjElementContext is an interface to support dynamic dispatch.
type IObjElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjElementContext differentiates from other interfaces.
	IsObjElementContext()
}

type ObjElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjElementContext() *ObjElementContext {
	var p = new(ObjElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_objElement
	return p
}

func (*ObjElementContext) IsObjElementContext() {}

func NewObjElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjElementContext {
	var p = new(ObjElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_objElement

	return p
}

func (s *ObjElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjElementContext) DOT() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserDOT, 0)
}

func (s *ObjElementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserIDENTIFIER, 0)
}

func (s *ObjElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterObjElement(s)
	}
}

func (s *ObjElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitObjElement(s)
	}
}

func (p *FuncWrapParser) ObjElement() (localctx IObjElementContext) {
	localctx = NewObjElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FuncWrapParserRULE_objElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(FuncWrapParserDOT)
	}
	{
		p.SetState(211)
		p.Match(FuncWrapParserIDENTIFIER)
	}

	return localctx
}

// IObjAssignContext is an interface to support dynamic dispatch.
type IObjAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjAssignContext differentiates from other interfaces.
	IsObjAssignContext()
}

type ObjAssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjAssignContext() *ObjAssignContext {
	var p = new(ObjAssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuncWrapParserRULE_objAssign
	return p
}

func (*ObjAssignContext) IsObjAssignContext() {}

func NewObjAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjAssignContext {
	var p = new(ObjAssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuncWrapParserRULE_objAssign

	return p
}

func (s *ObjAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjAssignContext) ObjMemberRef() IObjMemberRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjMemberRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjMemberRefContext)
}

func (s *ObjAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FuncWrapParserASSIGN, 0)
}

func (s *ObjAssignContext) AsnVal() IAsnValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsnValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsnValContext)
}

func (s *ObjAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.EnterObjAssign(s)
	}
}

func (s *ObjAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuncWrapParserListener); ok {
		listenerT.ExitObjAssign(s)
	}
}

func (p *FuncWrapParser) ObjAssign() (localctx IObjAssignContext) {
	localctx = NewObjAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, FuncWrapParserRULE_objAssign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.ObjMemberRef()
	}
	{
		p.SetState(214)
		p.Match(FuncWrapParserASSIGN)
	}
	{
		p.SetState(215)
		p.AsnVal()
	}

	return localctx
}

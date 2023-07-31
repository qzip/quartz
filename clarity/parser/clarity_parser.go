// Code generated from Clarity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Clarity

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ClarityParser struct {
	*antlr.BaseParser
}

var clarityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func clarityParserInit() {
	staticData := &clarityParserStaticData
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
		"sexpr", "item", "list_", "touple", "touples", "atom",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 59, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 5, 0, 14, 8, 0, 10, 0, 12, 0, 17, 9, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 31, 8,
		1, 1, 2, 1, 2, 5, 2, 35, 8, 2, 10, 2, 12, 2, 38, 9, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 50, 8, 4, 10, 4, 12, 4,
		53, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 0, 0, 6, 0, 2, 4, 6, 8, 10, 0,
		1, 5, 0, 5, 6, 8, 8, 10, 10, 12, 12, 15, 15, 59, 0, 15, 1, 0, 0, 0, 2,
		30, 1, 0, 0, 0, 4, 32, 1, 0, 0, 0, 6, 41, 1, 0, 0, 0, 8, 45, 1, 0, 0, 0,
		10, 56, 1, 0, 0, 0, 12, 14, 3, 2, 1, 0, 13, 12, 1, 0, 0, 0, 14, 17, 1,
		0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 18, 1, 0, 0, 0, 17,
		15, 1, 0, 0, 0, 18, 19, 5, 0, 0, 1, 19, 1, 1, 0, 0, 0, 20, 31, 3, 10, 5,
		0, 21, 31, 5, 11, 0, 0, 22, 31, 3, 4, 2, 0, 23, 31, 3, 8, 4, 0, 24, 25,
		5, 13, 0, 0, 25, 26, 3, 2, 1, 0, 26, 27, 5, 15, 0, 0, 27, 28, 3, 2, 1,
		0, 28, 29, 5, 14, 0, 0, 29, 31, 1, 0, 0, 0, 30, 20, 1, 0, 0, 0, 30, 21,
		1, 0, 0, 0, 30, 22, 1, 0, 0, 0, 30, 23, 1, 0, 0, 0, 30, 24, 1, 0, 0, 0,
		31, 3, 1, 0, 0, 0, 32, 36, 5, 13, 0, 0, 33, 35, 3, 2, 1, 0, 34, 33, 1,
		0, 0, 0, 35, 38, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37,
		39, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 39, 40, 5, 14, 0, 0, 40, 5, 1, 0, 0,
		0, 41, 42, 3, 10, 5, 0, 42, 43, 5, 1, 0, 0, 43, 44, 3, 2, 1, 0, 44, 7,
		1, 0, 0, 0, 45, 46, 5, 2, 0, 0, 46, 51, 3, 6, 3, 0, 47, 48, 5, 3, 0, 0,
		48, 50, 3, 6, 3, 0, 49, 47, 1, 0, 0, 0, 50, 53, 1, 0, 0, 0, 51, 49, 1,
		0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 54,
		55, 5, 4, 0, 0, 55, 9, 1, 0, 0, 0, 56, 57, 7, 0, 0, 0, 57, 11, 1, 0, 0,
		0, 4, 15, 30, 36, 51,
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

// ClarityParserInit initializes any static state used to implement ClarityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewClarityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClarityParserInit() {
	staticData := &clarityParserStaticData
	staticData.once.Do(clarityParserInit)
}

// NewClarityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewClarityParser(input antlr.TokenStream) *ClarityParser {
	ClarityParserInit()
	this := new(ClarityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &clarityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Clarity.g4"

	return this
}

// ClarityParser tokens.
const (
	ClarityParserEOF          = antlr.TokenEOF
	ClarityParserT__0         = 1
	ClarityParserT__1         = 2
	ClarityParserT__2         = 3
	ClarityParserT__3         = 4
	ClarityParserSTRING       = 5
	ClarityParserUTF8         = 6
	ClarityParserWHITESPACE   = 7
	ClarityParserNUMBER       = 8
	ClarityParserHEX          = 9
	ClarityParserSYMBOL       = 10
	ClarityParserBUILTIN      = 11
	ClarityParserPRINCIPAL    = 12
	ClarityParserLPAREN       = 13
	ClarityParserRPAREN       = 14
	ClarityParserDOT          = 15
	ClarityParserBlockComment = 16
	ClarityParserLineComment  = 17
)

// ClarityParser rules.
const (
	ClarityParserRULE_sexpr   = 0
	ClarityParserRULE_item    = 1
	ClarityParserRULE_list_   = 2
	ClarityParserRULE_touple  = 3
	ClarityParserRULE_touples = 4
	ClarityParserRULE_atom    = 5
)

// ISexprContext is an interface to support dynamic dispatch.
type ISexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSexprContext differentiates from other interfaces.
	IsSexprContext()
}

type SexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySexprContext() *SexprContext {
	var p = new(SexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClarityParserRULE_sexpr
	return p
}

func (*SexprContext) IsSexprContext() {}

func NewSexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SexprContext {
	var p = new(SexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClarityParserRULE_sexpr

	return p
}

func (s *SexprContext) GetParser() antlr.Parser { return s.parser }

func (s *SexprContext) EOF() antlr.TerminalNode {
	return s.GetToken(ClarityParserEOF, 0)
}

func (s *SexprContext) AllItem() []IItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IItemContext); ok {
			len++
		}
	}

	tst := make([]IItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IItemContext); ok {
			tst[i] = t.(IItemContext)
			i++
		}
	}

	return tst
}

func (s *SexprContext) Item(i int) IItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *SexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.EnterSexpr(s)
	}
}

func (s *SexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.ExitSexpr(s)
	}
}

func (p *ClarityParser) Sexpr() (localctx ISexprContext) {
	this := p
	_ = this

	localctx = NewSexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ClarityParserRULE_sexpr)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ClarityParserT__1)|(1<<ClarityParserSTRING)|(1<<ClarityParserUTF8)|(1<<ClarityParserNUMBER)|(1<<ClarityParserSYMBOL)|(1<<ClarityParserBUILTIN)|(1<<ClarityParserPRINCIPAL)|(1<<ClarityParserLPAREN)|(1<<ClarityParserDOT))) != 0 {
		{
			p.SetState(12)
			p.Item()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(18)
		p.Match(ClarityParserEOF)
	}

	return localctx
}

// IItemContext is an interface to support dynamic dispatch.
type IItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItemContext differentiates from other interfaces.
	IsItemContext()
}

type ItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItemContext() *ItemContext {
	var p = new(ItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClarityParserRULE_item
	return p
}

func (*ItemContext) IsItemContext() {}

func NewItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemContext {
	var p = new(ItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClarityParserRULE_item

	return p
}

func (s *ItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ItemContext) BUILTIN() antlr.TerminalNode {
	return s.GetToken(ClarityParserBUILTIN, 0)
}

func (s *ItemContext) List_() IList_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_Context)
}

func (s *ItemContext) Touples() ITouplesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITouplesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITouplesContext)
}

func (s *ItemContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ClarityParserLPAREN, 0)
}

func (s *ItemContext) AllItem() []IItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IItemContext); ok {
			len++
		}
	}

	tst := make([]IItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IItemContext); ok {
			tst[i] = t.(IItemContext)
			i++
		}
	}

	return tst
}

func (s *ItemContext) Item(i int) IItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *ItemContext) DOT() antlr.TerminalNode {
	return s.GetToken(ClarityParserDOT, 0)
}

func (s *ItemContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ClarityParserRPAREN, 0)
}

func (s *ItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.EnterItem(s)
	}
}

func (s *ItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.ExitItem(s)
	}
}

func (p *ClarityParser) Item() (localctx IItemContext) {
	this := p
	_ = this

	localctx = NewItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ClarityParserRULE_item)

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

	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Atom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.Match(ClarityParserBUILTIN)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.List_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(23)
			p.Touples()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(24)
			p.Match(ClarityParserLPAREN)
		}
		{
			p.SetState(25)
			p.Item()
		}
		{
			p.SetState(26)
			p.Match(ClarityParserDOT)
		}
		{
			p.SetState(27)
			p.Item()
		}
		{
			p.SetState(28)
			p.Match(ClarityParserRPAREN)
		}

	}

	return localctx
}

// IList_Context is an interface to support dynamic dispatch.
type IList_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_Context differentiates from other interfaces.
	IsList_Context()
}

type List_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_Context() *List_Context {
	var p = new(List_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClarityParserRULE_list_
	return p
}

func (*List_Context) IsList_Context() {}

func NewList_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_Context {
	var p = new(List_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClarityParserRULE_list_

	return p
}

func (s *List_Context) GetParser() antlr.Parser { return s.parser }

func (s *List_Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(ClarityParserLPAREN, 0)
}

func (s *List_Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(ClarityParserRPAREN, 0)
}

func (s *List_Context) AllItem() []IItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IItemContext); ok {
			len++
		}
	}

	tst := make([]IItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IItemContext); ok {
			tst[i] = t.(IItemContext)
			i++
		}
	}

	return tst
}

func (s *List_Context) Item(i int) IItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *List_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.EnterList_(s)
	}
}

func (s *List_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.ExitList_(s)
	}
}

func (p *ClarityParser) List_() (localctx IList_Context) {
	this := p
	_ = this

	localctx = NewList_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ClarityParserRULE_list_)
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
		p.SetState(32)
		p.Match(ClarityParserLPAREN)
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ClarityParserT__1)|(1<<ClarityParserSTRING)|(1<<ClarityParserUTF8)|(1<<ClarityParserNUMBER)|(1<<ClarityParserSYMBOL)|(1<<ClarityParserBUILTIN)|(1<<ClarityParserPRINCIPAL)|(1<<ClarityParserLPAREN)|(1<<ClarityParserDOT))) != 0 {
		{
			p.SetState(33)
			p.Item()
		}

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(39)
		p.Match(ClarityParserRPAREN)
	}

	return localctx
}

// IToupleContext is an interface to support dynamic dispatch.
type IToupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsToupleContext differentiates from other interfaces.
	IsToupleContext()
}

type ToupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyToupleContext() *ToupleContext {
	var p = new(ToupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClarityParserRULE_touple
	return p
}

func (*ToupleContext) IsToupleContext() {}

func NewToupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ToupleContext {
	var p = new(ToupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClarityParserRULE_touple

	return p
}

func (s *ToupleContext) GetParser() antlr.Parser { return s.parser }

func (s *ToupleContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ToupleContext) Item() IItemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *ToupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ToupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.EnterTouple(s)
	}
}

func (s *ToupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.ExitTouple(s)
	}
}

func (p *ClarityParser) Touple() (localctx IToupleContext) {
	this := p
	_ = this

	localctx = NewToupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ClarityParserRULE_touple)

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
		p.SetState(41)
		p.Atom()
	}
	{
		p.SetState(42)
		p.Match(ClarityParserT__0)
	}
	{
		p.SetState(43)
		p.Item()
	}

	return localctx
}

// ITouplesContext is an interface to support dynamic dispatch.
type ITouplesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTouplesContext differentiates from other interfaces.
	IsTouplesContext()
}

type TouplesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTouplesContext() *TouplesContext {
	var p = new(TouplesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClarityParserRULE_touples
	return p
}

func (*TouplesContext) IsTouplesContext() {}

func NewTouplesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TouplesContext {
	var p = new(TouplesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClarityParserRULE_touples

	return p
}

func (s *TouplesContext) GetParser() antlr.Parser { return s.parser }

func (s *TouplesContext) AllTouple() []IToupleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IToupleContext); ok {
			len++
		}
	}

	tst := make([]IToupleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IToupleContext); ok {
			tst[i] = t.(IToupleContext)
			i++
		}
	}

	return tst
}

func (s *TouplesContext) Touple(i int) IToupleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IToupleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IToupleContext)
}

func (s *TouplesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TouplesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TouplesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.EnterTouples(s)
	}
}

func (s *TouplesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.ExitTouples(s)
	}
}

func (p *ClarityParser) Touples() (localctx ITouplesContext) {
	this := p
	_ = this

	localctx = NewTouplesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ClarityParserRULE_touples)
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
		p.SetState(45)
		p.Match(ClarityParserT__1)
	}
	{
		p.SetState(46)
		p.Touple()
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ClarityParserT__2 {
		{
			p.SetState(47)
			p.Match(ClarityParserT__2)
		}
		{
			p.SetState(48)
			p.Touple()
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
		p.Match(ClarityParserT__3)
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClarityParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClarityParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(ClarityParserSTRING, 0)
}

func (s *AtomContext) UTF8() antlr.TerminalNode {
	return s.GetToken(ClarityParserUTF8, 0)
}

func (s *AtomContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(ClarityParserSYMBOL, 0)
}

func (s *AtomContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ClarityParserNUMBER, 0)
}

func (s *AtomContext) DOT() antlr.TerminalNode {
	return s.GetToken(ClarityParserDOT, 0)
}

func (s *AtomContext) PRINCIPAL() antlr.TerminalNode {
	return s.GetToken(ClarityParserPRINCIPAL, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *ClarityParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ClarityParserRULE_atom)
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
		p.SetState(56)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ClarityParserSTRING)|(1<<ClarityParserUTF8)|(1<<ClarityParserNUMBER)|(1<<ClarityParserSYMBOL)|(1<<ClarityParserPRINCIPAL)|(1<<ClarityParserDOT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

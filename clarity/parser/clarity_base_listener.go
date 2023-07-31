// Code generated from Clarity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Clarity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseClarityListener is a complete listener for a parse tree produced by ClarityParser.
type BaseClarityListener struct{}

var _ ClarityListener = &BaseClarityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseClarityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseClarityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseClarityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseClarityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSexpr is called when production sexpr is entered.
func (s *BaseClarityListener) EnterSexpr(ctx *SexprContext) {}

// ExitSexpr is called when production sexpr is exited.
func (s *BaseClarityListener) ExitSexpr(ctx *SexprContext) {}

// EnterItem is called when production item is entered.
func (s *BaseClarityListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BaseClarityListener) ExitItem(ctx *ItemContext) {}

// EnterList_ is called when production list_ is entered.
func (s *BaseClarityListener) EnterList_(ctx *List_Context) {}

// ExitList_ is called when production list_ is exited.
func (s *BaseClarityListener) ExitList_(ctx *List_Context) {}

// EnterTouple is called when production touple is entered.
func (s *BaseClarityListener) EnterTouple(ctx *ToupleContext) {}

// ExitTouple is called when production touple is exited.
func (s *BaseClarityListener) ExitTouple(ctx *ToupleContext) {}

// EnterTouples is called when production touples is entered.
func (s *BaseClarityListener) EnterTouples(ctx *TouplesContext) {}

// ExitTouples is called when production touples is exited.
func (s *BaseClarityListener) ExitTouples(ctx *TouplesContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseClarityListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseClarityListener) ExitAtom(ctx *AtomContext) {}

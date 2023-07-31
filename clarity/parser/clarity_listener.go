// Code generated from Clarity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Clarity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ClarityListener is a complete listener for a parse tree produced by ClarityParser.
type ClarityListener interface {
	antlr.ParseTreeListener

	// EnterSexpr is called when entering the sexpr production.
	EnterSexpr(c *SexprContext)

	// EnterItem is called when entering the item production.
	EnterItem(c *ItemContext)

	// EnterList_ is called when entering the list_ production.
	EnterList_(c *List_Context)

	// EnterTouple is called when entering the touple production.
	EnterTouple(c *ToupleContext)

	// EnterTouples is called when entering the touples production.
	EnterTouples(c *TouplesContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// ExitSexpr is called when exiting the sexpr production.
	ExitSexpr(c *SexprContext)

	// ExitItem is called when exiting the item production.
	ExitItem(c *ItemContext)

	// ExitList_ is called when exiting the list_ production.
	ExitList_(c *List_Context)

	// ExitTouple is called when exiting the touple production.
	ExitTouple(c *ToupleContext)

	// ExitTouples is called when exiting the touples production.
	ExitTouples(c *TouplesContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)
}

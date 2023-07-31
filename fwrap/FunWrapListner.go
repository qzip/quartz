/*
Copyright (c) 2019-20, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>


*/

package fwrap

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FuncWrapListener is a complete listener for a parse tree produced by FuncWrapParser.
type FuncWrapListener struct {
}

// VisitTerminal is called when a terminal node is visited.
func (s *FuncWrapListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *FuncWrapListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *FuncWrapListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *FuncWrapListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *FuncWrapListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *FuncWrapListener) ExitProg(ctx *ProgContext) {}

// EnterPstatement is called when production pstatement is entered.
func (s *FuncWrapListener) EnterPstatement(ctx *PstatementContext) {}

// ExitPstatement is called when production pstatement is exited.
func (s *FuncWrapListener) ExitPstatement(ctx *PstatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *FuncWrapListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *FuncWrapListener) ExitStatement(ctx *StatementContext) {}

// EnterVarAssign is called when production varAssign is entered.
func (s *FuncWrapListener) EnterVarAssign(ctx *VarAssignContext) {}

// ExitVarAssign is called when production varAssign is exited.
func (s *FuncWrapListener) ExitVarAssign(ctx *VarAssignContext) {}

// EnterFunCall is called when production funCall is entered.
func (s *FuncWrapListener) EnterFunCall(ctx *FunCallContext) {}

// ExitFunCall is called when production funCall is exited.
func (s *FuncWrapListener) ExitFunCall(ctx *FunCallContext) {}

// EnterFuncName is called when production funcName is entered.
func (s *FuncWrapListener) EnterFuncName(ctx *FuncNameContext) {}

// ExitFuncName is called when production funcName is exited.
func (s *FuncWrapListener) ExitFuncName(ctx *FuncNameContext) {}

// EnterFunDef is called when production funDef is entered.
func (s *FuncWrapListener) EnterFunDef(ctx *FunDefContext) {}

// ExitFunDef is called when production funDef is exited.
func (s *FuncWrapListener) ExitFunDef(ctx *FunDefContext) {}

// EnterRetStmt is called when production retStmt is entered.
func (s *FuncWrapListener) EnterRetStmt(ctx *RetStmtContext) {}

// ExitRetStmt is called when production retStmt is exited.
func (s *FuncWrapListener) ExitRetStmt(ctx *RetStmtContext) {}

// EnterCond is called when production cond is entered.
func (s *FuncWrapListener) EnterCond(ctx *CondContext) {}

// ExitCond is called when production cond is exited.
func (s *FuncWrapListener) ExitCond(ctx *CondContext) {}

// EnterEval is called when production eval is entered.
func (s *FuncWrapListener) EnterEval(ctx *EvalContext) {}

// ExitEval is called when production eval is exited.
func (s *FuncWrapListener) ExitEval(ctx *EvalContext) {}

// EnterRVal is called when production rVal is entered.
func (s *FuncWrapListener) EnterRVal(ctx *RValContext) {}

// ExitRVal is called when production rVal is exited.
func (s *FuncWrapListener) ExitRVal(ctx *RValContext) {}

// EnterRVals is called when production rVals is entered.
func (s *FuncWrapListener) EnterRVals(ctx *RValsContext) {}

// ExitRVals is called when production rVals is exited.
func (s *FuncWrapListener) ExitRVals(ctx *RValsContext) {}

// EnterVariable is called when production variable is entered.
func (s *FuncWrapListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *FuncWrapListener) ExitVariable(ctx *VariableContext) {}

// EnterArray is called when production array is entered.
func (s *FuncWrapListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *FuncWrapListener) ExitArray(ctx *ArrayContext) {}

// EnterOp is called when production op is entered.
func (s *FuncWrapListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *FuncWrapListener) ExitOp(ctx *OpContext) {}

// EnterLiterals is called when production literals is entered.
func (s *FuncWrapListener) EnterLiterals(ctx *LiteralsContext) {}

// ExitLiterals is called when production literals is exited.
func (s *FuncWrapListener) ExitLiterals(ctx *LiteralsContext) {}

// EnterNumbers is called when production numbers is entered.
func (s *FuncWrapListener) EnterNumbers(ctx *NumbersContext) {}

// ExitNumbers is called when production numbers is exited.
func (s *FuncWrapListener) ExitNumbers(ctx *NumbersContext) {}

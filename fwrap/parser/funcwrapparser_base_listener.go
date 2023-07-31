// Code generated from FuncWrapParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // FuncWrapParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFuncWrapParserListener is a complete listener for a parse tree produced by FuncWrapParser.
type BaseFuncWrapParserListener struct{}

var _ FuncWrapParserListener = &BaseFuncWrapParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFuncWrapParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFuncWrapParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFuncWrapParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFuncWrapParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseFuncWrapParserListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseFuncWrapParserListener) ExitProg(ctx *ProgContext) {}

// EnterPstatement is called when production pstatement is entered.
func (s *BaseFuncWrapParserListener) EnterPstatement(ctx *PstatementContext) {}

// ExitPstatement is called when production pstatement is exited.
func (s *BaseFuncWrapParserListener) ExitPstatement(ctx *PstatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseFuncWrapParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseFuncWrapParserListener) ExitStatement(ctx *StatementContext) {}

// EnterVarAssign is called when production varAssign is entered.
func (s *BaseFuncWrapParserListener) EnterVarAssign(ctx *VarAssignContext) {}

// ExitVarAssign is called when production varAssign is exited.
func (s *BaseFuncWrapParserListener) ExitVarAssign(ctx *VarAssignContext) {}

// EnterAsnVal is called when production asnVal is entered.
func (s *BaseFuncWrapParserListener) EnterAsnVal(ctx *AsnValContext) {}

// ExitAsnVal is called when production asnVal is exited.
func (s *BaseFuncWrapParserListener) ExitAsnVal(ctx *AsnValContext) {}

// EnterFunCall is called when production funCall is entered.
func (s *BaseFuncWrapParserListener) EnterFunCall(ctx *FunCallContext) {}

// ExitFunCall is called when production funCall is exited.
func (s *BaseFuncWrapParserListener) ExitFunCall(ctx *FunCallContext) {}

// EnterFunValName is called when production funValName is entered.
func (s *BaseFuncWrapParserListener) EnterFunValName(ctx *FunValNameContext) {}

// ExitFunValName is called when production funValName is exited.
func (s *BaseFuncWrapParserListener) ExitFunValName(ctx *FunValNameContext) {}

// EnterFuncName is called when production funcName is entered.
func (s *BaseFuncWrapParserListener) EnterFuncName(ctx *FuncNameContext) {}

// ExitFuncName is called when production funcName is exited.
func (s *BaseFuncWrapParserListener) ExitFuncName(ctx *FuncNameContext) {}

// EnterFunDef is called when production funDef is entered.
func (s *BaseFuncWrapParserListener) EnterFunDef(ctx *FunDefContext) {}

// ExitFunDef is called when production funDef is exited.
func (s *BaseFuncWrapParserListener) ExitFunDef(ctx *FunDefContext) {}

// EnterRetStmt is called when production retStmt is entered.
func (s *BaseFuncWrapParserListener) EnterRetStmt(ctx *RetStmtContext) {}

// ExitRetStmt is called when production retStmt is exited.
func (s *BaseFuncWrapParserListener) ExitRetStmt(ctx *RetStmtContext) {}

// EnterFVal is called when production fVal is entered.
func (s *BaseFuncWrapParserListener) EnterFVal(ctx *FValContext) {}

// ExitFVal is called when production fVal is exited.
func (s *BaseFuncWrapParserListener) ExitFVal(ctx *FValContext) {}

// EnterRetVal is called when production retVal is entered.
func (s *BaseFuncWrapParserListener) EnterRetVal(ctx *RetValContext) {}

// ExitRetVal is called when production retVal is exited.
func (s *BaseFuncWrapParserListener) ExitRetVal(ctx *RetValContext) {}

// EnterRVal is called when production rVal is entered.
func (s *BaseFuncWrapParserListener) EnterRVal(ctx *RValContext) {}

// ExitRVal is called when production rVal is exited.
func (s *BaseFuncWrapParserListener) ExitRVal(ctx *RValContext) {}

// EnterRVals is called when production rVals is entered.
func (s *BaseFuncWrapParserListener) EnterRVals(ctx *RValsContext) {}

// ExitRVals is called when production rVals is exited.
func (s *BaseFuncWrapParserListener) ExitRVals(ctx *RValsContext) {}

// EnterCond is called when production cond is entered.
func (s *BaseFuncWrapParserListener) EnterCond(ctx *CondContext) {}

// ExitCond is called when production cond is exited.
func (s *BaseFuncWrapParserListener) ExitCond(ctx *CondContext) {}

// EnterEval is called when production eval is entered.
func (s *BaseFuncWrapParserListener) EnterEval(ctx *EvalContext) {}

// ExitEval is called when production eval is exited.
func (s *BaseFuncWrapParserListener) ExitEval(ctx *EvalContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseFuncWrapParserListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseFuncWrapParserListener) ExitVariable(ctx *VariableContext) {}

// EnterArray is called when production array is entered.
func (s *BaseFuncWrapParserListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseFuncWrapParserListener) ExitArray(ctx *ArrayContext) {}

// EnterArrRef is called when production arrRef is entered.
func (s *BaseFuncWrapParserListener) EnterArrRef(ctx *ArrRefContext) {}

// ExitArrRef is called when production arrRef is exited.
func (s *BaseFuncWrapParserListener) ExitArrRef(ctx *ArrRefContext) {}

// EnterOp is called when production op is entered.
func (s *BaseFuncWrapParserListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseFuncWrapParserListener) ExitOp(ctx *OpContext) {}

// EnterLiterals is called when production literals is entered.
func (s *BaseFuncWrapParserListener) EnterLiterals(ctx *LiteralsContext) {}

// ExitLiterals is called when production literals is exited.
func (s *BaseFuncWrapParserListener) ExitLiterals(ctx *LiteralsContext) {}

// EnterNumbers is called when production numbers is entered.
func (s *BaseFuncWrapParserListener) EnterNumbers(ctx *NumbersContext) {}

// ExitNumbers is called when production numbers is exited.
func (s *BaseFuncWrapParserListener) ExitNumbers(ctx *NumbersContext) {}

// EnterObjectDef is called when production objectDef is entered.
func (s *BaseFuncWrapParserListener) EnterObjectDef(ctx *ObjectDefContext) {}

// ExitObjectDef is called when production objectDef is exited.
func (s *BaseFuncWrapParserListener) ExitObjectDef(ctx *ObjectDefContext) {}

// EnterObjMemberRef is called when production objMemberRef is entered.
func (s *BaseFuncWrapParserListener) EnterObjMemberRef(ctx *ObjMemberRefContext) {}

// ExitObjMemberRef is called when production objMemberRef is exited.
func (s *BaseFuncWrapParserListener) ExitObjMemberRef(ctx *ObjMemberRefContext) {}

// EnterObjElement is called when production objElement is entered.
func (s *BaseFuncWrapParserListener) EnterObjElement(ctx *ObjElementContext) {}

// ExitObjElement is called when production objElement is exited.
func (s *BaseFuncWrapParserListener) ExitObjElement(ctx *ObjElementContext) {}

// EnterObjAssign is called when production objAssign is entered.
func (s *BaseFuncWrapParserListener) EnterObjAssign(ctx *ObjAssignContext) {}

// ExitObjAssign is called when production objAssign is exited.
func (s *BaseFuncWrapParserListener) ExitObjAssign(ctx *ObjAssignContext) {}

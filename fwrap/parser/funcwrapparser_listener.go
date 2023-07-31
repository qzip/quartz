// Code generated from FuncWrapParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // FuncWrapParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FuncWrapParserListener is a complete listener for a parse tree produced by FuncWrapParser.
type FuncWrapParserListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterPstatement is called when entering the pstatement production.
	EnterPstatement(c *PstatementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVarAssign is called when entering the varAssign production.
	EnterVarAssign(c *VarAssignContext)

	// EnterAsnVal is called when entering the asnVal production.
	EnterAsnVal(c *AsnValContext)

	// EnterFunCall is called when entering the funCall production.
	EnterFunCall(c *FunCallContext)

	// EnterFunValName is called when entering the funValName production.
	EnterFunValName(c *FunValNameContext)

	// EnterFuncName is called when entering the funcName production.
	EnterFuncName(c *FuncNameContext)

	// EnterFunDef is called when entering the funDef production.
	EnterFunDef(c *FunDefContext)

	// EnterRetStmt is called when entering the retStmt production.
	EnterRetStmt(c *RetStmtContext)

	// EnterFVal is called when entering the fVal production.
	EnterFVal(c *FValContext)

	// EnterRetVal is called when entering the retVal production.
	EnterRetVal(c *RetValContext)

	// EnterRVal is called when entering the rVal production.
	EnterRVal(c *RValContext)

	// EnterRVals is called when entering the rVals production.
	EnterRVals(c *RValsContext)

	// EnterCond is called when entering the cond production.
	EnterCond(c *CondContext)

	// EnterEval is called when entering the eval production.
	EnterEval(c *EvalContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterArrRef is called when entering the arrRef production.
	EnterArrRef(c *ArrRefContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// EnterLiterals is called when entering the literals production.
	EnterLiterals(c *LiteralsContext)

	// EnterNumbers is called when entering the numbers production.
	EnterNumbers(c *NumbersContext)

	// EnterObjectDef is called when entering the objectDef production.
	EnterObjectDef(c *ObjectDefContext)

	// EnterObjMemberRef is called when entering the objMemberRef production.
	EnterObjMemberRef(c *ObjMemberRefContext)

	// EnterObjElement is called when entering the objElement production.
	EnterObjElement(c *ObjElementContext)

	// EnterObjAssign is called when entering the objAssign production.
	EnterObjAssign(c *ObjAssignContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitPstatement is called when exiting the pstatement production.
	ExitPstatement(c *PstatementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVarAssign is called when exiting the varAssign production.
	ExitVarAssign(c *VarAssignContext)

	// ExitAsnVal is called when exiting the asnVal production.
	ExitAsnVal(c *AsnValContext)

	// ExitFunCall is called when exiting the funCall production.
	ExitFunCall(c *FunCallContext)

	// ExitFunValName is called when exiting the funValName production.
	ExitFunValName(c *FunValNameContext)

	// ExitFuncName is called when exiting the funcName production.
	ExitFuncName(c *FuncNameContext)

	// ExitFunDef is called when exiting the funDef production.
	ExitFunDef(c *FunDefContext)

	// ExitRetStmt is called when exiting the retStmt production.
	ExitRetStmt(c *RetStmtContext)

	// ExitFVal is called when exiting the fVal production.
	ExitFVal(c *FValContext)

	// ExitRetVal is called when exiting the retVal production.
	ExitRetVal(c *RetValContext)

	// ExitRVal is called when exiting the rVal production.
	ExitRVal(c *RValContext)

	// ExitRVals is called when exiting the rVals production.
	ExitRVals(c *RValsContext)

	// ExitCond is called when exiting the cond production.
	ExitCond(c *CondContext)

	// ExitEval is called when exiting the eval production.
	ExitEval(c *EvalContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitArrRef is called when exiting the arrRef production.
	ExitArrRef(c *ArrRefContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)

	// ExitLiterals is called when exiting the literals production.
	ExitLiterals(c *LiteralsContext)

	// ExitNumbers is called when exiting the numbers production.
	ExitNumbers(c *NumbersContext)

	// ExitObjectDef is called when exiting the objectDef production.
	ExitObjectDef(c *ObjectDefContext)

	// ExitObjMemberRef is called when exiting the objMemberRef production.
	ExitObjMemberRef(c *ObjMemberRefContext)

	// ExitObjElement is called when exiting the objElement production.
	ExitObjElement(c *ObjElementContext)

	// ExitObjAssign is called when exiting the objAssign production.
	ExitObjAssign(c *ObjAssignContext)
}

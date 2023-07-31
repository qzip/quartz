/*
Copyright (c) 2019, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>

version: 28-june-20: added object.
*/

/*
fwrap defines a minimlist DSL (Domain Specific Language) to cater
to blckchain API calls & smart contract.
It is just a builtin function wrapper variable (& array) declaration
and mechanism for buildin function calls.
*/

//ref: https://github.com/antlr/grammars-v4/blob/master/graphql/GraphQL.g4

parser grammar FuncWrapParser;

options {
	tokenVocab = FuncWrapLexer;
}

prog: pstatement+;

pstatement: funDef | statement;
statement: variable | varAssign | funCall | cond | objectDef | objAssign;

varAssign: variable ASSIGN asnVal ;

asnVal: rVal| funcName;

funCall: funValName L_PAREN rVals? R_PAREN;
funValName: funcName| variable;

funcName: IDENTIFIER;
// Function name cannot clash with internal built-in functions
funDef:
	FUNC funcName L_PAREN rVals? R_PAREN fVal? 
             L_CURLY statement* retStmt? R_CURLY;

retStmt: RETURN retVal?;
fVal: literals | funcName;
retVal: rVal| funcName;
rVal: variable| array | funCall | literals | objMemberRef;
rVals: rVal ( COMMA rVals)*;

cond:
	IF L_PAREN eval R_PAREN L_CURLY statement* R_CURLY (
		ELSE L_CURLY statement* R_CURLY
	)?;
//eval: rVal op rVal;
// eval must evaluate to boolean 
eval : rVal ( op eval)* | L_PAREN eval R_PAREN;


// variable can be of type Array, Object or Litrals
variable: VAR IDENTIFIER;

array: arrRef? L_BRACKET rVals R_BRACKET;
arrRef: variable | objMemberRef;

op:
	LOGICAL_OR
	| LOGICAL_AND
	| EQUALS
	| NOT_EQUALS
	| LESS
	| LESS_OR_EQUALS
	| GREATER
	| GREATER_OR_EQUALS;

literals:
	BOOLEAN
	| RAW_STRING_LIT
	| INTERPRETED_STRING_LIT
	| numbers;

numbers: DECIMAL_LIT | OCTAL_LIT | HEX_LIT | FLOAT_LIT;

/**** Object defs ****/
objectDef: OBJ variable;
objMemberRef: variable objElement+;
objElement : DOT IDENTIFIER ;
objAssign: objMemberRef ASSIGN asnVal;

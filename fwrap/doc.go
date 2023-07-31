/*
Copyright (c) 2019-20, QzIP Blockchain Technology LLP

All Rights Reserved.

Author: Ashish Banerjee, <ashish@qzip.in>

example:
https://github.com/antlr/antlr4/blob/master/doc/go-target.md
*/

/*
Package fwrap defines a minimlist DSL (Domain Specific Language) to cater
to blckchain API calls & smart contract.
It is just a builtin function wrapper variable (& array) declaration
and mechanism for buildin function calls.

Generate : antlr4 -Dlanguage=Go FuncWrapParser.g4 FuncWrapLexer.g4

*/
package fwrap

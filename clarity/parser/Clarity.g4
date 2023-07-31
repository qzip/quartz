/*
The MIT License

Copyright (c) 2008 Robert Stehwien

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

*/

/*
Port to Antlr4 by Tom Everett
https://raw.githubusercontent.com/antlr/grammars-v4/master/sexpression/sexpression.g4

Modified by Ashish Banerjee 28-may-2022

https://github.com/clarity-lang/reference/blob/master/reference.md


*/
grammar Clarity;

sexpr
   : item* EOF
   ;

item
   : atom
   | BUILTIN
   | list_
   | touples
   | LPAREN item DOT item RPAREN
   ;



list_
   : LPAREN item* RPAREN
   ;

touple 
   : atom ':' item
   ;

touples
   : '{' touple (',' touple)* '}'
   ;

atom
   : STRING
   | UTF8
   | SYMBOL
   | NUMBER
   | DOT
   | PRINCIPAL
   
   ;

STRING
   : '"' ('\\' . | ~ ('\\' | '"'))* '"'
   ;

UTF8 
   : 'u' STRING
   ;

WHITESPACE
   : (' ' | '\n' | '\t' | '\r')+ -> channel (HIDDEN); // skip
   

NUMBER
   : ('+' | '-')? (DIGIT)+ ('.' (DIGIT)+)?
   | HEX
   ;

HEX
  : '0' 'x'(DIGIT | ('a' .. 'f') | ('A' .. 'F'))?

  ;

SYMBOL
   : SYMBOL_START (SYMBOL_START | DIGIT)*
   ;

BUILTIN
   : SYMBOL ('!' | '?')
   
   ;

PRINCIPAL 
   : '\'' SYMBOL ;

LPAREN
   : '('
   ;

RPAREN
   : ')'
   ;

DOT
   : '.'
   ;

fragment SYMBOL_START
   : ('a' .. 'z')
   | ('A' .. 'Z')
   | '+'
   | '-'
   | '*'
   | '/'
   | '.'
   | '_'
   ;

fragment DIGIT
   : ('0' .. '9')
   ;

BlockComment
    :   '/*' .*? '*/'
        -> channel (HIDDEN);

LineComment
    :   ';;' ~[\r\n]*
        -> channel (HIDDEN); // skip

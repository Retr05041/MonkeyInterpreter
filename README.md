# MonkeyInterpreter
Interperter for a fictional "Monkey" language

The code in this repo follows a book called: "Writing an Interpreter in Go" by Thorsten Ball

Current Section: 2.5


# Steps for an Interpreter
Tokenize - Every char that comes in needs to be tokenized
Lexer - Refine the tokens to specific keywords, allowing words, double operators, single oerators, etc.
Parser & AST - Take the bigger tokens (Identifiers, statements, etc.) and make an AST with them 
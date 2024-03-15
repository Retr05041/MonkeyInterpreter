package parser

import (
	"MonkeyInterpreter/ast"
	"MonkeyInterpreter/lexer"
	"MonkeyInterpreter/token"
	"fmt"
)

type Parser struct {
	l *lexer.Lexer // Instance of the lexer

	// like position and readPosition but for tokens instead of characters
	curToken  token.Token
	peekToken token.Token

	// Error handleing
	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set at the beginning
	p.nextToken()
	p.nextToken()

	return p
}

// Returns errors of current program
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// So far empty
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// Slide down the tokens from the lexer
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Constructs a *ast.LetStatement node with the token its currently on
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// Checks to make ure the next token is an Identifier, i.e. foobar
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// Then makes an Identifier node
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Then it expects an '='
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Checks if the current token is t
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// Checks if next token is t
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// Checks if the next token is t then moves to it if its true - enforces correctness of structure
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

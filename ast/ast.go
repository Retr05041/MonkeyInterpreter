package ast

import (
	"MonkeyInterpreter/token"
	"bytes"
)

// Node - a spot in the tree
// Statment - don't produce values
// Expression - produce values
// e.x. let <identifier> = <expression>;

// Every node in our AST needs to implement a Node
// Has to provide a TokenLiteral that returns the literal value of the token its associated with
// TokenLiteral is only for debugging
// === NODE TYPES ===
type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// === IDENTIFIER + STATEMENT NODES ===
// struct for the given identifier
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      // the var name
}

// Satisfies the Statement and Node interfaces
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// struct for building valid LET statements
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // variable name
	Value Expression  // The expression following the '='
}

func (i *LetStatement) statementNode()       {}
func (i *LetStatement) TokenLiteral() string { return i.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// === PROGRAM STRUCT ===
// Root Node for every AST
// Every valid Monkey program is just a series of statments
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

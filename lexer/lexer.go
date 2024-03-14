package lexer

import "MonkeyInterpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char (ch))
	readPosition int  // current reading position in input (after current char (ch index +1))
	ch           byte // current char under examination from input string
}

// Create a new Lexer struct and return a pointer to it
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // initialise l.ch, l.position, and l.readPosition
	return l
}

// Lexer pointer function - give us next char and advance our position in the input string
func (l *Lexer) readChar() {
	// EOI check
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 = ASCII Code "NUL"
	} else {
		l.ch = l.input[l.readPosition] // sets char to be the next position from our current one
	}
	l.position = l.readPosition // set our position index to be the next as well
	l.readPosition += 1         // now move our readPos up one as well
}

//
func (l *Lexer) NextToken() token.Token {
	// Create an uninitialised token
	var tok token.Token

	// Skip whitespaces
	l.skipWhitespace()

	// Depending on what the current ch from the lexers input, create a new token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		// If token is NUL (remember ascii code form readChar), establish EOF
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// If not a symbole, we need to check if it is a idenitifier, keyword, or number
		// if it is, read the entire thing and build our token from it
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			// If it's outside of ASCII (No unicode support)
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// Move the lexer down one and return the new token so we can call this function again
	l.readChar()
	return tok
}

// While the char is a number, continue reading until it is not then return it, like readIdentifier
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Is the character a number
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// If the char is a whitespace of any kind, move down one (ignore it)
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Reads in an identifier and advances our lexers positions until it encounters a non-letter-character
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]

}

// Check if a char is a-zA-Z or _, so we can have something like: foo_bar
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Create and return a token based off the token type and the character
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

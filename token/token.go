package token

// type: string - ease of tokenization and ease of debugging - better perfomance would be int or byte
type TokenType string

type Token struct {
	Type    TokenType // type of token
	Literal string    // the token itself
}

const (
	ILLEGAL = "ILLEGAL" // tokens we don't know about
	EOF     = "EOF"     // stop at end of file

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Table for mapping strings to token types
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// takes a string to test and assigns it to a tokentype or returns a uesr-defined identifier
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

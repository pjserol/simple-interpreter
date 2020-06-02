package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiels + literals

	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"

	// Operators

	ASSIGN = "="
	PLUS   = "+"

	// Delimiters

	COMMA     = "."
	SEMICOLON = ":"
	LPARENT   = "("
	RPARENT   = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

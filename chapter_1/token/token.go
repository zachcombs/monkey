package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special Types
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers and Literals
	IDENT  = "IDENT"
	INT    = "INT"
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

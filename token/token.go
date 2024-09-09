package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

// Determining all possible token types

const(
	ILLEGAL="ILLEGAL"
	EOF = "EOF"

	// Identifier, Literal
	IDNET = "IDENT"
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"

)
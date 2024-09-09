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
	IDENT = "IDENT"
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

// Defines the correct Token type for keywords
var keywords = map[string]TokenType{
	"fn":FUNCTION,
	"let":LET,
}

// Checks if string detected is a keyword and returns the type of token
func LookupIdent(ident string) TokenType{
	if tok,ok := keywords[ident];ok{ // The ok variable holds bool value which signinfies if element is there in map or not
		return tok // return the value of the key
	}
	return IDENT // if it doesnt exists in keyword map, then it is of IDENT type


}
package lexer

import "github.com/tren03/go-interpreter/token"

type Lexer struct {
	input       string
	position    int  // Current input postition
	readPostion int  // Current reading position, next char after input
	ch          byte // current char under examination
}

// Initialize all fields of the lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Read next char of the input, and Handle EOF, returns nothing
func (l *Lexer) readChar() {
	if l.readPostion >= len(l.input) {
		l.ch = 0 // ASCII code for NIL
	} else {
		l.ch = l.input[l.readPostion]
	}
	l.position = l.readPostion
	l.readPostion += 1
}

// Returns the token struct of the parsed char
func (l *Lexer) NextToken() token.Token {

	var tok token.Token
	l.skipWhiteSpace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)} // we do this instead of calling new token since '==' is a string and cant be converted to byte
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
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
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isletter(l.ch) {
			tok.Literal = l.readIdentifier()          // read the entire string
			tok.Type = token.LookupIdent(tok.Literal) // returns IDENT if its a literal, else the type of keyword
			return tok                                // return it here instead of after the switch, since the above two move the readpostition and postition, we do not need to call l.readChar() after switch
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// Used to read the full content of a string, if we encounter a letter, to determine the full lenght of the identifier
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isletter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position] // Slice the input the only return the string

}

// Defined all the types of char allowed in our identifiers, or keywords
func isletter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' // Since _ can also be a part of the string -> which may be a identifier/keyword
}

// Creates a new token struct, given the type and char
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

// Skips over the whitespaces of the input
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}

}

// Checks if current char the lexer points to it a int
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// If the number is multidigit, it needs to be read
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPostion >= len(l.input) {
		return 0
	}
	return l.input[l.readPostion]
}

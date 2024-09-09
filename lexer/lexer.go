package lexer

import "github.com/tren03/go-interpreter/token"

type Lexer struct {
	input       string
	position    int  // Current input postition
	readPostion int  // Current reading position, next char after input
	ch          byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Initialize all fields of the lexer
	return l
}

func (l *Lexer) readChar() {
	if l.readPostion >= len(l.input) {
		l.ch = 0 // ASCII code for NIL
	} else {
		l.ch = l.input[l.readPostion]
	}
	l.position = l.readPostion
	l.readPostion += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
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
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

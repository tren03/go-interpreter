package parser

import (
	"fmt"

	"github.com/tren03/go-interpreter/ast"
	"github.com/tren03/go-interpreter/lexer"
	"github.com/tren03/go-interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer // Pointer to instance of a lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token

    prefixParseFns map[token.TokenType]prefixParsefn
    infixParseFns map[token.TokenType]infixParsefn

}

type (
	prefixParsefn func() ast.Expression
	infixParsefn  func(ast.Expression) ast.Expression // the parameter taken in is the left side of the expression, since infix takes is between two things
)

// Helper functions to add the prefix and infix func for a specific token
func(p *Parser)registerPefix(tokenType token.TokenType,fn prefixParsefn){
    p.prefixParseFns[tokenType] = fn
}

// Initialized the Parser, by taking in lexer struct, initializing it in parser and calling nextoken func for pareser to intialize the cur and peek token of parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// call Parser nextToken() twice to initialize.
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
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
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}

}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// for now we skip any other expressions till we reach semicolon

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt

}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()
	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t

}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

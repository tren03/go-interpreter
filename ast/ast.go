package ast

import "github.com/tren03/go-interpreter/token"

// Every node in ast implements the node interface, meaning it has to provide a TokenLiteral() method, -> just to debug. Some of the nodes implement statement interface and some expression interface.
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

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


// Ast node struct for Identifier -> Identifier may produce an expression, so they implement they expression interaface
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}
func(i *Identifier) expressionNode(){}
func(i *Identifier)TokenLiteral() string {return i.Token.Literal} // just so that let can be considered in the node interface


// Ast node stuct for let statement - let <identifier> = <expression> (the Right side can be a literal or expression, we take experssion to cover all cases)
type LetStatement struct {
	Token token.Token // The token.LET token
	Name  *Identifier
	Value Expression
}
func(ls *LetStatement)statementNode(){} // just so that let can be considered in the statement interface
func(ls *LetStatement)TokenLiteral() string {return ls.Token.Literal} // just so that let can be considered in the node interface

// Ast node struct for return statements - return <expression>
type ReturnStatement struct{
    Token token.Token // return token
    ReturnValue Expression
}
func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

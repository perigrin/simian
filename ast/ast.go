package ast

import "github.com/perigrin/simian/token"

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

type MyStatement struct {
	Token token.Token // "LET"
	Name  *Identifier
	Value Expression
}

func (ls *MyStatement) statementNode()       {}
func (ls *MyStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // "IDENT"
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type IntegerLiteral struct {
	Token token.Token
	Value string
}

func (i *IntegerLiteral) expressionNode()      {}
func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Literal }

type TokenNode struct {
	Token token.Token
}

func (t *TokenNode) expressionNode()      {}
func (t *TokenNode) TokenLiteral() string { return t.Token.Literal }

func TokenToAstNode(t token.Token) Node {
	switch t.Type {
	case token.IDENTIFIER:
		return &Identifier{
			Token: t,
			Value: t.Literal,
		}
	case token.DIGIT:
		// Assuming you have an IntegerLiteral type
		return &IntegerLiteral{
			Token: t,
			Value: t.Literal,
		}
	default:
		// For other token types, create a generic node or handle as needed
		return &TokenNode{Token: t}
	}
}

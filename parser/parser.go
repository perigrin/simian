package parser

import (
	"fmt"

	"github.com/perigrin/simian/ast"
	"github.com/perigrin/simian/lexer"
	"github.com/perigrin/simian/token"
)

type Parser interface {
	Errors() []string
	ParseProgram() *ast.Program
}

type parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) Parser {
	p := &parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *parser) Errors() []string {
	return p.errors
}

func (p *parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token %s got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.MY:
		return p.parseMyStatement()
	default:
		return nil
	}
}

func (p *parser) parseMyStatement() *ast.MyStatement {
	stmt := &ast.MyStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: string(p.curToken.Literal)}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: skip the expression until we get a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

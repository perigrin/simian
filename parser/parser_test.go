package parser_test

import (
	"testing"

	"github.com/perigrin/simian/ast"
	"github.com/perigrin/simian/lexer"
	"github.com/perigrin/simian/parser"
)

func TestLetStatments(t *testing.T) {
	input := `
	my $x = 5;
	my $y = 10;
	my $foobar = 838383;
	`
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("not enough statements: got %d expected %d", len(program.Statements), 3)
	}

	tests := []struct {
		Identifier string
	}{
		{"$x"},
		{"$y"},
		{"$foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testMyStatement(t, stmt, tt.Identifier) {
			return
		}
	}
}

func testMyStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "my" {
		t.Errorf("s.TokenLiteral not 'my': got %s", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement: got %T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value != '%s': got %s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() != '%s': got %s", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParseErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parse error: %s", msg)
	}
	t.FailNow()
}

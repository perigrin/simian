package lexer

import (
	"testing"

	"github.com/perigrin/simian/token"
)

func newToken(t token.TokenType, literal string) token.Token {
	return token.Token{Type: t, Literal: []byte(literal)}
}

func TestReaders(t *testing.T) {
	tests := []struct {
		input    string
		reader   func(*Lexer) token.Token
		expected token.Token
	}{
		{"my", (*Lexer).readIdentifier, newToken(token.MY, "my")},
		{"$five", (*Lexer).readIdentifier, newToken(token.IDENTIFIER, "$five")},
		{"=", (*Lexer).readOperator, newToken(token.ASSIGN, "=")},
		{"5", (*Lexer).readNumber, newToken(token.DIGIT, "5")},
		{";", (*Lexer).readSingleToken, newToken(token.SEMICOLON, ";")},
		{"$ten", (*Lexer).readIdentifier, newToken(token.IDENTIFIER, "$ten")},
		{"sub", (*Lexer).readIdentifier, newToken(token.SUB, "sub")},
		{"add", (*Lexer).readIdentifier, newToken(token.IDENTIFIER, "add")},
		{",", (*Lexer).readOperator, newToken(token.COMMA, ",")},
		{"(", (*Lexer).readSingleToken, newToken(token.LPAREN, "(")},
		{")", (*Lexer).readSingleToken, newToken(token.RPAREN, ")")},
		{"{", (*Lexer).readSingleToken, newToken(token.LBRACE, "{")},
		{"}", (*Lexer).readSingleToken, newToken(token.RBRACE, "}")},
		{"+", (*Lexer).readOperator, newToken(token.PLUS, "+")},
		{"!", (*Lexer).readOperator, newToken(token.NOT, "!")},
		{"-", (*Lexer).readOperator, newToken(token.MINUS, "-")},
		{"/", (*Lexer).readOperator, newToken(token.SLASH, "/")},
		{"*", (*Lexer).readOperator, newToken(token.ASTERISK, "*")},
		{"<", (*Lexer).readOperator, newToken(token.LT, "<")},
		{">", (*Lexer).readOperator, newToken(token.GT, ">")},
		{"==", (*Lexer).readOperator, newToken(token.EQUAL, "==")},
		{"false", (*Lexer).readIdentifier, newToken(token.FALSE, "false")},
		{"true", (*Lexer).readIdentifier, newToken(token.TRUE, "true")},
		{"class", (*Lexer).readIdentifier, newToken(token.CLASS, "class")},
		{"field", (*Lexer).readIdentifier, newToken(token.FIELD, "field")},
		{"method", (*Lexer).readIdentifier, newToken(token.METHOD, "method")},
		{"state", (*Lexer).readIdentifier, newToken(token.STATE, "state")},

		// Extra tests
		{
			"something_with_underscores",
			(*Lexer).readIdentifier,
			newToken(token.IDENTIFIER, "something_with_underscores"),
		},
		{"%hash", (*Lexer).readIdentifier, newToken(token.IDENTIFIER, "%hash")},
		{"@array", (*Lexer).readIdentifier, newToken(token.IDENTIFIER, "@array")},
		{"&sub", (*Lexer).readIdentifier, newToken(token.IDENTIFIER, "&sub")},
		{"*glob", (*Lexer).readIdentifier, newToken(token.IDENTIFIER, "*glob")},
		{"10", (*Lexer).readNumber, newToken(token.DIGIT, "10")},
		{"**", (*Lexer).readOperator, newToken(token.OP_POWER, "**")},
	}

	for i, tt := range tests {
		l := New([]byte(tt.input))
		tok := tt.reader(l)
		if tok.Type != tt.expected.Type {
			t.Fatalf("tests[%d] - token.Type wrong. expected=%q, got=%q",
				i, tt.expected.Type, tok.Type)
		}
		if string(tok.Literal) != string(tt.expected.Literal) {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expected.Literal, tok.Literal)
		}
	}
}

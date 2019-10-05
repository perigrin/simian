package lexer

import (
	"testing"

	"github.com/perigrin/simian/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x,y) {
		x + y
	};

	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if ( 5 < 10 ) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;
	`

	tests := []struct {
		Type    token.TokenType
		Literal string
	}{
		// let five = 5;
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		// let ten = 10;
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		// let add = fn(x,y) {
		// 		x + y
		// 	};
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		// let result = add(five, ten);
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		// !-/*5;
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		// 5 < 10 > 5;
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		// if ( 5 < 10 ) {
		// 		return true;
		// 	} else {
		// 		return false;
		// 	}
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		// 10 == 10;
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		// 10 != 9;
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		// EOF
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.Type {
			t.Fatalf("tests[%d] - token.type wrong, expected %q, got %q", i, tt.Type, tok.Type)
		}

		if tok.Literal != tt.Literal {
			t.Fatalf("tests[%d] - token.literal wrong, expected %q, got %q", i, tt.Literal, tok.Literal)
		}
	}
}

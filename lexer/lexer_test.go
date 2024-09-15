package lexer_test

import (
	"testing"

	"github.com/perigrin/simian/token"
	"github.com/perigrin/simian/lexer"

)

func TestNextToken(t *testing.T) {
	input := `my $five = 5;
    my $ten = 10;

    sub add($x, $y=0) { $x + $y }

    my $result = add($five, $ten);
    !-/*5;
    $five < $ten > 5;

    if ( 5 < $ten ) {
        return true;
    } else {
        return false;
    }

    $ten == 10;
    $ten != 9;
    `

	tests := []struct {
		Type    token.TokenType
		Literal string
	}{
		// my $fiver = 5;
		{token.MY, "my"},
		{token.IDENT, "$five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		// my $ten = 10;
		{token.MY, "my"},
		{token.IDENT, "$ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		// sub add	($x, $y) { $x + $y }
		{token.FUNCTION, "sub"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "$x"},
		{token.COMMA, ","},
		{token.IDENT, "$y"},
		{token.ASSIGN, "="},
		{token.INT, "0"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "$x"},
		{token.PLUS, "+"},
		{token.IDENT, "$y"},
		{token.RBRACE, "}"},
		// let result = add(five, ten);
		{token.MY, "my"},
		{token.IDENT, "$result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "$five"},
		{token.COMMA, ","},
		{token.IDENT, "$ten"},
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
		{token.IDENT, "$five"},
		{token.LT, "<"},
		{token.IDENT, "$ten"},
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
		{token.IDENT, "$ten"},
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
		{token.IDENT, "$ten"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		// 10 != 9;
		{token.IDENT, "$ten"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		// EOF
		{token.EOF, ""},
	}

	l := lexer.New(input)

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

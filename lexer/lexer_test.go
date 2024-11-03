package lexer_test

import (
	"testing"

	"github.com/perigrin/simian/lexer"
	"github.com/perigrin/simian/token"
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

    class Foo :isa(Bar) {
        field $id :reader = state $i++;

        method set_count($i) { $count = $i }
    }
    `

	tests := []struct {
		Type    token.TokenType
		Literal string
	}{
		// my $fiver = 5;
		{token.MY, "my"},
		{token.IDENTIFIER, "$five"},
		{token.ASSIGN, "="},
		{token.DIGIT, "5"},
		{token.SEMICOLON, ";"},
		// my $ten = 10;
		{token.MY, "my"},
		{token.IDENTIFIER, "$ten"},
		{token.ASSIGN, "="},
		{token.DIGIT, "10"},
		{token.SEMICOLON, ";"},
		// sub add	($x, $y) { $x + $y }
		{token.SUB, "sub"},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "$x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "$y"},
		{token.ASSIGN, "="},
		{token.DIGIT, "0"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "$x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "$y"},
		{token.RBRACE, "}"},
		// let result = add(five, ten);
		{token.MY, "my"},
		{token.IDENTIFIER, "$result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "$five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "$ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		// !-/*5;
		{token.NOT, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.DIGIT, "5"},
		{token.SEMICOLON, ";"},
		// 5 < 10 > 5;
		{token.IDENTIFIER, "$five"},
		{token.LT, "<"},
		{token.IDENTIFIER, "$ten"},
		{token.GT, ">"},
		{token.DIGIT, "5"},
		{token.SEMICOLON, ";"},
		// if ( 5 < 10 ) {
		// 		return true;
		// 	} else {
		// 		return false;
		// 	}
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.DIGIT, "5"},
		{token.LT, "<"},
		{token.IDENTIFIER, "$ten"},
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
		{token.IDENTIFIER, "$ten"},
		{token.EQUAL, "=="},
		{token.DIGIT, "10"},
		{token.SEMICOLON, ";"},
		// 10 != 9;
		{token.IDENTIFIER, "$ten"},
		{token.NOT_EQUAL, "!="},
		{token.DIGIT, "9"},
		{token.SEMICOLON, ";"},
		// class
		{token.CLASS, "class"},
		{token.IDENTIFIER, "Foo"},
		{token.IDENTIFIER, ":isa"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "Bar"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.FIELD, "field"},
		{token.IDENTIFIER, "$id"},
		{token.IDENTIFIER, ":reader"},
		{token.ASSIGN, "="},
		{token.STATE, "state"},
		{token.IDENTIFIER, "$i"},
		{token.PLUS, "++"},
		{token.SEMICOLON, ";"},

		{token.METHOD, "method"},
		{token.IDENTIFIER, "set_count"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "$i"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "$count"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "$i"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		// EOF
		{token.EOF, ""},
	}

	l := lexer.New([]byte(input))

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.Type {
			t.Fatalf("tests[%d] (%v) - token.type wrong, expected %+v, got %+v", i, tt.Literal, tt.Type, tok.Type)
		}

		if string(tok.Literal) != tt.Literal {
			t.Fatalf("tests[%d] - token.literal wrong, expected %+v, got %+v", i, tt.Literal, string(tok.Literal))
		}
	}
}

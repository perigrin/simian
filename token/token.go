package token

import (
	"log"
	"strconv"
	"unicode"
)

type TokenType string

const (
	LETTER      = "LETTER"      // Alphabet or underscore (for identifiers)
	DIGIT       = "DIGIT"       // Digits (for numbers)
	SIGIL       = "SIGIL"       // $, @, % symbols
	QUOTE       = "QUOTE"       // ' or " for string literals
	HASH        = "HASH"        // # for comments
	EQUAL       = "EQUAL"       // '=' character
	NOT_EQUAL   = "NOT_EQUAL"   // '!='
	PLUS        = "PLUS"        // '+'
	MINUS       = "MINUS"       // '-'
	SLASH       = "SLASH"       // '/'
	ASTERISK    = "ASTERISK"    // '*'
	AMPERSAND   = "AMPERSAND"   // '&'
	PIPE        = "PIPE"        // '|' bitwise OR or regex alternation
	CARET       = "CARET"       // '^' bitwise XOR
	TILDE       = "TILDE"       // '~' bitwise NOT
	LT          = "LT"          // '<'
	GT          = "GT"          // '>'
	LPAREN      = "LPAREN"      // '('
	RPAREN      = "RPAREN"      // ')'
	LBRACE      = "LBRACE"      // '{'
	RBRACE      = "RBRACE"      // '}'
	SEMICOLON   = "SEMICOLON"   // ';'
	COMMA       = "COMMA"       // ','
	FATCOMMA    = "FATCOMMA"    // '=>' (fat comma)
	DOT         = "DOT"         // '.' (string concatenation)
	DOUBLESTAR  = "DOUBLESTAR"  // '**' (exponentiation)
	NOT         = "NOT"         // '!' (logical NOT)
	ASSIGN      = "ASSIGN"      // '=' assignment
	AND         = "AND"         // 'and' logical AND
	OR          = "OR"          // 'or' logical OR
	XOR         = "XOR"         // 'xor' logical XOR
	MOD         = "MOD"         // '%' modulo
	INCREMENT   = "INCREMENT"   // '++' increment
	DECREMENT   = "DECREMENT"   // '--' decrement
	BITWISE_AND = "BITWISE_AND" // '&' bitwise AND
	BITWISE_OR  = "BITWISE_OR"  // '|' bitwise OR
	BITWISE_NOT = "BITWISE_NOT" // '~' bitwise NOT
	BITWISE_XOR = "BITWISE_XOR" // '^' bitwise XOR
	LEFT_SHIFT  = "LEFT_SHIFT"  // '<<' bitwise left shift
	RIGHT_SHIFT = "RIGHT_SHIFT" // '>>' bitwise right shift
	EOF         = "EOF"         // End of file
	ILLEGAL     = "ILLEGAL"     // Illegal character

	// Specific Perl keywords for lexing
	IF      = "IF"      // 'if'
	ELSIF   = "ELSIF"   // 'elsif'
	ELSE    = "ELSE"    // 'else'
	UNLESS  = "UNLESS"  // 'unless'
	WHILE   = "WHILE"   // 'while'
	UNTIL   = "UNTIL"   // 'until'
	FOR     = "FOR"     // 'for'
	FOREACH = "FOREACH" // 'foreach'
	DO      = "DO"      // 'do'
	NEXT    = "NEXT"    // 'next'
	LAST    = "LAST"    // 'last'
	REDO    = "REDO"    // 'redo'
	GOTO    = "GOTO"    // 'goto'
	MY      = "MY"      // 'my'
	OUR     = "OUR"     // 'our'
	LOCAL   = "LOCAL"   // 'local'
	STATE   = "STATE"   // 'state'
	SUB     = "SUB"     // 'sub'
	RETURN  = "RETURN"  // 'return'
	PACKAGE = "PACKAGE" // 'package'
	USE     = "USE"     // 'use'
	REQUIRE = "REQUIRE" // 'require'
	NO      = "NO"      // 'no'
	BEGIN   = "BEGIN"   // 'BEGIN'
	END     = "END"     // 'END'
	TRUE    = "TRUE"    // 'true' (boolean)
	FALSE   = "FALSE"   // 'false' (boolean)
	PRINT   = "PRINT"   // 'print' function
	SAY     = "SAY"     // 'say' function (if enabled)
	CHOMP   = "CHOMP"   // 'chomp' function
	CHOP    = "CHOP"    // 'chop' function
	PUSH    = "PUSH"    // 'push' function for arrays
	POP     = "POP"     // 'pop' function for arrays
	SHIFT   = "SHIFT"   // 'shift' function for arrays
	UNSHIFT = "UNSHIFT" // 'unshift' function for arrays
	SPLIT   = "SPLIT"   // 'split' function
	JOIN    = "JOIN"    // 'join' function

	IDENTIFIER = "IDENTIFIER"

	CLASS  = "CLASS"
	FIELD  = "FIELD"
	METHOD = "METHOD"

	WHITESPACE = "WHITESPACE"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"sub":    SUB,
	"my":     MY,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"class":  CLASS,
	"field":  FIELD,
	"method": METHOD,
	"state":  STATE,
}

func LookupIdent(ident string) TokenType {
	log.Printf("LookupIdent: %s", ident)
	if t, ok := keywords[ident]; ok {
		return t
	}
	return IDENTIFIER
}

// getCharType determines the type of token based on the character provided.
func GetCharType(ch rune, next rune) TokenType {
	log.Printf("GetCharType: %s %s", string(ch), string(next))
	switch {

	// Symbols and Operators
	case ch == '$' || ch == '@' || ch == '%':
		return SIGIL
	case ch == '\'' || ch == '"':
		return QUOTE
	case ch == '#':
		return HASH
	case ch == '=':
		if next == '>' {
			return FATCOMMA // `=>` fat comma
		}
		return EQUAL
	case ch == '+':
		if next == '+' {
			return INCREMENT // `++`
		}
		return PLUS
	case ch == '-':
		if next == '-' {
			return DECREMENT // `--`
		}
		return MINUS
	case ch == '/':
		return SLASH
	case ch == '*':
		if next == '*' {
			return DOUBLESTAR // `**` for exponentiation
		}
		return ASTERISK
	case ch == '&':
		return AMPERSAND
	case ch == '|':
		return PIPE
	case ch == '^':
		return CARET
	case ch == '~':
		return TILDE
	case ch == '<':
		if next == '<' {
			return LEFT_SHIFT // `<<`
		}
		return LT
	case ch == '>':
		if next == '>' {
			return RIGHT_SHIFT // `>>`
		}
		return GT
	case ch == '(':
		return LPAREN
	case ch == ')':
		return RPAREN
	case ch == '{':
		return LBRACE
	case ch == '}':
		return RBRACE
	case ch == ';':
		return SEMICOLON
	case ch == ',':
		return COMMA
	case ch == '.':
		return DOT
	case ch == '!':
		return NOT
	case ch == '%':
		return MOD
	case ch == 'a':
		if next == 'n' {
			return AND // `and`
		}
	case ch == 'o':
		if next == 'r' {
			return OR // `or`
		}
	case ch == 'x':
		if next == 'o' {
			return XOR // `xor`
		}
		// THESE ARE GENERIC THEY NEED TO GO AT THE BOTTOM
		// Identifiers and keywords
	case unicode.IsLetter(ch) || ch == '_':
		return LETTER

	// Digits
	case unicode.IsDigit(ch):
		return DIGIT

	case unicode.IsSpace(ch):
		return WHITESPACE

	// End of file or input
	case ch == '\x00':
		return EOF
	}

	// Illegal character
	log.Printf("illegal character %q", strconv.QuoteRune(ch))
	return ILLEGAL
}

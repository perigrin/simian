package token

import (
	"fmt"
	"unicode"
)

type TokenType string

const (
	INVALID = "INVALID" // Invalid token

	LETTER      = "LETTER"      // Alphabet or underscore (for identifiers)
	DIGIT       = "DIGIT"       // Digits (for numbers)
	NUMBER      = "NUMBER"      // Currently a sequence of digits
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

	// Operators
	OPERATOR                  = "OPERATOR"
	OP_UNARY                  = "UNARY OP"
	OP_MULTI                  = "MULTI OP"
	OP_REGEX                  = "REGEX OP"
	OP_SHIFT                  = "SHIFT OP"
	OP_NUMERIC_COMPARE        = "NUMERIC COMPARE OP"
	OP_BITWISE                = "BITWISE OP"
	OP_LOGICAL                = "LOGICAL OP"
	OP_LOGICAL_LOW_PRECEDENCE = "LOGICAL OP LOW PRECEDENCE"
	OP_TRI_THEN               = "TRI THEN OP"
	OP_TRI_ELSE               = "TRI ELSE OP"
	COLON                     = "COLON (:)"
	OP_REPEAT                 = "REPEAT OP (x)"
	OP_LESS_THAN_EQUAL        = "OP_LESS_THAN_EQUAL (<=)"
	OP_GREATER_THAN_EQUAL     = "OP_GREATER_THAN_EQUAL (>=)"
	OP_ARROW                  = "OP_ARROW (->)"
	OP_INC                    = "OP_INC (++)"
	OP_DEC                    = "OP_DEC (--)"
	OP_POWER                  = "OP_POWER (**)"
	OP_NOT                    = "OP_NOT (!)"
	OP_LOGICAL_DEFINED_OR     = "OP_LOGICAL_DEFINED_OR (//)"

	OP_LOGICAL_AND_LOW_PRECEDENCE = "OP_LOGICAL_AND_LOW_PRECEDENCE (and)"
	OP_LOGICAL_OR_LOW_PRECEDENCE  = "OP_LOGICAL_OR_LOW_PRECEDENCE (or)"
	OP_LOGICAL_XOR_LOW_PRECEDENCE = "OP_LOGICAL_XOR_LOW_PRECEDENCE (xor)"
	OP_LOGICAL_NOT_LOW_PRECEDENCE = "OP_LOGICAL_NOT_LOW_PRECEDENCE (not)"

	OP_COMPLEMENT            = "OP_COMPLEMENT (~)"
	OP_ADD                   = "ADD (+)"
	OP_MINUS                 = "OP_MINUS (-)"
	OP_MULTIPLY              = "OP_MULTIPLY (*)"
	OP_DIVIDE                = "OP_DIVIDE (/)"
	OP_MODULUS               = "OP_MODULUS (%)"
	ASSIGN                   = "ASSIGN (=)"
	OP_ADD_ASSIGN            = "+="
	OP_SUB_ASSIGN            = "-="
	OP_MUL_ASSIGN            = "*="
	OP_DIV_ASSIGN            = "/="
	OP_MOD_ASSIGN            = "%="
	OP_POWER_ASSIGN          = "**="
	OP_REPEAT_ASSIGN         = "OP_REPEAT_ASSIGN (x=)"
	OP_LEFT_SHIFT_ASSIGN     = "<<="
	OP_RIGHT_SHIFT_ASSIGN    = ">>="
	OP_BITWISE_AND_ASSIGN    = "&="
	OP_BITWISE_OR_ASSIGN     = "|="
	OP_BITWISE_XOR_ASSIGN    = "^="
	OP_LOGICAL_AND_ASSIGN    = "&&="
	OP_LOGICAL_OR_ASSIGN     = "||="
	OP_LEFT_SHIFT            = "<<"
	OP_RIGHT_SHIFT           = ">>"
	OP_LESS_THAN             = "<"
	OP_GREATER_THAN          = ">"
	OP_LESS_THAN_OR_EQUAL    = "<="
	OP_GREATER_THAN_OR_EQUAL = ">="
	OP_EQUAL                 = "=="
	OP_NOT_EQUAL             = "OP_NOT_EQUAL (!=)"
	OP_COMPARE               = "<=>"
	OP_BITWISE_AND           = "&"
	OP_BITWISE_OR            = "|"
	OP_BITWISE_XOR           = "^"
	OP_LOGICAL_AND           = "&&"
	OP_LOGICAL_OR            = "||"
	OP_RANGE                 = ".."
	OP_RANGE_INCLUSIVE       = "..."
	OP_TERNARY               = "?"
	OP_COLON                 = ":"
	OP_NOT_KEYWORD           = "not"
	OP_AND_KEYWORD           = "and"
	OP_OR_KEYWORD            = "or"
	OP_XOR_KEYWORD           = "xor"

	OP_MATCH   = "=~"
	OP_NOMATCH = "!~"
)

type Token struct {
	Type    TokenType
	Literal []byte
}

func (t *Token) String() string {
	return fmt.Sprintf("%s(%q)", t.Type, string(t.Literal))
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

func LookupIdent(ident []byte) TokenType {
	if t, ok := keywords[string(ident)]; ok {
		return t
	}
	return IDENTIFIER
}

func LookupSingleToken(ch byte) TokenType {
	switch {
	case string(ch) == "{":
		return LBRACE
	case string(ch) == "}":
		return RBRACE
	case string(ch) == "(":
		return LPAREN
	case string(ch) == ")":
		return RPAREN
	case string(ch) == ";":
		return SEMICOLON
	case string(ch) == "*":
		return ASTERISK
	case string(ch) == ":":
		return COLON
	case IsLetter(ch):
		return LETTER
	case IsSigil(ch):
		return SIGIL
	case IsDigit(ch):
		return DIGIT
	case IsWhitespace(ch):
		return WHITESPACE
	case IsOperator([]byte{ch}):
		return OPERATOR
	default:
		return INVALID
	}
}

var operators map[string]TokenType = map[string]TokenType{
	"->": OP_ARROW,
	"++": PLUS, // TODO OP_INC
	"--": OP_DEC,
	"**": OP_POWER,
	"+":  PLUS,  // TODO OP_ADD
	"-":  MINUS, // TODO OP_MINUS
	"=~": OP_MATCH,
	"!~": OP_NOMATCH,
	"/":  SLASH,    // TODO OP_DIVIDE
	"*":  ASTERISK, // TODO OP_MULTIPLY
	"%":  OP_MODULUS,
	"x":  OP_REPEAT,

	"==": EQUAL, // TODO OP_EQUAL
	"!=": NOT_EQUAL,
	"<=": OP_LESS_THAN_EQUAL,
	">=": OP_GREATER_THAN_EQUAL,
	"<":  LT, // TODO OP_LESS_THAN,
	">":  GT, // TODO OP_GREATER_THAN

	"&":  OP_BITWISE_AND,
	"|":  OP_BITWISE_OR,
	"^":  OP_BITWISE_XOR,
	"<<": OP_LEFT_SHIFT,
	">>": OP_RIGHT_SHIFT,
	"~":  OP_COMPLEMENT,

	"!":  NOT, // TODO OP_NOT
	"&&": OP_LOGICAL_AND,
	"||": OP_LOGICAL_OR,
	"//": OP_LOGICAL_DEFINED_OR,

	"and": OP_LOGICAL_AND_LOW_PRECEDENCE,
	"or":  OP_LOGICAL_OR_LOW_PRECEDENCE,
	"xor": OP_LOGICAL_XOR_LOW_PRECEDENCE,
	"not": OP_LOGICAL_NOT_LOW_PRECEDENCE,

	"=":   ASSIGN, // TODO OP_ASSIGN
	"+=":  OP_ADD_ASSIGN,
	"-=":  OP_SUB_ASSIGN,
	"*=":  OP_MUL_ASSIGN,
	"/=":  OP_DIV_ASSIGN,
	"%=":  OP_MOD_ASSIGN,
	"**=": OP_POWER_ASSIGN,
	"x=":  OP_REPEAT_ASSIGN,
	"<<=": OP_LEFT_SHIFT_ASSIGN,
	">>=": OP_RIGHT_SHIFT_ASSIGN,
	"&=":  OP_BITWISE_AND_ASSIGN,
	"|=":  OP_BITWISE_OR_ASSIGN,
	"^=":  OP_BITWISE_XOR_ASSIGN,
	"..":  OP_RANGE,
	"...": OP_RANGE_INCLUSIVE,
	"?":   OP_TRI_THEN,
	":":   OP_TRI_ELSE,
	",":   COMMA, // TODO OP_COMMA
	"=>":  COMMA, // TODO OP_COMMA or OP_FAT_ARROW
}

func LookupOperator(op []byte) TokenType {
	if t, ok := operators[string(op)]; ok {
		return t
	}
	return INVALID
}

func IsLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func IsDigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}

func IsSigil(ch byte) bool {
	switch rune(ch) {
	case '$', '@', '%', '&', '*':
		return true
	default:
		return false
	}
}

func IsOperator(buf []byte) bool {
	return LookupOperator(buf) != INVALID
}

// getCharType determines the type of token based on the character provided.
func GetCharType(ch rune, next rune) TokenType {
	switch {

	// Symbols and Operators
	case ch == '$' || ch == '@' || ch == '%':
		return SIGIL
	case ch == '\'' || ch == '"':
		return QUOTE
	case ch == '#':
		return HASH
	case ch == '=':
		return EQUAL
	case ch == '+':
		return PLUS
	case ch == '-':
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
	case ch == 'a' && next == 'n':
		return AND // `and`
	case ch == 'o' && next == 'r':
		return OR // `or`
	case ch == 'x' && next == 'o':
		return XOR // `xor`

	// THESE ARE GENERIC THEY NEED TO GO AT THE BOTTOM
	// Identifiers and keywords
	case unicode.IsLetter(ch) || ch == '_':
		return LETTER

	// Digits
	case unicode.IsDigit(ch):
		return NUMBER

	case unicode.IsSpace(ch):
		return WHITESPACE

	// End of file or input
	case ch == '\x00':
		return EOF
	}

	// Illegal character
	return ILLEGAL
}

func IsWhitespace(ch byte) bool {
	return unicode.IsSpace(rune(ch))
}

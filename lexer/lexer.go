package lexer

import (
	"github.com/perigrin/simian/token"
)

type (
	lexerFunction func(l *Lexer) token.Token
	state         struct {
		name string
		run  lexerFunction
	}
)

func (s state) String() string {
	return s.name
}

var stateTable = map[token.TokenType]state{
	token.SIGIL:      {name: "readIdentifier", run: (*Lexer).readIdentifier},
	token.LETTER:     {name: "readIdentifier", run: (*Lexer).readIdentifier},
	token.DIGIT:      {name: "readNumber", run: (*Lexer).readNumber},
	token.WHITESPACE: {name: "readWhitespace", run: (*Lexer).readWhitespace},
	token.OPERATOR:   {name: "readOperator", run: (*Lexer).readOperator},
	token.COLON:      {name: "readIdentifier", run: (*Lexer).readIdentifier},
}

func readerForToken(t token.TokenType) state {
	if t, ok := stateTable[t]; ok {
		return t
	}
	s := state{name: "readSingleToken", run: (*Lexer).readSingleToken}
	return s
}

type Lexer struct {
	input        []byte
	position     int
	readPosition int
	ch           byte
}

func New(input []byte) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	if l.isAtEnd() {
		return token.Token{Type: token.EOF}
	}

	nextToken := token.LookupSingleToken(l.ch)
	reader := readerForToken(nextToken)
	tok := reader.run(l)

	// skip whitespace
	if tok.Type == token.WHITESPACE {
		return l.NextToken()
	}
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readChar() {
	l.ch = l.peekChar()
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) isAtEnd() bool {
	return l.readPosition >= len(l.input)
}

func (l *Lexer) readSequence(check func(byte) bool) []byte {
	position := l.position
	for check(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() token.Token {
	matcher := func(ch byte) bool {
		switch {
		case token.IsLetter(ch):
			return true
		case token.IsSigil(ch):
			return true
		case token.IsDigit(ch):
			return true
		case ch == '_':
			return true
		case ch == ':': // start of an attribute
			return true
		default:
			return false
		}
	}

	tok := token.Token{}
	tok.Literal = l.readSequence(matcher)
	tok.Type = token.LookupIdent(tok.Literal)
	return tok
}

func (l *Lexer) readNumber() token.Token {
	matcher := func(ch byte) bool {
		return token.IsDigit(ch)
	}

	tok := token.Token{}
	tok.Literal = l.readSequence(matcher)
	tok.Type = token.DIGIT
	return tok
}

func (l *Lexer) readOperator() token.Token {
	buf := make([]byte, 0)
	matcher := func(ch byte) bool {
		buf = append(buf, ch)
		return token.IsOperator(buf)
	}

	tok := token.Token{}
	tok.Literal = l.readSequence(matcher)
	tok.Type = token.LookupOperator(tok.Literal)
	return tok
}

func (l *Lexer) readSingleToken() token.Token {
	tok := token.Token{}
	// we only need the one character
	tok.Literal = []byte{l.ch}
	tok.Type = token.LookupSingleToken(l.ch)
	l.readChar()
	return tok
}

func (l *Lexer) readWhitespace() token.Token {
	tok := token.Token{}
	tok.Literal = l.readSequence(func(ch byte) bool {
		return token.IsWhitespace(ch)
	})
	tok.Type = token.WHITESPACE
	return tok
}

func (l *Lexer) Tokens() []token.Token {
	tokens := []token.Token{}
	for {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			break
		}
		tokens = append(tokens, tok)
	}
	return tokens
}

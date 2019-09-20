package template

import (
	"errors"
)

// Lexer
type Lexer struct {
	Chars     CharIterator
	automatas map[string]Automata
}

func NewLexer(chars CharIterator) *Lexer {
	return &Lexer{
		Chars: chars,
		automatas: map[string]Automata{
			"DOG": &DOG{},
		},
	}
}

func (l *Lexer) NextToken() (t *Token, err error) {
	if l.Chars.Peek() == 0 {
		return nil, errors.New("reach EOF")
	}
	l.skipWhitespace()
	idx := l.Chars.CurrentIndex()

	for k, a := range l.automatas {
		err = a.RunGreedy(l.Chars)
		if err != nil {
			l.Chars.SetIndex(idx)
		} else {
			var tv = l.Chars.SubString(idx, l.Chars.CurrentIndex())
			idx = l.Chars.CurrentIndex()

			t = &Token{TokenType(k), TokenValue(tv)}
			err = nil
			break
		}
	}

	l.skipWhitespace()

	return
}

func (l *Lexer) HasNext() bool {
	return l.Chars.Peek() != 0
}

func (l *Lexer) skipWhitespace() {
	c := l.Chars.Peek()
	for c == ' ' || c == '\t' || c == '\n' {
		l.Chars.NextChar()
		c = l.Chars.Peek()
	}
}

// Token
type TokenType string
type TokenValue string
type Token struct {
	Type  TokenType
	Token TokenValue
}

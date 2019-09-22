package _go

import (
	"errors"
	"unicode"
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
	if Peek() == 0 {
		return nil, errors.New("reach EOF")
	}
	l.skipWhitespace()
	idx := CurrentIndex()

	for k, a := range l.automatas {
		err = RunGreedy(l.Chars)
		if err != nil {
			SetIndex(idx)
		} else {
			var tv = SubString(idx, CurrentIndex())
			idx = CurrentIndex()

			t = &Token{TokenType(k), TokenValue(tv)}
			err = nil
			break
		}
	}

	l.skipWhitespace()

	return
}

func (l *Lexer) HasNext() bool {
	return Peek() != 0
}

func (l *Lexer) skipWhitespace() {
	c := Peek()
	for unicode.IsSpace(c) {
		NextChar()
		c = Peek()
	}
}

// Token
type TokenType string
type TokenValue string
type Token struct {
	Type  TokenType
	Token TokenValue
}

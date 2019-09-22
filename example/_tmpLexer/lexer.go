package lexer

import (
	"errors"
	"unicode"
)

// Lexer
type Lexer struct {
	Chars     CharIterator
	automatas []struct{
		Name string
		Automata Automata
	}
}

func NewLexer(chars CharIterator) *Lexer {
	return &Lexer{
		Chars: chars,
		automatas: []struct{
			Name string
			Automata Automata
		}{
			{"IF", &IF{}},
			{"ELSE", &ELSE{}},
			{"LP", &LP{}},
			{"RP", &RP{}},
			{"LB", &LB{}},
			{"RB", &RB{}},
			{"SEM", &SEM{}},
			{"VAR", &VAR{}},
			{"RETURN", &RETURN{}},
			{"EQ", &EQ{}},
			{"ASSIGN", &ASSIGN{}},
			{"NUM", &NUM{}},
			{"IDENT", &IDENT{}},
			{"COMMENT", &COMMENT{}},
		},
	}
}

func (l *Lexer) NextToken() (t *Token, err error) {
	if l.Chars.Peek() == 0 {
		return nil, errors.New("reach EOF")
	}
	l.skipWhitespace()
	idx := l.Chars.CurrentIndex()

	for _, a := range l.automatas {
		err = a.Automata.RunGreedy(l.Chars)
		if err != nil {
			l.Chars.SetIndex(idx)
		} else {
			var tv = l.Chars.SubString(idx, l.Chars.CurrentIndex())
			idx = l.Chars.CurrentIndex()

			t = &Token{TokenType(a.Name), TokenValue(tv)}
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
	for unicode.IsSpace(c) {
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

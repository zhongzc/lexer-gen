package template

import (
	"errors"
	"fmt"
)

// Character reader
type CharIterator interface {
	CurrentIndex() int
	NextChar() rune
	SetIndex(i int)
	SubString(from int, limit int) []rune
}

// Character reader implement
type CharStream struct {
	curIdx int
	runes  []rune
}

func NewStream(input string) *CharStream {
	return &CharStream{0, []rune(input)}
}

func (cs *CharStream) CurrentIndex() int {
	return cs.curIdx
}

func (cs *CharStream) NextChar() rune {
	if cs.curIdx >= len(cs.runes) {
		return 0
	}
	r := cs.runes[cs.curIdx]
	cs.curIdx++
	return r
}

func (cs *CharStream) SetIndex(i int) {
	cs.curIdx = i
}

func (cs *CharStream) SubString(from int, limit int) []rune {
	return cs.runes[from:limit]
}

// Token
type TokenType string
type TokenValue string
type Token struct {
	Type  TokenType
	Token TokenValue
}

// Lexer
type Lexer struct {
	Chars     CharIterator
	automatas map[string]Automata
}

func NewLexer(chars CharIterator) *Lexer {
	return &Lexer{
		Chars: chars,
		automatas: map[string]Automata{
			"A": &A{},
			"B": &B{},
		},
	}
}

func (l *Lexer) NextToken() (t *Token, err error) {
	if l.Chars.NextChar() == 0 {
		return nil, errors.New("reach EOF")
	}
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
	return
}

// Automata
type Automata interface {
	RunGreedy(iter CharIterator) error
}

type A struct{}

func (t *A) RunGreedy(iter CharIterator) error {
	currentState := 1
	acceptState := map[int]bool{
		3: true,
	}

outer:
	for c := iter.NextChar(); ; c = iter.NextChar() {
		switch currentState {
		case 1:
			switch c {
			case 'a':
				currentState = 2
			case 'b':
				currentState = 1
			default:
				break outer
			}
		case 2:
			switch c {
			case 'a':
				currentState = 2
			case 'b':
				currentState = 3
			default:
				break outer
			}
		case 3:
			switch c {
			case 'a':
				currentState = 3
			case 'b':
				currentState = 3
			default:
				break outer
			}
		default:
			break outer
		}
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", t)
	return errors.New(msg)
}

type B struct{}

func (t *B) RunGreedy(iter CharIterator) error {
	currentState := 1
	acceptState := map[int]bool{
		3: true,
	}

outer:
	for c := iter.NextChar(); ; c = iter.NextChar() {
		switch currentState {
		case 1:
			switch c {
			case 'a':
				currentState = 2
			case 'b':
				currentState = 1
			default:
				break outer
			}
		case 2:
			switch c {
			case 'a':
				currentState = 2
			case 'b':
				currentState = 3
			default:
				break outer
			}
		case 3:
			switch c {
			case 'a':
				currentState = 3
			case 'b':
				currentState = 3
			default:
				break outer
			}
		default:
			break outer
		}
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", t)
	return errors.New(msg)
}

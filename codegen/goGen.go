package codegen

import (
	"bytes"
	"fmt"
	"github.com/zhongzc/lexerGen/fa"
	"io"
)

type GoGen struct{}

func (*GoGen) Generate(dfas []*NamedDFA) map[string]func(io.Writer) error {
	res := make(map[string]func(io.Writer) error)

	res["lexer.go"] = func(writer io.Writer) (err error) {
		_, err = io.WriteString(writer, lexerHeader)
		if err != nil {
			return
		}

		var nameBuf bytes.Buffer
		for _, lx := range dfas {
			nameBuf.WriteString(fmt.Sprintf(
				"\n\t\t\t\"%s\": &%s{},", lx.Name, lx.Name,
			))
		}
		_, err = io.WriteString(writer, fmt.Sprintf(lexer, nameBuf.String()))
		if err != nil {
			return
		}

		_, err = io.WriteString(writer, token)
		if err != nil {
			return
		}
		return
	}

	res["charIterator.go"] = func(writer io.Writer) (err error) {
		_, err = io.WriteString(writer, charIterHeader)
		if err != nil {
			return
		}

		_, err = io.WriteString(writer, charReader)
		return
	}

	res["automata.go"] = func(writer io.Writer) (err error) {
		_, err = io.WriteString(writer, automataHeader)
		if err != nil {
			return
		}

		_, err = io.WriteString(writer, automata)
		if err != nil {
			return
		}

		for _, lx := range dfas {
			c := buildLexerCode(lx)
			_, err = io.WriteString(writer, c)
			if err != nil {
				return
			}
		}
		return
	}

	res["lexer_test.go"] = func(writer io.Writer) (err error) {
		_, err = io.WriteString(writer, testTemplate)
		return
	}

	return res
}

func buildLexerCode(dfa *NamedDFA) string {
	var buf bytes.Buffer
	buf.WriteString("\ntype " + dfa.Name + " struct{}")

	ruleFromState := make(map[int][]*fa.Rule)
	for _, r := range dfa.DFA.RuleBook.Rules {
		ruleFromState[r.From] = append(ruleFromState[r.From], r)
	}

	var stateTransfers bytes.Buffer
	for from, transfer := range ruleFromState {
		var transferRules bytes.Buffer
		for _, rule := range transfer {
			transferRules.WriteString(fmt.Sprintf(
				transferRuleTemplate, rule.By, rule.To,
			))
		}
		stateTransfers.WriteString(fmt.Sprintf(
			stateTransferTemplate, from, transferRules.String(),
		))
	}

	var acceptStates bytes.Buffer
	for k := range dfa.DFA.AcceptStates {
		acceptStates.WriteString(fmt.Sprintf("\t\t%d: true,\n", k))
	}

	buf.WriteString(fmt.Sprintf(
		automataTemplate,
		dfa.Name,
		dfa.DFA.CurrentState,
		acceptStates.String(),
		stateTransfers.String(),
	))

	return buf.String()
}

const (
	lexerHeader = `package lexer

import (
	"errors"
)
`
	charIterHeader = `package lexer
`
	automataHeader = `package lexer

import (
	"errors"
	"fmt"
)
`

	charReader = `
// character reader
type CharIterator interface {
	CurrentIndex() int
	NextChar() rune
	Peek() rune
	SetIndex(i int)
	SubString(l int, r int) []rune
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

func (cs *CharStream) Peek() rune {
	if cs.curIdx >= len(cs.runes) {
		return 0
	}
	return cs.runes[cs.curIdx]
}

func (cs *CharStream) SetIndex(i int) {
	cs.curIdx = i
}

func (cs *CharStream) SubString(from int, limit int) []rune {
	return cs.runes[from:limit]
}
`

	token = `
// Token
type TokenType string
type TokenValue string
type Token struct {
	Type  TokenType
	Token TokenValue
}
`

	// lexer(
	//     []name,
	// )
	lexer = `
// Lexer
type Lexer struct {
	Chars     CharIterator
	automatas map[string]Automata
}

func NewLexer(chars CharIterator) *Lexer {
	return &Lexer{
		Chars: chars,
		automatas: map[string]Automata{%s
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
`

	automata = `
// Automata
type Automata interface {
	RunGreedy(iter CharIterator) error
}
`

	// automataTemplate(
	//     name,
	//     startState,
	//     acceptStates,
	//     []stateTransfer,
	// )
	automataTemplate = `
func (a *%s) RunGreedy(iter CharIterator) error {
	currentState := %d
	acceptState := map[int]bool{
%s	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
%s
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%%T run failed", a)
	return errors.New(msg)
}
`
	// stateTransfer(
	//     fromState,
	//     []transferRule,
	// )
	stateTransferTemplate = `
		case %d:
			switch c {
%s			default:
				break outer
			}
`

	// rule(
	//     by,
	//     to,
	// )
	transferRuleTemplate = `			case '%c':
				currentState = %d
`

	testTemplate = `package lexer

import (
	"testing"
)

func TestLexer(t *testing.T) {
	cs := NewStream(` + "`" +`
xxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxx
` + "`" +
		`)
	l := NewLexer(cs)
	for l.HasNext() {
		tk, err := l.NextToken()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(tk)
	}
}
`
)

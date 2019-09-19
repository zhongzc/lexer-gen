package codegen

import (
	"bytes"
	"fmt"
	"github.com/zhongzc/lexerGen/fa"
	"io"
	"log"
)

type GoGen struct{}

func (*GoGen) Generate(writer io.Writer, lxs []*Lexer) {
	mustWrite(writer, header)
	mustWrite(writer, charReader)
	mustWrite(writer, token)

	var nameBuf bytes.Buffer
	for _, lx := range lxs {
		nameBuf.WriteString(fmt.Sprintf(
			"\n\t\t\t\"%s\": &%s{},", lx.Name, lx.Name,
		))
	}
	mustWrite(writer, fmt.Sprintf(lexer, nameBuf.String()))

	mustWrite(writer, automata)
	for _, lx := range lxs {
		c := buildLexerCode(lx)
		mustWrite(writer, c)
	}
}

func buildLexerCode(lexer *Lexer) string {
	var buf bytes.Buffer
	buf.WriteString("\ntype " + lexer.Name + " struct{}")

	ruleFromState := make(map[int][]*fa.Rule)
	for _, r := range lexer.DFA.RuleBook.Rules {
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
	for k := range lexer.DFA.AcceptStates {
		acceptStates.WriteString(fmt.Sprintf("\t\t%d: true,", k))
	}

	buf.WriteString(fmt.Sprintf(
		automataTemplate,
		lexer.Name,
		lexer.DFA.CurrentState,
		acceptStates.String(),
		stateTransfers.String(),
	))

	return buf.String()
}

func mustWrite(writer io.Writer, s string) {
	_, err := io.WriteString(writer, s)
	if err != nil {
		log.Fatal(err)
	}
}

const (
	header = `
package lexer

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
%s
	}

outer:
	for c := iter.NextChar(); ; c = iter.NextChar() {
		switch currentState {
%s
		default:
			break outer
		}
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
)

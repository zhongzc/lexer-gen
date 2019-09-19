package dfa

import (
	"errors"
	"fmt"
	. "github.com/zhongzc/lexerGen/fa"
)

type DFA struct {
	RuleBook     *RuleBook
	CurrentState int
	AcceptStates StateSet
}

func New(ruleBook *RuleBook, startState int, acceptStates ...int) *DFA {
	as := make(map[int]bool)
	for _, s := range acceptStates {
		as[s] = true
	}
	return &DFA{ruleBook, startState, as}
}

func (dfa *DFA) CanAccept() bool {
	_, ok := dfa.AcceptStates[dfa.CurrentState]
	return ok
}

func (dfa *DFA) ReadChar(by rune) (err error) {
	dfa.CurrentState, err = dfa.RuleBook.nextState(dfa.CurrentState, by)
	return
}

func (dfa *DFA) ReadString(input string) (err error) {
	for _, by := range []rune(input) {
		err = dfa.ReadChar(by)
		if err != nil {
			return
		}
	}
	return nil
}

type RuleBook struct {
	Rules []*Rule
}

func (rb *RuleBook) nextState(state int, by rune) (nextState int, err error) {
	for _, r := range rb.Rules {
		if next, ok := r.Apply(state, by); ok {
			return next, nil
		}
	}
	msg := fmt.Sprintf("DFA.nextState() can not accept %c", by)
	return -1, errors.New(msg)
}

package dfa

import (
	. "github.com/zhongzc/lexerGen/fa"
)

type DFA struct {
	RuleBook     *RuleBook
	CurrentState int
	AcceptStates StateSet
}

func New(ruleBook *RuleBook, startState int, acceptStates StateSet) *DFA {
	return &DFA{ruleBook, startState, acceptStates}
}

func (dfa *DFA) Match(input string) bool {
	old := dfa.CurrentState
	defer func() { dfa.CurrentState = old }()
	err := dfa.ReadString(input)
	if err != nil || !dfa.CanAccept() {
		return false
	}
	return true
}

func (dfa *DFA) CanAccept() bool {
	_, ok := dfa.AcceptStates[dfa.CurrentState]
	return ok
}

func (dfa *DFA) ReadChar(by rune) (err error) {
	dfa.CurrentState, err = dfa.RuleBook.NextState(dfa.CurrentState, OneChar(by))
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

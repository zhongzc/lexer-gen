package nfa

import (
	. "github.com/zhongzc/lexerGen/fa"
)

type NFA struct {
	RuleBook      *RuleBook
	currentStates StateSet
	AcceptStates  StateSet
}

func New(ruleBook *RuleBook, currentStates StateSet, acceptStates StateSet) *NFA {
	return &NFA{ruleBook, currentStates, acceptStates}
}

func (nfa *NFA) CurrentStates() StateSet {
	return nfa.RuleBook.FreeMove(nfa.currentStates)
}

func (nfa *NFA) CanAccept() bool {
	for k := range nfa.CurrentStates() {
		if nfa.AcceptStates[k] {
			return true
		}
	}
	return false
}

func (nfa *NFA) ReadChar(by rune) (err error) {
	nfa.currentStates, err = nfa.RuleBook.NextStates(nfa.CurrentStates(), by)
	return
}

func (nfa *NFA) ReadString(input string) (err error) {
	for _, by := range []rune(input) {
		err = nfa.ReadChar(by)
		if err != nil {
			return
		}
	}
	return nil
}

package dfa

import (
	"errors"
	"fmt"
)

type DFA struct {
	RuleBook *RuleBook
	CurrentState int
	AcceptStates map[int]bool
}

func (dfa *DFA) CanAccept() bool {
	_, ok := dfa.AcceptStates[dfa.CurrentState]
	return ok
}

func (dfa *DFA) ReadChar(by rune) (err error) {
	dfa.CurrentState, err = dfa.RuleBook.NextState(dfa.CurrentState, by)
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

type Rule struct {
	From int
	By   rune
	To   int
}

func (r *Rule) apply(from int, by rune) (next int, ok bool) {
	if r.From == from && r.By == by {
		return r.To, true
	} else {
		return r.From, false
	}
}

func (r *Rule) String() string {
	return fmt.Sprintf("%d --%c--> %d", r.From, r.By, r.To)
}

type RuleBook struct {
	Rules []*Rule
}

func (rb *RuleBook) NextState(state int, by rune) (nextState int, err error) {
	for _, r := range rb.Rules {
		if next, ok := r.apply(state, by); ok {
			return next, nil
		}
	}
	msg := fmt.Sprintf("DFA.NextState() can not accept %c", by)
	return -1, errors.New(msg)
}


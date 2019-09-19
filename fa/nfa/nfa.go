package nfa

import (
	"errors"
	"fmt"
	. "github.com/zhongzc/lexerGen/fa"
)

type NFA struct {
	RuleBook      *RuleBook
	currentStates StateSet
	AcceptStates  StateSet
}

func New(ruleBook *RuleBook, currentStates []int, acceptStates []int) *NFA {
	cs := make(map[int]bool)
	for _, s := range currentStates {
		cs[s] = true
	}
	as := make(map[int]bool)
	for _, s := range acceptStates {
		as[s] = true
	}
	return &NFA{ruleBook, cs, as}
}

func (nfa *NFA) CurrentStates() StateSet {
	return nfa.RuleBook.freeMove(nfa.currentStates)
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
	nfa.currentStates, err = nfa.RuleBook.nextStates(nfa.CurrentStates(), by)
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

type RuleBook struct {
	Rules []*Rule
}

func (rb *RuleBook) nextStates(ss StateSet, by rune) (nextStateSet StateSet, err error) {
	nextStateSet = make(map[int]bool)
	for _, r := range rb.Rules {
		if ss[r.From] {
			if r.By == by {
				nextStateSet[r.To] = true
			}
		}
	}
	if len(nextStateSet) == 0 {
		msg := fmt.Sprintf("NFA.nextStates() can not accept %c", by)
		return nextStateSet, errors.New(msg)
	}
	return
}

func (rb *RuleBook) freeMove(ss StateSet) StateSet {
	more, _ := rb.nextStates(ss, 0)
	if more.LE(ss) {
		return ss
	} else {
		ss.Add(more)
		return rb.freeMove(ss)
	}
}

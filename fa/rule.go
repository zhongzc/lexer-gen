package fa

import (
	"errors"
	"fmt"
)

type Rule struct {
	From int
	By   rune
	To   int
}

func (r *Rule) Apply(from int, by rune) (next int, ok bool) {
	if r.From == from && r.By == by {
		return r.To, true
	} else {
		return -1, false
	}
}

type RuleBook struct {
	Rules []*Rule
}

func (rb *RuleBook) Alphabet() (a map[rune]bool) {
	a = make(map[rune]bool)
	for _, r := range rb.Rules {
		a[r.By] = true
	}
	delete(a, 0)
	return
}

func (rb *RuleBook) NextState(state int, by rune) (nextState int, err error) {
	for _, r := range rb.Rules {
		if next, ok := r.Apply(state, by); ok {
			return next, nil
		}
	}
	msg := fmt.Sprintf("DFA.nextState() can not accept %c", by)
	return -1, errors.New(msg)
}

func (rb *RuleBook) NextStates(ss StateSet, by rune) (nextStateSet StateSet, err error) {
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

func (rb *RuleBook) FreeMove(ss StateSet) StateSet {
	more, _ := rb.NextStates(ss, 0)
	if more.LE(ss) {
		return ss
	} else {
		ss.Add(more)
		return rb.FreeMove(ss)
	}
}

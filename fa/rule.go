package fa

import (
	"errors"
	"fmt"
)

type Rule struct {
	From int
	By   Charset
	To   int
}

func NewRule(from int, by Charset, to int) *Rule {
	return &Rule{From: from, By: by, To: to}
}

type Charset struct {
	LeftMost  rune
	RightMost rune
}

func (r *Rule) Apply(from int, by Charset) (next int, ok bool) {
	if r.From == from && r.By.LeftMost <= by.LeftMost && by.RightMost <= r.By.RightMost {
		return r.To, true
	} else {
		return -1, false
	}
}

type RuleBook struct {
	Rules []*Rule
}

func (rb *RuleBook) Alphabet() (a map[Charset]bool) {
	a = make(map[Charset]bool)
	for _, r := range rb.Rules {
		a[r.By] = true
	}
	delete(a, Charset{0, 0})
	return
}

func (rb *RuleBook) NextState(state int, by Charset) (nextState int, err error) {
	for _, r := range rb.Rules {
		if next, ok := r.Apply(state, by); ok {
			return next, nil
		}
	}
	msg := fmt.Sprintf("DFA.nextState() can not accept %c", by)
	return -1, errors.New(msg)
}

func (rb *RuleBook) NextStates(ss StateSet, by Charset) (nextStateSet StateSet, err error) {
	nextStateSet = make(map[int]bool)
	for _, r := range rb.Rules {
		if ss[r.From] {
			if next, ok := r.Apply(r.From, by); ok {
				nextStateSet[next] = true
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
	more, _ := rb.NextStates(ss, Charset{0, 0})
	if more.LE(ss) {
		return ss
	} else {
		ss.Add(more)
		return rb.FreeMove(ss)
	}
}

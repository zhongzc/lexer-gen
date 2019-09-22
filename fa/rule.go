package fa

import (
	"errors"
	"fmt"
	"sort"
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

func Free() Charset {
	return Charset{0, 0}
}
func OneChar(c rune) Charset {
	return Charset{c, c}
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

func NewRuleBook(rules []*Rule) *RuleBook {
	splitPoints := make(map[rune]bool)
	for _, r := range rules {
		splitPoints[r.By.LeftMost] = true
		splitPoints[r.By.RightMost] = true
	}
	ordered := make([]int, 0, len(splitPoints))
	for p := range splitPoints {
		ordered = append(ordered, int(p))
	}
	sort.Ints(ordered)

	newRB := make([]*Rule, 0, len(rules))
	for _, r := range rules {
		lt := sort.SearchInts(ordered, int(r.By.LeftMost))
		rt := sort.SearchInts(ordered, int(r.By.RightMost))
		if lt == rt || lt+1 == rt {
			newRB = append(newRB, r)
			continue
		}
		for i := lt; i < rt; i++ {
			cs := Charset{rune(ordered[i]), rune(ordered[i+1])}
			newRule := NewRule(r.From, cs, r.To)
			newRB = append(newRB, newRule)
		}
	}
	return &RuleBook{Rules: newRB}
}

func (rb *RuleBook) Alphabet() (a map[Charset]bool) {
	a = make(map[Charset]bool)
	for _, r := range rb.Rules {
		a[r.By] = true
	}
	delete(a, Free())
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
	more, _ := rb.NextStates(ss, Free())
	if more.LE(ss) {
		return ss
	} else {
		ss.Add(more)
		return rb.FreeMove(ss)
	}
}

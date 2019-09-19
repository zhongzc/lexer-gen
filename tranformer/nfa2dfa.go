package tranformer

import (
	"github.com/zhongzc/lexerGen/fa"
	"github.com/zhongzc/lexerGen/fa/dfa"
	"github.com/zhongzc/lexerGen/fa/nfa"
)

func ToDFA(nfa *nfa.NFA) *dfa.DFA {
	ns := &NFASimulation{NFA: nfa}
	return ns.ToDFA()
}

type SRule struct {
	From fa.StateSet
	By   rune
	To   fa.StateSet
}
type NFASimulation struct {
	NFA *nfa.NFA
}

func (ns *NFASimulation) ToDFA() *dfa.DFA {
	sss, rs := ns.DiscoverStatesAndRules(NewSSS(ns.NFA.CurrentStates()))
	findIdx := func(ss fa.StateSet) int {
		for i := range sss {
			if sss[i].Equal(ss) {
				return i
			}
		}
		return -1
	}

	// buildStartState
	var startState = findIdx(ns.NFA.CurrentStates())

	// buildRules
	var rules []*fa.Rule
	for _, r := range rs {
		rules = append(rules,
			&fa.Rule{From: findIdx(r.From), By: r.By, To: findIdx(r.To)})
	}

	// buildAcceptStates
	var acceptStates = make([]int, 0)
	for _, ss := range sss {
		if nfa.New(ns.NFA.RuleBook, ss, ns.NFA.AcceptStates).CanAccept() {
			acceptStates = append(acceptStates, findIdx(ss))
		}
	}

	return dfa.New(&fa.RuleBook{Rules: rules}, startState, fa.NewSet(acceptStates...))
}

func (ns *NFASimulation) NextState(from fa.StateSet, by rune) fa.StateSet {
	n := nfa.New(ns.NFA.RuleBook, from, fa.NewSet())

	_ = n.ReadChar(by)
	return n.CurrentStates()
}

func (ns *NFASimulation) RulesFor(ss fa.StateSet) (res []*SRule) {
	alphabet := ns.NFA.RuleBook.Alphabet()
	for a := range alphabet {
		t := ns.NextState(ss, a)
		if len(t) != 0 {
			res = append(res, &SRule{From: ss, By: a, To: t})
		}
	}
	return res
}

func (ns *NFASimulation) DiscoverStatesAndRules(sss StateSetSet) (res StateSetSet, rs []*SRule) {
	for _, ss := range sss {
		rs = append(rs, ns.RulesFor(ss)...)
	}
	moreStates := NewSSS()
	for _, r := range rs {
		moreStates = moreStates.append(r.To)
	}
	if sss.gt(moreStates) {
		return sss, rs
	} else {
		return ns.DiscoverStatesAndRules(sss.add(moreStates))
	}
}

// bunch of things for go's poor infrastructure
type StateSetSet []fa.StateSet

func NewSSS(ss ...fa.StateSet) StateSetSet {
	return uniq(ss)
}
func (sss StateSetSet) append(ss fa.StateSet) StateSetSet {
	if sss.contain(ss) {
		return sss
	}
	return append(sss, ss)
}
func (sss StateSetSet) equal(o StateSetSet) bool {
	if len(sss) != len(o) {
		return false
	}
	for _, ss := range sss {
		if !o.contain(ss) {
			return false
		}
	}
	return true
}
func (sss StateSetSet) add(o StateSetSet) StateSetSet {
	return uniq(append(sss, o...))
}
func (sss StateSetSet) gt(o StateSetSet) bool {
	if len(sss) < len(o) {
		return false
	}
	for _, ss := range o {
		if !sss.contain(ss) {
			return false
		}
	}
	return true
}
func uniq(sss StateSetSet) (res StateSetSet) {
	res = make([]fa.StateSet, 0)
	for _, ss := range sss {
		if !res.contain(ss) {
			res = append(res, ss)
		}
	}
	return
}
func (sss StateSetSet) contain(ss fa.StateSet) bool {
	for _, s := range sss {
		if s.Equal(ss) {
			return true
		}
	}
	return false
}

package tranformer

import (
	"github.com/zhongzc/lexerGen/fa"
	"github.com/zhongzc/lexerGen/fa/nfa"
	"github.com/zhongzc/lexerGen/reast"
	"log"
)

func ToNFA(re reast.RegEx) *nfa.NFA {
	return toNFA(re, NewIdGenerator(0))
}

func toNFA(re reast.RegEx, ig *IdGenerator) *nfa.NFA {
	switch re.(type) {
	case *reast.Primitive:
		return primitiveToNFA(re.(*reast.Primitive), ig)
	case *reast.Sequence:
		return sequenceToNFA(re.(*reast.Sequence), ig)
	case *reast.Choose:
		return chooseToNFA(re.(*reast.Choose), ig)
	case *reast.Repeat:
		return repeatToNFA(re.(*reast.Repeat), ig)
	default:
		return nil
	}
}

func primitiveToNFA(l *reast.Primitive, ig *IdGenerator) *nfa.NFA {
	startState := ig.NextId()
	acceptState := ig.NextId()
	rule := &fa.Rule{From: startState, By: l.Rune, To: acceptState}
	rb := &fa.RuleBook{Rules: []*fa.Rule{rule}}
	return nfa.New(rb, fa.NewSet(startState), fa.NewSet(acceptState))
}

func sequenceToNFA(s *reast.Sequence, ig *IdGenerator) *nfa.NFA {
	cur := toNFA(s.RegExs[0], ig)
	startState := cur.CurrentStates()

	rules := cur.RuleBook.Rules
	var next *nfa.NFA
	for _, re := range s.RegExs[1:] {
		next = toNFA(re, ig)
		rules = append(rules, next.RuleBook.Rules...)
		for a := range cur.AcceptStates {
			for s := range next.CurrentStates() {
				rules = append(rules, &fa.Rule{From: a, By: 0, To: s})
			}
		}
		cur = next
	}
	if next == nil {
		log.Fatalf("reast.Sequence.RegExs length should greater than 1")
	}
	acceptStates := next.AcceptStates
	return nfa.New(&fa.RuleBook{Rules: rules}, startState, acceptStates)
}

func chooseToNFA(c *reast.Choose, ig *IdGenerator) *nfa.NFA {
	startState := ig.NextId()

	rules := make([]*fa.Rule, 0)
	acceptStates := fa.NewSet()
	for _, re := range c.RegExs {
		n := toNFA(re, ig)
		acceptStates.Add(n.AcceptStates)
		rules = append(rules, n.RuleBook.Rules...)
		for k := range n.CurrentStates() {
			rules = append(rules, &fa.Rule{From: startState, By: 0, To: k})
		}
	}

	rb := &fa.RuleBook{Rules: rules}
	return nfa.New(rb, fa.NewSet(startState), acceptStates)
}

func repeatToNFA(r *reast.Repeat, ig *IdGenerator) *nfa.NFA {
	a := toNFA(r.RegEx, ig)

	startState := ig.NextId()
	acceptStates := a.AcceptStates
	acceptStates.Add(fa.NewSet(startState))
	rules := a.RuleBook.Rules
	for sta := range a.CurrentStates() {
		for acc := range a.AcceptStates {
			rules = append(rules, &fa.Rule{From: acc, By: 0, To: sta})
		}
		rules = append(rules, &fa.Rule{From: startState, By: 0, To: sta})
	}

	rb := &fa.RuleBook{Rules: rules}
	return nfa.New(rb, fa.NewSet(startState), acceptStates)
}

type IdGenerator struct {
	curId int
}

func NewIdGenerator(curId int) *IdGenerator {
	return &IdGenerator{curId: curId}
}

func (ig *IdGenerator) NextId() int {
	res := ig.curId
	ig.curId++
	return res
}
package tranformer

import (
	"github.com/zhongzc/lexerGen/fa"
	"github.com/zhongzc/lexerGen/fa/nfa"
	"github.com/zhongzc/lexerGen/reast"
	"log"
)

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

func ToNFA(re reast.RegEx, ig *IdGenerator) *nfa.NFA {
	switch re.(type) {
	case *reast.Primitive:
		return primitiveToNFA(re.(*reast.Primitive), ig)
	case *reast.Sequence:
		return sequenceToNFA(re.(*reast.Sequence), ig)

		// TODO: other rules

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
	cur := ToNFA(s.RegExs[0], ig)
	startState := cur.CurrentStates()

	rules := cur.RuleBook.Rules
	var next *nfa.NFA
	for _, re := range s.RegExs[1:] {
		next = ToNFA(re, ig)
		rules = append(rules, next.RuleBook.Rules...)
		for a := range cur.AcceptStates {
			for s := range next.CurrentStates() {
				rules = append(rules, &fa.Rule{From: a, By: 0, To: s})
			}
		}
	}

	if next == nil {
		log.Fatalf("reast.Sequence.RegExs length should greater than 1")
	}
	acceptStates := next.AcceptStates
	return nfa.New(&fa.RuleBook{Rules: rules}, startState, acceptStates)
}

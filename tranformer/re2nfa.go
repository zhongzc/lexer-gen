package tranformer

import (
	"github.com/zhongzc/lexerGen/fa"
	"github.com/zhongzc/lexerGen/fa/nfa"
	"github.com/zhongzc/lexerGen/reast"
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
	rule := &fa.Rule{From: startState, By: fa.Charset{LeftMost: l.From, RightMost: l.To}, To: acceptState}
	rb := fa.NewRuleBook([]*fa.Rule{rule})
	return nfa.New(rb, fa.NewSet(startState), fa.NewSet(acceptState))
}

func sequenceToNFA(s *reast.Sequence, ig *IdGenerator) *nfa.NFA {
	cur := toNFA(s.Left, ig)
	startState := cur.CurrentStates()
	rules := cur.RuleBook.Rules
	next := toNFA(s.Right, ig)
	rules = append(rules, next.RuleBook.Rules...)
	for a := range cur.AcceptStates {
		for s := range next.CurrentStates() {
			rules = append(rules, fa.NewRule(a, fa.Free(), s))
		}
	}
	acceptStates := next.AcceptStates
	return nfa.New(fa.NewRuleBook(rules), startState, acceptStates)
}

func chooseToNFA(c *reast.Choose, ig *IdGenerator) *nfa.NFA {
	startState := ig.NextId()
	acceptStates := fa.NewSet()

	c1 := toNFA(c.Left, ig)
	acceptStates.Add(c1.AcceptStates)
	c2 := toNFA(c.Right, ig)
	acceptStates.Add(c2.AcceptStates)
	rules := append(c1.RuleBook.Rules, c2.RuleBook.Rules...)
	for k := range c1.CurrentStates() {
		rules = append(rules, fa.NewRule(startState, fa.Free(), k))
	}
	for k := range c2.CurrentStates() {
		rules = append(rules, fa.NewRule(startState, fa.Free(), k))
	}

	rb := fa.NewRuleBook(rules)
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
			rules = append(rules, fa.NewRule(acc, fa.Free(), sta))
		}
		rules = append(rules, fa.NewRule(startState, fa.Free(), sta))
	}

	rb := fa.NewRuleBook(rules)
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

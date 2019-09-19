package tranformer

import (
	"github.com/zhongzc/lexerGen/fa"
	"github.com/zhongzc/lexerGen/fa/nfa"
	"testing"
)

var ns *NFASimulation

func init() {
	rb := &fa.RuleBook{Rules: []*fa.Rule{
		{1, 'a', 1},
		{1, 'a', 2},
		{1, 0, 2},
		{2, 'b', 3},
		{3, 'b', 1},
		{3, 0, 2},
	}}
	n := nfa.New(rb, fa.NewSet(1), fa.NewSet(3))
	ns = &NFASimulation{n}
}

func TestNFASimulation_NextState(t *testing.T) {

	if !ns.NextState(fa.NewSet(1, 2), 'a').Equal(fa.NewSet(1, 2)) {
		t.Errorf("ns.NextState() expected %v. got %v", fa.NewSet(1, 2), ns.NextState(fa.NewSet(1, 2), 'a'))
	}
	if !ns.NextState(fa.NewSet(1, 2), 'b').Equal(fa.NewSet(2, 3)) {
		t.Errorf("ns.NextState() expected %v. got %v", fa.NewSet(2, 3), ns.NextState(fa.NewSet(1, 2), 'b'))
	}
	if !ns.NextState(fa.NewSet(3, 2), 'b').Equal(fa.NewSet(1, 2, 3)) {
		t.Errorf("ns.NextState() expected %v. got %v", fa.NewSet(1, 2, 3), ns.NextState(fa.NewSet(3, 2), 'b'))
	}
	if !ns.NextState(fa.NewSet(1, 3, 2), 'b').Equal(fa.NewSet(1, 2, 3)) {
		t.Errorf("ns.NextState() expected %v. got %v", fa.NewSet(1, 2, 3), ns.NextState(fa.NewSet(1, 3, 2), 'b'))
	}
	if !ns.NextState(fa.NewSet(1, 3, 2), 'a').Equal(fa.NewSet(1, 2)) {
		t.Errorf("ns.NextState() expected %v. got %v", fa.NewSet(1, 2), ns.NextState(fa.NewSet(1, 3, 2), 'a'))
	}
}

func TestNFASimulation_RulesFor(t *testing.T) {
	if len(ns.RulesFor(fa.NewSet(1, 2))) != 2 {
		t.Errorf("ns.RulesFor() length expected %d. got %d",
			2, len(ns.RulesFor(fa.NewSet(1, 2))))
	}
	if len(ns.RulesFor(fa.NewSet(3, 2))) != 1 {
		t.Errorf("ns.RulesFor() length expected %d. got %d",
			1, len(ns.RulesFor(fa.NewSet(3, 2))))
	}
}

func TestNFASimulation_DiscoverStatesAndRules(t *testing.T) {
	startStateSet := ns.NFA.CurrentStates()
	sss, rs := ns.DiscoverStatesAndRules(NewSSS(startStateSet))

	if !sss.equal(NewSSS(fa.NewSet(1, 2), fa.NewSet(2, 3), fa.NewSet(1, 2, 3))) {
		t.Errorf("ns.DiscoverStatesAndRules() expected %v. got %v",
			NewSSS(fa.NewSet(1, 2), fa.NewSet(2, 3), fa.NewSet(1, 2, 3)),
			sss)
	}

	if len(rs) != 5 {
		t.Errorf("ns.DiscoverStatesAndRules() rules length expected %d. got %d",
			5, len(rs))
	}
}

func TestNFASimulation_ToDFA(t *testing.T) {
	dfa := ns.ToDFA()

	if len(dfa.AcceptStates) != 2 {
		t.Errorf("ns.ToDFA().AcceptStates length expected %d. got %d", 2, len(dfa.AcceptStates))
	}
	if len(dfa.RuleBook.Rules) != 5 {
		t.Errorf("ns.ToDFA().Rules length expected %d. got %d", 5, len(dfa.RuleBook.Rules))
	}

	err := dfa.ReadString("aaa")
	if err != nil {
		t.Fatal(err)
	}
	if dfa.CanAccept() {
		t.Errorf("ns.ToDFA().CanAccept() expected %t. got %t", false, true)
	}

	dfa = ns.ToDFA()
	err = dfa.ReadString("aab")
	if err != nil {
		t.Fatal(err)
	}
	if !dfa.CanAccept() {
		t.Errorf("ns.ToDFA().CanAccept() expected %t. got %t", true, false)
	}

	dfa = ns.ToDFA()
	err = dfa.ReadString("bbbabb")
	if err != nil {
		t.Fatal(err)
	}
	if !dfa.CanAccept() {
		t.Errorf("ns.ToDFA().CanAccept() expected %t. got %t", true, false)
	}
}
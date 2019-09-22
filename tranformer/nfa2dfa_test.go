package tranformer

import (
	"github.com/zhongzc/lexerGen/fa"
	"github.com/zhongzc/lexerGen/fa/nfa"
	"testing"
)

var ns *NFASimulation

func init() {
	rb := &fa.RuleBook{Rules: []*fa.Rule{
		{1, fa.OneChar('a'), 1},
		{1, fa.OneChar('a'), 2},
		{1, fa.Free(), 2},
		{2, fa.OneChar('b'), 3},
		{3, fa.OneChar('b'), 1},
		{3, fa.Free(), 2},
	}}
	n := nfa.New(rb, fa.NewSet(1), fa.NewSet(3))
	ns = &NFASimulation{n}
}

func TestNFASimulation_NextState(t *testing.T) {
	expected := fa.NewSet(1, 2)
	got := ns.NextState(fa.NewSet(1, 2), fa.OneChar('a'))
	if !got.Equal(expected) {
		t.Errorf("ns.NextState() expected %v. got %v", expected, got)
	}
	expected = fa.NewSet(2, 3)
	got = ns.NextState(fa.NewSet(1, 2), fa.OneChar('b'))
	if !got.Equal(expected) {
		t.Errorf("ns.NextState() expected %v. got %v", expected, got)
	}
	expected = fa.NewSet(1, 2, 3)
	got = ns.NextState(fa.NewSet(3, 2), fa.OneChar('b'))
	if !got.Equal(expected) {
		t.Errorf("ns.NextState() expected %v. got %v", expected, got)
	}
	expected = fa.NewSet(1, 2, 3)
	got = ns.NextState(fa.NewSet(1, 3, 2), fa.OneChar('b'))
	if !got.Equal(expected) {
		t.Errorf("ns.NextState() expected %v. got %v", expected, got)
	}
	expected = fa.NewSet(1, 2)
	got = ns.NextState(fa.NewSet(1, 3, 2), fa.OneChar('a'))
	if !got.Equal(expected) {
		t.Errorf("ns.NextState() expected %v. got %v", expected, got)
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
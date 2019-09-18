package dfa

import (
	"github.com/zhongzc/lexerGen/fa"
	"testing"
)

var rb *RuleBook

func init() {
	rb = &RuleBook{Rules: []*fa.Rule{
		{1, 'a', 2},
		{1, 'b', 1},
		{2, 'a', 2},
		{2, 'b', 3},
		{3, 'a', 3},
		{3, 'b', 3},
	}}
}

func TestRuleBook_NextState(t *testing.T) {
	var s int
	var err error
	s, err = rb.nextState(1, 'a')
	if err != nil {
		t.Fatalf("rb.nextState() failed: %s", err)
	}
	if s != 2 {
		t.Errorf("rb.nextState() expected %d, got %d", 2, s)
	}

	s, err = rb.nextState(1, 'b')
	if err != nil {
		t.Fatalf("rb.nextState() failed: %s", err)
	}
	if s != 1 {
		t.Errorf("rb.nextState() expected %d, got %d", 1, s)
	}

	s, err = rb.nextState(2, 'b')
	if err != nil {
		t.Fatalf("rb.nextState() failed: %s", err)
	}
	if s != 3 {
		t.Errorf("rb.nextState() expected %d, got %d", 2, s)
	}
}

func TestDFA_CanAccept(t *testing.T) {
	dfa := &DFA{
		RuleBook:     rb,
		AcceptStates: map[int]bool{1: true, 2: true},
		CurrentState: 1,
	}
	if !dfa.CanAccept() {
		t.Errorf("dfa.CanAccept() expected %t, got %t", true, false)
	}

	dfa = &DFA{
		RuleBook:     rb,
		AcceptStates: map[int]bool{3: true},
		CurrentState: 1,
	}
	if dfa.CanAccept() {
		t.Errorf("dfa.CanAccept() expected %t, got %t", false, true)
	}
}

func TestRule_String(t *testing.T) {
	dfa := &DFA{
		RuleBook:     rb,
		AcceptStates: map[int]bool{3: true},
		CurrentState: 1,
	}
	if dfa.CanAccept() {
		t.Errorf("dfa.CanAccept() expected %t, got %t", false, true)
	}

	err := dfa.ReadString("baaab")
	if err != nil {
		t.Fatalf("dfa.ReadString() failed: %s", err)
	}
	if !dfa.CanAccept() {
		t.Errorf("dfa.CanAccept() expected %t, got %t", true, false)
	}
}

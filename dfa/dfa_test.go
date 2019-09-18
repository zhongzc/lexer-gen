package dfa

import "testing"

func TestRuleBook_NextState(t *testing.T) {
	rb := &RuleBook{Rules: []*Rule{
		{1, 'a', 2},
		{1, 'b', 1},
		{2, 'a', 2},
		{2, 'b', 3},
		{3, 'a', 3},
		{3, 'b', 3},
	}}

	var s int
	var err error
	s, err = rb.NextState(1, 'a')
	if err != nil {
		t.Fatalf("rb.NextState() failed: %s", err)
	}
	if s != 2 {
		t.Errorf("rb.NextState() expected %d, got %d", 2, s)
	}

	s, err = rb.NextState(1, 'b')
	if err != nil {
		t.Fatalf("rb.NextState() failed: %s", err)
	}
	if s != 1 {
		t.Errorf("rb.NextState() expected %d, got %d", 1, s)
	}

	s, err = rb.NextState(2, 'b')
	if err != nil {
		t.Fatalf("rb.NextState() failed: %s", err)
	}
	if s != 3 {
		t.Errorf("rb.NextState() expected %d, got %d", 2, s)
	}
}

func TestDFA_CanAccept(t *testing.T) {
	rb := &RuleBook{Rules: []*Rule{
		{1, 'a', 2},
		{1, 'b', 1},
		{2, 'a', 2},
		{2, 'b', 3},
		{3, 'a', 3},
		{3, 'b', 3},
	}}

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
	rb := &RuleBook{Rules: []*Rule{
		{1, 'a', 2},
		{1, 'b', 1},
		{2, 'a', 2},
		{2, 'b', 3},
		{3, 'a', 3},
		{3, 'b', 3},
	}}

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
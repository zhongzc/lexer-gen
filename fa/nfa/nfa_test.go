package nfa

import (
	. "github.com/zhongzc/lexerGen/fa"
	"testing"
)

var rb *RuleBook

func init() {
	rb = &RuleBook{Rules: []*Rule{
		{1, 'a', 1},
		{1, 'b', 1},
		{1, 'b', 2},
		{2, 'a', 3},
		{2, 'b', 3},
		{3, 'a', 4},
		{3, 'b', 4},
	}}
}

func TestNFA_CanAccept(t *testing.T) {
	var nfa *NFA
	nfa = &NFA{
		RuleBook:      rb,
		currentStates: NewSet(1),
		AcceptStates:  NewSet(4),
	}
	if nfa.CanAccept() {
		t.Errorf("nfa.CanAccept() expected %t. got %t", false, true)
	}

	nfa = &NFA{
		RuleBook:      rb,
		currentStates: NewSet(1, 2, 4),
		AcceptStates:  NewSet(4),
	}
	if !nfa.CanAccept() {
		t.Errorf("nfa.CanAccept() expected %t. got %t", true, false)
	}
}

func TestNFA_ReadChar(t *testing.T) {
	nfa := &NFA{
		RuleBook:      rb,
		currentStates: NewSet(1),
		AcceptStates:  NewSet(4),
	}
	if nfa.CanAccept() {
		t.Errorf("nfa.CanAccept() expected %t. got %t", false, true)
	}

	var err error
	err = nfa.ReadChar('b')
	if err != nil {
		t.Fatalf("nfa.ReadChar() failed: %s", err)
	}
	if nfa.CanAccept() {
		t.Errorf("nfa.CanAccept() expected %t. got %t", false, true)
	}

	err = nfa.ReadChar('a')
	if err != nil {
		t.Fatalf("nfa.ReadChar() failed: %s", err)
	}
	if nfa.CanAccept() {
		t.Errorf("nfa.CanAccept() expected %t. got %t", false, true)
	}

	err = nfa.ReadChar('b')
	if err != nil {
		t.Fatalf("nfa.ReadChar() failed: %s", err)
	}
	if !nfa.CanAccept() {
		t.Errorf("nfa.CanAccept() expected %t. got %t", true, false)
	}
}

func TestNFA_ReadString(t *testing.T) {
	nfa := &NFA{
		RuleBook:      rb,
		currentStates: NewSet(1),
		AcceptStates:  NewSet(4),
	}
	if nfa.CanAccept() {
		t.Errorf("nfa.CanAccept() expected %t. got %t", false, true)
	}

	var err error
	err = nfa.ReadString("bbbbb")
	if err != nil {
		t.Fatalf("nfa.ReadString() failed: %s", err)
	}
	if !nfa.CanAccept() {
		t.Errorf("nfa.CanAccept() expected %t. got %t", true, false)
	}
}

func TestRuleBook_NextStates(t *testing.T) {
	var ss StateSet
	var err error
	ss, err = rb.nextStates(NewSet(1), 'b')
	if err != nil {
		t.Fatalf("rb.nextState() failed: %s", err)
	}
	if !ss.Equal(NewSet(1, 2)) {
		t.Errorf("rb.nextStates() expected %v. got %v", NewSet(1, 2), ss)
	}

	ss, err = rb.nextStates(NewSet(1, 2), 'a')
	if err != nil {
		t.Fatalf("rb.nextState() failed: %s", err)
	}
	if !ss.Equal(NewSet(1, 3)) {
		t.Errorf("rb.nextStates() expected %v. got %v", NewSet(1, 3), ss)
	}

	ss, err = rb.nextStates(NewSet(1, 3), 'b')
	if err != nil {
		t.Fatalf("rb.nextState() failed: %s", err)
	}
	if !ss.Equal(NewSet(1, 2, 4)) {
		t.Errorf("rb.nextStates() expected %v. got %v", NewSet(1, 2, 4), ss)
	}
}

func TestRuleBook_FreeMove(t *testing.T) {
	rb := &RuleBook{Rules: []*Rule{
		{1, 0, 2},
		{1, 0, 4},
		{2, 'a', 3},
		{3, 'a', 2},
		{4, 'a', 5},
		{5, 'a', 6},
		{6, 'a', 4},
	}}
	ss, err := rb.nextStates(NewSet(1), 0)
	if err != nil {
		t.Fatalf("rb.nextState() failed: %s", err)
	}
	if !ss.Equal(NewSet(2, 4)) {
		t.Errorf("rb.nextStates() expected %v. got %v", NewSet( 2, 4), ss)
	}

	ss = rb.freeMove(NewSet(1))
	if !ss.Equal(NewSet(1, 2, 4)) {
		t.Errorf("rb.freeMove() expected %v. got %v", NewSet( 1, 2, 4), ss)
	}
}

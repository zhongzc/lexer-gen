package tranformer

import (
	"github.com/zhongzc/lexerGen/reparser"
	"testing"
)

func TestSequenceToNFA(t *testing.T) {
	re, err := reparser.Parse("ab")
	if err != nil {
		t.Fatal(err)
	}
	nfa := ToNFA(re)
	if !nfa.Match("ab") {
		t.Errorf("toNFA().Match() expected %t. got %t", true, false)
	}
	if nfa.Match("ac") {
		t.Errorf("toNFA().Match() expected %t. got %t", false, true)
	}
	dfa := ToDFA(nfa)
	if !dfa.Match("ab") {
		t.Errorf("toNFA().ToDFA().Match() expected %t. got %t", true, false)
	}
	if dfa.Match("ac") {
		t.Errorf("toNFA().Match() expected %t. got %t", false, true)
	}
}


func TestChooseToNFA(t *testing.T) {
	re, err := reparser.Parse("a|b|d")
	if err != nil {
		t.Fatal(err)
	}
	nfa := ToNFA(re)
	if !nfa.Match("a") {
		t.Errorf("toNFA().Match() expected %t. got %t", true, false)
	}
	if !nfa.Match("b") {
		t.Errorf("toNFA().Match() expected %t. got %t", true, false)
	}
	if nfa.Match("c") {
		t.Errorf("toNFA().Match() expected %t. got %t", false, true)
	}
	dfa := ToDFA(nfa)
	if !dfa.Match("a") {
		t.Errorf("toNFA().ToDFA().Match() expected %t. got %t", true, false)
	}
	if !dfa.Match("b") {
		t.Errorf("toNFA().ToDFA().Match() expected %t. got %t", true, false)
	}
	if dfa.Match("c") {
		t.Errorf("toNFA().ToDFA().Match() expected %t. got %t", false, true)
	}
}

func TestRepeatToNFA(t *testing.T) {
	re, err := reparser.Parse("a*")
	if err != nil {
		t.Fatal(err)
	}
	nfa := ToNFA(re)
	if !nfa.Match("") {
		t.Errorf("toNFA().Match() expected %t. got %t", true, false)
	}
	if !nfa.Match("a") {
		t.Errorf("toNFA().Match() expected %t. got %t", true, false)
	}
	if !nfa.Match("aaa") {
		t.Errorf("toNFA().Match() expected %t. got %t", true, false)
	}
	if nfa.Match("b") {
		t.Errorf("toNFA().Match() expected %t. got %t", false, true)
	}
	if nfa.Match("c") {
		t.Errorf("toNFA().Match() expected %t. got %t", false, true)
	}
	dfa := ToDFA(nfa)
	if !dfa.Match("") {
		t.Errorf("toNFA().ToDFA().Match() expected %t. got %t", true, false)
	}
	if !dfa.Match("a") {
		t.Errorf("toNFA().ToDFA().Match() expected %t. got %t", true, false)
	}
	if !dfa.Match("aaa") {
		t.Errorf("toNFA().ToDFA().Match() expected %t. got %t", true, false)
	}
	if dfa.Match("b") {
		t.Errorf("toNFA().ToDFA().Match() expected %t. got %t", false, true)
	}
	if dfa.Match("c") {
		t.Errorf("toNFA().ToDFA().Match() expected %t. got %t", false, true)
	}
}

func TestToNFA(t *testing.T) {
	re, err := reparser.Parse("abc*((cat)*|dog)")
	if err != nil {
		t.Fatal(err)
	}
	nfa := ToNFA(re)
	if !nfa.Match("abdog") {
		t.Errorf("toNFA().Match() expected %t. got %t", true, false)
	}
	if !nfa.Match("abccccatcat") {
		t.Errorf("toNFA().Match() expected %t. got %t", true, false)
	}
	dfa := ToDFA(nfa)
	if !dfa.Match("abdog") {
		t.Errorf("toNFA().ToDFA().Match() expected %t. got %t", true, false)
	}
	if !dfa.Match("abccccatcat") {
		t.Errorf("toNFA().ToDFA().Match() expected %t. got %t", true, false)
	}
}
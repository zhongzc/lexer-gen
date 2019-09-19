package tranformer

import (
	"github.com/zhongzc/lexerGen/reparser"
	"testing"
)

func TestToNFA(t *testing.T) {
	re, err := reparser.Parse("ab")
	if err != nil {
		t.Fatal(err)
	}
	nfa := ToNFA(re, NewIdGenerator(0))
	err = nfa.ReadString("ab")
	if err != nil {
		t.Fatal(err)
	}
	if !nfa.CanAccept() {
		t.Errorf("ToNFA().CanAccept() expected %t. got %t", true, false)
	}

	nfa = ToNFA(re, NewIdGenerator(0))
	dfa := ToDFA(nfa)
	err = dfa.ReadString("ab")
	if err != nil {
		t.Fatal(err)
	}
	if !dfa.CanAccept() {
		t.Errorf("ToNFA().ToDFA().CanAccept() expected %t. got %t", true, false)
	}
}

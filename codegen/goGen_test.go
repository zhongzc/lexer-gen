package codegen

import (
	"github.com/zhongzc/lexerGen/fa"
	"github.com/zhongzc/lexerGen/fa/dfa"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	rb := &fa.RuleBook{Rules: []*fa.Rule{
		{1, fa.OneChar('a'), 2},
		{1, fa.OneChar('b'), 1},
		{2, fa.OneChar('a'), 2},
		{2, fa.OneChar('b'), 3},
		{3, fa.OneChar('a'), 3},
		{3, fa.OneChar('b'), 3},
	}}
	d := dfa.New(rb, 1, fa.NewSet(3))

	err := Gen(&GoGen{}, "lexer", []*NamedDFA{
		{Name: "A", DFA: d},
		{Name: "B", DFA: d},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = os.RemoveAll("lexer")
	if err != nil {
		t.Fatal(err)
	}
}

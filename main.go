package main

import (
	"github.com/zhongzc/lexerGen/codegen"
	"github.com/zhongzc/lexerGen/reparser"
	"github.com/zhongzc/lexerGen/tranformer"
)

func main() {
	re := "abc*((cat)*|dog)"
	ast, _ := reparser.Parse(re)
	nfa := tranformer.ToNFA(ast)
	dfa := tranformer.ToDFA(nfa)
	_ = codegen.Gen(&codegen.GoGen{}, "tmplexer", []*codegen.NamedDFA{
		{Name: "DOG", DFA: dfa},
	})
}

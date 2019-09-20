package ruledef

import (
	"github.com/zhongzc/lexerGen/codegen"
	"github.com/zhongzc/lexerGen/reast"
	"github.com/zhongzc/lexerGen/reparser"
	"github.com/zhongzc/lexerGen/tranformer"
	"io"
)

func BuildAll(input io.Reader, outputPath string, g codegen.Generator) (err error) {
	re := Parse(input)

	var namedDFA []*codegen.NamedDFA
	var rex reast.RegEx
	for _, r := range re {
		rex, err = reparser.Parse(r.RE)
		if err != nil {
			return
		}
		nfa := tranformer.ToNFA(rex)
		dfa := tranformer.ToDFA(nfa)
		namedDFA = append(namedDFA, &codegen.NamedDFA{Name: r.Name, DFA: dfa})
	}

	err = codegen.Gen(g, outputPath, namedDFA)
	return
}

package lexerGen

import (
	"bufio"
	"bytes"
	"github.com/zhongzc/lexerGen/codegen"
	"github.com/zhongzc/lexerGen/fa"
	"github.com/zhongzc/lexerGen/reast"
	"github.com/zhongzc/lexerGen/reparser"
	"github.com/zhongzc/lexerGen/tranformer"
	"io"
	"unicode"
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

		rulx := make(map[int][]struct {
			By fa.Charset
			To int
		})

		for _, r := range dfa.RuleBook.Rules {
			rulx[r.From] = append(rulx[r.From], struct {
				By fa.Charset
				To int
			}{By: r.By, To: r.To})
		}

		namedDFA = append(namedDFA, &codegen.NamedDFA{
			Name:         r.Name,
			StartState:   dfa.CurrentState,
			AcceptStates: dfa.AcceptStates,
			Rules:        rulx,
		})
	}

	err = codegen.Gen(g, outputPath, namedDFA)
	return
}

type NamedRE struct {
	Name string
	RE   string
}

func Parse(reader io.Reader) (res []*NamedRE) {
	rd := bufio.NewReader(reader)
	var err error
	var buf []byte
	for err == nil {
		buf, _, err = rd.ReadLine()
		buf = bytes.TrimSpace(buf)
		if len(buf) == 0 {
			continue
		}

		i := bytes.IndexFunc(buf, unicode.IsSpace)
		name := buf[:i]
		buf = bytes.TrimSpace(buf[i:])
		res = append(res, &NamedRE{string(name), string(buf)})
	}
	return
}

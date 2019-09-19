package codegen

import (
	"github.com/zhongzc/lexerGen/fa/dfa"
	"io"
	"os"
)

type Generator interface {
	Generate(writer io.Writer, lxs []*Lexer)
}

func Gen(g Generator, filename string, lxs []*Lexer) (err error) {
	err = os.Mkdir("lexer", 0755)
	if err != nil {
		return
	}

	var f *os.File
	f, err = os.OpenFile("lexer\\"+filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	g.Generate(f, lxs)
	return
}

type Lexer struct {
	Name string
	DFA  *dfa.DFA
}

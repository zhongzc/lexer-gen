package codegen

import (
	"github.com/zhongzc/lexerGen/fa/dfa"
	"io"
	"os"
	"path/filepath"
)

type Generator interface {
	Generate(dfas []*NamedDFA) map[string]func(io.Writer) error
}

func Gen(g Generator, path string, dfas []*NamedDFA) (err error) {
	err = os.Mkdir(path, 0755)
	if err != nil {
		return
	}
	for filename, w := range g.Generate(dfas) {
		var f *os.File
		f, err = os.OpenFile(filepath.Join(path, filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return
		}
		err = w(f)
		if err != nil {
			return
		}
	}

	return
}

type NamedDFA struct {
	Name string
	DFA  *dfa.DFA
}

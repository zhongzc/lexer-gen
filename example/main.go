package main

import (
	"github.com/zhongzc/lexerGen/codegen"
	"github.com/zhongzc/lexerGen/ruledef"
	"os"
)

func main() {
	file, err := os.Open("example/tmp.lx")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = ruledef.BuildAll(file, "example/_tmpLexer", codegen.NewGoGen("lexer"))
	if err != nil {
		panic(err)
	}
}

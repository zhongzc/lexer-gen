package main

import (
	"github.com/zhongzc/lexerGen"
	"github.com/zhongzc/lexerGen/codegen"
	"os"
)

func main() {
	file, err := os.Open("example/scutlab1/sample.lx")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = lexerGen.BuildAll(file, "example/scutlab1/sample/lexer", codegen.NewGoGen("lexer"))
	if err != nil {
		panic(err)
	}
}

package main

import (
	"github.com/zhongzc/lexerGen"
	"github.com/zhongzc/lexerGen/codegen"
	"os"
)

func main() {
	file, err := os.Open("example/simpleCpp/tmp.lx")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = lexerGen.BuildAll(file, "example/simpleCpp/_tmpLexer", codegen.NewCppGen())
	if err != nil {
		panic(err)
	}
}

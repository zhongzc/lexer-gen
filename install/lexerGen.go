package main

import (
	"flag"
	"github.com/zhongzc/lexerGen"
	"github.com/zhongzc/lexerGen/codegen"
	"os"
)

func main() {
	f := flag.String("i", "", "the input rule file")
	p := flag.String("o", "lexer", "the output path")
	n := flag.String("go", "", "generate golang code, need specify package name")
	flag.Parse()

	if *f == "" {
		panic("should specify the input file")
	}
	file, err := os.Open(*f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var g codegen.Generator
	if *n != "" {
		g = codegen.NewGoGen(*n)
	}
	if g == nil {
		panic("should specify a code generator")
	}

	err = lexerGen.BuildAll(file, *p, g)
	if err != nil {
		panic(err)
	}
}

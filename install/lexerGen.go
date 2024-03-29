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
	goo := flag.String("go", "", "generate golang code, need specify package name")
	cpp := flag.Bool("cpp", false, "generate c++11 code")
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
	if *goo != "" {
		g = codegen.NewGoGen(*goo)
	}
	if *cpp {
		g = codegen.NewCppGen()
	}
	if g == nil {
		panic("should specify a code generator")
	}

	err = lexerGen.BuildAll(file, *p, g)
	if err != nil {
		panic(err)
	}
}

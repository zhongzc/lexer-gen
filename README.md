# lexerGen

**LexerGen** is a code generation tool that converts regular expression rules into Golang version of Lexer code. Also, with custom generation rules, users can customize the Go code that generates other languages, see: `codogen.Generator`.

## Usage

- Specify a rule input file. The left side of each rule is the name, and the right side is the regular expression. Like 
`input.lx`:
```text
NUM     0|[1-9][0-9]*
IDENT   [A-Za-z_][A-Za-z0-9_]*
COMMENT /\*([^\*]*\*[^/])*[^\*]*\*/
``` 

- Run the command line program `lexerGen`, in Windows, `lexerGen.exe`:
```shell
./lexerGen -i input.lx -o lexer -go lexer
```
Usage of `./lexerGen`:
```text
-cpp
    generate c++11 code
-go string
    generate golang code, need specify package name
-i string
    the input rule file
-o string
    the output path (default "lexer")
```

- Get the generated code from the specified path.
```
$ tree lexer
lexer
├── automata.go
├── charIterator.go
├── lexer.go
└── lexer_test.go
```

## Feature

### v0.3

- Support generating C++11 code.
- Fix some bugs.

### v0.2

- Significantly reduce the amount of generated code.
- Supported regular expression:
    - negated group


### v0.1

- Can specify **multiple** regular expression rules for the lexer.
- Support UTF-8.
- The generated lexer has well-defined APIs.
- The codebase is self-explanatory and organized to be learner friendly.
- Easy to customize to generate codes for other languages.
- Supported regular expression:
    - alternation
    - concatenate
    - quantifier
    - character set

## TODO

- Support escape group, like `\w`,`\d`,`\s`...

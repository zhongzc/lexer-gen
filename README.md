# lexerGen

**LexerGen** is a code generation tool that converts regular expression rules into Golang version of Lexer code. Also, with custom generation rules, users can customize the Go code that generates other languages, see: `codogen.Generator`.

## v0.1

### Feature

- Can specify **multiple** regular expression rules for the lexer.
- The generated lexer has well-defined APIs.
- The codebase is self-explanatory and organized to be learner friendly.
- Easy to customize to generate codes for other languages.


## TODO

- Reduce the amount of codes generated.
- Supports more complete version of regular expression.

package reast

import (
	"bytes"
	"fmt"
)

// RegEx :
//   AST(Abstract Syntax Tree) Node for constructing Regular
//   Expression's AST;
//   REString() returns string represented by the AST
type RegEx interface {
	REString() string
}

// Choose :
//   Matches one of regular expression in `RegExs`;
//   It represents "r0|r1|r2|...|rn"
type Choose struct {
	RegExs []RegEx
}

func (c *Choose) REString() string {
	if len(c.RegExs) == 0 {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteString("(")
	buf.WriteString(c.RegExs[0].REString())
	for _, cc := range c.RegExs[1:] {
		buf.WriteString("|")
		buf.WriteString(cc.REString())
	}
	buf.WriteString(")")
	return buf.String()
}

// Repeat :
//   Matches zero or more regular expressions represented
//   by `RegEx`;
//   It represents "r*"
type Repeat struct {
	RegEx RegEx
}

func (r *Repeat) REString() string {
	return fmt.Sprintf("(%s)*", r.RegEx.REString())
}

// Sequence :
//   Matches all regular expressions concatenated Sequentially
//   in `RegExs`;
//   It represents "r1r2r3...rn"
type Sequence struct {
	RegExs []RegEx
}

func (c *Sequence) REString() string {
	var buf bytes.Buffer
	for _, g := range c.RegExs {
		buf.WriteString(g.REString())
	}
	return buf.String()
}

// Primitive :
//   Matches a utf-8 character equals to `Rune`
type Primitive struct {
	Rune rune
}

func (l *Primitive) REString() string {
	return string(l.Rune)
}

package rgex

import (
	"bytes"
	"fmt"
)

// ASTNode :
//   AST(Abstract Syntax Tree) Node for constructing Regular
//   Expression's AST;
//   ToRegex() returns string represented by the AST
type ASTNode interface {
	ToRegex() string
}

// Choose :
//   Matches one of regular expression in `Concats`;
//   It represents "c0 | c1 | c2 | ... | cn"
type Choose struct {
	Concats []Concat
}
func (c Choose) ToRegex() string {
	if len(c.Concats) == 0 {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteString(c.Concats[0].ToRegex())
	for _, cc := range c.Concats[1:] {
		buf.WriteString("|")
		buf.WriteString(cc.ToRegex())
	}
	return buf.String()
}

// Repeat :
//   Matches zero or more regular expressions represented
//   by `Choose`;
//   It represents "c*"
type Repeat struct {
	Choose Choose
}
func (r Repeat) ToRegex() string {
	return fmt.Sprintf("(%s)*", r.Choose.ToRegex())
}

// Concat :
//   Matches all regular expressions concatenated Sequentially
//   in `Grounds`;
//   It represents "c*"
type Concat struct {
	Grounds []Ground
}
func (c Concat) ToRegex() string {
	var buf bytes.Buffer
	for _, g := range c.Grounds {
		buf.WriteString(g.ToRegex())
	}
	return buf.String()
}

// CharacterSet :
//   Matches one utf-8 character in one of character ranges;
//   It represents "[r0-r1 r2-r3 r4-r5 ... rn-rm]"
type CharacterSet struct {
	Ranges []Range
}
func (cs CharacterSet) ToRegex() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for _, r := range cs.Ranges {
		buf.WriteString(fmt.Sprintf("%c-%c", r.From, r.To))
	}
	buf.WriteString("]")
	return buf.String()
}
type Range struct {
	From rune
	To rune
}

// Literal :
//   Matches a utf-8 character equals to `Rune`
type Literal struct {
	Rune rune
}
func (l Literal) ToRegex() string {
	return string(l.Rune)
}

// Ground :
//   Just for sub-typing
type Ground interface {
	ASTNode
	GType() string
}
func (Repeat) GType() string {
	return "repeat"
}
func (CharacterSet) GType() string {
	return "characterSet"
}
func (Literal) GType() string {
	return "literal"
}

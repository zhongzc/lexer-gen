package ast

import (
	"bytes"
	"fmt"
)

// ASTNode :
//   AST(Abstract Syntax Tree) Node for constructing Regular
//   Expression's AST;
//   ToRegex() returns string represented by the AST
type Node interface {
	ToRegex() string
}

// Choose :
//   Matches one of regular expression in `Concats`;
//   It represents "c0 | c1 | c2 | ... | cn"
type Choose struct {
	nodes []Node
}

func (c *Choose) ToRegex() string {
	if len(c.nodes) == 0 {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteString("(")
	buf.WriteString(c.nodes[0].ToRegex())
	for _, cc := range c.nodes[1:] {
		buf.WriteString("|")
		buf.WriteString(cc.ToRegex())
	}
	buf.WriteString(")")
	return buf.String()
}

// Repeat :
//   Matches zero or more regular expressions represented
//   by `Choose`;
//   It represents "c*"
type Repeat struct {
	node Node
}

func (r *Repeat) ToRegex() string {
	return fmt.Sprintf("(%s)*", r.node.ToRegex())
}

// Concat :
//   Matches all regular expressions concatenated Sequentially
//   in `Grounds`;
//   It represents "c*"
type Concat struct {
	nodes []Node
}

func (c *Concat) ToRegex() string {
	var buf bytes.Buffer
	for _, g := range c.nodes {
		buf.WriteString(g.ToRegex())
	}
	return buf.String()
}

// Literal :
//   Matches a utf-8 character equals to `Rune`
type Literal struct {
	Rune rune
}

func (l *Literal) ToRegex() string {
	return string(l.Rune)
}

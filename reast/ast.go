package reast

import (
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
	Left  RegEx
	Right RegEx
}

func NewChoose(left RegEx, right RegEx) *Choose {
	return &Choose{Left: left, Right: right}
}

func (c *Choose) REString() string {
	return "(" + c.Left.REString() + "|" + c.Right.REString() + ")"
}

// Repeat :
//   Matches zero or more regular expressions represented
//   by `RegEx`;
//   It represents "r*"
type Repeat struct {
	RegEx RegEx
}

func NewRepeat(regEx RegEx) *Repeat {
	return &Repeat{RegEx: regEx}
}

func (r *Repeat) REString() string {
	return fmt.Sprintf("(%s)*", r.RegEx.REString())
}

// Sequence :
//   Matches all regular expressions concatenated Sequentially
//   in `RegExs`;
//   It represents "r1r2r3...rn"
type Sequence struct {
	Left  RegEx
	Right RegEx
}

func NewSequence(left RegEx, right RegEx) *Sequence {
	return &Sequence{Left: left, Right: right}
}

func (s *Sequence) REString() string {
	return s.Left.REString() + s.Right.REString()
}

// Primitive :
//   Matches a utf-8 character equals to `Rune`
type Primitive struct {
	From rune
	To   rune
}

func NewPrimitive(from rune, to rune) *Primitive {
	return &Primitive{From: from, To: to}
}

func (l *Primitive) REString() string {
	if l.To == l.From {
		return string(l.From)
	}
	return fmt.Sprintf("%c-%c", l.From, l.To)
}

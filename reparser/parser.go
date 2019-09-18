package reparser

import (
	"errors"
	"fmt"
	. "github.com/zhongzc/lexerGen/reast"
)

type Parser struct {
	input  []rune
	curPos int
}

func Parse(input string) (re RegEx, err error) {
	p := Parser{input: []rune(input), curPos: -1}
	return p.regex()
}

func (p *Parser) peek() rune {
	if p.curPos+1 >= len(p.input) {
		return 0
	}
	return p.input[p.curPos+1]
}

func (p *Parser) cur() rune {
	return p.input[p.curPos]
}

func (p *Parser) eat(r rune) error {
	if r == p.peek() {
		p.curPos++
		return nil
	}

	msg := fmt.Sprintf("Parser.eat() expected: %s, but got %s", show(r), show(p.peek()))
	return errors.New(msg)
}

// regex := <term> '|' <regex> | <term>
func (p *Parser) regex() (re RegEx, err error) {
	var t RegEx
	t, err = p.term()
	if err != nil {
		return
	}

	c := &Choose{RegExs: []RegEx{t}}
	for ; p.peek() == '|'; {
		_ = p.eat('|')
		t, err = p.term()
		if err != nil {
			return
		}
		c.RegExs = append(c.RegExs, t)
	}
	if len(c.RegExs) == 1 {
		re = c.RegExs[0]
	} else {
		re = c
	}
	return
}

// term := <factor> <term> | <factor>
func (p *Parser) term() (re RegEx, err error) {
	var f RegEx
	f, err = p.factor()
	if err != nil {
		return
	}

	s := &Sequence{RegExs: []RegEx{f}}
	for ; p.peek() != 0 && p.peek() != ')' && p.peek() != '|'; {
		f, err = p.factor()
		if err != nil {
			return
		}

		s.RegExs = append(s.RegExs, f)
	}

	if len(s.RegExs) == 1 {
		re = s.RegExs[0]
	} else {
		re = s
	}
	return
}

// factor := <base> | <base> '*'
func (p *Parser) factor() (re RegEx, err error) {
	var b RegEx
	b, err = p.base()
	if err != nil {
		return
	}

	if p.peek() == '*' {
		_ = p.eat('*')
		b = &Repeat{RegEx: b}
	}
	return b, nil
}

// base := '('   <regex>  )'  |
//         '[' <charsets> ']' |
//         <primitive>
func (p *Parser) base() (re RegEx, err error) {
	switch p.peek() {
	case '(':
		_ = p.eat('(')
		re, err = p.regex()
		if err != nil {
			return
		}
		err = p.eat(')')
		return
	case '\\':
		_ = p.eat('\\')
		c := p.peek()
		_ = p.eat(c)
		return &Primitive{Rune: c}, nil
	case '[':
		_ = p.eat('[')
		re, err = p.charsets()
		if err != nil {
			return
		}
		err = p.eat(']')
		return
	default:
		c := p.peek()
		if c == 0 {
			return nil, errors.New("unexpected: EOF")
		} else if c == '|' || c == ')' || c == '*' || c == ']' {
			return nil, errors.New(fmt.Sprintf("expected: \\%s, but got %s", show(c), show(c)))
		}
		_ = p.eat(c)
		return &Primitive{Rune: c}, nil
	}
}

// charsets := <primitive> '-' <primitive> <charsets> |
//             <primitive> '-' <primitive>            |
//             <primitive>                 <charsets> |
//             <primitive>                            |
func (p *Parser) charsets() (re RegEx, err error) {
	res := make([]RegEx, 0)

	for ; p.peek() != ']'; {
		var f rune
		var t rune

		f, err = eatEscOrOrd(p)
		if err != nil {
			return
		}
		if p.peek() == '-' {
			_ = p.eat('-')
			t, err = eatEscOrOrd(p)
			if err != nil {
				return
			}
		} else {
			t = f
		}

		for i := f; i <= t; i++ {
			res = append(res, &Primitive{Rune: i})
		}
	}

	return &Choose{RegExs: res}, nil
}

// primitive := '\' . | .
func eatEscOrOrd(p *Parser) (r rune, err error) {
	if p.peek() == '\\' {
		_ = p.eat('\\')
	}
	r = p.peek()
	if r == 0 {
		err = errors.New("unexpected: EOF")
		return
	}
	_ = p.eat(r)
	return
}

func show(r rune) string {
	if r == 0 {
		return "EOF"
	} else {
		return string(r)
	}
}

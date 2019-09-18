package reparser

import (
	"fmt"
	. "github.com/zhongzc/lexerGen/reast"
)

type Parser struct {
	input  []rune
	curPos int
}

func Parse(input string) RegEx {
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

func (p *Parser) eat(r rune) {
	if r == p.peek() {
		p.curPos++
	} else {
		panic(fmt.Sprintf("Parser.eat() expected: %s, but got %s", show(r), show(p.peek())))
	}
}

// regex := <term> '|' <regex> | <term>
func (p *Parser) regex() RegEx {
	term := p.term()

	c := &Choose{RegExs: []RegEx{term}}
	for ; p.peek() == '|'; {
		p.eat('|')
		c.RegExs = append(c.RegExs, p.term())
	}
	if len(c.RegExs) == 1 {
		return c.RegExs[0]
	}
	return c
}

// term := <factor> <term> | <factor>
func (p *Parser) term() RegEx {
	factor := p.factor()

	s := &Sequence{RegExs: []RegEx{factor}}
	for ; p.peek() != 0 && p.peek() != ')' && p.peek() != '|'; {
		factor = p.factor()
		s.RegExs = append(s.RegExs, factor)
	}

	if len(s.RegExs) == 1 {
		return s.RegExs[0]
	}
	return s
}

// factor := <base> | <base> '*'
func (p *Parser) factor() RegEx {
	base := p.base()

	if p.peek() == '*' {
		p.eat('*')
		base = &Repeat{RegEx: base}
	}
	return base
}

// base := '('   <regex>  )'  |
//         '[' <charsets> ']' |
//         <primitive>
func (p *Parser) base() RegEx {
	switch p.peek() {
	case '(':
		p.eat('(')
		r := p.regex()
		p.eat(')')
		return r
	case '\\':
		p.eat('\\')
		c := p.peek()
		p.eat(c)
		return &Primitive{Rune: c}
	case '[':
		p.eat('[')
		cs := p.charsets()
		p.eat(']')
		return cs
	default:
		c := p.peek()
		if c == 0 {
			panic("unexpected: EOF")
		} else if c == '|' || c == ')' || c == '*' || c == ']' {
			panic(fmt.Sprintf("expected: \\%s, but got %s", show(c), show(c)))
		}
		p.eat(c)
		return &Primitive{Rune: c}
	}
}

// charsets := <primitive> '-' <primitive> <charsets> |
//             <primitive> '-' <primitive>            |
//             <primitive>                 <charsets> |
//             <primitive>                            |
func (p *Parser) charsets() RegEx {
	res := make([]RegEx, 0)

	for ; p.peek() != ']'; {
		var f rune
		var t rune

		f = eatEscOrOrd(p)
		if p.peek() == '-' {
			p.eat('-')
			t = eatEscOrOrd(p)
		} else {
			t = f
		}

		for i := f; i <= t; i++ {
			res = append(res, &Primitive{Rune: i})
		}
	}

	return &Choose{RegExs: res}
}

// primitive := '\' . | .
func eatEscOrOrd(p *Parser) (r rune) {
	if p.peek() == '\\' {
		p.eat('\\')
		r = p.peek()
		p.eat(r)
	} else {
		r = p.peek()
		p.eat(r)
	}
	return
}

func show(r rune) string {
	if r == 0 {
		return "EOF"
	} else {
		return string(r)
	}
}
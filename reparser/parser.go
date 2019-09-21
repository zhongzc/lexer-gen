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
func (p *Parser) regex() (RegEx, error) {
	t, err := p.term()
	if err != nil {
		return nil, errors.New("p.regex() expected a term\n" + err.Error())
	}

	if p.peek() == '|' {
		_ = p.eat('|')
		r, err := p.regex()
		if err != nil {
			return nil, errors.New("p.regex() expected a regex\n" + err.Error())
		}
		return NewChoose(t, r), nil
	} else {
		return t, nil
	}
}

// term := <factor> <term> | <factor>
func (p *Parser) term() (RegEx, error) {
	var f RegEx
	f, err := p.factor()
	if err != nil {
		return nil, errors.New("p.term() expected a factor\n" + err.Error())
	}

	if p.peek() != 0 && p.peek() != ')' && p.peek() != '|' {
		r, err := p.term()
		if err != nil {
			return nil, errors.New("p.term() expected a term\n" + err.Error())
		}
		return NewSequence(f, r), nil
	} else {
		return f, nil
	}
}

// factor := <base> | <base> '*'
func (p *Parser) factor() (RegEx, error) {
	b, err := p.base()
	if err != nil {
		return nil, errors.New("p.factor() expected a base\n" + err.Error())
	}

	if p.peek() == '*' {
		_ = p.eat('*')
		b = NewRepeat(b)
	}
	return b, nil
}

// base := '('   <regex>  )'  |
//         '[' <charsets> ']' |
//         <primitive>
func (p *Parser) base() (RegEx, error) {
	switch p.peek() {
	case '(':
		_ = p.eat('(')
		re, err := p.regex()
		if err != nil {
			return nil, errors.New("p.base() expected a regex\n" + err.Error())
		}
		err = p.eat(')')
		if err != nil {
			return nil, err
		}
		return re, nil
	case '\\':
		_ = p.eat('\\')
		c := p.peek()
		_ = p.eat(c)
		return NewPrimitive(c, c), nil
	case '[':
		_ = p.eat('[')
		re, err := p.charsets()
		if err != nil {
			return nil, errors.New("p.base() expected a charsets\n" + err.Error())
		}
		err = p.eat(']')
		if err != nil {
			return nil, err
		}
		return re, nil
	default:
		c := p.peek()
		if c == 0 {
			return nil, errors.New("unexpected: EOF")
		} else if c == '|' || c == ')' || c == '*' || c == ']' {
			return nil, errors.New(fmt.Sprintf("expected: \\%s, but got %s", show(c), show(c)))
		}
		_ = p.eat(c)
		return NewPrimitive(c, c), nil
	}
}

// charsets := <primitive> '-' <primitive> <charsets> |
//             <primitive> '-' <primitive>            |
//             <primitive>                 <charsets> |
//             <primitive>                            |
func (p *Parser) charsets() (res RegEx, err error) {
	for ; p.peek() != ']'; {
		var f rune
		var t rune

		f, err = eatEscOrOrd(p)
		if err != nil {
			return nil, errors.New("p.charsets(): failed\n" + err.Error())
		}
		if p.peek() == '-' {
			_ = p.eat('-')
			t, err = eatEscOrOrd(p)
			if err != nil {
				return nil, errors.New("p.charsets(): failed\n" + err.Error())
			}
		} else {
			t = f
		}

		if f <= t {
			if res == nil {
				res = NewPrimitive(f, t)
			} else {
				res = NewChoose(res, NewPrimitive(f, t))
			}
		} else {
			return nil, errors.New("p.charsets(): rhs can not be greater than lhs")
		}
	}

	return
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

// just for display EOF
func show(r rune) string {
	if r == 0 {
		return "EOF"
	} else {
		return string(r)
	}
}

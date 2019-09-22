package lexer

import (
	"errors"
	"fmt"
)

// Automata
type Automata interface {
	RunGreedy(iter CharIterator) error
}

type IF struct{}
func (a *IF) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		2: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 1:
			switch {
			case c == 'f':
				currentState = 2
			default:
				break outer
			}
		case 0:
			switch {
			case c == 'i':
				currentState = 1
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type ELSE struct{}
func (a *ELSE) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		4: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 0:
			switch {
			case c == 'e':
				currentState = 1
			default:
				break outer
			}
		case 1:
			switch {
			case c == 'l':
				currentState = 2
			default:
				break outer
			}
		case 2:
			switch {
			case c == 's':
				currentState = 3
			default:
				break outer
			}
		case 3:
			switch {
			case c == 'e':
				currentState = 4
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type LPAREN struct{}
func (a *LPAREN) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		1: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 0:
			switch {
			case c == '(':
				currentState = 1
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type RPAREN struct{}
func (a *RPAREN) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		1: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 0:
			switch {
			case c == ')':
				currentState = 1
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type LBRACKET struct{}
func (a *LBRACKET) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		1: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 0:
			switch {
			case c == '{':
				currentState = 1
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type RBRACKET struct{}
func (a *RBRACKET) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		1: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 0:
			switch {
			case c == '}':
				currentState = 1
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type SEMICOLON struct{}
func (a *SEMICOLON) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		1: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 0:
			switch {
			case c == ';':
				currentState = 1
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type VAR struct{}
func (a *VAR) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		3: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 0:
			switch {
			case c == 'v':
				currentState = 1
			default:
				break outer
			}
		case 1:
			switch {
			case c == 'a':
				currentState = 2
			default:
				break outer
			}
		case 2:
			switch {
			case c == 'r':
				currentState = 3
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type RETURN struct{}
func (a *RETURN) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		6: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 2:
			switch {
			case c == 't':
				currentState = 3
			default:
				break outer
			}
		case 3:
			switch {
			case c == 'u':
				currentState = 4
			default:
				break outer
			}
		case 4:
			switch {
			case c == 'r':
				currentState = 5
			default:
				break outer
			}
		case 5:
			switch {
			case c == 'n':
				currentState = 6
			default:
				break outer
			}
		case 0:
			switch {
			case c == 'r':
				currentState = 1
			default:
				break outer
			}
		case 1:
			switch {
			case c == 'e':
				currentState = 2
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type EQ struct{}
func (a *EQ) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		2: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 0:
			switch {
			case c == '=':
				currentState = 1
			default:
				break outer
			}
		case 1:
			switch {
			case c == '=':
				currentState = 2
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type ASSIGN struct{}
func (a *ASSIGN) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		1: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 0:
			switch {
			case c == '=':
				currentState = 1
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type NUM struct{}
func (a *NUM) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		3: true,
		1: true,
		2: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 0:
			switch {
			case c == '0':
				currentState = 1
			case '1' <= c && c <= '9':
				currentState = 2
			default:
				break outer
			}
		case 2:
			switch {
			case '0' <= c && c <= '9':
				currentState = 3
			default:
				break outer
			}
		case 3:
			switch {
			case '0' <= c && c <= '9':
				currentState = 3
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type IDENT struct{}
func (a *IDENT) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
		7: true,
		1: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 7:
			switch {
			case 'a' <= c && c <= 'z':
				currentState = 5
			case '0' <= c && c <= '9':
				currentState = 7
			case 'A' <= c && c <= 'Z':
				currentState = 4
			case c == '_':
				currentState = 6
			default:
				break outer
			}
		case 1:
			switch {
			case 'a' <= c && c <= 'z':
				currentState = 5
			case 'A' <= c && c <= 'Z':
				currentState = 4
			case c == '_':
				currentState = 6
			case '0' <= c && c <= '9':
				currentState = 7
			default:
				break outer
			}
		case 6:
			switch {
			case c == '_':
				currentState = 6
			case '0' <= c && c <= '9':
				currentState = 7
			case 'a' <= c && c <= 'z':
				currentState = 5
			case 'A' <= c && c <= 'Z':
				currentState = 4
			default:
				break outer
			}
		case 4:
			switch {
			case '0' <= c && c <= '9':
				currentState = 7
			case c == '_':
				currentState = 6
			case 'a' <= c && c <= 'z':
				currentState = 5
			case 'A' <= c && c <= 'Z':
				currentState = 4
			default:
				break outer
			}
		case 0:
			switch {
			case 'A' <= c && c <= 'Z':
				currentState = 2
			case c == '_':
				currentState = 1
			case 'a' <= c && c <= 'z':
				currentState = 3
			default:
				break outer
			}
		case 2:
			switch {
			case '0' <= c && c <= '9':
				currentState = 7
			case 'a' <= c && c <= 'z':
				currentState = 5
			case 'A' <= c && c <= 'Z':
				currentState = 4
			case c == '_':
				currentState = 6
			default:
				break outer
			}
		case 3:
			switch {
			case 'a' <= c && c <= 'z':
				currentState = 5
			case 'A' <= c && c <= 'Z':
				currentState = 4
			case c == '_':
				currentState = 6
			case '0' <= c && c <= '9':
				currentState = 7
			default:
				break outer
			}
		case 5:
			switch {
			case '0' <= c && c <= '9':
				currentState = 7
			case c == '_':
				currentState = 6
			case 'a' <= c && c <= 'z':
				currentState = 5
			case 'A' <= c && c <= 'Z':
				currentState = 4
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type MCOMMENT struct{}
func (a *MCOMMENT) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		8: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 2:
			switch {
			case c == '*':
				currentState = 5
			case '\x01' <= c && c <= ')':
				currentState = 4
			case '+' <= c && c <= '\U0010ffff':
				currentState = 3
			default:
				break outer
			}
		case 4:
			switch {
			case '+' <= c && c <= '\U0010ffff':
				currentState = 3
			case c == '*':
				currentState = 5
			case '\x01' <= c && c <= ')':
				currentState = 4
			default:
				break outer
			}
		case 1:
			switch {
			case c == '*':
				currentState = 2
			default:
				break outer
			}
		case 5:
			switch {
			case c == '/':
				currentState = 8
			case '0' <= c && c <= '\U0010ffff':
				currentState = 6
			case '\x01' <= c && c <= '.':
				currentState = 7
			default:
				break outer
			}
		case 7:
			switch {
			case '\x01' <= c && c <= ')':
				currentState = 4
			case '+' <= c && c <= '\U0010ffff':
				currentState = 3
			case c == '*':
				currentState = 5
			default:
				break outer
			}
		case 0:
			switch {
			case c == '/':
				currentState = 1
			default:
				break outer
			}
		case 3:
			switch {
			case '+' <= c && c <= '\U0010ffff':
				currentState = 3
			case '\x01' <= c && c <= ')':
				currentState = 4
			case c == '*':
				currentState = 5
			default:
				break outer
			}
		case 6:
			switch {
			case '\x01' <= c && c <= ')':
				currentState = 4
			case '+' <= c && c <= '\U0010ffff':
				currentState = 3
			case c == '*':
				currentState = 5
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

type SCOMMENT struct{}
func (a *SCOMMENT) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		2: true,
		3: true,
		4: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {
		case 3:
			switch {
			case '\v' <= c && c <= '\U0010ffff':
				currentState = 3
			case '\x01' <= c && c <= '\t':
				currentState = 4
			default:
				break outer
			}
		case 4:
			switch {
			case '\v' <= c && c <= '\U0010ffff':
				currentState = 3
			case '\x01' <= c && c <= '\t':
				currentState = 4
			default:
				break outer
			}
		case 0:
			switch {
			case c == '/':
				currentState = 1
			default:
				break outer
			}
		case 1:
			switch {
			case c == '/':
				currentState = 2
			default:
				break outer
			}
		case 2:
			switch {
			case '\x01' <= c && c <= '\t':
				currentState = 4
			case '\v' <= c && c <= '\U0010ffff':
				currentState = 3
			default:
				break outer
			}
		default:
			break outer
		}
		iter.NextChar()
		c = iter.Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

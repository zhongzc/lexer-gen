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

		case 0:
			switch {
			case c == 'i':
				currentState = 1
			default:
				break outer
			}

		case 1:
			switch {
			case c == 'f':
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

type LP struct{}
func (a *LP) RunGreedy(iter CharIterator) error {
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

type RP struct{}
func (a *RP) RunGreedy(iter CharIterator) error {
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

type LB struct{}
func (a *LB) RunGreedy(iter CharIterator) error {
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

type RB struct{}
func (a *RB) RunGreedy(iter CharIterator) error {
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

type SEM struct{}
func (a *SEM) RunGreedy(iter CharIterator) error {
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
		1: true,
		2: true,
		3: true,
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
			case c == '0':
				currentState = 3
			case '1' <= c && c <= '9':
				currentState = 3
			case '0' <= c && c <= '1':
				currentState = 3
			default:
				break outer
			}

		case 3:
			switch {
			case c == '0':
				currentState = 3
			case '1' <= c && c <= '9':
				currentState = 3
			case '0' <= c && c <= '1':
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
		3: true,
		4: true,
		5: true,
		6: true,
		7: true,
		1: true,
		2: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {

		case 6:
			switch {
			case 'A' <= c && c <= 'Z':
				currentState = 4
			case 'a' <= c && c <= 'z':
				currentState = 5
			case c == '_':
				currentState = 6
			case '0' <= c && c <= '9':
				currentState = 7
			default:
				break outer
			}

		case 7:
			switch {
			case 'A' <= c && c <= 'Z':
				currentState = 4
			case 'a' <= c && c <= 'z':
				currentState = 5
			case c == '_':
				currentState = 6
			case '0' <= c && c <= '9':
				currentState = 7
			default:
				break outer
			}

		case 0:
			switch {
			case 'A' <= c && c <= 'Z':
				currentState = 1
			case 'a' <= c && c <= 'z':
				currentState = 2
			case c == '_':
				currentState = 3
			default:
				break outer
			}

		case 1:
			switch {
			case c == '_':
				currentState = 6
			case '0' <= c && c <= '9':
				currentState = 7
			case 'A' <= c && c <= 'Z':
				currentState = 4
			case 'a' <= c && c <= 'z':
				currentState = 5
			default:
				break outer
			}

		case 2:
			switch {
			case 'A' <= c && c <= 'Z':
				currentState = 4
			case 'a' <= c && c <= 'z':
				currentState = 5
			case c == '_':
				currentState = 6
			case '0' <= c && c <= '9':
				currentState = 7
			default:
				break outer
			}

		case 3:
			switch {
			case 'A' <= c && c <= 'Z':
				currentState = 4
			case 'a' <= c && c <= 'z':
				currentState = 5
			case c == '_':
				currentState = 6
			case '0' <= c && c <= '9':
				currentState = 7
			default:
				break outer
			}

		case 4:
			switch {
			case 'A' <= c && c <= 'Z':
				currentState = 4
			case 'a' <= c && c <= 'z':
				currentState = 5
			case c == '_':
				currentState = 6
			case '0' <= c && c <= '9':
				currentState = 7
			default:
				break outer
			}

		case 5:
			switch {
			case c == '_':
				currentState = 6
			case '0' <= c && c <= '9':
				currentState = 7
			case 'A' <= c && c <= 'Z':
				currentState = 4
			case 'a' <= c && c <= 'z':
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

type COMMENT struct{}
func (a *COMMENT) RunGreedy(iter CharIterator) error {
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
			case c == '/':
				currentState = 3
			case '\x01' <= c && c <= ')':
				currentState = 5
			case '+' <= c && c <= '.':
				currentState = 3
			case '0' <= c && c <= '\U0010ffff':
				currentState = 3
			case c == '*':
				currentState = 4
			case '.' <= c && c <= '/':
				currentState = 3
			case '/' <= c && c <= '0':
				currentState = 3
			default:
				break outer
			}

		case 3:
			switch {
			case c == '/':
				currentState = 3
			case '\x01' <= c && c <= ')':
				currentState = 5
			case '.' <= c && c <= '/':
				currentState = 3
			case '/' <= c && c <= '0':
				currentState = 3
			case '0' <= c && c <= '\U0010ffff':
				currentState = 3
			case c == '*':
				currentState = 4
			case '+' <= c && c <= '.':
				currentState = 3
			default:
				break outer
			}

		case 4:
			switch {
			case '\x01' <= c && c <= ')':
				currentState = 7
			case '+' <= c && c <= '.':
				currentState = 7
			case c == '/':
				currentState = 6
			case c == '*':
				currentState = 7
			case ')' <= c && c <= '*':
				currentState = 7
			case '*' <= c && c <= '+':
				currentState = 7
			case '0' <= c && c <= '\U0010ffff':
				currentState = 8
			default:
				break outer
			}

		case 5:
			switch {
			case c == '*':
				currentState = 4
			case '\x01' <= c && c <= ')':
				currentState = 5
			case '.' <= c && c <= '/':
				currentState = 3
			case '0' <= c && c <= '\U0010ffff':
				currentState = 3
			case c == '/':
				currentState = 3
			case '+' <= c && c <= '.':
				currentState = 3
			case '/' <= c && c <= '0':
				currentState = 3
			default:
				break outer
			}

		case 7:
			switch {
			case c == '/':
				currentState = 3
			case c == '*':
				currentState = 4
			case '\x01' <= c && c <= ')':
				currentState = 5
			case '+' <= c && c <= '.':
				currentState = 3
			case '.' <= c && c <= '/':
				currentState = 3
			case '0' <= c && c <= '\U0010ffff':
				currentState = 3
			case '/' <= c && c <= '0':
				currentState = 3
			default:
				break outer
			}

		case 8:
			switch {
			case '+' <= c && c <= '.':
				currentState = 3
			case '/' <= c && c <= '0':
				currentState = 3
			case c == '*':
				currentState = 4
			case '\x01' <= c && c <= ')':
				currentState = 5
			case '.' <= c && c <= '/':
				currentState = 3
			case '0' <= c && c <= '\U0010ffff':
				currentState = 3
			case c == '/':
				currentState = 3
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
			case c == '*':
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

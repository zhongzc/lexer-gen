package _tmpLexer

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
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {

		case 0:
			switch {
			case '0' <= c && c <= '9':
				currentState = 1
			default:
				break outer
			}

		case 1:
			switch {
			case '0' <= c && c <= '9':
				currentState = 2
			default:
				break outer
			}

		case 2:
			switch {
			case '0' <= c && c <= '9':
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

type IDENT struct{}
func (a *IDENT) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {

		case 3:
			switch {
			case 'A' <= c && c <= 'Z':
				currentState = 5
			case 'a' <= c && c <= 'z':
				currentState = 6
			case c == '_':
				currentState = 4
			default:
				break outer
			}

		case 4:
			switch {
			case 'a' <= c && c <= 'z':
				currentState = 6
			case c == '_':
				currentState = 4
			case 'A' <= c && c <= 'Z':
				currentState = 5
			default:
				break outer
			}

		case 5:
			switch {
			case 'A' <= c && c <= 'Z':
				currentState = 5
			case 'a' <= c && c <= 'z':
				currentState = 6
			case c == '_':
				currentState = 4
			default:
				break outer
			}

		case 6:
			switch {
			case c == '_':
				currentState = 4
			case 'A' <= c && c <= 'Z':
				currentState = 5
			case 'a' <= c && c <= 'z':
				currentState = 6
			default:
				break outer
			}

		case 0:
			switch {
			case c == '_':
				currentState = 3
			case 'A' <= c && c <= 'Z':
				currentState = 1
			case 'a' <= c && c <= 'z':
				currentState = 2
			default:
				break outer
			}

		case 1:
			switch {
			case c == '_':
				currentState = 4
			case 'A' <= c && c <= 'Z':
				currentState = 5
			case 'a' <= c && c <= 'z':
				currentState = 6
			default:
				break outer
			}

		case 2:
			switch {
			case 'A' <= c && c <= 'Z':
				currentState = 5
			case 'a' <= c && c <= 'z':
				currentState = 6
			case c == '_':
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

package template

import (
	"errors"
	"fmt"
)

// Automata
type Automata interface {
	RunGreedy(iter CharIterator) error
}

type DOG struct{}
func (a *DOG) RunGreedy(iter CharIterator) error {
	currentState := 0
	acceptState := map[int]bool{
		2: true,
		4: true,
		7: true,
		8: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {

		case 8:
			switch c {
			case 'c':
				currentState = 9
			default:
				break outer
			}

		case 9:
			switch c {
			case 'a':
				currentState = 6
			default:
				break outer
			}

		case 3:
			switch c {
			case 'o':
				currentState = 5
			default:
				break outer
			}

		case 4:
			switch c {
			case 'c':
				currentState = 4
			case 'd':
				currentState = 3
			case 'a':
				currentState = 6
			default:
				break outer
			}

		case 5:
			switch c {
			case 'g':
				currentState = 7
			default:
				break outer
			}

		case 6:
			switch c {
			case 't':
				currentState = 8
			default:
				break outer
			}

		case 0:
			switch c {
			case 'a':
				currentState = 1
			default:
				break outer
			}

		case 1:
			switch c {
			case 'b':
				currentState = 2
			default:
				break outer
			}

		case 2:
			switch c {
			case 'd':
				currentState = 3
			case 'c':
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

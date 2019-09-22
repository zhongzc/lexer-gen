package _go

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

	c := Peek()
outer:
	for {
		switch currentState {

		case 8:
			switch {
			case 'c' <= c && c <= 'c':
				currentState = 9
			default:
				break outer
			}


		case 9:
			switch {
			case 'a' <= c && c <= 'a':
				currentState = 6
			default:
				break outer
			}

		case 3:
			switch {
			case 'o' <= c && c <= 'o':
				currentState = 5
			default:
				break outer
			}

		case 4:
			switch {
			case 'c' <= c && c <= 'c':
				currentState = 4
			case 'd' <= c && c <= 'd':
				currentState = 3
			case 'a' <= c && c <= 'a':
				currentState = 6
			default:
				break outer
			}

		case 5:
			switch {
			case 'g' <= c && c <= 'g':
				currentState = 7
			default:
				break outer
			}

		case 6:
			switch {
			case 't' <= c && c <= 't':
				currentState = 8
			default:
				break outer
			}

		case 0:
			switch {
			case 'a' <= c && c <= 'a':
				currentState = 1
			default:
				break outer
			}

		case 1:
			switch {
			case 'b' <= c && c <= 'b':
				currentState = 2
			default:
				break outer
			}

		case 2:
			switch {
			case 'd' <= c && c <= 'd':
				currentState = 3
			case 'c' <= c && c <= 'c':
				currentState = 4
			default:
				break outer
			}

		default:
			break outer
		}
		NextChar()
		c = Peek()
	}

	if acceptState[currentState] {
		return nil
	}
	msg := fmt.Sprintf("%T run failed", a)
	return errors.New(msg)
}

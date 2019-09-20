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
			switch c {
			case 'i':
				currentState = 1
			default:
				break outer
			}

		case 1:
			switch c {
			case 'f':
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

		case 3:
			switch c {
			case 'e':
				currentState = 4
			default:
				break outer
			}

		case 0:
			switch c {
			case 'e':
				currentState = 1
			default:
				break outer
			}

		case 1:
			switch c {
			case 'l':
				currentState = 2
			default:
				break outer
			}

		case 2:
			switch c {
			case 's':
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
			switch c {
			case '(':
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
			switch c {
			case ')':
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
			switch c {
			case '{':
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
			switch c {
			case '}':
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
			switch c {
			case ';':
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

		case 1:
			switch c {
			case 'a':
				currentState = 2
			default:
				break outer
			}

		case 2:
			switch c {
			case 'r':
				currentState = 3
			default:
				break outer
			}

		case 0:
			switch c {
			case 'v':
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

		case 1:
			switch c {
			case 'e':
				currentState = 2
			default:
				break outer
			}

		case 2:
			switch c {
			case 't':
				currentState = 3
			default:
				break outer
			}

		case 3:
			switch c {
			case 'u':
				currentState = 4
			default:
				break outer
			}

		case 4:
			switch c {
			case 'r':
				currentState = 5
			default:
				break outer
			}

		case 5:
			switch c {
			case 'n':
				currentState = 6
			default:
				break outer
			}

		case 0:
			switch c {
			case 'r':
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

		case 1:
			switch c {
			case '=':
				currentState = 2
			default:
				break outer
			}

		case 0:
			switch c {
			case '=':
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
			switch c {
			case '=':
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
		8: true,
		11: true,
		15: true,
		16: true,
		18: true,
		19: true,
		5: true,
		6: true,
		10: true,
		12: true,
		17: true,
		7: true,
		9: true,
		13: true,
		20: true,
		2: true,
		3: true,
		4: true,
		14: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {

		case 4:
			switch c {
			case '0':
				currentState = 13
			case '1':
				currentState = 14
			case '3':
				currentState = 16
			case '6':
				currentState = 18
			case '9':
				currentState = 20
			case '2':
				currentState = 15
			case '4':
				currentState = 17
			case '5':
				currentState = 11
			case '7':
				currentState = 12
			case '8':
				currentState = 19
			default:
				break outer
			}

		case 8:
			switch c {
			case '2':
				currentState = 15
			case '3':
				currentState = 16
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '9':
				currentState = 20
			case '0':
				currentState = 13
			case '1':
				currentState = 14
			case '4':
				currentState = 17
			case '5':
				currentState = 11
			case '8':
				currentState = 19
			default:
				break outer
			}

		case 9:
			switch c {
			case '6':
				currentState = 18
			case '9':
				currentState = 20
			case '0':
				currentState = 13
			case '3':
				currentState = 16
			case '2':
				currentState = 15
			case '4':
				currentState = 17
			case '5':
				currentState = 11
			case '7':
				currentState = 12
			case '8':
				currentState = 19
			case '1':
				currentState = 14
			default:
				break outer
			}

		case 10:
			switch c {
			case '5':
				currentState = 11
			case '8':
				currentState = 19
			case '1':
				currentState = 14
			case '3':
				currentState = 16
			case '4':
				currentState = 17
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '9':
				currentState = 20
			case '0':
				currentState = 13
			case '2':
				currentState = 15
			default:
				break outer
			}

		case 11:
			switch c {
			case '1':
				currentState = 14
			case '4':
				currentState = 17
			case '2':
				currentState = 15
			case '3':
				currentState = 16
			case '5':
				currentState = 11
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '8':
				currentState = 19
			case '9':
				currentState = 20
			case '0':
				currentState = 13
			default:
				break outer
			}

		case 19:
			switch c {
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '8':
				currentState = 19
			case '9':
				currentState = 20
			case '4':
				currentState = 17
			case '5':
				currentState = 11
			case '3':
				currentState = 16
			case '0':
				currentState = 13
			case '1':
				currentState = 14
			case '2':
				currentState = 15
			default:
				break outer
			}

		case 1:
			switch c {
			case '9':
				currentState = 20
			case '4':
				currentState = 17
			case '2':
				currentState = 15
			case '3':
				currentState = 16
			case '5':
				currentState = 11
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '8':
				currentState = 19
			case '0':
				currentState = 13
			case '1':
				currentState = 14
			default:
				break outer
			}

		case 5:
			switch c {
			case '5':
				currentState = 11
			case '6':
				currentState = 18
			case '8':
				currentState = 19
			case '9':
				currentState = 20
			case '0':
				currentState = 13
			case '1':
				currentState = 14
			case '3':
				currentState = 16
			case '4':
				currentState = 17
			case '2':
				currentState = 15
			case '7':
				currentState = 12
			default:
				break outer
			}

		case 6:
			switch c {
			case '3':
				currentState = 16
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '8':
				currentState = 19
			case '9':
				currentState = 20
			case '2':
				currentState = 15
			case '4':
				currentState = 17
			case '5':
				currentState = 11
			case '0':
				currentState = 13
			case '1':
				currentState = 14
			default:
				break outer
			}

		case 7:
			switch c {
			case '8':
				currentState = 19
			case '0':
				currentState = 13
			case '2':
				currentState = 15
			case '3':
				currentState = 16
			case '5':
				currentState = 11
			case '7':
				currentState = 12
			case '1':
				currentState = 14
			case '4':
				currentState = 17
			case '6':
				currentState = 18
			case '9':
				currentState = 20
			default:
				break outer
			}

		case 16:
			switch c {
			case '9':
				currentState = 20
			case '1':
				currentState = 14
			case '2':
				currentState = 15
			case '3':
				currentState = 16
			case '4':
				currentState = 17
			case '5':
				currentState = 11
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '0':
				currentState = 13
			case '8':
				currentState = 19
			default:
				break outer
			}

		case 20:
			switch c {
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '9':
				currentState = 20
			case '2':
				currentState = 15
			case '3':
				currentState = 16
			case '4':
				currentState = 17
			case '5':
				currentState = 11
			case '0':
				currentState = 13
			case '1':
				currentState = 14
			case '8':
				currentState = 19
			default:
				break outer
			}

		case 2:
			switch c {
			case '1':
				currentState = 14
			case '2':
				currentState = 15
			case '5':
				currentState = 11
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '8':
				currentState = 19
			case '9':
				currentState = 20
			case '0':
				currentState = 13
			case '3':
				currentState = 16
			case '4':
				currentState = 17
			default:
				break outer
			}

		case 12:
			switch c {
			case '1':
				currentState = 14
			case '2':
				currentState = 15
			case '4':
				currentState = 17
			case '7':
				currentState = 12
			case '8':
				currentState = 19
			case '9':
				currentState = 20
			case '0':
				currentState = 13
			case '3':
				currentState = 16
			case '5':
				currentState = 11
			case '6':
				currentState = 18
			default:
				break outer
			}

		case 14:
			switch c {
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '9':
				currentState = 20
			case '0':
				currentState = 13
			case '1':
				currentState = 14
			case '2':
				currentState = 15
			case '3':
				currentState = 16
			case '5':
				currentState = 11
			case '4':
				currentState = 17
			case '8':
				currentState = 19
			default:
				break outer
			}

		case 17:
			switch c {
			case '0':
				currentState = 13
			case '1':
				currentState = 14
			case '3':
				currentState = 16
			case '5':
				currentState = 11
			case '8':
				currentState = 19
			case '2':
				currentState = 15
			case '4':
				currentState = 17
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '9':
				currentState = 20
			default:
				break outer
			}

		case 18:
			switch c {
			case '0':
				currentState = 13
			case '1':
				currentState = 14
			case '5':
				currentState = 11
			case '2':
				currentState = 15
			case '3':
				currentState = 16
			case '4':
				currentState = 17
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '8':
				currentState = 19
			case '9':
				currentState = 20
			default:
				break outer
			}

		case 0:
			switch c {
			case '4':
				currentState = 8
			case '6':
				currentState = 10
			case '7':
				currentState = 4
			case '8':
				currentState = 5
			case '0':
				currentState = 1
			case '1':
				currentState = 6
			case '2':
				currentState = 7
			case '3':
				currentState = 2
			case '5':
				currentState = 9
			case '9':
				currentState = 3
			default:
				break outer
			}

		case 3:
			switch c {
			case '0':
				currentState = 13
			case '1':
				currentState = 14
			case '2':
				currentState = 15
			case '3':
				currentState = 16
			case '5':
				currentState = 11
			case '6':
				currentState = 18
			case '8':
				currentState = 19
			case '4':
				currentState = 17
			case '7':
				currentState = 12
			case '9':
				currentState = 20
			default:
				break outer
			}

		case 13:
			switch c {
			case '2':
				currentState = 15
			case '9':
				currentState = 20
			case '1':
				currentState = 14
			case '3':
				currentState = 16
			case '4':
				currentState = 17
			case '5':
				currentState = 11
			case '6':
				currentState = 18
			case '7':
				currentState = 12
			case '8':
				currentState = 19
			case '0':
				currentState = 13
			default:
				break outer
			}

		case 15:
			switch c {
			case '8':
				currentState = 19
			case '9':
				currentState = 20
			case '4':
				currentState = 17
			case '5':
				currentState = 11
			case '6':
				currentState = 18
			case '2':
				currentState = 15
			case '3':
				currentState = 16
			case '7':
				currentState = 12
			case '0':
				currentState = 13
			case '1':
				currentState = 14
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
		4: true,
		17: true,
		65: true,
		68: true,
		70: true,
		3: true,
		35: true,
		52: true,
		89: true,
		101: true,
		5: true,
		10: true,
		72: true,
		12: true,
		67: true,
		78: true,
		104: true,
		82: true,
		41: true,
		45: true,
		98: true,
		62: true,
		103: true,
		2: true,
		76: true,
		94: true,
		1: true,
		31: true,
		73: true,
		6: true,
		26: true,
		66: true,
		50: true,
		61: true,
		92: true,
		77: true,
		8: true,
		18: true,
		25: true,
		30: true,
		33: true,
		53: true,
		22: true,
		46: true,
		56: true,
		75: true,
		84: true,
		13: true,
		37: true,
		80: true,
		85: true,
		7: true,
		54: true,
		71: true,
		14: true,
		43: true,
		57: true,
		91: true,
		102: true,
		55: true,
		74: true,
		44: true,
		99: true,
		79: true,
		9: true,
		48: true,
		51: true,
		69: true,
		97: true,
		11: true,
		100: true,
		23: true,
		58: true,
		39: true,
		83: true,
		96: true,
		105: true,
		90: true,
		20: true,
		38: true,
		88: true,
		27: true,
		59: true,
		81: true,
		16: true,
		93: true,
		28: true,
		36: true,
		47: true,
		60: true,
		21: true,
		24: true,
		40: true,
		49: true,
		64: true,
		87: true,
		106: true,
		19: true,
		29: true,
		32: true,
		63: true,
		86: true,
		95: true,
		15: true,
		34: true,
		42: true,
	}

	c := iter.Peek()
outer:
	for {
		switch currentState {

		case 39:
			switch c {
			case 'O':
				currentState = 73
			case 'n':
				currentState = 85
			case 'A':
				currentState = 77
			case 'j':
				currentState = 84
			case 'm':
				currentState = 54
			case 'v':
				currentState = 96
			case 'Y':
				currentState = 75
			case 'I':
				currentState = 62
			case 'k':
				currentState = 57
			case 'E':
				currentState = 95
			case 'W':
				currentState = 56
			case 'i':
				currentState = 80
			case 'l':
				currentState = 106
			case 'q':
				currentState = 81
			case '_':
				currentState = 97
			case 'D':
				currentState = 94
			case 'T':
				currentState = 74
			case 'C':
				currentState = 86
			case 'B':
				currentState = 88
			case 'H':
				currentState = 98
			case 'N':
				currentState = 99
			case 'U':
				currentState = 67
			case 'h':
				currentState = 79
			case 'r':
				currentState = 83
			case 'S':
				currentState = 78
			case 'a':
				currentState = 104
			case 'c':
				currentState = 69
			case 'd':
				currentState = 66
			case 'y':
				currentState = 72
			case 'R':
				currentState = 103
			case 'Z':
				currentState = 101
			case 'p':
				currentState = 70
			case 'F':
				currentState = 89
			case 'V':
				currentState = 92
			case 't':
				currentState = 76
			case 'M':
				currentState = 90
			case 'G':
				currentState = 100
			case 'L':
				currentState = 61
			case 'u':
				currentState = 55
			case 'w':
				currentState = 71
			case 'J':
				currentState = 59
			case 'e':
				currentState = 105
			case 'g':
				currentState = 82
			case 'o':
				currentState = 58
			case 's':
				currentState = 87
			case 'z':
				currentState = 64
			case 'b':
				currentState = 60
			case 'Q':
				currentState = 91
			case 'x':
				currentState = 93
			case 'P':
				currentState = 65
			case 'X':
				currentState = 68
			case 'f':
				currentState = 102
			case 'K':
				currentState = 63
			default:
				break outer
			}

		case 51:
			switch c {
			case 'n':
				currentState = 85
			case 'p':
				currentState = 70
			case 'D':
				currentState = 94
			case 'T':
				currentState = 74
			case 'Y':
				currentState = 75
			case 'g':
				currentState = 82
			case 'y':
				currentState = 72
			case 'X':
				currentState = 68
			case 'b':
				currentState = 60
			case 'e':
				currentState = 105
			case 'F':
				currentState = 89
			case 'S':
				currentState = 78
			case 'c':
				currentState = 69
			case 'v':
				currentState = 96
			case 'E':
				currentState = 95
			case 'H':
				currentState = 98
			case 'Q':
				currentState = 91
			case 'C':
				currentState = 86
			case 'm':
				currentState = 54
			case 'd':
				currentState = 66
			case 'j':
				currentState = 84
			case 'A':
				currentState = 77
			case 'N':
				currentState = 99
			case 'U':
				currentState = 67
			case 'R':
				currentState = 103
			case 'l':
				currentState = 106
			case 't':
				currentState = 76
			case 'J':
				currentState = 59
			case 'K':
				currentState = 63
			case 'O':
				currentState = 73
			case 'M':
				currentState = 90
			case 'k':
				currentState = 57
			case 'o':
				currentState = 58
			case 'w':
				currentState = 71
			case 's':
				currentState = 87
			case 'z':
				currentState = 64
			case 'a':
				currentState = 104
			case 'i':
				currentState = 80
			case 'V':
				currentState = 92
			case 'W':
				currentState = 56
			case 'f':
				currentState = 102
			case 'G':
				currentState = 100
			case '_':
				currentState = 97
			case 'B':
				currentState = 88
			case 'Z':
				currentState = 101
			case 'h':
				currentState = 79
			case 'r':
				currentState = 83
			case 'u':
				currentState = 55
			case 'q':
				currentState = 81
			case 'x':
				currentState = 93
			case 'I':
				currentState = 62
			case 'L':
				currentState = 61
			case 'P':
				currentState = 65
			default:
				break outer
			}

		case 65:
			switch c {
			case 'x':
				currentState = 93
			case 'H':
				currentState = 98
			case 'K':
				currentState = 63
			case 'L':
				currentState = 61
			case 'l':
				currentState = 106
			case 'B':
				currentState = 88
			case 'I':
				currentState = 62
			case 'Q':
				currentState = 91
			case 't':
				currentState = 76
			case 'w':
				currentState = 71
			case 'S':
				currentState = 78
			case 'e':
				currentState = 105
			case 'p':
				currentState = 70
			case 'z':
				currentState = 64
			case 'F':
				currentState = 89
			case 'G':
				currentState = 100
			case 's':
				currentState = 87
			case 'y':
				currentState = 72
			case 'A':
				currentState = 77
			case 'M':
				currentState = 90
			case 'o':
				currentState = 58
			case 'C':
				currentState = 86
			case 'g':
				currentState = 82
			case 'f':
				currentState = 102
			case 'j':
				currentState = 84
			case 'J':
				currentState = 59
			case 'O':
				currentState = 73
			case 'R':
				currentState = 103
			case 'h':
				currentState = 79
			case 'm':
				currentState = 54
			case 'n':
				currentState = 85
			case 'u':
				currentState = 55
			case 'D':
				currentState = 94
			case 'E':
				currentState = 95
			case 'X':
				currentState = 68
			case 'c':
				currentState = 69
			case 'N':
				currentState = 99
			case 'W':
				currentState = 56
			case 'b':
				currentState = 60
			case 'V':
				currentState = 92
			case '_':
				currentState = 97
			case 'q':
				currentState = 81
			case 'r':
				currentState = 83
			case 'T':
				currentState = 74
			case 'Y':
				currentState = 75
			case 'Z':
				currentState = 101
			case 'U':
				currentState = 67
			case 'v':
				currentState = 96
			case 'P':
				currentState = 65
			case 'd':
				currentState = 66
			case 'a':
				currentState = 104
			case 'i':
				currentState = 80
			case 'k':
				currentState = 57
			default:
				break outer
			}

		case 93:
			switch c {
			case 'C':
				currentState = 86
			case 'h':
				currentState = 79
			case 'x':
				currentState = 93
			case 'A':
				currentState = 77
			case 'd':
				currentState = 66
			case 'i':
				currentState = 80
			case 'u':
				currentState = 55
			case 'L':
				currentState = 61
			case 'v':
				currentState = 96
			case 'm':
				currentState = 54
			case 'F':
				currentState = 89
			case 'J':
				currentState = 59
			case 'M':
				currentState = 90
			case 'T':
				currentState = 74
			case 'V':
				currentState = 92
			case 'W':
				currentState = 56
			case 'P':
				currentState = 65
			case 'S':
				currentState = 78
			case 't':
				currentState = 76
			case 'E':
				currentState = 95
			case 'Y':
				currentState = 75
			case '_':
				currentState = 97
			case 'B':
				currentState = 88
			case 'e':
				currentState = 105
			case 'j':
				currentState = 84
			case 's':
				currentState = 87
			case 'G':
				currentState = 100
			case 'K':
				currentState = 63
			case 'Q':
				currentState = 91
			case 'Z':
				currentState = 101
			case 'b':
				currentState = 60
			case 'k':
				currentState = 57
			case 'r':
				currentState = 83
			case 'y':
				currentState = 72
			case 'w':
				currentState = 71
			case 'z':
				currentState = 64
			case 'R':
				currentState = 103
			case 'U':
				currentState = 67
			case 'l':
				currentState = 106
			case 'o':
				currentState = 58
			case 'q':
				currentState = 81
			case 'a':
				currentState = 104
			case 'n':
				currentState = 85
			case 'H':
				currentState = 98
			case 'I':
				currentState = 62
			case 'O':
				currentState = 73
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case 'p':
				currentState = 70
			case 'D':
				currentState = 94
			case 'N':
				currentState = 99
			case 'X':
				currentState = 68
			case 'c':
				currentState = 69
			default:
				break outer
			}

		case 30:
			switch c {
			case 'A':
				currentState = 77
			case 'N':
				currentState = 99
			case 'a':
				currentState = 104
			case 'c':
				currentState = 69
			case 't':
				currentState = 76
			case 'g':
				currentState = 82
			case 'm':
				currentState = 54
			case 'p':
				currentState = 70
			case 'y':
				currentState = 72
			case 'K':
				currentState = 63
			case 'S':
				currentState = 78
			case 'U':
				currentState = 67
			case 'j':
				currentState = 84
			case 'R':
				currentState = 103
			case 'T':
				currentState = 74
			case 'b':
				currentState = 60
			case 'n':
				currentState = 85
			case 'u':
				currentState = 55
			case 'F':
				currentState = 89
			case 'k':
				currentState = 57
			case 's':
				currentState = 87
			case 'x':
				currentState = 93
			case 'I':
				currentState = 62
			case 'M':
				currentState = 90
			case 'W':
				currentState = 56
			case 'h':
				currentState = 79
			case 'l':
				currentState = 106
			case 'C':
				currentState = 86
			case 'L':
				currentState = 61
			case 'O':
				currentState = 73
			case 'X':
				currentState = 68
			case 'v':
				currentState = 96
			case 'D':
				currentState = 94
			case 'Q':
				currentState = 91
			case 'Y':
				currentState = 75
			case 'o':
				currentState = 58
			case 'z':
				currentState = 64
			case 'd':
				currentState = 66
			case 'e':
				currentState = 105
			case 'J':
				currentState = 59
			case 'q':
				currentState = 81
			case 'r':
				currentState = 83
			case 'w':
				currentState = 71
			case 'E':
				currentState = 95
			case 'H':
				currentState = 98
			case 'P':
				currentState = 65
			case 'G':
				currentState = 100
			case 'Z':
				currentState = 101
			case 'i':
				currentState = 80
			case 'B':
				currentState = 88
			case 'f':
				currentState = 102
			case 'V':
				currentState = 92
			case '_':
				currentState = 97
			default:
				break outer
			}

		case 80:
			switch c {
			case 'Y':
				currentState = 75
			case 'e':
				currentState = 105
			case 'l':
				currentState = 106
			case 'x':
				currentState = 93
			case 'F':
				currentState = 89
			case 'J':
				currentState = 59
			case 'y':
				currentState = 72
			case 'C':
				currentState = 86
			case 'c':
				currentState = 69
			case 'u':
				currentState = 55
			case '_':
				currentState = 97
			case 'P':
				currentState = 65
			case 'S':
				currentState = 78
			case 'i':
				currentState = 80
			case 'm':
				currentState = 54
			case 'n':
				currentState = 85
			case 't':
				currentState = 76
			case 'X':
				currentState = 68
			case 'a':
				currentState = 104
			case 'g':
				currentState = 82
			case 'q':
				currentState = 81
			case 'K':
				currentState = 63
			case 'Q':
				currentState = 91
			case 'b':
				currentState = 60
			case 'f':
				currentState = 102
			case 'k':
				currentState = 57
			case 'z':
				currentState = 64
			case 'L':
				currentState = 61
			case 'U':
				currentState = 67
			case 'p':
				currentState = 70
			case 'r':
				currentState = 83
			case 's':
				currentState = 87
			case 'D':
				currentState = 94
			case 'N':
				currentState = 99
			case 'W':
				currentState = 56
			case 'o':
				currentState = 58
			case 'E':
				currentState = 95
			case 'T':
				currentState = 74
			case 'G':
				currentState = 100
			case 'H':
				currentState = 98
			case 'Z':
				currentState = 101
			case 'h':
				currentState = 79
			case 'v':
				currentState = 96
			case 'w':
				currentState = 71
			case 'A':
				currentState = 77
			case 'B':
				currentState = 88
			case 'V':
				currentState = 92
			case 'j':
				currentState = 84
			case 'I':
				currentState = 62
			case 'M':
				currentState = 90
			case 'O':
				currentState = 73
			case 'd':
				currentState = 66
			case 'R':
				currentState = 103
			default:
				break outer
			}

		case 75:
			switch c {
			case 'J':
				currentState = 59
			case 'U':
				currentState = 67
			case 'a':
				currentState = 104
			case 'z':
				currentState = 64
			case 'D':
				currentState = 94
			case 'T':
				currentState = 74
			case 'Z':
				currentState = 101
			case 'i':
				currentState = 80
			case 'm':
				currentState = 54
			case 'b':
				currentState = 60
			case 'g':
				currentState = 82
			case 't':
				currentState = 76
			case 'c':
				currentState = 69
			case 'u':
				currentState = 55
			case 'A':
				currentState = 77
			case 'W':
				currentState = 56
			case 'd':
				currentState = 66
			case 'K':
				currentState = 63
			case 'Y':
				currentState = 75
			case 'L':
				currentState = 61
			case 'k':
				currentState = 57
			case 's':
				currentState = 87
			case 'C':
				currentState = 86
			case 'G':
				currentState = 100
			case 'V':
				currentState = 92
			case 'e':
				currentState = 105
			case 'y':
				currentState = 72
			case 'B':
				currentState = 88
			case 'N':
				currentState = 99
			case 'R':
				currentState = 103
			case 'j':
				currentState = 84
			case 'r':
				currentState = 83
			case 'w':
				currentState = 71
			case 'q':
				currentState = 81
			case 'x':
				currentState = 93
			case '_':
				currentState = 97
			case 'E':
				currentState = 95
			case 'H':
				currentState = 98
			case 'I':
				currentState = 62
			case 'M':
				currentState = 90
			case 'p':
				currentState = 70
			case 'P':
				currentState = 65
			case 'h':
				currentState = 79
			case 'o':
				currentState = 58
			case 'O':
				currentState = 73
			case 'Q':
				currentState = 91
			case 'f':
				currentState = 102
			case 'n':
				currentState = 85
			case 'v':
				currentState = 96
			case 'X':
				currentState = 68
			case 'F':
				currentState = 89
			case 'l':
				currentState = 106
			case 'S':
				currentState = 78
			default:
				break outer
			}

		case 28:
			switch c {
			case 'j':
				currentState = 84
			case 'p':
				currentState = 70
			case 'z':
				currentState = 64
			case 'P':
				currentState = 65
			case 'W':
				currentState = 56
			case 'c':
				currentState = 69
			case 'g':
				currentState = 82
			case 'v':
				currentState = 96
			case 'x':
				currentState = 93
			case 'J':
				currentState = 59
			case 'F':
				currentState = 89
			case 'T':
				currentState = 74
			case 'V':
				currentState = 92
			case 'I':
				currentState = 62
			case 'h':
				currentState = 79
			case 't':
				currentState = 76
			case 'u':
				currentState = 55
			case '_':
				currentState = 97
			case 'B':
				currentState = 88
			case 'N':
				currentState = 99
			case 'r':
				currentState = 83
			case 'K':
				currentState = 63
			case 'U':
				currentState = 67
			case 'f':
				currentState = 102
			case 'X':
				currentState = 68
			case 'l':
				currentState = 106
			case 'm':
				currentState = 54
			case 'D':
				currentState = 94
			case 'E':
				currentState = 95
			case 'L':
				currentState = 61
			case 'C':
				currentState = 86
			case 'G':
				currentState = 100
			case 'H':
				currentState = 98
			case 'k':
				currentState = 57
			case 'M':
				currentState = 90
			case 'S':
				currentState = 78
			case 'w':
				currentState = 71
			case 'Y':
				currentState = 75
			case 'e':
				currentState = 105
			case 'o':
				currentState = 58
			case 'b':
				currentState = 60
			case 'Q':
				currentState = 91
			case 'R':
				currentState = 103
			case 'Z':
				currentState = 101
			case 'i':
				currentState = 80
			case 'n':
				currentState = 85
			case 'q':
				currentState = 81
			case 's':
				currentState = 87
			case 'y':
				currentState = 72
			case 'A':
				currentState = 77
			case 'O':
				currentState = 73
			case 'd':
				currentState = 66
			case 'a':
				currentState = 104
			default:
				break outer
			}

		case 54:
			switch c {
			case 'O':
				currentState = 73
			case 'f':
				currentState = 102
			case 'k':
				currentState = 57
			case 's':
				currentState = 87
			case 'w':
				currentState = 71
			case 'F':
				currentState = 89
			case 'J':
				currentState = 59
			case 'l':
				currentState = 106
			case 'W':
				currentState = 56
			case 'd':
				currentState = 66
			case 'u':
				currentState = 55
			case 'N':
				currentState = 99
			case 'e':
				currentState = 105
			case 'q':
				currentState = 81
			case 'C':
				currentState = 86
			case 'Q':
				currentState = 91
			case 'g':
				currentState = 82
			case 'p':
				currentState = 70
			case 'G':
				currentState = 100
			case 'V':
				currentState = 92
			case '_':
				currentState = 97
			case 'B':
				currentState = 88
			case 'o':
				currentState = 58
			case 'i':
				currentState = 80
			case 'x':
				currentState = 93
			case 'h':
				currentState = 79
			case 'n':
				currentState = 85
			case 'v':
				currentState = 96
			case 'L':
				currentState = 61
			case 'r':
				currentState = 83
			case 'D':
				currentState = 94
			case 'I':
				currentState = 62
			case 'R':
				currentState = 103
			case 'Y':
				currentState = 75
			case 'b':
				currentState = 60
			case 'y':
				currentState = 72
			case 'E':
				currentState = 95
			case 'K':
				currentState = 63
			case 'M':
				currentState = 90
			case 'P':
				currentState = 65
			case 'm':
				currentState = 54
			case 't':
				currentState = 76
			case 'z':
				currentState = 64
			case 'a':
				currentState = 104
			case 'A':
				currentState = 77
			case 'U':
				currentState = 67
			case 'X':
				currentState = 68
			case 'j':
				currentState = 84
			case 'T':
				currentState = 74
			case 'S':
				currentState = 78
			case 'Z':
				currentState = 101
			case 'c':
				currentState = 69
			case 'H':
				currentState = 98
			default:
				break outer
			}

		case 22:
			switch c {
			case 'N':
				currentState = 99
			case 'b':
				currentState = 60
			case 'p':
				currentState = 70
			case 'A':
				currentState = 77
			case 'V':
				currentState = 92
			case 'i':
				currentState = 80
			case 'o':
				currentState = 58
			case 'r':
				currentState = 83
			case 'R':
				currentState = 103
			case 'X':
				currentState = 68
			case 'c':
				currentState = 69
			case 't':
				currentState = 76
			case 'E':
				currentState = 95
			case 'G':
				currentState = 100
			case 'P':
				currentState = 65
			case 'g':
				currentState = 82
			case 'n':
				currentState = 85
			case 'K':
				currentState = 63
			case 'S':
				currentState = 78
			case 'x':
				currentState = 93
			case 'C':
				currentState = 86
			case 'Q':
				currentState = 91
			case 'U':
				currentState = 67
			case 'j':
				currentState = 84
			case 'J':
				currentState = 59
			case 'W':
				currentState = 56
			case 'd':
				currentState = 66
			case 's':
				currentState = 87
			case 'B':
				currentState = 88
			case 'T':
				currentState = 74
			case 'k':
				currentState = 57
			case 'z':
				currentState = 64
			case 'D':
				currentState = 94
			case 'O':
				currentState = 73
			case 'Z':
				currentState = 101
			case 'f':
				currentState = 102
			case 'q':
				currentState = 81
			case 'H':
				currentState = 98
			case 'm':
				currentState = 54
			case '_':
				currentState = 97
			case 'F':
				currentState = 89
			case 'L':
				currentState = 61
			case 'M':
				currentState = 90
			case 'Y':
				currentState = 75
			case 'a':
				currentState = 104
			case 'e':
				currentState = 105
			case 'I':
				currentState = 62
			case 'l':
				currentState = 106
			case 'v':
				currentState = 96
			case 'h':
				currentState = 79
			case 'u':
				currentState = 55
			case 'w':
				currentState = 71
			case 'y':
				currentState = 72
			default:
				break outer
			}

		case 88:
			switch c {
			case 'f':
				currentState = 102
			case 'K':
				currentState = 63
			case 'P':
				currentState = 65
			case 'a':
				currentState = 104
			case 'i':
				currentState = 80
			case 'p':
				currentState = 70
			case 'M':
				currentState = 90
			case 'c':
				currentState = 69
			case 'r':
				currentState = 83
			case 'v':
				currentState = 96
			case 'x':
				currentState = 93
			case 'y':
				currentState = 72
			case 'g':
				currentState = 82
			case 'l':
				currentState = 106
			case 'X':
				currentState = 68
			case 's':
				currentState = 87
			case 'Z':
				currentState = 101
			case 'u':
				currentState = 55
			case 'T':
				currentState = 74
			case 'k':
				currentState = 57
			case 'n':
				currentState = 85
			case 'q':
				currentState = 81
			case 'o':
				currentState = 58
			case '_':
				currentState = 97
			case 'e':
				currentState = 105
			case 'G':
				currentState = 100
			case 'I':
				currentState = 62
			case 'N':
				currentState = 99
			case 'Q':
				currentState = 91
			case 'S':
				currentState = 78
			case 'b':
				currentState = 60
			case 'C':
				currentState = 86
			case 'F':
				currentState = 89
			case 'H':
				currentState = 98
			case 'h':
				currentState = 79
			case 't':
				currentState = 76
			case 'w':
				currentState = 71
			case 'z':
				currentState = 64
			case 'B':
				currentState = 88
			case 'R':
				currentState = 103
			case 'U':
				currentState = 67
			case 'm':
				currentState = 54
			case 'D':
				currentState = 94
			case 'L':
				currentState = 61
			case 'O':
				currentState = 73
			case 'V':
				currentState = 92
			case 'd':
				currentState = 66
			case 'j':
				currentState = 84
			case 'A':
				currentState = 77
			case 'J':
				currentState = 59
			case 'Y':
				currentState = 75
			case 'E':
				currentState = 95
			case 'W':
				currentState = 56
			default:
				break outer
			}

		case 97:
			switch c {
			case 'p':
				currentState = 70
			case 'P':
				currentState = 65
			case 'r':
				currentState = 83
			case 'u':
				currentState = 55
			case 'J':
				currentState = 59
			case 'f':
				currentState = 102
			case 'H':
				currentState = 98
			case 'E':
				currentState = 95
			case 'a':
				currentState = 104
			case 'b':
				currentState = 60
			case 'C':
				currentState = 86
			case 's':
				currentState = 87
			case 't':
				currentState = 76
			case 'k':
				currentState = 57
			case 'S':
				currentState = 78
			case 'Z':
				currentState = 101
			case 'l':
				currentState = 106
			case 'n':
				currentState = 85
			case 'Q':
				currentState = 91
			case 'N':
				currentState = 99
			case 'c':
				currentState = 69
			case 'j':
				currentState = 84
			case 'w':
				currentState = 71
			case 'M':
				currentState = 90
			case 'x':
				currentState = 93
			case 'z':
				currentState = 64
			case 'K':
				currentState = 63
			case 'L':
				currentState = 61
			case 'T':
				currentState = 74
			case 'U':
				currentState = 67
			case 'X':
				currentState = 68
			case 'm':
				currentState = 54
			case 'o':
				currentState = 58
			case 'G':
				currentState = 100
			case 'F':
				currentState = 89
			case 'W':
				currentState = 56
			case 'i':
				currentState = 80
			case 'D':
				currentState = 94
			case 'R':
				currentState = 103
			case 'd':
				currentState = 66
			case 'g':
				currentState = 82
			case 'q':
				currentState = 81
			case 'v':
				currentState = 96
			case 'y':
				currentState = 72
			case 'I':
				currentState = 62
			case 'O':
				currentState = 73
			case 'V':
				currentState = 92
			case 'Y':
				currentState = 75
			case 'e':
				currentState = 105
			case 'h':
				currentState = 79
			case '_':
				currentState = 97
			case 'A':
				currentState = 77
			case 'B':
				currentState = 88
			default:
				break outer
			}

		case 18:
			switch c {
			case 'N':
				currentState = 99
			case 'B':
				currentState = 88
			case 'e':
				currentState = 105
			case 's':
				currentState = 87
			case 'K':
				currentState = 63
			case 'z':
				currentState = 64
			case 'Z':
				currentState = 101
			case 'M':
				currentState = 90
			case 'l':
				currentState = 106
			case 'y':
				currentState = 72
			case 'C':
				currentState = 86
			case 'U':
				currentState = 67
			case 'Q':
				currentState = 91
			case 'F':
				currentState = 89
			case 'O':
				currentState = 73
			case 'm':
				currentState = 54
			case 'A':
				currentState = 77
			case 'b':
				currentState = 60
			case 'j':
				currentState = 84
			case 'n':
				currentState = 85
			case 'r':
				currentState = 83
			case 'E':
				currentState = 95
			case 'P':
				currentState = 65
			case 'a':
				currentState = 104
			case 'g':
				currentState = 82
			case 'p':
				currentState = 70
			case 'u':
				currentState = 55
			case '_':
				currentState = 97
			case 'H':
				currentState = 98
			case 'W':
				currentState = 56
			case 'f':
				currentState = 102
			case 't':
				currentState = 76
			case 'V':
				currentState = 92
			case 'X':
				currentState = 68
			case 'k':
				currentState = 57
			case 'w':
				currentState = 71
			case 'S':
				currentState = 78
			case 'c':
				currentState = 69
			case 'o':
				currentState = 58
			case 'v':
				currentState = 96
			case 'x':
				currentState = 93
			case 'D':
				currentState = 94
			case 'Y':
				currentState = 75
			case 'L':
				currentState = 61
			case 'R':
				currentState = 103
			case 'h':
				currentState = 79
			case 'i':
				currentState = 80
			case 'G':
				currentState = 100
			case 'J':
				currentState = 59
			case 'T':
				currentState = 74
			case 'd':
				currentState = 66
			case 'q':
				currentState = 81
			case 'I':
				currentState = 62
			default:
				break outer
			}

		case 56:
			switch c {
			case 'B':
				currentState = 88
			case '_':
				currentState = 97
			case 'P':
				currentState = 65
			case 'S':
				currentState = 78
			case 'U':
				currentState = 67
			case 'o':
				currentState = 58
			case 'J':
				currentState = 59
			case 'O':
				currentState = 73
			case 'b':
				currentState = 60
			case 'g':
				currentState = 82
			case 'n':
				currentState = 85
			case 'M':
				currentState = 90
			case 'C':
				currentState = 86
			case 'Q':
				currentState = 91
			case 'Z':
				currentState = 101
			case 'c':
				currentState = 69
			case 'l':
				currentState = 106
			case 'q':
				currentState = 81
			case 'd':
				currentState = 66
			case 'm':
				currentState = 54
			case 'R':
				currentState = 103
			case 'T':
				currentState = 74
			case 'e':
				currentState = 105
			case 'h':
				currentState = 79
			case 's':
				currentState = 87
			case 'f':
				currentState = 102
			case 't':
				currentState = 76
			case 'x':
				currentState = 93
			case 'z':
				currentState = 64
			case 'D':
				currentState = 94
			case 'G':
				currentState = 100
			case 'I':
				currentState = 62
			case 'N':
				currentState = 99
			case 'A':
				currentState = 77
			case 'F':
				currentState = 89
			case 'L':
				currentState = 61
			case 'X':
				currentState = 68
			case 'i':
				currentState = 80
			case 'p':
				currentState = 70
			case 'v':
				currentState = 96
			case 'V':
				currentState = 92
			case 'j':
				currentState = 84
			case 'y':
				currentState = 72
			case 'W':
				currentState = 56
			case 'r':
				currentState = 83
			case 'u':
				currentState = 55
			case 'w':
				currentState = 71
			case 'k':
				currentState = 57
			case 'H':
				currentState = 98
			case 'Y':
				currentState = 75
			case 'E':
				currentState = 95
			case 'K':
				currentState = 63
			case 'a':
				currentState = 104
			default:
				break outer
			}

		case 15:
			switch c {
			case 'P':
				currentState = 65
			case 'i':
				currentState = 80
			case 'Z':
				currentState = 101
			case 'w':
				currentState = 71
			case 'p':
				currentState = 70
			case 'y':
				currentState = 72
			case 'I':
				currentState = 62
			case 'K':
				currentState = 63
			case 'L':
				currentState = 61
			case 'N':
				currentState = 99
			case 'x':
				currentState = 93
			case 'z':
				currentState = 64
			case 'g':
				currentState = 82
			case '_':
				currentState = 97
			case 'S':
				currentState = 78
			case 'j':
				currentState = 84
			case 'G':
				currentState = 100
			case 'b':
				currentState = 60
			case 'u':
				currentState = 55
			case 'v':
				currentState = 96
			case 'Q':
				currentState = 91
			case 'a':
				currentState = 104
			case 'd':
				currentState = 66
			case 'q':
				currentState = 81
			case 'M':
				currentState = 90
			case 'O':
				currentState = 73
			case 'm':
				currentState = 54
			case 'o':
				currentState = 58
			case 'A':
				currentState = 77
			case 'F':
				currentState = 89
			case 'W':
				currentState = 56
			case 'e':
				currentState = 105
			case 'h':
				currentState = 79
			case 'r':
				currentState = 83
			case 'k':
				currentState = 57
			case 'B':
				currentState = 88
			case 'E':
				currentState = 95
			case 'T':
				currentState = 74
			case 'U':
				currentState = 67
			case 'Y':
				currentState = 75
			case 'c':
				currentState = 69
			case 'C':
				currentState = 86
			case 'H':
				currentState = 98
			case 'R':
				currentState = 103
			case 'l':
				currentState = 106
			case 's':
				currentState = 87
			case 'D':
				currentState = 94
			case 'J':
				currentState = 59
			case 'f':
				currentState = 102
			case 't':
				currentState = 76
			case 'V':
				currentState = 92
			case 'X':
				currentState = 68
			case 'n':
				currentState = 85
			default:
				break outer
			}

		case 24:
			switch c {
			case 'C':
				currentState = 86
			case 'j':
				currentState = 84
			case 'v':
				currentState = 96
			case 'K':
				currentState = 63
			case 'W':
				currentState = 56
			case 'X':
				currentState = 68
			case 'I':
				currentState = 62
			case 'q':
				currentState = 81
			case 'B':
				currentState = 88
			case 'f':
				currentState = 102
			case 'p':
				currentState = 70
			case 'P':
				currentState = 65
			case 'V':
				currentState = 92
			case 'o':
				currentState = 58
			case 'u':
				currentState = 55
			case 'k':
				currentState = 57
			case 't':
				currentState = 76
			case 'R':
				currentState = 103
			case 'i':
				currentState = 80
			case 'y':
				currentState = 72
			case 'L':
				currentState = 61
			case 'M':
				currentState = 90
			case 'Q':
				currentState = 91
			case 'a':
				currentState = 104
			case 'F':
				currentState = 89
			case 'S':
				currentState = 78
			case 'U':
				currentState = 67
			case 'w':
				currentState = 71
			case 'H':
				currentState = 98
			case 'c':
				currentState = 69
			case 'g':
				currentState = 82
			case 'z':
				currentState = 64
			case '_':
				currentState = 97
			case 'n':
				currentState = 85
			case 's':
				currentState = 87
			case 'A':
				currentState = 77
			case 'T':
				currentState = 74
			case 'e':
				currentState = 105
			case 'h':
				currentState = 79
			case 'm':
				currentState = 54
			case 'D':
				currentState = 94
			case 'N':
				currentState = 99
			case 'b':
				currentState = 60
			case 'x':
				currentState = 93
			case 'E':
				currentState = 95
			case 'O':
				currentState = 73
			case 'Y':
				currentState = 75
			case 'l':
				currentState = 106
			case 'r':
				currentState = 83
			case 'Z':
				currentState = 101
			case 'd':
				currentState = 66
			case 'G':
				currentState = 100
			case 'J':
				currentState = 59
			default:
				break outer
			}

		case 25:
			switch c {
			case 'S':
				currentState = 78
			case 'f':
				currentState = 102
			case 'h':
				currentState = 79
			case 'o':
				currentState = 58
			case 'A':
				currentState = 77
			case 'K':
				currentState = 63
			case 'L':
				currentState = 61
			case 'k':
				currentState = 57
			case 'y':
				currentState = 72
			case 'M':
				currentState = 90
			case 'U':
				currentState = 67
			case 'W':
				currentState = 56
			case 'Z':
				currentState = 101
			case 'm':
				currentState = 54
			case 'I':
				currentState = 62
			case 'V':
				currentState = 92
			case 'd':
				currentState = 66
			case 'D':
				currentState = 94
			case 'H':
				currentState = 98
			case 'N':
				currentState = 99
			case 'i':
				currentState = 80
			case 'l':
				currentState = 106
			case 'u':
				currentState = 55
			case 'Y':
				currentState = 75
			case 'j':
				currentState = 84
			case 'T':
				currentState = 74
			case 'z':
				currentState = 64
			case 'B':
				currentState = 88
			case 'P':
				currentState = 65
			case 'R':
				currentState = 103
			case 'X':
				currentState = 68
			case 'e':
				currentState = 105
			case 'q':
				currentState = 81
			case 'C':
				currentState = 86
			case 'J':
				currentState = 59
			case 'a':
				currentState = 104
			case 't':
				currentState = 76
			case 'v':
				currentState = 96
			case 'E':
				currentState = 95
			case 'p':
				currentState = 70
			case 'c':
				currentState = 69
			case '_':
				currentState = 97
			case 'F':
				currentState = 89
			case 'b':
				currentState = 60
			case 'g':
				currentState = 82
			case 'r':
				currentState = 83
			case 'x':
				currentState = 93
			case 'G':
				currentState = 100
			case 'O':
				currentState = 73
			case 'Q':
				currentState = 91
			case 's':
				currentState = 87
			case 'n':
				currentState = 85
			case 'w':
				currentState = 71
			default:
				break outer
			}

		case 44:
			switch c {
			case 'E':
				currentState = 95
			case 'W':
				currentState = 56
			case 'o':
				currentState = 58
			case 'w':
				currentState = 71
			case 'p':
				currentState = 70
			case 'z':
				currentState = 64
			case 'e':
				currentState = 105
			case 'l':
				currentState = 106
			case 'Z':
				currentState = 101
			case 't':
				currentState = 76
			case 'i':
				currentState = 80
			case 'u':
				currentState = 55
			case 'F':
				currentState = 89
			case 'M':
				currentState = 90
			case 'X':
				currentState = 68
			case 's':
				currentState = 87
			case 'H':
				currentState = 98
			case 'c':
				currentState = 69
			case 'n':
				currentState = 85
			case 'D':
				currentState = 94
			case 'V':
				currentState = 92
			case 'q':
				currentState = 81
			case 'N':
				currentState = 99
			case 'S':
				currentState = 78
			case 'v':
				currentState = 96
			case 'P':
				currentState = 65
			case 'a':
				currentState = 104
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case '_':
				currentState = 97
			case 'A':
				currentState = 77
			case 'C':
				currentState = 86
			case 'R':
				currentState = 103
			case 'T':
				currentState = 74
			case 'Y':
				currentState = 75
			case 'h':
				currentState = 79
			case 'x':
				currentState = 93
			case 'B':
				currentState = 88
			case 'L':
				currentState = 61
			case 'O':
				currentState = 73
			case 'd':
				currentState = 66
			case 'K':
				currentState = 63
			case 'b':
				currentState = 60
			case 'J':
				currentState = 59
			case 'Q':
				currentState = 91
			case 'j':
				currentState = 84
			case 'k':
				currentState = 57
			case 'I':
				currentState = 62
			case 'U':
				currentState = 67
			case 'm':
				currentState = 54
			case 'y':
				currentState = 72
			case 'G':
				currentState = 100
			case 'r':
				currentState = 83
			default:
				break outer
			}

		case 57:
			switch c {
			case 'K':
				currentState = 63
			case 'Z':
				currentState = 101
			case 'y':
				currentState = 72
			case 'A':
				currentState = 77
			case 'T':
				currentState = 74
			case 'm':
				currentState = 54
			case 'F':
				currentState = 89
			case 'b':
				currentState = 60
			case 'n':
				currentState = 85
			case 'S':
				currentState = 78
			case '_':
				currentState = 97
			case 'Q':
				currentState = 91
			case 'l':
				currentState = 106
			case 'X':
				currentState = 68
			case 'I':
				currentState = 62
			case 'M':
				currentState = 90
			case 'R':
				currentState = 103
			case 'g':
				currentState = 82
			case 'o':
				currentState = 58
			case 'w':
				currentState = 71
			case 'G':
				currentState = 100
			case 'L':
				currentState = 61
			case 'a':
				currentState = 104
			case 'r':
				currentState = 83
			case 's':
				currentState = 87
			case 'x':
				currentState = 93
			case 'k':
				currentState = 57
			case 'u':
				currentState = 55
			case 'B':
				currentState = 88
			case 'd':
				currentState = 66
			case 'j':
				currentState = 84
			case 'J':
				currentState = 59
			case 'z':
				currentState = 64
			case 'D':
				currentState = 94
			case 'E':
				currentState = 95
			case 'N':
				currentState = 99
			case 't':
				currentState = 76
			case 'v':
				currentState = 96
			case 'Y':
				currentState = 75
			case 'c':
				currentState = 69
			case 'p':
				currentState = 70
			case 'i':
				currentState = 80
			case 'q':
				currentState = 81
			case 'C':
				currentState = 86
			case 'H':
				currentState = 98
			case 'U':
				currentState = 67
			case 'e':
				currentState = 105
			case 'f':
				currentState = 102
			case 'h':
				currentState = 79
			case 'P':
				currentState = 65
			case 'V':
				currentState = 92
			case 'W':
				currentState = 56
			case 'O':
				currentState = 73
			default:
				break outer
			}

		case 83:
			switch c {
			case 'N':
				currentState = 99
			case 'n':
				currentState = 85
			case 'r':
				currentState = 83
			case 'P':
				currentState = 65
			case 'o':
				currentState = 58
			case 'G':
				currentState = 100
			case 'I':
				currentState = 62
			case 'J':
				currentState = 59
			case 'y':
				currentState = 72
			case 'u':
				currentState = 55
			case 'M':
				currentState = 90
			case 'T':
				currentState = 74
			case 'm':
				currentState = 54
			case 'p':
				currentState = 70
			case 'v':
				currentState = 96
			case 'D':
				currentState = 94
			case 'V':
				currentState = 92
			case 'b':
				currentState = 60
			case 'c':
				currentState = 69
			case 'q':
				currentState = 81
			case 'F':
				currentState = 89
			case 'W':
				currentState = 56
			case 'd':
				currentState = 66
			case 'k':
				currentState = 57
			case 'E':
				currentState = 95
			case 'w':
				currentState = 71
			case 'Q':
				currentState = 91
			case 'Y':
				currentState = 75
			case 'j':
				currentState = 84
			case 'O':
				currentState = 73
			case 'i':
				currentState = 80
			case 'X':
				currentState = 68
			case 'l':
				currentState = 106
			case 's':
				currentState = 87
			case 'L':
				currentState = 61
			case 'R':
				currentState = 103
			case 'K':
				currentState = 63
			case 'z':
				currentState = 64
			case 't':
				currentState = 76
			case 'S':
				currentState = 78
			case 'a':
				currentState = 104
			case 'e':
				currentState = 105
			case 'f':
				currentState = 102
			case 'C':
				currentState = 86
			case 'H':
				currentState = 98
			case 'g':
				currentState = 82
			case 'B':
				currentState = 88
			case 'U':
				currentState = 67
			case 'x':
				currentState = 93
			case '_':
				currentState = 97
			case 'A':
				currentState = 77
			case 'Z':
				currentState = 101
			case 'h':
				currentState = 79
			default:
				break outer
			}

		case 89:
			switch c {
			case 'D':
				currentState = 94
			case 'I':
				currentState = 62
			case '_':
				currentState = 97
			case 'A':
				currentState = 77
			case 'y':
				currentState = 72
			case 'M':
				currentState = 90
			case 'Q':
				currentState = 91
			case 'k':
				currentState = 57
			case 'v':
				currentState = 96
			case 'c':
				currentState = 69
			case 'q':
				currentState = 81
			case 'V':
				currentState = 92
			case 'a':
				currentState = 104
			case 'o':
				currentState = 58
			case 'B':
				currentState = 88
			case 'E':
				currentState = 95
			case 'F':
				currentState = 89
			case 'H':
				currentState = 98
			case 'C':
				currentState = 86
			case 'K':
				currentState = 63
			case 'i':
				currentState = 80
			case 'P':
				currentState = 65
			case 'h':
				currentState = 79
			case 't':
				currentState = 76
			case 'N':
				currentState = 99
			case 'm':
				currentState = 54
			case 's':
				currentState = 87
			case 'J':
				currentState = 59
			case 'S':
				currentState = 78
			case 'T':
				currentState = 74
			case 'X':
				currentState = 68
			case 'l':
				currentState = 106
			case 'p':
				currentState = 70
			case 'U':
				currentState = 67
			case 'Y':
				currentState = 75
			case 'd':
				currentState = 66
			case 'b':
				currentState = 60
			case 'g':
				currentState = 82
			case 'u':
				currentState = 55
			case 'x':
				currentState = 93
			case 'R':
				currentState = 103
			case 'W':
				currentState = 56
			case 'Z':
				currentState = 101
			case 'f':
				currentState = 102
			case 'r':
				currentState = 83
			case 'w':
				currentState = 71
			case 'z':
				currentState = 64
			case 'G':
				currentState = 100
			case 'e':
				currentState = 105
			case 'j':
				currentState = 84
			case 'n':
				currentState = 85
			case 'L':
				currentState = 61
			case 'O':
				currentState = 73
			default:
				break outer
			}

		case 17:
			switch c {
			case 'P':
				currentState = 65
			case 'w':
				currentState = 71
			case 'x':
				currentState = 93
			case 'F':
				currentState = 89
			case 'O':
				currentState = 73
			case 'S':
				currentState = 78
			case 'c':
				currentState = 69
			case 'r':
				currentState = 83
			case 'z':
				currentState = 64
			case 'M':
				currentState = 90
			case 'U':
				currentState = 67
			case 'l':
				currentState = 106
			case 'j':
				currentState = 84
			case 'n':
				currentState = 85
			case 'E':
				currentState = 95
			case 'N':
				currentState = 99
			case 'f':
				currentState = 102
			case 'A':
				currentState = 77
			case 'b':
				currentState = 60
			case 'd':
				currentState = 66
			case '_':
				currentState = 97
			case 'L':
				currentState = 61
			case 'V':
				currentState = 92
			case 's':
				currentState = 87
			case 'k':
				currentState = 57
			case 'y':
				currentState = 72
			case 'J':
				currentState = 59
			case 'X':
				currentState = 68
			case 'a':
				currentState = 104
			case 'u':
				currentState = 55
			case 'C':
				currentState = 86
			case 'I':
				currentState = 62
			case 'K':
				currentState = 63
			case 'R':
				currentState = 103
			case 'W':
				currentState = 56
			case 'h':
				currentState = 79
			case 'o':
				currentState = 58
			case 'p':
				currentState = 70
			case 'B':
				currentState = 88
			case 'D':
				currentState = 94
			case 'm':
				currentState = 54
			case 't':
				currentState = 76
			case 'Z':
				currentState = 101
			case 'e':
				currentState = 105
			case 'g':
				currentState = 82
			case 'i':
				currentState = 80
			case 'Q':
				currentState = 91
			case 'T':
				currentState = 74
			case 'Y':
				currentState = 75
			case 'G':
				currentState = 100
			case 'q':
				currentState = 81
			case 'v':
				currentState = 96
			case 'H':
				currentState = 98
			default:
				break outer
			}

		case 49:
			switch c {
			case 'Q':
				currentState = 91
			case 'm':
				currentState = 54
			case 'y':
				currentState = 72
			case 'A':
				currentState = 77
			case 'D':
				currentState = 94
			case 'J':
				currentState = 59
			case 'N':
				currentState = 99
			case 'O':
				currentState = 73
			case 'M':
				currentState = 90
			case 'U':
				currentState = 67
			case 'I':
				currentState = 62
			case 'r':
				currentState = 83
			case 'j':
				currentState = 84
			case 'k':
				currentState = 57
			case 'h':
				currentState = 79
			case 'H':
				currentState = 98
			case 'R':
				currentState = 103
			case 'c':
				currentState = 69
			case 'd':
				currentState = 66
			case 'f':
				currentState = 102
			case 'G':
				currentState = 100
			case 'L':
				currentState = 61
			case 'i':
				currentState = 80
			case 'x':
				currentState = 93
			case '_':
				currentState = 97
			case 'K':
				currentState = 63
			case 'T':
				currentState = 74
			case 's':
				currentState = 87
			case 'F':
				currentState = 89
			case 'u':
				currentState = 55
			case 'v':
				currentState = 96
			case 'E':
				currentState = 95
			case 'z':
				currentState = 64
			case 'V':
				currentState = 92
			case 'Y':
				currentState = 75
			case 'l':
				currentState = 106
			case 'C':
				currentState = 86
			case 'X':
				currentState = 68
			case 'p':
				currentState = 70
			case 'P':
				currentState = 65
			case 'a':
				currentState = 104
			case 't':
				currentState = 76
			case 'w':
				currentState = 71
			case 'W':
				currentState = 56
			case 'e':
				currentState = 105
			case 'n':
				currentState = 85
			case 'Z':
				currentState = 101
			case 'g':
				currentState = 82
			case 'B':
				currentState = 88
			case 'S':
				currentState = 78
			case 'b':
				currentState = 60
			case 'o':
				currentState = 58
			case 'q':
				currentState = 81
			default:
				break outer
			}

		case 94:
			switch c {
			case 'k':
				currentState = 57
			case 'Q':
				currentState = 91
			case 'R':
				currentState = 103
			case 'n':
				currentState = 85
			case 'p':
				currentState = 70
			case 't':
				currentState = 76
			case '_':
				currentState = 97
			case 'L':
				currentState = 61
			case 'H':
				currentState = 98
			case 'd':
				currentState = 66
			case 'j':
				currentState = 84
			case 'u':
				currentState = 55
			case 'B':
				currentState = 88
			case 'N':
				currentState = 99
			case 'm':
				currentState = 54
			case 'D':
				currentState = 94
			case 'F':
				currentState = 89
			case 'K':
				currentState = 63
			case 'b':
				currentState = 60
			case 'h':
				currentState = 79
			case 's':
				currentState = 87
			case 'v':
				currentState = 96
			case 'E':
				currentState = 95
			case 'M':
				currentState = 90
			case 'e':
				currentState = 105
			case 'o':
				currentState = 58
			case 'z':
				currentState = 64
			case 'G':
				currentState = 100
			case 'T':
				currentState = 74
			case 'W':
				currentState = 56
			case 'a':
				currentState = 104
			case 'i':
				currentState = 80
			case 'I':
				currentState = 62
			case 'P':
				currentState = 65
			case 'J':
				currentState = 59
			case 'U':
				currentState = 67
			case 'S':
				currentState = 78
			case 'V':
				currentState = 92
			case 'r':
				currentState = 83
			case 'C':
				currentState = 86
			case 'O':
				currentState = 73
			case 'X':
				currentState = 68
			case 'Y':
				currentState = 75
			case 'Z':
				currentState = 101
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case 'l':
				currentState = 106
			case 'w':
				currentState = 71
			case 'x':
				currentState = 93
			case 'y':
				currentState = 72
			case 'A':
				currentState = 77
			case 'q':
				currentState = 81
			case 'c':
				currentState = 69
			default:
				break outer
			}

		case 26:
			switch c {
			case 'J':
				currentState = 59
			case 'd':
				currentState = 66
			case 'm':
				currentState = 54
			case 'p':
				currentState = 70
			case 'w':
				currentState = 71
			case 'x':
				currentState = 93
			case 'B':
				currentState = 88
			case 'j':
				currentState = 84
			case 'u':
				currentState = 55
			case 'z':
				currentState = 64
			case 'a':
				currentState = 104
			case 'r':
				currentState = 83
			case 'b':
				currentState = 60
			case 'D':
				currentState = 94
			case 'L':
				currentState = 61
			case 'N':
				currentState = 99
			case 'Y':
				currentState = 75
			case 't':
				currentState = 76
			case 'V':
				currentState = 92
			case 'Z':
				currentState = 101
			case 'e':
				currentState = 105
			case 'q':
				currentState = 81
			case 'K':
				currentState = 63
			case 'T':
				currentState = 74
			case 's':
				currentState = 87
			case 'y':
				currentState = 72
			case 'G':
				currentState = 100
			case 'g':
				currentState = 82
			case 'S':
				currentState = 78
			case 'U':
				currentState = 67
			case 'c':
				currentState = 69
			case 'C':
				currentState = 86
			case 'I':
				currentState = 62
			case 'O':
				currentState = 73
			case 'i':
				currentState = 80
			case 'F':
				currentState = 89
			case 'P':
				currentState = 65
			case 'X':
				currentState = 68
			case 'h':
				currentState = 79
			case 'k':
				currentState = 57
			case 'E':
				currentState = 95
			case 'H':
				currentState = 98
			case 'W':
				currentState = 56
			case 'v':
				currentState = 96
			case '_':
				currentState = 97
			case 'M':
				currentState = 90
			case 'n':
				currentState = 85
			case 'o':
				currentState = 58
			case 'A':
				currentState = 77
			case 'Q':
				currentState = 91
			case 'R':
				currentState = 103
			case 'f':
				currentState = 102
			case 'l':
				currentState = 106
			default:
				break outer
			}

		case 50:
			switch c {
			case 'n':
				currentState = 85
			case 'E':
				currentState = 95
			case 'I':
				currentState = 62
			case 'N':
				currentState = 99
			case 'S':
				currentState = 78
			case 'e':
				currentState = 105
			case 'q':
				currentState = 81
			case '_':
				currentState = 97
			case 'a':
				currentState = 104
			case 'b':
				currentState = 60
			case 'i':
				currentState = 80
			case 'k':
				currentState = 57
			case 'v':
				currentState = 96
			case 'h':
				currentState = 79
			case 'P':
				currentState = 65
			case 'U':
				currentState = 67
			case 'c':
				currentState = 69
			case 'j':
				currentState = 84
			case 'u':
				currentState = 55
			case 'x':
				currentState = 93
			case 't':
				currentState = 76
			case 'y':
				currentState = 72
			case 'z':
				currentState = 64
			case 'Z':
				currentState = 101
			case 'm':
				currentState = 54
			case 'r':
				currentState = 83
			case 'X':
				currentState = 68
			case 'T':
				currentState = 74
			case 'C':
				currentState = 86
			case 'G':
				currentState = 100
			case 'H':
				currentState = 98
			case 'Y':
				currentState = 75
			case 'B':
				currentState = 88
			case 'M':
				currentState = 90
			case 'O':
				currentState = 73
			case 'D':
				currentState = 94
			case 'L':
				currentState = 61
			case 'Q':
				currentState = 91
			case 'R':
				currentState = 103
			case 'W':
				currentState = 56
			case 'w':
				currentState = 71
			case 'F':
				currentState = 89
			case 'l':
				currentState = 106
			case 'o':
				currentState = 58
			case 'd':
				currentState = 66
			case 'g':
				currentState = 82
			case 'p':
				currentState = 70
			case 'A':
				currentState = 77
			case 'K':
				currentState = 63
			case 'V':
				currentState = 92
			case 'J':
				currentState = 59
			case 'f':
				currentState = 102
			case 's':
				currentState = 87
			default:
				break outer
			}

		case 60:
			switch c {
			case 'M':
				currentState = 90
			case 'Y':
				currentState = 75
			case 'c':
				currentState = 69
			case 'l':
				currentState = 106
			case 'A':
				currentState = 77
			case 'C':
				currentState = 86
			case 'J':
				currentState = 59
			case 'a':
				currentState = 104
			case 'h':
				currentState = 79
			case 'K':
				currentState = 63
			case 'X':
				currentState = 68
			case 'n':
				currentState = 85
			case 'B':
				currentState = 88
			case 'W':
				currentState = 56
			case 'q':
				currentState = 81
			case 'f':
				currentState = 102
			case 'r':
				currentState = 83
			case 'x':
				currentState = 93
			case 'R':
				currentState = 103
			case 'p':
				currentState = 70
			case 'y':
				currentState = 72
			case 'D':
				currentState = 94
			case 'H':
				currentState = 98
			case 'N':
				currentState = 99
			case 'I':
				currentState = 62
			case 'L':
				currentState = 61
			case 'm':
				currentState = 54
			case 'k':
				currentState = 57
			case 'w':
				currentState = 71
			case 'P':
				currentState = 65
			case 'g':
				currentState = 82
			case 'F':
				currentState = 89
			case 'O':
				currentState = 73
			case '_':
				currentState = 97
			case 'V':
				currentState = 92
			case 'v':
				currentState = 96
			case 'z':
				currentState = 64
			case 't':
				currentState = 76
			case 'o':
				currentState = 58
			case 's':
				currentState = 87
			case 'Q':
				currentState = 91
			case 'b':
				currentState = 60
			case 'd':
				currentState = 66
			case 'i':
				currentState = 80
			case 'E':
				currentState = 95
			case 'G':
				currentState = 100
			case 'U':
				currentState = 67
			case 'j':
				currentState = 84
			case 'u':
				currentState = 55
			case 'e':
				currentState = 105
			case 'S':
				currentState = 78
			case 'T':
				currentState = 74
			case 'Z':
				currentState = 101
			default:
				break outer
			}

		case 64:
			switch c {
			case 'n':
				currentState = 85
			case 'r':
				currentState = 83
			case 'R':
				currentState = 103
			case 'Y':
				currentState = 75
			case 'Z':
				currentState = 101
			case 'b':
				currentState = 60
			case 'C':
				currentState = 86
			case 'I':
				currentState = 62
			case 'Q':
				currentState = 91
			case 'U':
				currentState = 67
			case 'w':
				currentState = 71
			case 'f':
				currentState = 102
			case 'i':
				currentState = 80
			case 's':
				currentState = 87
			case 'S':
				currentState = 78
			case 'H':
				currentState = 98
			case 'M':
				currentState = 90
			case 'P':
				currentState = 65
			case 'D':
				currentState = 94
			case 'k':
				currentState = 57
			case 'q':
				currentState = 81
			case 'W':
				currentState = 56
			case 'G':
				currentState = 100
			case 'a':
				currentState = 104
			case 'z':
				currentState = 64
			case 'A':
				currentState = 77
			case 'u':
				currentState = 55
			case 'e':
				currentState = 105
			case 'E':
				currentState = 95
			case 'N':
				currentState = 99
			case 'p':
				currentState = 70
			case 'B':
				currentState = 88
			case 'v':
				currentState = 96
			case 'm':
				currentState = 54
			case 't':
				currentState = 76
			case 'g':
				currentState = 82
			case 'O':
				currentState = 73
			case 'T':
				currentState = 74
			case 'V':
				currentState = 92
			case 'h':
				currentState = 79
			case 'j':
				currentState = 84
			case 'l':
				currentState = 106
			case 'F':
				currentState = 89
			case 'd':
				currentState = 66
			case 'o':
				currentState = 58
			case 'X':
				currentState = 68
			case '_':
				currentState = 97
			case 'y':
				currentState = 72
			case 'L':
				currentState = 61
			case 'K':
				currentState = 63
			case 'c':
				currentState = 69
			case 'x':
				currentState = 93
			case 'J':
				currentState = 59
			default:
				break outer
			}

		case 67:
			switch c {
			case 'R':
				currentState = 103
			case 'm':
				currentState = 54
			case '_':
				currentState = 97
			case 'G':
				currentState = 100
			case 'K':
				currentState = 63
			case 'Y':
				currentState = 75
			case 'e':
				currentState = 105
			case 'o':
				currentState = 58
			case 'Q':
				currentState = 91
			case 'f':
				currentState = 102
			case 'h':
				currentState = 79
			case 'C':
				currentState = 86
			case 'H':
				currentState = 98
			case 'g':
				currentState = 82
			case 'k':
				currentState = 57
			case 's':
				currentState = 87
			case 'O':
				currentState = 73
			case 'l':
				currentState = 106
			case 'w':
				currentState = 71
			case 'B':
				currentState = 88
			case 'p':
				currentState = 70
			case 'F':
				currentState = 89
			case 'a':
				currentState = 104
			case 'V':
				currentState = 92
			case 't':
				currentState = 76
			case 'x':
				currentState = 93
			case 'L':
				currentState = 61
			case 'Z':
				currentState = 101
			case 'n':
				currentState = 85
			case 'q':
				currentState = 81
			case 'y':
				currentState = 72
			case 'M':
				currentState = 90
			case 'P':
				currentState = 65
			case 'r':
				currentState = 83
			case 'z':
				currentState = 64
			case 'c':
				currentState = 69
			case 'i':
				currentState = 80
			case 'D':
				currentState = 94
			case 'E':
				currentState = 95
			case 'I':
				currentState = 62
			case 'J':
				currentState = 59
			case 'S':
				currentState = 78
			case 'X':
				currentState = 68
			case 'u':
				currentState = 55
			case 'j':
				currentState = 84
			case 'v':
				currentState = 96
			case 'A':
				currentState = 77
			case 'N':
				currentState = 99
			case 'U':
				currentState = 67
			case 'd':
				currentState = 66
			case 'T':
				currentState = 74
			case 'W':
				currentState = 56
			case 'b':
				currentState = 60
			default:
				break outer
			}

		case 78:
			switch c {
			case 'o':
				currentState = 58
			case 's':
				currentState = 87
			case 'w':
				currentState = 71
			case 'O':
				currentState = 73
			case 'U':
				currentState = 67
			case 'f':
				currentState = 102
			case 'C':
				currentState = 86
			case 'D':
				currentState = 94
			case 'c':
				currentState = 69
			case 'k':
				currentState = 57
			case 'r':
				currentState = 83
			case 'u':
				currentState = 55
			case 'A':
				currentState = 77
			case 'Q':
				currentState = 91
			case 'a':
				currentState = 104
			case 'F':
				currentState = 89
			case 'K':
				currentState = 63
			case 'y':
				currentState = 72
			case 'd':
				currentState = 66
			case 'z':
				currentState = 64
			case 'L':
				currentState = 61
			case 'R':
				currentState = 103
			case 'p':
				currentState = 70
			case 'E':
				currentState = 95
			case 'G':
				currentState = 100
			case 'b':
				currentState = 60
			case 'h':
				currentState = 79
			case 'q':
				currentState = 81
			case 'I':
				currentState = 62
			case 'T':
				currentState = 74
			case 'g':
				currentState = 82
			case 'H':
				currentState = 98
			case '_':
				currentState = 97
			case 'B':
				currentState = 88
			case 'Z':
				currentState = 101
			case 'v':
				currentState = 96
			case 'V':
				currentState = 92
			case 'j':
				currentState = 84
			case 'Y':
				currentState = 75
			case 'n':
				currentState = 85
			case 'J':
				currentState = 59
			case 'N':
				currentState = 99
			case 'W':
				currentState = 56
			case 'l':
				currentState = 106
			case 'x':
				currentState = 93
			case 'M':
				currentState = 90
			case 'S':
				currentState = 78
			case 'e':
				currentState = 105
			case 'i':
				currentState = 80
			case 't':
				currentState = 76
			case 'P':
				currentState = 65
			case 'X':
				currentState = 68
			case 'm':
				currentState = 54
			default:
				break outer
			}

		case 48:
			switch c {
			case 'Y':
				currentState = 75
			case 'R':
				currentState = 103
			case 'H':
				currentState = 98
			case 'T':
				currentState = 74
			case 't':
				currentState = 76
			case 'x':
				currentState = 93
			case 'F':
				currentState = 89
			case 'n':
				currentState = 85
			case 'C':
				currentState = 86
			case 'I':
				currentState = 62
			case 'K':
				currentState = 63
			case 'M':
				currentState = 90
			case 'h':
				currentState = 79
			case 'A':
				currentState = 77
			case 'U':
				currentState = 67
			case 'W':
				currentState = 56
			case 'X':
				currentState = 68
			case 'f':
				currentState = 102
			case 'S':
				currentState = 78
			case 'b':
				currentState = 60
			case 'G':
				currentState = 100
			case 'u':
				currentState = 55
			case 'a':
				currentState = 104
			case 'd':
				currentState = 66
			case 'c':
				currentState = 69
			case 'D':
				currentState = 94
			case 'L':
				currentState = 61
			case 'g':
				currentState = 82
			case 'i':
				currentState = 80
			case 'o':
				currentState = 58
			case 'r':
				currentState = 83
			case 'v':
				currentState = 96
			case 'P':
				currentState = 65
			case 'k':
				currentState = 57
			case 'm':
				currentState = 54
			case 'N':
				currentState = 99
			case 'l':
				currentState = 106
			case 'j':
				currentState = 84
			case 'p':
				currentState = 70
			case 'V':
				currentState = 92
			case 'J':
				currentState = 59
			case 'Z':
				currentState = 101
			case 'E':
				currentState = 95
			case 's':
				currentState = 87
			case 'B':
				currentState = 88
			case 'Q':
				currentState = 91
			case 'e':
				currentState = 105
			case 'q':
				currentState = 81
			case 'w':
				currentState = 71
			case 'y':
				currentState = 72
			case 'z':
				currentState = 64
			case '_':
				currentState = 97
			case 'O':
				currentState = 73
			default:
				break outer
			}

		case 33:
			switch c {
			case 'J':
				currentState = 59
			case 'a':
				currentState = 104
			case 'O':
				currentState = 73
			case 'D':
				currentState = 94
			case 'I':
				currentState = 62
			case 'b':
				currentState = 60
			case 'f':
				currentState = 102
			case 'm':
				currentState = 54
			case 'w':
				currentState = 71
			case 'M':
				currentState = 90
			case 'V':
				currentState = 92
			case 'r':
				currentState = 83
			case 'U':
				currentState = 67
			case 'p':
				currentState = 70
			case '_':
				currentState = 97
			case 'B':
				currentState = 88
			case 'Q':
				currentState = 91
			case 'S':
				currentState = 78
			case 'T':
				currentState = 74
			case 'L':
				currentState = 61
			case 'c':
				currentState = 69
			case 'o':
				currentState = 58
			case 'q':
				currentState = 81
			case 'h':
				currentState = 79
			case 'n':
				currentState = 85
			case 's':
				currentState = 87
			case 't':
				currentState = 76
			case 'K':
				currentState = 63
			case 'W':
				currentState = 56
			case 'C':
				currentState = 86
			case 'j':
				currentState = 84
			case 'l':
				currentState = 106
			case 'z':
				currentState = 64
			case 'F':
				currentState = 89
			case 'X':
				currentState = 68
			case 'v':
				currentState = 96
			case 'y':
				currentState = 72
			case 'x':
				currentState = 93
			case 'A':
				currentState = 77
			case 'G':
				currentState = 100
			case 'N':
				currentState = 99
			case 'R':
				currentState = 103
			case 'e':
				currentState = 105
			case 'g':
				currentState = 82
			case 'k':
				currentState = 57
			case 'i':
				currentState = 80
			case 'u':
				currentState = 55
			case 'E':
				currentState = 95
			case 'H':
				currentState = 98
			case 'P':
				currentState = 65
			case 'Y':
				currentState = 75
			case 'Z':
				currentState = 101
			case 'd':
				currentState = 66
			default:
				break outer
			}

		case 59:
			switch c {
			case 'B':
				currentState = 88
			case 'Z':
				currentState = 101
			case 'j':
				currentState = 84
			case 'u':
				currentState = 55
			case 'K':
				currentState = 63
			case 'Y':
				currentState = 75
			case '_':
				currentState = 97
			case 'd':
				currentState = 66
			case 't':
				currentState = 76
			case 'k':
				currentState = 57
			case 'E':
				currentState = 95
			case 'O':
				currentState = 73
			case 'n':
				currentState = 85
			case 'y':
				currentState = 72
			case 'g':
				currentState = 82
			case 'm':
				currentState = 54
			case 'r':
				currentState = 83
			case 'l':
				currentState = 106
			case 'C':
				currentState = 86
			case 'F':
				currentState = 89
			case 'b':
				currentState = 60
			case 'Q':
				currentState = 91
			case 'T':
				currentState = 74
			case 'c':
				currentState = 69
			case 'p':
				currentState = 70
			case 's':
				currentState = 87
			case 'x':
				currentState = 93
			case 'J':
				currentState = 59
			case 'S':
				currentState = 78
			case 'h':
				currentState = 79
			case 'A':
				currentState = 77
			case 'D':
				currentState = 94
			case 'H':
				currentState = 98
			case 'a':
				currentState = 104
			case 'w':
				currentState = 71
			case 'G':
				currentState = 100
			case 'X':
				currentState = 68
			case 'V':
				currentState = 92
			case 'W':
				currentState = 56
			case 'i':
				currentState = 80
			case 'q':
				currentState = 81
			case 'e':
				currentState = 105
			case 'o':
				currentState = 58
			case 'z':
				currentState = 64
			case 'N':
				currentState = 99
			case 'R':
				currentState = 103
			case 'U':
				currentState = 67
			case 'f':
				currentState = 102
			case 'v':
				currentState = 96
			case 'L':
				currentState = 61
			case 'P':
				currentState = 65
			case 'I':
				currentState = 62
			case 'M':
				currentState = 90
			default:
				break outer
			}

		case 69:
			switch c {
			case 'J':
				currentState = 59
			case 'N':
				currentState = 99
			case 'D':
				currentState = 94
			case 'U':
				currentState = 67
			case 'j':
				currentState = 84
			case 'C':
				currentState = 86
			case 'K':
				currentState = 63
			case 'g':
				currentState = 82
			case 'v':
				currentState = 96
			case 'I':
				currentState = 62
			case 'V':
				currentState = 92
			case 'b':
				currentState = 60
			case 'd':
				currentState = 66
			case 'o':
				currentState = 58
			case 'q':
				currentState = 81
			case 'P':
				currentState = 65
			case 'c':
				currentState = 69
			case '_':
				currentState = 97
			case 'e':
				currentState = 105
			case 'n':
				currentState = 85
			case 'h':
				currentState = 79
			case 'k':
				currentState = 57
			case 'L':
				currentState = 61
			case 'S':
				currentState = 78
			case 't':
				currentState = 76
			case 'F':
				currentState = 89
			case 'O':
				currentState = 73
			case 'T':
				currentState = 74
			case 'G':
				currentState = 100
			case 'R':
				currentState = 103
			case 'f':
				currentState = 102
			case 'r':
				currentState = 83
			case 'y':
				currentState = 72
			case 'B':
				currentState = 88
			case 'H':
				currentState = 98
			case 'Y':
				currentState = 75
			case 'a':
				currentState = 104
			case 'i':
				currentState = 80
			case 'u':
				currentState = 55
			case 'w':
				currentState = 71
			case 'z':
				currentState = 64
			case 'E':
				currentState = 95
			case 'l':
				currentState = 106
			case 's':
				currentState = 87
			case 'M':
				currentState = 90
			case 'm':
				currentState = 54
			case 'x':
				currentState = 93
			case 'A':
				currentState = 77
			case 'W':
				currentState = 56
			case 'X':
				currentState = 68
			case 'p':
				currentState = 70
			case 'Q':
				currentState = 91
			case 'Z':
				currentState = 101
			default:
				break outer
			}

		case 98:
			switch c {
			case 'Z':
				currentState = 101
			case 'b':
				currentState = 60
			case 'e':
				currentState = 105
			case 'B':
				currentState = 88
			case 'G':
				currentState = 100
			case 'Y':
				currentState = 75
			case 'W':
				currentState = 56
			case 'K':
				currentState = 63
			case 'M':
				currentState = 90
			case 'Q':
				currentState = 91
			case 'y':
				currentState = 72
			case 'L':
				currentState = 61
			case 'S':
				currentState = 78
			case 'v':
				currentState = 96
			case 'I':
				currentState = 62
			case 'k':
				currentState = 57
			case 'q':
				currentState = 81
			case 's':
				currentState = 87
			case 'w':
				currentState = 71
			case 'O':
				currentState = 73
			case 'c':
				currentState = 69
			case 'g':
				currentState = 82
			case 'd':
				currentState = 66
			case 'h':
				currentState = 79
			case 'J':
				currentState = 59
			case 'V':
				currentState = 92
			case 'a':
				currentState = 104
			case 'D':
				currentState = 94
			case 'n':
				currentState = 85
			case 'F':
				currentState = 89
			case 'H':
				currentState = 98
			case 'X':
				currentState = 68
			case '_':
				currentState = 97
			case 'i':
				currentState = 80
			case 'l':
				currentState = 106
			case 'z':
				currentState = 64
			case 'A':
				currentState = 77
			case 'E':
				currentState = 95
			case 'U':
				currentState = 67
			case 'N':
				currentState = 99
			case 'R':
				currentState = 103
			case 'j':
				currentState = 84
			case 'm':
				currentState = 54
			case 't':
				currentState = 76
			case 'T':
				currentState = 74
			case 'f':
				currentState = 102
			case 'r':
				currentState = 83
			case 'p':
				currentState = 70
			case 'u':
				currentState = 55
			case 'x':
				currentState = 93
			case 'C':
				currentState = 86
			case 'P':
				currentState = 65
			case 'o':
				currentState = 58
			default:
				break outer
			}

		case 6:
			switch c {
			case 'D':
				currentState = 94
			case 'F':
				currentState = 89
			case 'N':
				currentState = 99
			case 'B':
				currentState = 88
			case 'T':
				currentState = 74
			case 'a':
				currentState = 104
			case 'd':
				currentState = 66
			case 'm':
				currentState = 54
			case 'o':
				currentState = 58
			case 'O':
				currentState = 73
			case 'y':
				currentState = 72
			case '_':
				currentState = 97
			case 'L':
				currentState = 61
			case 'P':
				currentState = 65
			case 'R':
				currentState = 103
			case 'Z':
				currentState = 101
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case 'x':
				currentState = 93
			case 'j':
				currentState = 84
			case 'E':
				currentState = 95
			case 'Q':
				currentState = 91
			case 'b':
				currentState = 60
			case 'k':
				currentState = 57
			case 'n':
				currentState = 85
			case 'q':
				currentState = 81
			case 'w':
				currentState = 71
			case 'J':
				currentState = 59
			case 'h':
				currentState = 79
			case 'A':
				currentState = 77
			case 'M':
				currentState = 90
			case 'Y':
				currentState = 75
			case 'c':
				currentState = 69
			case 's':
				currentState = 87
			case 'U':
				currentState = 67
			case 'G':
				currentState = 100
			case 'V':
				currentState = 92
			case 'W':
				currentState = 56
			case 'X':
				currentState = 68
			case 'H':
				currentState = 98
			case 'u':
				currentState = 55
			case 'K':
				currentState = 63
			case 'i':
				currentState = 80
			case 'v':
				currentState = 96
			case 'r':
				currentState = 83
			case 't':
				currentState = 76
			case 'C':
				currentState = 86
			case 'I':
				currentState = 62
			case 'S':
				currentState = 78
			case 'e':
				currentState = 105
			case 'l':
				currentState = 106
			case 'p':
				currentState = 70
			case 'z':
				currentState = 64
			default:
				break outer
			}

		case 79:
			switch c {
			case 'b':
				currentState = 60
			case 'h':
				currentState = 79
			case 'e':
				currentState = 105
			case 'k':
				currentState = 57
			case 's':
				currentState = 87
			case 't':
				currentState = 76
			case 'D':
				currentState = 94
			case 'G':
				currentState = 100
			case 'K':
				currentState = 63
			case 'M':
				currentState = 90
			case 'v':
				currentState = 96
			case 'Q':
				currentState = 91
			case 'A':
				currentState = 77
			case 'C':
				currentState = 86
			case 'H':
				currentState = 98
			case 'X':
				currentState = 68
			case 'c':
				currentState = 69
			case 'z':
				currentState = 64
			case 'E':
				currentState = 95
			case 'j':
				currentState = 84
			case '_':
				currentState = 97
			case 'u':
				currentState = 55
			case 'L':
				currentState = 61
			case 'Y':
				currentState = 75
			case 'Z':
				currentState = 101
			case 'p':
				currentState = 70
			case 'r':
				currentState = 83
			case 'S':
				currentState = 78
			case 'T':
				currentState = 74
			case 'w':
				currentState = 71
			case 'N':
				currentState = 99
			case 'o':
				currentState = 58
			case 'y':
				currentState = 72
			case 'B':
				currentState = 88
			case 'F':
				currentState = 89
			case 'a':
				currentState = 104
			case 'n':
				currentState = 85
			case 'm':
				currentState = 54
			case 'P':
				currentState = 65
			case 'd':
				currentState = 66
			case 'i':
				currentState = 80
			case 'W':
				currentState = 56
			case 'l':
				currentState = 106
			case 'x':
				currentState = 93
			case 'I':
				currentState = 62
			case 'O':
				currentState = 73
			case 'U':
				currentState = 67
			case 'V':
				currentState = 92
			case 'q':
				currentState = 81
			case 'J':
				currentState = 59
			case 'R':
				currentState = 103
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			default:
				break outer
			}

		case 45:
			switch c {
			case 'E':
				currentState = 95
			case 'M':
				currentState = 90
			case 'k':
				currentState = 57
			case 'v':
				currentState = 96
			case 'B':
				currentState = 88
			case 'G':
				currentState = 100
			case 'J':
				currentState = 59
			case 'Z':
				currentState = 101
			case 'H':
				currentState = 98
			case 'I':
				currentState = 62
			case 'S':
				currentState = 78
			case 'T':
				currentState = 74
			case 'Y':
				currentState = 75
			case 'j':
				currentState = 84
			case 'a':
				currentState = 104
			case 'c':
				currentState = 69
			case 'l':
				currentState = 106
			case '_':
				currentState = 97
			case 'L':
				currentState = 61
			case 'b':
				currentState = 60
			case 'f':
				currentState = 102
			case 'F':
				currentState = 89
			case 'e':
				currentState = 105
			case 'A':
				currentState = 77
			case 'O':
				currentState = 73
			case 'r':
				currentState = 83
			case 'x':
				currentState = 93
			case 'R':
				currentState = 103
			case 'U':
				currentState = 67
			case 'V':
				currentState = 92
			case 'd':
				currentState = 66
			case 'm':
				currentState = 54
			case 't':
				currentState = 76
			case 'P':
				currentState = 65
			case 'X':
				currentState = 68
			case 'y':
				currentState = 72
			case 'C':
				currentState = 86
			case 'h':
				currentState = 79
			case 'i':
				currentState = 80
			case 'p':
				currentState = 70
			case 'z':
				currentState = 64
			case 'N':
				currentState = 99
			case 's':
				currentState = 87
			case 'u':
				currentState = 55
			case 'w':
				currentState = 71
			case 'D':
				currentState = 94
			case 'K':
				currentState = 63
			case 'n':
				currentState = 85
			case 'g':
				currentState = 82
			case 'o':
				currentState = 58
			case 'q':
				currentState = 81
			case 'Q':
				currentState = 91
			case 'W':
				currentState = 56
			default:
				break outer
			}

		case 10:
			switch c {
			case 'Y':
				currentState = 75
			case 'o':
				currentState = 58
			case 'p':
				currentState = 70
			case 's':
				currentState = 87
			case 'J':
				currentState = 59
			case 'S':
				currentState = 78
			case 'm':
				currentState = 54
			case 'W':
				currentState = 56
			case 'Z':
				currentState = 101
			case 'a':
				currentState = 104
			case 'h':
				currentState = 79
			case 'F':
				currentState = 89
			case 'M':
				currentState = 90
			case 'U':
				currentState = 67
			case 'y':
				currentState = 72
			case 'X':
				currentState = 68
			case 'e':
				currentState = 105
			case 't':
				currentState = 76
			case 'B':
				currentState = 88
			case 'i':
				currentState = 80
			case 'n':
				currentState = 85
			case 'd':
				currentState = 66
			case 'C':
				currentState = 86
			case 'E':
				currentState = 95
			case 'K':
				currentState = 63
			case 'I':
				currentState = 62
			case 'k':
				currentState = 57
			case 'v':
				currentState = 96
			case 'x':
				currentState = 93
			case 'D':
				currentState = 94
			case 'G':
				currentState = 100
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case 'z':
				currentState = 64
			case 'N':
				currentState = 99
			case 'Q':
				currentState = 91
			case 'b':
				currentState = 60
			case 'q':
				currentState = 81
			case 'w':
				currentState = 71
			case '_':
				currentState = 97
			case 'j':
				currentState = 84
			case 'l':
				currentState = 106
			case 'r':
				currentState = 83
			case 'A':
				currentState = 77
			case 'R':
				currentState = 103
			case 'c':
				currentState = 69
			case 'P':
				currentState = 65
			case 'T':
				currentState = 74
			case 'V':
				currentState = 92
			case 'u':
				currentState = 55
			case 'H':
				currentState = 98
			case 'L':
				currentState = 61
			case 'O':
				currentState = 73
			default:
				break outer
			}

		case 19:
			switch c {
			case 'D':
				currentState = 94
			case 'q':
				currentState = 81
			case 'm':
				currentState = 54
			case 'E':
				currentState = 95
			case 'S':
				currentState = 78
			case 'd':
				currentState = 66
			case 'g':
				currentState = 82
			case 'h':
				currentState = 79
			case 'o':
				currentState = 58
			case 'L':
				currentState = 61
			case 'V':
				currentState = 92
			case 'u':
				currentState = 55
			case 'v':
				currentState = 96
			case 'R':
				currentState = 103
			case 'Z':
				currentState = 101
			case 'r':
				currentState = 83
			case 'A':
				currentState = 77
			case 'C':
				currentState = 86
			case 'F':
				currentState = 89
			case 'K':
				currentState = 63
			case 'j':
				currentState = 84
			case 'J':
				currentState = 59
			case 'b':
				currentState = 60
			case 'e':
				currentState = 105
			case 's':
				currentState = 87
			case 'w':
				currentState = 71
			case 'G':
				currentState = 100
			case 'H':
				currentState = 98
			case 'a':
				currentState = 104
			case 'k':
				currentState = 57
			case 'N':
				currentState = 99
			case 'O':
				currentState = 73
			case 'f':
				currentState = 102
			case 'p':
				currentState = 70
			case 'z':
				currentState = 64
			case '_':
				currentState = 97
			case 'I':
				currentState = 62
			case 'T':
				currentState = 74
			case 'X':
				currentState = 68
			case 't':
				currentState = 76
			case 'x':
				currentState = 93
			case 'y':
				currentState = 72
			case 'B':
				currentState = 88
			case 'P':
				currentState = 65
			case 'i':
				currentState = 80
			case 'l':
				currentState = 106
			case 'Y':
				currentState = 75
			case 'c':
				currentState = 69
			case 'n':
				currentState = 85
			case 'M':
				currentState = 90
			case 'Q':
				currentState = 91
			case 'U':
				currentState = 67
			case 'W':
				currentState = 56
			default:
				break outer
			}

		case 40:
			switch c {
			case 'S':
				currentState = 78
			case 'f':
				currentState = 102
			case 'a':
				currentState = 104
			case 'g':
				currentState = 82
			case 'i':
				currentState = 80
			case 'u':
				currentState = 55
			case 'W':
				currentState = 56
			case 'm':
				currentState = 54
			case 'A':
				currentState = 77
			case 'B':
				currentState = 88
			case 'L':
				currentState = 61
			case 'T':
				currentState = 74
			case 'Y':
				currentState = 75
			case 'D':
				currentState = 94
			case 'H':
				currentState = 98
			case 'J':
				currentState = 59
			case 'x':
				currentState = 93
			case 'w':
				currentState = 71
			case 'M':
				currentState = 90
			case 'N':
				currentState = 99
			case 'X':
				currentState = 68
			case 'o':
				currentState = 58
			case 'q':
				currentState = 81
			case 'Z':
				currentState = 101
			case 'l':
				currentState = 106
			case 'O':
				currentState = 73
			case 'C':
				currentState = 86
			case 'P':
				currentState = 65
			case 'Q':
				currentState = 91
			case 'G':
				currentState = 100
			case 'b':
				currentState = 60
			case 'n':
				currentState = 85
			case 'r':
				currentState = 83
			case 'I':
				currentState = 62
			case 'K':
				currentState = 63
			case 'e':
				currentState = 105
			case 'k':
				currentState = 57
			case 'v':
				currentState = 96
			case 'U':
				currentState = 67
			case 'p':
				currentState = 70
			case 'd':
				currentState = 66
			case 's':
				currentState = 87
			case 'z':
				currentState = 64
			case 'E':
				currentState = 95
			case 'h':
				currentState = 79
			case '_':
				currentState = 97
			case 'F':
				currentState = 89
			case 'y':
				currentState = 72
			case 'R':
				currentState = 103
			case 'V':
				currentState = 92
			case 'c':
				currentState = 69
			case 'j':
				currentState = 84
			case 't':
				currentState = 76
			default:
				break outer
			}

		case 53:
			switch c {
			case 'C':
				currentState = 86
			case 'I':
				currentState = 62
			case 'p':
				currentState = 70
			case 'q':
				currentState = 81
			case 'G':
				currentState = 100
			case 'T':
				currentState = 74
			case 'g':
				currentState = 82
			case 'm':
				currentState = 54
			case 'u':
				currentState = 55
			case 'A':
				currentState = 77
			case 'c':
				currentState = 69
			case 'f':
				currentState = 102
			case 'D':
				currentState = 94
			case 'k':
				currentState = 57
			case 'r':
				currentState = 83
			case 'B':
				currentState = 88
			case 'U':
				currentState = 67
			case 'a':
				currentState = 104
			case 'b':
				currentState = 60
			case 'd':
				currentState = 66
			case 'E':
				currentState = 95
			case 'e':
				currentState = 105
			case 'P':
				currentState = 65
			case 'Q':
				currentState = 91
			case 'R':
				currentState = 103
			case 'W':
				currentState = 56
			case 'Z':
				currentState = 101
			case 'i':
				currentState = 80
			case 'x':
				currentState = 93
			case 'M':
				currentState = 90
			case 'Y':
				currentState = 75
			case 'o':
				currentState = 58
			case 'S':
				currentState = 78
			case 'X':
				currentState = 68
			case 'l':
				currentState = 106
			case 'n':
				currentState = 85
			case 'w':
				currentState = 71
			case 'y':
				currentState = 72
			case 'F':
				currentState = 89
			case 'J':
				currentState = 59
			case 'L':
				currentState = 61
			case 'j':
				currentState = 84
			case 't':
				currentState = 76
			case 'H':
				currentState = 98
			case 'O':
				currentState = 73
			case 'h':
				currentState = 79
			case 'V':
				currentState = 92
			case 's':
				currentState = 87
			case 'K':
				currentState = 63
			case 'v':
				currentState = 96
			case 'N':
				currentState = 99
			case 'z':
				currentState = 64
			case '_':
				currentState = 97
			default:
				break outer
			}

		case 81:
			switch c {
			case 'R':
				currentState = 103
			case 'X':
				currentState = 68
			case 'c':
				currentState = 69
			case 'k':
				currentState = 57
			case 'y':
				currentState = 72
			case '_':
				currentState = 97
			case 'L':
				currentState = 61
			case 'N':
				currentState = 99
			case 'l':
				currentState = 106
			case 'u':
				currentState = 55
			case 'z':
				currentState = 64
			case 'T':
				currentState = 74
			case 'i':
				currentState = 80
			case 'D':
				currentState = 94
			case 'e':
				currentState = 105
			case 'G':
				currentState = 100
			case 'j':
				currentState = 84
			case 'r':
				currentState = 83
			case 'V':
				currentState = 92
			case 'o':
				currentState = 58
			case 'I':
				currentState = 62
			case 'g':
				currentState = 82
			case 'C':
				currentState = 86
			case 'Q':
				currentState = 91
			case 'S':
				currentState = 78
			case 'h':
				currentState = 79
			case 'B':
				currentState = 88
			case 'U':
				currentState = 67
			case 'P':
				currentState = 65
			case 'J':
				currentState = 59
			case 's':
				currentState = 87
			case 'v':
				currentState = 96
			case 'x':
				currentState = 93
			case 'A':
				currentState = 77
			case 'a':
				currentState = 104
			case 'H':
				currentState = 98
			case 'q':
				currentState = 81
			case 'b':
				currentState = 60
			case 'd':
				currentState = 66
			case 'n':
				currentState = 85
			case 'Y':
				currentState = 75
			case 'f':
				currentState = 102
			case 't':
				currentState = 76
			case 'w':
				currentState = 71
			case 'E':
				currentState = 95
			case 'F':
				currentState = 89
			case 'O':
				currentState = 73
			case 'W':
				currentState = 56
			case 'Z':
				currentState = 101
			case 'm':
				currentState = 54
			case 'p':
				currentState = 70
			case 'K':
				currentState = 63
			case 'M':
				currentState = 90
			default:
				break outer
			}

		case 5:
			switch c {
			case 'Y':
				currentState = 75
			case 'F':
				currentState = 89
			case 'V':
				currentState = 92
			case 's':
				currentState = 87
			case 'L':
				currentState = 61
			case 'S':
				currentState = 78
			case 'a':
				currentState = 104
			case 'd':
				currentState = 66
			case 'k':
				currentState = 57
			case 'p':
				currentState = 70
			case 'x':
				currentState = 93
			case 'O':
				currentState = 73
			case 'B':
				currentState = 88
			case 'Q':
				currentState = 91
			case 'j':
				currentState = 84
			case 'n':
				currentState = 85
			case 'w':
				currentState = 71
			case 'G':
				currentState = 100
			case 'T':
				currentState = 74
			case 'W':
				currentState = 56
			case 'i':
				currentState = 80
			case 'o':
				currentState = 58
			case 'z':
				currentState = 64
			case 'N':
				currentState = 99
			case 'Z':
				currentState = 101
			case 'h':
				currentState = 79
			case 'l':
				currentState = 106
			case 'r':
				currentState = 83
			case 'u':
				currentState = 55
			case 'c':
				currentState = 69
			case 't':
				currentState = 76
			case 'D':
				currentState = 94
			case 'K':
				currentState = 63
			case '_':
				currentState = 97
			case 'R':
				currentState = 103
			case 'g':
				currentState = 82
			case 'A':
				currentState = 77
			case 'H':
				currentState = 98
			case 'b':
				currentState = 60
			case 'f':
				currentState = 102
			case 'J':
				currentState = 59
			case 'X':
				currentState = 68
			case 'e':
				currentState = 105
			case 'q':
				currentState = 81
			case 'C':
				currentState = 86
			case 'E':
				currentState = 95
			case 'v':
				currentState = 96
			case 'y':
				currentState = 72
			case 'U':
				currentState = 67
			case 'm':
				currentState = 54
			case 'I':
				currentState = 62
			case 'M':
				currentState = 90
			case 'P':
				currentState = 65
			default:
				break outer
			}

		case 41:
			switch c {
			case 'Y':
				currentState = 75
			case 'V':
				currentState = 92
			case 'g':
				currentState = 82
			case 'R':
				currentState = 103
			case 'N':
				currentState = 99
			case 'a':
				currentState = 104
			case 'j':
				currentState = 84
			case 'k':
				currentState = 57
			case 'X':
				currentState = 68
			case 'O':
				currentState = 73
			case 'K':
				currentState = 63
			case 'U':
				currentState = 67
			case 'd':
				currentState = 66
			case 'v':
				currentState = 96
			case 'C':
				currentState = 86
			case 'J':
				currentState = 59
			case 'Q':
				currentState = 91
			case 'h':
				currentState = 79
			case 'E':
				currentState = 95
			case 'T':
				currentState = 74
			case 'W':
				currentState = 56
			case 'b':
				currentState = 60
			case 'n':
				currentState = 85
			case 'q':
				currentState = 81
			case 'F':
				currentState = 89
			case 's':
				currentState = 87
			case '_':
				currentState = 97
			case 'L':
				currentState = 61
			case 'I':
				currentState = 62
			case 'e':
				currentState = 105
			case 'w':
				currentState = 71
			case 'A':
				currentState = 77
			case 'm':
				currentState = 54
			case 'u':
				currentState = 55
			case 'x':
				currentState = 93
			case 'M':
				currentState = 90
			case 'f':
				currentState = 102
			case 'o':
				currentState = 58
			case 't':
				currentState = 76
			case 'y':
				currentState = 72
			case 'H':
				currentState = 98
			case 'c':
				currentState = 69
			case 'i':
				currentState = 80
			case 'p':
				currentState = 70
			case 'r':
				currentState = 83
			case 'z':
				currentState = 64
			case 'P':
				currentState = 65
			case 'D':
				currentState = 94
			case 'G':
				currentState = 100
			case 'l':
				currentState = 106
			case 'B':
				currentState = 88
			case 'Z':
				currentState = 101
			case 'S':
				currentState = 78
			default:
				break outer
			}

		case 38:
			switch c {
			case 'v':
				currentState = 96
			case 'Q':
				currentState = 91
			case 'Z':
				currentState = 101
			case 'l':
				currentState = 106
			case 'r':
				currentState = 83
			case '_':
				currentState = 97
			case 'I':
				currentState = 62
			case 'V':
				currentState = 92
			case 'x':
				currentState = 93
			case 'y':
				currentState = 72
			case 'd':
				currentState = 66
			case 'f':
				currentState = 102
			case 'm':
				currentState = 54
			case 'S':
				currentState = 78
			case 'T':
				currentState = 74
			case 'X':
				currentState = 68
			case 'Y':
				currentState = 75
			case 'F':
				currentState = 89
			case 'L':
				currentState = 61
			case 'z':
				currentState = 64
			case 'R':
				currentState = 103
			case 'c':
				currentState = 69
			case 'q':
				currentState = 81
			case 's':
				currentState = 87
			case 'e':
				currentState = 105
			case 'i':
				currentState = 80
			case 'o':
				currentState = 58
			case 'j':
				currentState = 84
			case 'p':
				currentState = 70
			case 'w':
				currentState = 71
			case 'D':
				currentState = 94
			case 'E':
				currentState = 95
			case 'P':
				currentState = 65
			case 'W':
				currentState = 56
			case 'C':
				currentState = 86
			case 'H':
				currentState = 98
			case 'N':
				currentState = 99
			case 'h':
				currentState = 79
			case 't':
				currentState = 76
			case 'n':
				currentState = 85
			case 'J':
				currentState = 59
			case 'b':
				currentState = 60
			case 'k':
				currentState = 57
			case 'B':
				currentState = 88
			case 'G':
				currentState = 100
			case 'A':
				currentState = 77
			case 'O':
				currentState = 73
			case 'g':
				currentState = 82
			case 'K':
				currentState = 63
			case 'M':
				currentState = 90
			case 'u':
				currentState = 55
			case 'U':
				currentState = 67
			case 'a':
				currentState = 104
			default:
				break outer
			}

		case 43:
			switch c {
			case 'L':
				currentState = 61
			case 'P':
				currentState = 65
			case 'c':
				currentState = 69
			case 'F':
				currentState = 89
			case 'J':
				currentState = 59
			case 'O':
				currentState = 73
			case 'b':
				currentState = 60
			case 's':
				currentState = 87
			case 'g':
				currentState = 82
			case 'k':
				currentState = 57
			case 'l':
				currentState = 106
			case 'o':
				currentState = 58
			case 'p':
				currentState = 70
			case 'w':
				currentState = 71
			case 'D':
				currentState = 94
			case 'K':
				currentState = 63
			case 'R':
				currentState = 103
			case 'X':
				currentState = 68
			case 'z':
				currentState = 64
			case 'I':
				currentState = 62
			case 'r':
				currentState = 83
			case 'H':
				currentState = 98
			case 'N':
				currentState = 99
			case 'v':
				currentState = 96
			case 'Y':
				currentState = 75
			case 'n':
				currentState = 85
			case 't':
				currentState = 76
			case '_':
				currentState = 97
			case 'U':
				currentState = 67
			case 'j':
				currentState = 84
			case 'h':
				currentState = 79
			case 'V':
				currentState = 92
			case 'Z':
				currentState = 101
			case 'G':
				currentState = 100
			case 'W':
				currentState = 56
			case 'a':
				currentState = 104
			case 'd':
				currentState = 66
			case 'm':
				currentState = 54
			case 'A':
				currentState = 77
			case 'E':
				currentState = 95
			case 'C':
				currentState = 86
			case 'f':
				currentState = 102
			case 'T':
				currentState = 74
			case 'u':
				currentState = 55
			case 'q':
				currentState = 81
			case 'M':
				currentState = 90
			case 'i':
				currentState = 80
			case 'S':
				currentState = 78
			case 'e':
				currentState = 105
			case 'x':
				currentState = 93
			case 'y':
				currentState = 72
			case 'B':
				currentState = 88
			case 'Q':
				currentState = 91
			default:
				break outer
			}

		case 68:
			switch c {
			case 'W':
				currentState = 56
			case 'I':
				currentState = 62
			case 'M':
				currentState = 90
			case 'R':
				currentState = 103
			case 'j':
				currentState = 84
			case 'w':
				currentState = 71
			case 'y':
				currentState = 72
			case 'L':
				currentState = 61
			case 'P':
				currentState = 65
			case 's':
				currentState = 87
			case 't':
				currentState = 76
			case 'g':
				currentState = 82
			case 'v':
				currentState = 96
			case 'Q':
				currentState = 91
			case 'T':
				currentState = 74
			case 'X':
				currentState = 68
			case 'c':
				currentState = 69
			case 'u':
				currentState = 55
			case 'A':
				currentState = 77
			case 'J':
				currentState = 59
			case 'D':
				currentState = 94
			case 'N':
				currentState = 99
			case 'Z':
				currentState = 101
			case 'a':
				currentState = 104
			case 'E':
				currentState = 95
			case 'O':
				currentState = 73
			case 'h':
				currentState = 79
			case 'b':
				currentState = 60
			case 'f':
				currentState = 102
			case 'm':
				currentState = 54
			case 'n':
				currentState = 85
			case 'B':
				currentState = 88
			case 'V':
				currentState = 92
			case 'p':
				currentState = 70
			case 'F':
				currentState = 89
			case 'q':
				currentState = 81
			case 'x':
				currentState = 93
			case 'z':
				currentState = 64
			case 'C':
				currentState = 86
			case 'K':
				currentState = 63
			case 'Y':
				currentState = 75
			case 'i':
				currentState = 80
			case 'k':
				currentState = 57
			case 'r':
				currentState = 83
			case 'H':
				currentState = 98
			case 'S':
				currentState = 78
			case 'U':
				currentState = 67
			case 'd':
				currentState = 66
			case 'G':
				currentState = 100
			case 'l':
				currentState = 106
			case 'e':
				currentState = 105
			case 'o':
				currentState = 58
			case '_':
				currentState = 97
			default:
				break outer
			}

		case 7:
			switch c {
			case 'r':
				currentState = 83
			case 'D':
				currentState = 94
			case 'I':
				currentState = 62
			case 'P':
				currentState = 65
			case 'A':
				currentState = 77
			case 'Z':
				currentState = 101
			case 't':
				currentState = 76
			case 'w':
				currentState = 71
			case 'L':
				currentState = 61
			case 'g':
				currentState = 82
			case 'l':
				currentState = 106
			case 'X':
				currentState = 68
			case 'E':
				currentState = 95
			case 'V':
				currentState = 92
			case 'T':
				currentState = 74
			case 'q':
				currentState = 81
			case 'H':
				currentState = 98
			case 'Q':
				currentState = 91
			case 'S':
				currentState = 78
			case 'G':
				currentState = 100
			case 'O':
				currentState = 73
			case 'i':
				currentState = 80
			case 'R':
				currentState = 103
			case 'W':
				currentState = 56
			case 'v':
				currentState = 96
			case 'F':
				currentState = 89
			case 'z':
				currentState = 64
			case 'j':
				currentState = 84
			case 'B':
				currentState = 88
			case 'J':
				currentState = 59
			case 'N':
				currentState = 99
			case 'b':
				currentState = 60
			case 'f':
				currentState = 102
			case 'k':
				currentState = 57
			case 'p':
				currentState = 70
			case 'y':
				currentState = 72
			case 'U':
				currentState = 67
			case 'h':
				currentState = 79
			case 'm':
				currentState = 54
			case 'o':
				currentState = 58
			case 'M':
				currentState = 90
			case 'a':
				currentState = 104
			case 'c':
				currentState = 69
			case 'n':
				currentState = 85
			case 'C':
				currentState = 86
			case 'K':
				currentState = 63
			case 'Y':
				currentState = 75
			case 'e':
				currentState = 105
			case 'u':
				currentState = 55
			case '_':
				currentState = 97
			case 'd':
				currentState = 66
			case 's':
				currentState = 87
			case 'x':
				currentState = 93
			default:
				break outer
			}

		case 9:
			switch c {
			case 'V':
				currentState = 92
			case 'Z':
				currentState = 101
			case 'k':
				currentState = 57
			case 'b':
				currentState = 60
			case 'x':
				currentState = 93
			case 'P':
				currentState = 65
			case 'X':
				currentState = 68
			case 'i':
				currentState = 80
			case 'w':
				currentState = 71
			case 'e':
				currentState = 105
			case 'z':
				currentState = 64
			case 'n':
				currentState = 85
			case 'u':
				currentState = 55
			case 'D':
				currentState = 94
			case 'J':
				currentState = 59
			case 'Q':
				currentState = 91
			case 'T':
				currentState = 74
			case 'j':
				currentState = 84
			case 'Y':
				currentState = 75
			case 'a':
				currentState = 104
			case 'H':
				currentState = 98
			case 'f':
				currentState = 102
			case 'q':
				currentState = 81
			case 't':
				currentState = 76
			case 'B':
				currentState = 88
			case 'L':
				currentState = 61
			case 'm':
				currentState = 54
			case 'O':
				currentState = 73
			case 'U':
				currentState = 67
			case 'W':
				currentState = 56
			case 'c':
				currentState = 69
			case 'g':
				currentState = 82
			case 'E':
				currentState = 95
			case 'G':
				currentState = 100
			case 'S':
				currentState = 78
			case 'o':
				currentState = 58
			case 's':
				currentState = 87
			case 'v':
				currentState = 96
			case 'M':
				currentState = 90
			case 'N':
				currentState = 99
			case 'h':
				currentState = 79
			case 'l':
				currentState = 106
			case 'p':
				currentState = 70
			case 'F':
				currentState = 89
			case 'K':
				currentState = 63
			case 'r':
				currentState = 83
			case 'y':
				currentState = 72
			case '_':
				currentState = 97
			case 'A':
				currentState = 77
			case 'C':
				currentState = 86
			case 'I':
				currentState = 62
			case 'R':
				currentState = 103
			case 'd':
				currentState = 66
			default:
				break outer
			}

		case 23:
			switch c {
			case 'I':
				currentState = 62
			case 'Z':
				currentState = 101
			case 'q':
				currentState = 81
			case 'W':
				currentState = 56
			case 't':
				currentState = 76
			case 'D':
				currentState = 94
			case 'N':
				currentState = 99
			case 'd':
				currentState = 66
			case '_':
				currentState = 97
			case 'B':
				currentState = 88
			case 'F':
				currentState = 89
			case 'T':
				currentState = 74
			case 'b':
				currentState = 60
			case 'Q':
				currentState = 91
			case 'S':
				currentState = 78
			case 'f':
				currentState = 102
			case 'o':
				currentState = 58
			case 'J':
				currentState = 59
			case 'a':
				currentState = 104
			case 'y':
				currentState = 72
			case 'U':
				currentState = 67
			case 'j':
				currentState = 84
			case 'k':
				currentState = 57
			case 'm':
				currentState = 54
			case 'r':
				currentState = 83
			case 'w':
				currentState = 71
			case 'x':
				currentState = 93
			case 'i':
				currentState = 80
			case 'l':
				currentState = 106
			case 's':
				currentState = 87
			case 'K':
				currentState = 63
			case 'Y':
				currentState = 75
			case 'v':
				currentState = 96
			case 'G':
				currentState = 100
			case 'H':
				currentState = 98
			case 'O':
				currentState = 73
			case 'p':
				currentState = 70
			case 'z':
				currentState = 64
			case 'C':
				currentState = 86
			case 'P':
				currentState = 65
			case 'u':
				currentState = 55
			case 'n':
				currentState = 85
			case 'A':
				currentState = 77
			case 'L':
				currentState = 61
			case 'V':
				currentState = 92
			case 'X':
				currentState = 68
			case 'g':
				currentState = 82
			case 'M':
				currentState = 90
			case 'c':
				currentState = 69
			case 'e':
				currentState = 105
			case 'h':
				currentState = 79
			case 'E':
				currentState = 95
			case 'R':
				currentState = 103
			default:
				break outer
			}

		case 27:
			switch c {
			case 'D':
				currentState = 94
			case 'F':
				currentState = 89
			case 'L':
				currentState = 61
			case 'U':
				currentState = 67
			case 'P':
				currentState = 65
			case 'j':
				currentState = 84
			case 'm':
				currentState = 54
			case 'G':
				currentState = 100
			case 'H':
				currentState = 98
			case 'o':
				currentState = 58
			case 'q':
				currentState = 81
			case 'c':
				currentState = 69
			case 'T':
				currentState = 74
			case 'n':
				currentState = 85
			case 'C':
				currentState = 86
			case 'E':
				currentState = 95
			case 'e':
				currentState = 105
			case 'a':
				currentState = 104
			case 'l':
				currentState = 106
			case 'I':
				currentState = 62
			case 'K':
				currentState = 63
			case 'O':
				currentState = 73
			case 'k':
				currentState = 57
			case 'A':
				currentState = 77
			case 'B':
				currentState = 88
			case 'J':
				currentState = 59
			case 'd':
				currentState = 66
			case 'r':
				currentState = 83
			case 'z':
				currentState = 64
			case 'R':
				currentState = 103
			case 'i':
				currentState = 80
			case 's':
				currentState = 87
			case 't':
				currentState = 76
			case 'W':
				currentState = 56
			case 'Z':
				currentState = 101
			case 'b':
				currentState = 60
			case 'f':
				currentState = 102
			case 'h':
				currentState = 79
			case '_':
				currentState = 97
			case 'Q':
				currentState = 91
			case 'g':
				currentState = 82
			case 'p':
				currentState = 70
			case 'w':
				currentState = 71
			case 'y':
				currentState = 72
			case 'M':
				currentState = 90
			case 'x':
				currentState = 93
			case 'N':
				currentState = 99
			case 'S':
				currentState = 78
			case 'V':
				currentState = 92
			case 'X':
				currentState = 68
			case 'Y':
				currentState = 75
			case 'u':
				currentState = 55
			case 'v':
				currentState = 96
			default:
				break outer
			}

		case 4:
			switch c {
			case 'J':
				currentState = 59
			case 'Q':
				currentState = 91
			case 'k':
				currentState = 57
			case 'l':
				currentState = 106
			case 'r':
				currentState = 83
			case 'a':
				currentState = 104
			case 'n':
				currentState = 85
			case 'U':
				currentState = 67
			case 'i':
				currentState = 80
			case 'j':
				currentState = 84
			case 'c':
				currentState = 69
			case 'o':
				currentState = 58
			case 'w':
				currentState = 71
			case 'G':
				currentState = 100
			case 'H':
				currentState = 98
			case 'N':
				currentState = 99
			case 'V':
				currentState = 92
			case 'Z':
				currentState = 101
			case 'e':
				currentState = 105
			case 'z':
				currentState = 64
			case 'C':
				currentState = 86
			case 'K':
				currentState = 63
			case 'M':
				currentState = 90
			case 'O':
				currentState = 73
			case 'R':
				currentState = 103
			case 'T':
				currentState = 74
			case 'u':
				currentState = 55
			case 'E':
				currentState = 95
			case 'q':
				currentState = 81
			case '_':
				currentState = 97
			case 'm':
				currentState = 54
			case 'X':
				currentState = 68
			case 'x':
				currentState = 93
			case 'D':
				currentState = 94
			case 'f':
				currentState = 102
			case 'L':
				currentState = 61
			case 'S':
				currentState = 78
			case 's':
				currentState = 87
			case 'I':
				currentState = 62
			case 'Y':
				currentState = 75
			case 'h':
				currentState = 79
			case 'P':
				currentState = 65
			case 'W':
				currentState = 56
			case 'F':
				currentState = 89
			case 'b':
				currentState = 60
			case 'd':
				currentState = 66
			case 'g':
				currentState = 82
			case 'v':
				currentState = 96
			case 'y':
				currentState = 72
			case 'A':
				currentState = 77
			case 'B':
				currentState = 88
			case 'p':
				currentState = 70
			case 't':
				currentState = 76
			default:
				break outer
			}

		case 102:
			switch c {
			case 'U':
				currentState = 67
			case 'S':
				currentState = 78
			case 'G':
				currentState = 100
			case 'Y':
				currentState = 75
			case 'a':
				currentState = 104
			case 'l':
				currentState = 106
			case '_':
				currentState = 97
			case 'D':
				currentState = 94
			case 'Z':
				currentState = 101
			case 'F':
				currentState = 89
			case 'T':
				currentState = 74
			case 'V':
				currentState = 92
			case 'z':
				currentState = 64
			case 'H':
				currentState = 98
			case 'e':
				currentState = 105
			case 'v':
				currentState = 96
			case 'B':
				currentState = 88
			case 'P':
				currentState = 65
			case 'n':
				currentState = 85
			case 'r':
				currentState = 83
			case 'E':
				currentState = 95
			case 'W':
				currentState = 56
			case 'X':
				currentState = 68
			case 'f':
				currentState = 102
			case 'i':
				currentState = 80
			case 'q':
				currentState = 81
			case 'w':
				currentState = 71
			case 'Q':
				currentState = 91
			case 'M':
				currentState = 90
			case 'x':
				currentState = 93
			case 'y':
				currentState = 72
			case 'I':
				currentState = 62
			case 'g':
				currentState = 82
			case 'C':
				currentState = 86
			case 't':
				currentState = 76
			case 'J':
				currentState = 59
			case 'j':
				currentState = 84
			case 'L':
				currentState = 61
			case 'b':
				currentState = 60
			case 'c':
				currentState = 69
			case 'd':
				currentState = 66
			case 'm':
				currentState = 54
			case 'o':
				currentState = 58
			case 'u':
				currentState = 55
			case 'N':
				currentState = 99
			case 'O':
				currentState = 73
			case 'K':
				currentState = 63
			case 'k':
				currentState = 57
			case 'A':
				currentState = 77
			case 'h':
				currentState = 79
			case 'p':
				currentState = 70
			case 's':
				currentState = 87
			case 'R':
				currentState = 103
			default:
				break outer
			}

		case 63:
			switch c {
			case 'g':
				currentState = 82
			case 'L':
				currentState = 61
			case 'n':
				currentState = 85
			case 'u':
				currentState = 55
			case 'i':
				currentState = 80
			case 'm':
				currentState = 54
			case 's':
				currentState = 87
			case 'v':
				currentState = 96
			case '_':
				currentState = 97
			case 'G':
				currentState = 100
			case 'P':
				currentState = 65
			case 'S':
				currentState = 78
			case 'e':
				currentState = 105
			case 'j':
				currentState = 84
			case 'H':
				currentState = 98
			case 'T':
				currentState = 74
			case 'W':
				currentState = 56
			case 'c':
				currentState = 69
			case 'f':
				currentState = 102
			case 'y':
				currentState = 72
			case 'B':
				currentState = 88
			case 'R':
				currentState = 103
			case 'V':
				currentState = 92
			case 'A':
				currentState = 77
			case 'k':
				currentState = 57
			case 'q':
				currentState = 81
			case 'r':
				currentState = 83
			case 'F':
				currentState = 89
			case 'a':
				currentState = 104
			case 'Z':
				currentState = 101
			case 'Y':
				currentState = 75
			case 'h':
				currentState = 79
			case 'I':
				currentState = 62
			case 'O':
				currentState = 73
			case 'p':
				currentState = 70
			case 'z':
				currentState = 64
			case 'N':
				currentState = 99
			case 'J':
				currentState = 59
			case 'U':
				currentState = 67
			case 'X':
				currentState = 68
			case 'l':
				currentState = 106
			case 'o':
				currentState = 58
			case 'E':
				currentState = 95
			case 'D':
				currentState = 94
			case 'x':
				currentState = 93
			case 'C':
				currentState = 86
			case 'M':
				currentState = 90
			case 'Q':
				currentState = 91
			case 'w':
				currentState = 71
			case 'K':
				currentState = 63
			case 'd':
				currentState = 66
			case 't':
				currentState = 76
			case 'b':
				currentState = 60
			default:
				break outer
			}

		case 77:
			switch c {
			case 'b':
				currentState = 60
			case 'X':
				currentState = 68
			case 'E':
				currentState = 95
			case 'O':
				currentState = 73
			case 'm':
				currentState = 54
			case 'F':
				currentState = 89
			case 'N':
				currentState = 99
			case 'a':
				currentState = 104
			case 'r':
				currentState = 83
			case 'L':
				currentState = 61
			case 'S':
				currentState = 78
			case 'Y':
				currentState = 75
			case 'B':
				currentState = 88
			case 'J':
				currentState = 59
			case 'k':
				currentState = 57
			case 'Z':
				currentState = 101
			case 'c':
				currentState = 69
			case 's':
				currentState = 87
			case '_':
				currentState = 97
			case 'K':
				currentState = 63
			case 'j':
				currentState = 84
			case 'v':
				currentState = 96
			case 'w':
				currentState = 71
			case 'A':
				currentState = 77
			case 'T':
				currentState = 74
			case 'R':
				currentState = 103
			case 't':
				currentState = 76
			case 'C':
				currentState = 86
			case 'I':
				currentState = 62
			case 'V':
				currentState = 92
			case 'd':
				currentState = 66
			case 'o':
				currentState = 58
			case 'q':
				currentState = 81
			case 'y':
				currentState = 72
			case 'H':
				currentState = 98
			case 'Q':
				currentState = 91
			case 'W':
				currentState = 56
			case 'n':
				currentState = 85
			case 'U':
				currentState = 67
			case 'z':
				currentState = 64
			case 'D':
				currentState = 94
			case 'P':
				currentState = 65
			case 'h':
				currentState = 79
			case 'l':
				currentState = 106
			case 'p':
				currentState = 70
			case 'u':
				currentState = 55
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case 'e':
				currentState = 105
			case 'i':
				currentState = 80
			case 'x':
				currentState = 93
			case 'G':
				currentState = 100
			case 'M':
				currentState = 90
			default:
				break outer
			}

		case 101:
			switch c {
			case 'b':
				currentState = 60
			case 'f':
				currentState = 102
			case 'v':
				currentState = 96
			case 'L':
				currentState = 61
			case 'W':
				currentState = 56
			case 'd':
				currentState = 66
			case 'g':
				currentState = 82
			case 'h':
				currentState = 79
			case 'G':
				currentState = 100
			case 'U':
				currentState = 67
			case 'e':
				currentState = 105
			case 'q':
				currentState = 81
			case 'I':
				currentState = 62
			case 'N':
				currentState = 99
			case 'x':
				currentState = 93
			case 'B':
				currentState = 88
			case 'Q':
				currentState = 91
			case 'a':
				currentState = 104
			case 'u':
				currentState = 55
			case 'M':
				currentState = 90
			case 'S':
				currentState = 78
			case 'k':
				currentState = 57
			case 'w':
				currentState = 71
			case 'F':
				currentState = 89
			case 'm':
				currentState = 54
			case 'n':
				currentState = 85
			case '_':
				currentState = 97
			case 'C':
				currentState = 86
			case 'J':
				currentState = 59
			case 'Z':
				currentState = 101
			case 'l':
				currentState = 106
			case 'r':
				currentState = 83
			case 'D':
				currentState = 94
			case 'O':
				currentState = 73
			case 'R':
				currentState = 103
			case 'T':
				currentState = 74
			case 'V':
				currentState = 92
			case 'y':
				currentState = 72
			case 'K':
				currentState = 63
			case 'Y':
				currentState = 75
			case 'i':
				currentState = 80
			case 'z':
				currentState = 64
			case 't':
				currentState = 76
			case 'H':
				currentState = 98
			case 'j':
				currentState = 84
			case 'P':
				currentState = 65
			case 'X':
				currentState = 68
			case 'o':
				currentState = 58
			case 's':
				currentState = 87
			case 'A':
				currentState = 77
			case 'E':
				currentState = 95
			case 'c':
				currentState = 69
			case 'p':
				currentState = 70
			default:
				break outer
			}

		case 3:
			switch c {
			case 'G':
				currentState = 100
			case 'T':
				currentState = 74
			case 'P':
				currentState = 65
			case 'Z':
				currentState = 101
			case 'b':
				currentState = 60
			case 'g':
				currentState = 82
			case 'K':
				currentState = 63
			case 'M':
				currentState = 90
			case 'e':
				currentState = 105
			case 'j':
				currentState = 84
			case 'k':
				currentState = 57
			case 'm':
				currentState = 54
			case 's':
				currentState = 87
			case 'N':
				currentState = 99
			case 'V':
				currentState = 92
			case 'z':
				currentState = 64
			case 'd':
				currentState = 66
			case 'h':
				currentState = 79
			case '_':
				currentState = 97
			case 'O':
				currentState = 73
			case 'a':
				currentState = 104
			case 'u':
				currentState = 55
			case 'L':
				currentState = 61
			case 'r':
				currentState = 83
			case 'w':
				currentState = 71
			case 'x':
				currentState = 93
			case 'y':
				currentState = 72
			case 'E':
				currentState = 95
			case 'S':
				currentState = 78
			case 'Y':
				currentState = 75
			case 'f':
				currentState = 102
			case 'n':
				currentState = 85
			case 'A':
				currentState = 77
			case 'B':
				currentState = 88
			case 'c':
				currentState = 69
			case 'o':
				currentState = 58
			case 'q':
				currentState = 81
			case 'H':
				currentState = 98
			case 'I':
				currentState = 62
			case 'C':
				currentState = 86
			case 'W':
				currentState = 56
			case 'J':
				currentState = 59
			case 'R':
				currentState = 103
			case 't':
				currentState = 76
			case 'v':
				currentState = 96
			case 'X':
				currentState = 68
			case 'i':
				currentState = 80
			case 'Q':
				currentState = 91
			case 'U':
				currentState = 67
			case 'l':
				currentState = 106
			case 'p':
				currentState = 70
			case 'D':
				currentState = 94
			case 'F':
				currentState = 89
			default:
				break outer
			}

		case 82:
			switch c {
			case 't':
				currentState = 76
			case 'x':
				currentState = 93
			case 'A':
				currentState = 77
			case 'o':
				currentState = 58
			case 'e':
				currentState = 105
			case 'v':
				currentState = 96
			case 'S':
				currentState = 78
			case 'U':
				currentState = 67
			case 'Y':
				currentState = 75
			case 'E':
				currentState = 95
			case 'M':
				currentState = 90
			case 'l':
				currentState = 106
			case 'q':
				currentState = 81
			case 'O':
				currentState = 73
			case 'a':
				currentState = 104
			case 'K':
				currentState = 63
			case 'w':
				currentState = 71
			case 'c':
				currentState = 69
			case 'u':
				currentState = 55
			case 'C':
				currentState = 86
			case 'b':
				currentState = 60
			case 'f':
				currentState = 102
			case 'Q':
				currentState = 91
			case 'W':
				currentState = 56
			case '_':
				currentState = 97
			case 'Z':
				currentState = 101
			case 'h':
				currentState = 79
			case 'n':
				currentState = 85
			case 'X':
				currentState = 68
			case 'g':
				currentState = 82
			case 'j':
				currentState = 84
			case 'p':
				currentState = 70
			case 'r':
				currentState = 83
			case 'R':
				currentState = 103
			case 'V':
				currentState = 92
			case 'N':
				currentState = 99
			case 'P':
				currentState = 65
			case 'd':
				currentState = 66
			case 'm':
				currentState = 54
			case 'B':
				currentState = 88
			case 'J':
				currentState = 59
			case 'k':
				currentState = 57
			case 'y':
				currentState = 72
			case 'H':
				currentState = 98
			case 'T':
				currentState = 74
			case 'L':
				currentState = 61
			case 'F':
				currentState = 89
			case 'G':
				currentState = 100
			case 's':
				currentState = 87
			case 'z':
				currentState = 64
			case 'D':
				currentState = 94
			case 'i':
				currentState = 80
			case 'I':
				currentState = 62
			default:
				break outer
			}

		case 37:
			switch c {
			case 'R':
				currentState = 103
			case 'a':
				currentState = 104
			case 'v':
				currentState = 96
			case 'J':
				currentState = 59
			case 'P':
				currentState = 65
			case 'Y':
				currentState = 75
			case 'u':
				currentState = 55
			case 'B':
				currentState = 88
			case 'M':
				currentState = 90
			case 'Z':
				currentState = 101
			case 'T':
				currentState = 74
			case 'V':
				currentState = 92
			case 'b':
				currentState = 60
			case 'H':
				currentState = 98
			case 'L':
				currentState = 61
			case 'w':
				currentState = 71
			case 'z':
				currentState = 64
			case 'S':
				currentState = 78
			case 'd':
				currentState = 66
			case 'j':
				currentState = 84
			case 'm':
				currentState = 54
			case 'E':
				currentState = 95
			case 'F':
				currentState = 89
			case 'p':
				currentState = 70
			case '_':
				currentState = 97
			case 'I':
				currentState = 62
			case 'N':
				currentState = 99
			case 'k':
				currentState = 57
			case 'n':
				currentState = 85
			case 'q':
				currentState = 81
			case 'D':
				currentState = 94
			case 'i':
				currentState = 80
			case 'o':
				currentState = 58
			case 'C':
				currentState = 86
			case 'h':
				currentState = 79
			case 'G':
				currentState = 100
			case 'e':
				currentState = 105
			case 'X':
				currentState = 68
			case 'K':
				currentState = 63
			case 'O':
				currentState = 73
			case 'y':
				currentState = 72
			case 'l':
				currentState = 106
			case 'r':
				currentState = 83
			case 'x':
				currentState = 93
			case 'Q':
				currentState = 91
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case 's':
				currentState = 87
			case 'A':
				currentState = 77
			case 'c':
				currentState = 69
			case 'U':
				currentState = 67
			case 't':
				currentState = 76
			case 'W':
				currentState = 56
			default:
				break outer
			}

		case 90:
			switch c {
			case 'B':
				currentState = 88
			case 'j':
				currentState = 84
			case 'm':
				currentState = 54
			case 'R':
				currentState = 103
			case 'q':
				currentState = 81
			case 'w':
				currentState = 71
			case 'x':
				currentState = 93
			case 'z':
				currentState = 64
			case 'V':
				currentState = 92
			case 'Z':
				currentState = 101
			case 'c':
				currentState = 69
			case 'f':
				currentState = 102
			case 'v':
				currentState = 96
			case '_':
				currentState = 97
			case 'u':
				currentState = 55
			case 'D':
				currentState = 94
			case 'G':
				currentState = 100
			case 'K':
				currentState = 63
			case 'Y':
				currentState = 75
			case 'h':
				currentState = 79
			case 'o':
				currentState = 58
			case 'U':
				currentState = 67
			case 'n':
				currentState = 85
			case 's':
				currentState = 87
			case 'L':
				currentState = 61
			case 'O':
				currentState = 73
			case 'T':
				currentState = 74
			case 'X':
				currentState = 68
			case 'i':
				currentState = 80
			case 't':
				currentState = 76
			case 'b':
				currentState = 60
			case 'k':
				currentState = 57
			case 'F':
				currentState = 89
			case 'H':
				currentState = 98
			case 'N':
				currentState = 99
			case 'Q':
				currentState = 91
			case 'd':
				currentState = 66
			case 'A':
				currentState = 77
			case 'I':
				currentState = 62
			case 'P':
				currentState = 65
			case 'W':
				currentState = 56
			case 'g':
				currentState = 82
			case 'r':
				currentState = 83
			case 'e':
				currentState = 105
			case 'J':
				currentState = 59
			case 'a':
				currentState = 104
			case 'l':
				currentState = 106
			case 'p':
				currentState = 70
			case 'y':
				currentState = 72
			case 'C':
				currentState = 86
			case 'M':
				currentState = 90
			case 'E':
				currentState = 95
			case 'S':
				currentState = 78
			default:
				break outer
			}

		case 87:
			switch c {
			case 'q':
				currentState = 81
			case 'F':
				currentState = 89
			case 'L':
				currentState = 61
			case 'd':
				currentState = 66
			case 'C':
				currentState = 86
			case 'J':
				currentState = 59
			case 'Z':
				currentState = 101
			case 'y':
				currentState = 72
			case 'c':
				currentState = 69
			case 'e':
				currentState = 105
			case 'i':
				currentState = 80
			case 'B':
				currentState = 88
			case 'I':
				currentState = 62
			case 'R':
				currentState = 103
			case 'X':
				currentState = 68
			case 'u':
				currentState = 55
			case 'w':
				currentState = 71
			case 'W':
				currentState = 56
			case 'Y':
				currentState = 75
			case 'p':
				currentState = 70
			case 'x':
				currentState = 93
			case '_':
				currentState = 97
			case 'U':
				currentState = 67
			case 'O':
				currentState = 73
			case 'j':
				currentState = 84
			case 'm':
				currentState = 54
			case 'D':
				currentState = 94
			case 'M':
				currentState = 90
			case 'k':
				currentState = 57
			case 'z':
				currentState = 64
			case 'G':
				currentState = 100
			case 'T':
				currentState = 74
			case 'a':
				currentState = 104
			case 'f':
				currentState = 102
			case 't':
				currentState = 76
			case 'l':
				currentState = 106
			case 'A':
				currentState = 77
			case 'H':
				currentState = 98
			case 'K':
				currentState = 63
			case 'P':
				currentState = 65
			case 'Q':
				currentState = 91
			case 'N':
				currentState = 99
			case 'h':
				currentState = 79
			case 'r':
				currentState = 83
			case 'E':
				currentState = 95
			case 'V':
				currentState = 92
			case 'b':
				currentState = 60
			case 'o':
				currentState = 58
			case 'S':
				currentState = 78
			case 'g':
				currentState = 82
			case 'n':
				currentState = 85
			case 's':
				currentState = 87
			case 'v':
				currentState = 96
			default:
				break outer
			}

		case 13:
			switch c {
			case 'u':
				currentState = 55
			case 'C':
				currentState = 86
			case 'K':
				currentState = 63
			case 'W':
				currentState = 56
			case 'X':
				currentState = 68
			case 't':
				currentState = 76
			case 'G':
				currentState = 100
			case 'h':
				currentState = 79
			case 'j':
				currentState = 84
			case 'x':
				currentState = 93
			case 'k':
				currentState = 57
			case 'w':
				currentState = 71
			case 'p':
				currentState = 70
			case 'B':
				currentState = 88
			case 'c':
				currentState = 69
			case 'g':
				currentState = 82
			case 'l':
				currentState = 106
			case 'N':
				currentState = 99
			case 'd':
				currentState = 66
			case '_':
				currentState = 97
			case 'H':
				currentState = 98
			case 'a':
				currentState = 104
			case 'U':
				currentState = 67
			case 'e':
				currentState = 105
			case 'm':
				currentState = 54
			case 'q':
				currentState = 81
			case 's':
				currentState = 87
			case 'M':
				currentState = 90
			case 'Z':
				currentState = 101
			case 'O':
				currentState = 73
			case 'n':
				currentState = 85
			case 'r':
				currentState = 83
			case 'R':
				currentState = 103
			case 'T':
				currentState = 74
			case 'o':
				currentState = 58
			case 'D':
				currentState = 94
			case 'L':
				currentState = 61
			case 'y':
				currentState = 72
			case 'Q':
				currentState = 91
			case 'S':
				currentState = 78
			case 'f':
				currentState = 102
			case 'v':
				currentState = 96
			case 'z':
				currentState = 64
			case 'J':
				currentState = 59
			case 'P':
				currentState = 65
			case 'Y':
				currentState = 75
			case 'b':
				currentState = 60
			case 'i':
				currentState = 80
			case 'A':
				currentState = 77
			case 'E':
				currentState = 95
			case 'F':
				currentState = 89
			case 'I':
				currentState = 62
			case 'V':
				currentState = 92
			default:
				break outer
			}

		case 34:
			switch c {
			case 'N':
				currentState = 99
			case 'P':
				currentState = 65
			case 'y':
				currentState = 72
			case 'R':
				currentState = 103
			case 'T':
				currentState = 74
			case 'e':
				currentState = 105
			case 'M':
				currentState = 90
			case 'Y':
				currentState = 75
			case 'f':
				currentState = 102
			case 'q':
				currentState = 81
			case 'I':
				currentState = 62
			case 'L':
				currentState = 61
			case 'm':
				currentState = 54
			case 'k':
				currentState = 57
			case 'l':
				currentState = 106
			case 's':
				currentState = 87
			case 'X':
				currentState = 68
			case 'Z':
				currentState = 101
			case 'h':
				currentState = 79
			case 'F':
				currentState = 89
			case 'd':
				currentState = 66
			case 'C':
				currentState = 86
			case 'O':
				currentState = 73
			case 'W':
				currentState = 56
			case 'b':
				currentState = 60
			case 'o':
				currentState = 58
			case 'p':
				currentState = 70
			case 'H':
				currentState = 98
			case 'Q':
				currentState = 91
			case 'n':
				currentState = 85
			case 'B':
				currentState = 88
			case 'S':
				currentState = 78
			case '_':
				currentState = 97
			case 'K':
				currentState = 63
			case 'V':
				currentState = 92
			case 'u':
				currentState = 55
			case 'J':
				currentState = 59
			case 'c':
				currentState = 69
			case 'g':
				currentState = 82
			case 'v':
				currentState = 96
			case 'E':
				currentState = 95
			case 'U':
				currentState = 67
			case 'a':
				currentState = 104
			case 'j':
				currentState = 84
			case 't':
				currentState = 76
			case 'w':
				currentState = 71
			case 'z':
				currentState = 64
			case 'A':
				currentState = 77
			case 'G':
				currentState = 100
			case 'i':
				currentState = 80
			case 'x':
				currentState = 93
			case 'D':
				currentState = 94
			case 'r':
				currentState = 83
			default:
				break outer
			}

		case 47:
			switch c {
			case 'b':
				currentState = 60
			case 'E':
				currentState = 95
			case 'a':
				currentState = 104
			case 'k':
				currentState = 57
			case 't':
				currentState = 76
			case 'J':
				currentState = 59
			case 'y':
				currentState = 72
			case 'X':
				currentState = 68
			case 'l':
				currentState = 106
			case 'v':
				currentState = 96
			case 'C':
				currentState = 86
			case 'I':
				currentState = 62
			case 'P':
				currentState = 65
			case 'e':
				currentState = 105
			case 'h':
				currentState = 79
			case 'w':
				currentState = 71
			case 'H':
				currentState = 98
			case 'O':
				currentState = 73
			case 'V':
				currentState = 92
			case 'c':
				currentState = 69
			case 'A':
				currentState = 77
			case 'j':
				currentState = 84
			case 'u':
				currentState = 55
			case 'D':
				currentState = 94
			case 'Q':
				currentState = 91
			case 'z':
				currentState = 64
			case 'K':
				currentState = 63
			case 'T':
				currentState = 74
			case 'W':
				currentState = 56
			case 'n':
				currentState = 85
			case 'p':
				currentState = 70
			case 'q':
				currentState = 81
			case 'r':
				currentState = 83
			case 'G':
				currentState = 100
			case 'd':
				currentState = 66
			case 'f':
				currentState = 102
			case 'm':
				currentState = 54
			case '_':
				currentState = 97
			case 'M':
				currentState = 90
			case 'i':
				currentState = 80
			case 'F':
				currentState = 89
			case 'o':
				currentState = 58
			case 'g':
				currentState = 82
			case 'x':
				currentState = 93
			case 'B':
				currentState = 88
			case 'L':
				currentState = 61
			case 'Y':
				currentState = 75
			case 'Z':
				currentState = 101
			case 'U':
				currentState = 67
			case 'N':
				currentState = 99
			case 'R':
				currentState = 103
			case 'S':
				currentState = 78
			case 's':
				currentState = 87
			default:
				break outer
			}

		case 1:
			switch c {
			case 'm':
				currentState = 54
			case 'P':
				currentState = 65
			case 'c':
				currentState = 69
			case 'Y':
				currentState = 75
			case 'e':
				currentState = 105
			case 'j':
				currentState = 84
			case 'q':
				currentState = 81
			case 'z':
				currentState = 64
			case 'W':
				currentState = 56
			case 'I':
				currentState = 62
			case 'p':
				currentState = 70
			case 'K':
				currentState = 63
			case 'f':
				currentState = 102
			case 'n':
				currentState = 85
			case 'B':
				currentState = 88
			case 'J':
				currentState = 59
			case 'C':
				currentState = 86
			case 'H':
				currentState = 98
			case 'd':
				currentState = 66
			case 'u':
				currentState = 55
			case 'w':
				currentState = 71
			case 'y':
				currentState = 72
			case 'D':
				currentState = 94
			case 'a':
				currentState = 104
			case 'U':
				currentState = 67
			case 's':
				currentState = 87
			case 'o':
				currentState = 58
			case 'E':
				currentState = 95
			case 'G':
				currentState = 100
			case 'i':
				currentState = 80
			case 't':
				currentState = 76
			case 'R':
				currentState = 103
			case 'Z':
				currentState = 101
			case 'S':
				currentState = 78
			case 'l':
				currentState = 106
			case '_':
				currentState = 97
			case 'N':
				currentState = 99
			case 'O':
				currentState = 73
			case 'b':
				currentState = 60
			case 'x':
				currentState = 93
			case 'Q':
				currentState = 91
			case 'X':
				currentState = 68
			case 'g':
				currentState = 82
			case 'A':
				currentState = 77
			case 'T':
				currentState = 74
			case 'k':
				currentState = 57
			case 'M':
				currentState = 90
			case 'V':
				currentState = 92
			case 'h':
				currentState = 79
			case 'r':
				currentState = 83
			case 'v':
				currentState = 96
			case 'F':
				currentState = 89
			case 'L':
				currentState = 61
			default:
				break outer
			}

		case 36:
			switch c {
			case 'K':
				currentState = 63
			case 'A':
				currentState = 77
			case 'c':
				currentState = 69
			case 'd':
				currentState = 66
			case 't':
				currentState = 76
			case 'D':
				currentState = 94
			case 'R':
				currentState = 103
			case 'W':
				currentState = 56
			case 'y':
				currentState = 72
			case 'C':
				currentState = 86
			case 'L':
				currentState = 61
			case 'Q':
				currentState = 91
			case 's':
				currentState = 87
			case 'w':
				currentState = 71
			case 'J':
				currentState = 59
			case 'M':
				currentState = 90
			case 'T':
				currentState = 74
			case 'V':
				currentState = 92
			case 'h':
				currentState = 79
			case 'n':
				currentState = 85
			case 'a':
				currentState = 104
			case 'f':
				currentState = 102
			case 'k':
				currentState = 57
			case 'p':
				currentState = 70
			case 'u':
				currentState = 55
			case 'm':
				currentState = 54
			case 'q':
				currentState = 81
			case 'G':
				currentState = 100
			case 'I':
				currentState = 62
			case 'O':
				currentState = 73
			case 'Z':
				currentState = 101
			case 'B':
				currentState = 88
			case 'N':
				currentState = 99
			case 'l':
				currentState = 106
			case 'S':
				currentState = 78
			case 'X':
				currentState = 68
			case 'o':
				currentState = 58
			case 'F':
				currentState = 89
			case 'H':
				currentState = 98
			case 'Y':
				currentState = 75
			case 'x':
				currentState = 93
			case 'E':
				currentState = 95
			case 'P':
				currentState = 65
			case 'e':
				currentState = 105
			case 'g':
				currentState = 82
			case 'i':
				currentState = 80
			case 'r':
				currentState = 83
			case 'b':
				currentState = 60
			case 'z':
				currentState = 64
			case '_':
				currentState = 97
			case 'U':
				currentState = 67
			case 'j':
				currentState = 84
			case 'v':
				currentState = 96
			default:
				break outer
			}

		case 58:
			switch c {
			case 'B':
				currentState = 88
			case 'N':
				currentState = 99
			case 'P':
				currentState = 65
			case 'T':
				currentState = 74
			case 'X':
				currentState = 68
			case 'a':
				currentState = 104
			case 'K':
				currentState = 63
			case 'q':
				currentState = 81
			case 'W':
				currentState = 56
			case 'k':
				currentState = 57
			case 'L':
				currentState = 61
			case 'p':
				currentState = 70
			case 'z':
				currentState = 64
			case 'R':
				currentState = 103
			case 'Z':
				currentState = 101
			case 'd':
				currentState = 66
			case 'E':
				currentState = 95
			case 'e':
				currentState = 105
			case 'h':
				currentState = 79
			case 'o':
				currentState = 58
			case 't':
				currentState = 76
			case '_':
				currentState = 97
			case 'S':
				currentState = 78
			case 'j':
				currentState = 84
			case 'F':
				currentState = 89
			case 'y':
				currentState = 72
			case 'O':
				currentState = 73
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case 'i':
				currentState = 80
			case 'G':
				currentState = 100
			case 'Q':
				currentState = 91
			case 'c':
				currentState = 69
			case 'v':
				currentState = 96
			case 'w':
				currentState = 71
			case 'A':
				currentState = 77
			case 'H':
				currentState = 98
			case 'V':
				currentState = 92
			case 'C':
				currentState = 86
			case 'J':
				currentState = 59
			case 'Y':
				currentState = 75
			case 'm':
				currentState = 54
			case 's':
				currentState = 87
			case 'I':
				currentState = 62
			case 'M':
				currentState = 90
			case 'D':
				currentState = 94
			case 'x':
				currentState = 93
			case 'U':
				currentState = 67
			case 'b':
				currentState = 60
			case 'l':
				currentState = 106
			case 'n':
				currentState = 85
			case 'r':
				currentState = 83
			case 'u':
				currentState = 55
			default:
				break outer
			}

		case 85:
			switch c {
			case 'm':
				currentState = 54
			case 'q':
				currentState = 81
			case 'a':
				currentState = 104
			case 'i':
				currentState = 80
			case 'v':
				currentState = 96
			case 'R':
				currentState = 103
			case 'd':
				currentState = 66
			case 'L':
				currentState = 61
			case 'V':
				currentState = 92
			case 'K':
				currentState = 63
			case 'T':
				currentState = 74
			case 'c':
				currentState = 69
			case 'w':
				currentState = 71
			case 'J':
				currentState = 59
			case 'N':
				currentState = 99
			case 'Q':
				currentState = 91
			case 'U':
				currentState = 67
			case 'z':
				currentState = 64
			case 'H':
				currentState = 98
			case 'p':
				currentState = 70
			case 'r':
				currentState = 83
			case 's':
				currentState = 87
			case 'W':
				currentState = 56
			case 'u':
				currentState = 55
			case 'D':
				currentState = 94
			case 'P':
				currentState = 65
			case 't':
				currentState = 76
			case 'E':
				currentState = 95
			case 'G':
				currentState = 100
			case 'h':
				currentState = 79
			case 'k':
				currentState = 57
			case 'o':
				currentState = 58
			case 'x':
				currentState = 93
			case 'B':
				currentState = 88
			case 'C':
				currentState = 86
			case 'Z':
				currentState = 101
			case 'y':
				currentState = 72
			case '_':
				currentState = 97
			case 'A':
				currentState = 77
			case 'S':
				currentState = 78
			case 'e':
				currentState = 105
			case 'f':
				currentState = 102
			case 'l':
				currentState = 106
			case 'O':
				currentState = 73
			case 'j':
				currentState = 84
			case 'n':
				currentState = 85
			case 'b':
				currentState = 60
			case 'M':
				currentState = 90
			case 'X':
				currentState = 68
			case 'g':
				currentState = 82
			case 'F':
				currentState = 89
			case 'Y':
				currentState = 75
			case 'I':
				currentState = 62
			default:
				break outer
			}

		case 91:
			switch c {
			case 'Y':
				currentState = 75
			case 'h':
				currentState = 79
			case 'm':
				currentState = 54
			case 'o':
				currentState = 58
			case 'F':
				currentState = 89
			case 'M':
				currentState = 90
			case 't':
				currentState = 76
			case 'z':
				currentState = 64
			case 'V':
				currentState = 92
			case 'r':
				currentState = 83
			case 's':
				currentState = 87
			case '_':
				currentState = 97
			case 'G':
				currentState = 100
			case 'k':
				currentState = 57
			case 'H':
				currentState = 98
			case 'R':
				currentState = 103
			case 'S':
				currentState = 78
			case 'X':
				currentState = 68
			case 'E':
				currentState = 95
			case 'K':
				currentState = 63
			case 'W':
				currentState = 56
			case 'a':
				currentState = 104
			case 'x':
				currentState = 93
			case 'T':
				currentState = 74
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case 'I':
				currentState = 62
			case 'n':
				currentState = 85
			case 'Z':
				currentState = 101
			case 'd':
				currentState = 66
			case 'j':
				currentState = 84
			case 'u':
				currentState = 55
			case 'U':
				currentState = 67
			case 'p':
				currentState = 70
			case 'q':
				currentState = 81
			case 'w':
				currentState = 71
			case 'L':
				currentState = 61
			case 'N':
				currentState = 99
			case 'O':
				currentState = 73
			case 'e':
				currentState = 105
			case 'l':
				currentState = 106
			case 'v':
				currentState = 96
			case 'y':
				currentState = 72
			case 'A':
				currentState = 77
			case 'C':
				currentState = 86
			case 'J':
				currentState = 59
			case 'Q':
				currentState = 91
			case 'b':
				currentState = 60
			case 'D':
				currentState = 94
			case 'P':
				currentState = 65
			case 'B':
				currentState = 88
			case 'c':
				currentState = 69
			case 'i':
				currentState = 80
			default:
				break outer
			}

		case 103:
			switch c {
			case 'F':
				currentState = 89
			case 'n':
				currentState = 85
			case 'B':
				currentState = 88
			case 'T':
				currentState = 74
			case 'e':
				currentState = 105
			case 's':
				currentState = 87
			case 'u':
				currentState = 55
			case 'S':
				currentState = 78
			case 'Q':
				currentState = 91
			case 'd':
				currentState = 66
			case 'o':
				currentState = 58
			case 't':
				currentState = 76
			case 'J':
				currentState = 59
			case 'K':
				currentState = 63
			case 'V':
				currentState = 92
			case 'Y':
				currentState = 75
			case 'y':
				currentState = 72
			case '_':
				currentState = 97
			case 'A':
				currentState = 77
			case 'E':
				currentState = 95
			case 'p':
				currentState = 70
			case 'z':
				currentState = 64
			case 'U':
				currentState = 67
			case 'l':
				currentState = 106
			case 'q':
				currentState = 81
			case 'i':
				currentState = 80
			case 'f':
				currentState = 102
			case 'I':
				currentState = 62
			case 'N':
				currentState = 99
			case 'P':
				currentState = 65
			case 'g':
				currentState = 82
			case 'D':
				currentState = 94
			case 'w':
				currentState = 71
			case 'C':
				currentState = 86
			case 'a':
				currentState = 104
			case 'h':
				currentState = 79
			case 'm':
				currentState = 54
			case 'L':
				currentState = 61
			case 'R':
				currentState = 103
			case 'c':
				currentState = 69
			case 'v':
				currentState = 96
			case 'G':
				currentState = 100
			case 'j':
				currentState = 84
			case 'H':
				currentState = 98
			case 'Z':
				currentState = 101
			case 'r':
				currentState = 83
			case 'O':
				currentState = 73
			case 'W':
				currentState = 56
			case 'x':
				currentState = 93
			case 'M':
				currentState = 90
			case 'b':
				currentState = 60
			case 'k':
				currentState = 57
			case 'X':
				currentState = 68
			default:
				break outer
			}

		case 104:
			switch c {
			case 'M':
				currentState = 90
			case 'Q':
				currentState = 91
			case 'p':
				currentState = 70
			case 'z':
				currentState = 64
			case 'I':
				currentState = 62
			case 'K':
				currentState = 63
			case 'L':
				currentState = 61
			case 'P':
				currentState = 65
			case 'm':
				currentState = 54
			case 's':
				currentState = 87
			case 'w':
				currentState = 71
			case 'x':
				currentState = 93
			case 'H':
				currentState = 98
			case 'o':
				currentState = 58
			case 'W':
				currentState = 56
			case 'j':
				currentState = 84
			case 'v':
				currentState = 96
			case 'J':
				currentState = 59
			case 'N':
				currentState = 99
			case 'V':
				currentState = 92
			case 'd':
				currentState = 66
			case 'T':
				currentState = 74
			case 'A':
				currentState = 77
			case 'F':
				currentState = 89
			case 'u':
				currentState = 55
			case 'y':
				currentState = 72
			case 'S':
				currentState = 78
			case 'l':
				currentState = 106
			case 'f':
				currentState = 102
			case 'k':
				currentState = 57
			case 'G':
				currentState = 100
			case 'Z':
				currentState = 101
			case 'E':
				currentState = 95
			case 'Y':
				currentState = 75
			case 'a':
				currentState = 104
			case 'r':
				currentState = 83
			case 'B':
				currentState = 88
			case 'D':
				currentState = 94
			case 'c':
				currentState = 69
			case 'g':
				currentState = 82
			case 'n':
				currentState = 85
			case 'C':
				currentState = 86
			case 'U':
				currentState = 67
			case 'e':
				currentState = 105
			case 'h':
				currentState = 79
			case 'i':
				currentState = 80
			case 'q':
				currentState = 81
			case 't':
				currentState = 76
			case 'X':
				currentState = 68
			case 'b':
				currentState = 60
			case '_':
				currentState = 97
			case 'O':
				currentState = 73
			case 'R':
				currentState = 103
			default:
				break outer
			}

		case 106:
			switch c {
			case 'M':
				currentState = 90
			case 'O':
				currentState = 73
			case 'y':
				currentState = 72
			case 'H':
				currentState = 98
			case 'w':
				currentState = 71
			case 'a':
				currentState = 104
			case 'i':
				currentState = 80
			case 'p':
				currentState = 70
			case 'X':
				currentState = 68
			case 'x':
				currentState = 93
			case 'E':
				currentState = 95
			case 'F':
				currentState = 89
			case 'R':
				currentState = 103
			case 'U':
				currentState = 67
			case 'Q':
				currentState = 91
			case 'j':
				currentState = 84
			case 'r':
				currentState = 83
			case 'B':
				currentState = 88
			case 'n':
				currentState = 85
			case 'v':
				currentState = 96
			case 'I':
				currentState = 62
			case 'u':
				currentState = 55
			case 'g':
				currentState = 82
			case 'm':
				currentState = 54
			case 'o':
				currentState = 58
			case 'J':
				currentState = 59
			case 'K':
				currentState = 63
			case 'V':
				currentState = 92
			case 'C':
				currentState = 86
			case 'P':
				currentState = 65
			case 'A':
				currentState = 77
			case 'd':
				currentState = 66
			case 'h':
				currentState = 79
			case 'z':
				currentState = 64
			case 'W':
				currentState = 56
			case 'q':
				currentState = 81
			case 't':
				currentState = 76
			case 'L':
				currentState = 61
			case 'S':
				currentState = 78
			case 'Y':
				currentState = 75
			case 'f':
				currentState = 102
			case 'D':
				currentState = 94
			case 'Z':
				currentState = 101
			case 'b':
				currentState = 60
			case 'c':
				currentState = 69
			case 's':
				currentState = 87
			case '_':
				currentState = 97
			case 'l':
				currentState = 106
			case 'G':
				currentState = 100
			case 'T':
				currentState = 74
			case 'e':
				currentState = 105
			case 'k':
				currentState = 57
			case 'N':
				currentState = 99
			default:
				break outer
			}

		case 20:
			switch c {
			case 'S':
				currentState = 78
			case 'i':
				currentState = 80
			case 'y':
				currentState = 72
			case 'L':
				currentState = 61
			case 'U':
				currentState = 67
			case 'c':
				currentState = 69
			case 'q':
				currentState = 81
			case 'w':
				currentState = 71
			case 'C':
				currentState = 86
			case 'r':
				currentState = 83
			case 'T':
				currentState = 74
			case 'H':
				currentState = 98
			case 'M':
				currentState = 90
			case 'e':
				currentState = 105
			case 'f':
				currentState = 102
			case 'o':
				currentState = 58
			case 'J':
				currentState = 59
			case 'Q':
				currentState = 91
			case 'V':
				currentState = 92
			case 'X':
				currentState = 68
			case 'm':
				currentState = 54
			case 'n':
				currentState = 85
			case 'G':
				currentState = 100
			case 'K':
				currentState = 63
			case 'P':
				currentState = 65
			case '_':
				currentState = 97
			case 'E':
				currentState = 95
			case 'a':
				currentState = 104
			case 'h':
				currentState = 79
			case 'Z':
				currentState = 101
			case 'k':
				currentState = 57
			case 't':
				currentState = 76
			case 'O':
				currentState = 73
			case 'N':
				currentState = 99
			case 'u':
				currentState = 55
			case 'z':
				currentState = 64
			case 'A':
				currentState = 77
			case 'F':
				currentState = 89
			case 'g':
				currentState = 82
			case 'B':
				currentState = 88
			case 'b':
				currentState = 60
			case 's':
				currentState = 87
			case 'x':
				currentState = 93
			case 'R':
				currentState = 103
			case 'd':
				currentState = 66
			case 'I':
				currentState = 62
			case 'j':
				currentState = 84
			case 'p':
				currentState = 70
			case 'v':
				currentState = 96
			case 'W':
				currentState = 56
			case 'D':
				currentState = 94
			case 'Y':
				currentState = 75
			case 'l':
				currentState = 106
			default:
				break outer
			}

		case 99:
			switch c {
			case 'T':
				currentState = 74
			case 'W':
				currentState = 56
			case 'i':
				currentState = 80
			case 'y':
				currentState = 72
			case '_':
				currentState = 97
			case 'M':
				currentState = 90
			case 'P':
				currentState = 65
			case 'Y':
				currentState = 75
			case 'l':
				currentState = 106
			case 'e':
				currentState = 105
			case 'o':
				currentState = 58
			case 'I':
				currentState = 62
			case 'S':
				currentState = 78
			case 'B':
				currentState = 88
			case 'J':
				currentState = 59
			case 'N':
				currentState = 99
			case 'v':
				currentState = 96
			case 'Q':
				currentState = 91
			case 'f':
				currentState = 102
			case 'A':
				currentState = 77
			case 'G':
				currentState = 100
			case 'x':
				currentState = 93
			case 'w':
				currentState = 71
			case 'g':
				currentState = 82
			case 'p':
				currentState = 70
			case 'z':
				currentState = 64
			case 'c':
				currentState = 69
			case 'd':
				currentState = 66
			case 'K':
				currentState = 63
			case 'U':
				currentState = 67
			case 'V':
				currentState = 92
			case 'C':
				currentState = 86
			case 'H':
				currentState = 98
			case 'n':
				currentState = 85
			case 'q':
				currentState = 81
			case 't':
				currentState = 76
			case 'E':
				currentState = 95
			case 'F':
				currentState = 89
			case 'Z':
				currentState = 101
			case 'b':
				currentState = 60
			case 'k':
				currentState = 57
			case 'r':
				currentState = 83
			case 's':
				currentState = 87
			case 'O':
				currentState = 73
			case 'X':
				currentState = 68
			case 'u':
				currentState = 55
			case 'D':
				currentState = 94
			case 'h':
				currentState = 79
			case 'm':
				currentState = 54
			case 'L':
				currentState = 61
			case 'a':
				currentState = 104
			case 'R':
				currentState = 103
			case 'j':
				currentState = 84
			default:
				break outer
			}

		case 12:
			switch c {
			case 'a':
				currentState = 104
			case 'k':
				currentState = 57
			case 'y':
				currentState = 72
			case 'z':
				currentState = 64
			case '_':
				currentState = 97
			case 'J':
				currentState = 59
			case 'K':
				currentState = 63
			case 'O':
				currentState = 73
			case 'W':
				currentState = 56
			case 's':
				currentState = 87
			case 'X':
				currentState = 68
			case 'f':
				currentState = 102
			case 'o':
				currentState = 58
			case 'w':
				currentState = 71
			case 'U':
				currentState = 67
			case 'g':
				currentState = 82
			case 'h':
				currentState = 79
			case 'Y':
				currentState = 75
			case 'r':
				currentState = 83
			case 'u':
				currentState = 55
			case 'E':
				currentState = 95
			case 'b':
				currentState = 60
			case 'j':
				currentState = 84
			case 'P':
				currentState = 65
			case 'V':
				currentState = 92
			case 'm':
				currentState = 54
			case 'D':
				currentState = 94
			case 'F':
				currentState = 89
			case 'M':
				currentState = 90
			case 'c':
				currentState = 69
			case 'p':
				currentState = 70
			case 'x':
				currentState = 93
			case 'L':
				currentState = 61
			case 'S':
				currentState = 78
			case 'l':
				currentState = 106
			case 'A':
				currentState = 77
			case 'H':
				currentState = 98
			case 'Q':
				currentState = 91
			case 'e':
				currentState = 105
			case 'd':
				currentState = 66
			case 'n':
				currentState = 85
			case 'C':
				currentState = 86
			case 'G':
				currentState = 100
			case 'Z':
				currentState = 101
			case 'R':
				currentState = 103
			case 'v':
				currentState = 96
			case 'I':
				currentState = 62
			case 't':
				currentState = 76
			case 'i':
				currentState = 80
			case 'q':
				currentState = 81
			case 'B':
				currentState = 88
			case 'N':
				currentState = 99
			case 'T':
				currentState = 74
			default:
				break outer
			}

		case 74:
			switch c {
			case 'M':
				currentState = 90
			case 'o':
				currentState = 58
			case 'p':
				currentState = 70
			case 'L':
				currentState = 61
			case 'U':
				currentState = 67
			case 'h':
				currentState = 79
			case 't':
				currentState = 76
			case 'B':
				currentState = 88
			case 'X':
				currentState = 68
			case 'f':
				currentState = 102
			case '_':
				currentState = 97
			case 'F':
				currentState = 89
			case 'S':
				currentState = 78
			case 'w':
				currentState = 71
			case 'x':
				currentState = 93
			case 'n':
				currentState = 85
			case 'J':
				currentState = 59
			case 'Q':
				currentState = 91
			case 'R':
				currentState = 103
			case 'V':
				currentState = 92
			case 'E':
				currentState = 95
			case 'm':
				currentState = 54
			case 'r':
				currentState = 83
			case 'C':
				currentState = 86
			case 'e':
				currentState = 105
			case 'g':
				currentState = 82
			case 'H':
				currentState = 98
			case 'I':
				currentState = 62
			case 'd':
				currentState = 66
			case 'z':
				currentState = 64
			case 'P':
				currentState = 65
			case 'T':
				currentState = 74
			case 'i':
				currentState = 80
			case 'D':
				currentState = 94
			case 'q':
				currentState = 81
			case 'v':
				currentState = 96
			case 'y':
				currentState = 72
			case 'k':
				currentState = 57
			case 'A':
				currentState = 77
			case 'G':
				currentState = 100
			case 'N':
				currentState = 99
			case 'Z':
				currentState = 101
			case 'b':
				currentState = 60
			case 'j':
				currentState = 84
			case 'l':
				currentState = 106
			case 'K':
				currentState = 63
			case 'Y':
				currentState = 75
			case 'a':
				currentState = 104
			case 'c':
				currentState = 69
			case 'O':
				currentState = 73
			case 'W':
				currentState = 56
			case 's':
				currentState = 87
			case 'u':
				currentState = 55
			default:
				break outer
			}

		case 46:
			switch c {
			case 'H':
				currentState = 98
			case 'J':
				currentState = 59
			case 'V':
				currentState = 92
			case 'Y':
				currentState = 75
			case 'b':
				currentState = 60
			case 'u':
				currentState = 55
			case 'W':
				currentState = 56
			case 'j':
				currentState = 84
			case 'o':
				currentState = 58
			case 'D':
				currentState = 94
			case 'R':
				currentState = 103
			case 'L':
				currentState = 61
			case 'K':
				currentState = 63
			case 'P':
				currentState = 65
			case 'S':
				currentState = 78
			case 'd':
				currentState = 66
			case 'I':
				currentState = 62
			case 'f':
				currentState = 102
			case 'k':
				currentState = 57
			case 'w':
				currentState = 71
			case 'C':
				currentState = 86
			case 'c':
				currentState = 69
			case 'm':
				currentState = 54
			case 'x':
				currentState = 93
			case 'z':
				currentState = 64
			case 'U':
				currentState = 67
			case 'e':
				currentState = 105
			case 'q':
				currentState = 81
			case 't':
				currentState = 76
			case 'E':
				currentState = 95
			case 'F':
				currentState = 89
			case 'a':
				currentState = 104
			case 'B':
				currentState = 88
			case 'T':
				currentState = 74
			case 'N':
				currentState = 99
			case 'O':
				currentState = 73
			case 'Z':
				currentState = 101
			case 'p':
				currentState = 70
			case 's':
				currentState = 87
			case '_':
				currentState = 97
			case 'Q':
				currentState = 91
			case 'l':
				currentState = 106
			case 'A':
				currentState = 77
			case 'X':
				currentState = 68
			case 'g':
				currentState = 82
			case 'n':
				currentState = 85
			case 'G':
				currentState = 100
			case 'h':
				currentState = 79
			case 'i':
				currentState = 80
			case 'r':
				currentState = 83
			case 'v':
				currentState = 96
			case 'y':
				currentState = 72
			case 'M':
				currentState = 90
			default:
				break outer
			}

		case 16:
			switch c {
			case 'Q':
				currentState = 91
			case '_':
				currentState = 97
			case 'L':
				currentState = 61
			case 'S':
				currentState = 78
			case 'j':
				currentState = 84
			case 'K':
				currentState = 63
			case 'Z':
				currentState = 101
			case 'i':
				currentState = 80
			case 'm':
				currentState = 54
			case 'n':
				currentState = 85
			case 'P':
				currentState = 65
			case 'e':
				currentState = 105
			case 'p':
				currentState = 70
			case 'w':
				currentState = 71
			case 'B':
				currentState = 88
			case 'C':
				currentState = 86
			case 'I':
				currentState = 62
			case 'T':
				currentState = 74
			case 'Y':
				currentState = 75
			case 'h':
				currentState = 79
			case 'E':
				currentState = 95
			case 'd':
				currentState = 66
			case 'k':
				currentState = 57
			case 'r':
				currentState = 83
			case 'u':
				currentState = 55
			case 'D':
				currentState = 94
			case 'H':
				currentState = 98
			case 'O':
				currentState = 73
			case 'R':
				currentState = 103
			case 'a':
				currentState = 104
			case 'y':
				currentState = 72
			case 'f':
				currentState = 102
			case 's':
				currentState = 87
			case 'v':
				currentState = 96
			case 'U':
				currentState = 67
			case 'c':
				currentState = 69
			case 'g':
				currentState = 82
			case 'A':
				currentState = 77
			case 'V':
				currentState = 92
			case 'l':
				currentState = 106
			case 'o':
				currentState = 58
			case 'z':
				currentState = 64
			case 'W':
				currentState = 56
			case 'G':
				currentState = 100
			case 'X':
				currentState = 68
			case 'q':
				currentState = 81
			case 't':
				currentState = 76
			case 'J':
				currentState = 59
			case 'N':
				currentState = 99
			case 'F':
				currentState = 89
			case 'M':
				currentState = 90
			case 'b':
				currentState = 60
			case 'x':
				currentState = 93
			default:
				break outer
			}

		case 31:
			switch c {
			case 'V':
				currentState = 92
			case 'e':
				currentState = 105
			case 'i':
				currentState = 80
			case 'k':
				currentState = 57
			case 'G':
				currentState = 100
			case 'T':
				currentState = 74
			case 't':
				currentState = 76
			case 'y':
				currentState = 72
			case 'z':
				currentState = 64
			case 'I':
				currentState = 62
			case 'j':
				currentState = 84
			case 'q':
				currentState = 81
			case 'A':
				currentState = 77
			case 'L':
				currentState = 61
			case 'd':
				currentState = 66
			case 'f':
				currentState = 102
			case '_':
				currentState = 97
			case 'W':
				currentState = 56
			case 'b':
				currentState = 60
			case 'p':
				currentState = 70
			case 'v':
				currentState = 96
			case 'D':
				currentState = 94
			case 'c':
				currentState = 69
			case 'B':
				currentState = 88
			case 'E':
				currentState = 95
			case 'm':
				currentState = 54
			case 'J':
				currentState = 59
			case 'Y':
				currentState = 75
			case 'h':
				currentState = 79
			case 'u':
				currentState = 55
			case 'S':
				currentState = 78
			case 'H':
				currentState = 98
			case 's':
				currentState = 87
			case 'C':
				currentState = 86
			case 'O':
				currentState = 73
			case 'o':
				currentState = 58
			case 'w':
				currentState = 71
			case 'x':
				currentState = 93
			case 'F':
				currentState = 89
			case 'Q':
				currentState = 91
			case 'a':
				currentState = 104
			case 'K':
				currentState = 63
			case 'R':
				currentState = 103
			case 'n':
				currentState = 85
			case 'N':
				currentState = 99
			case 'P':
				currentState = 65
			case 'Z':
				currentState = 101
			case 'g':
				currentState = 82
			case 'l':
				currentState = 106
			case 'M':
				currentState = 90
			case 'U':
				currentState = 67
			case 'X':
				currentState = 68
			case 'r':
				currentState = 83
			default:
				break outer
			}

		case 55:
			switch c {
			case 'N':
				currentState = 99
			case 'a':
				currentState = 104
			case 'r':
				currentState = 83
			case 'Y':
				currentState = 75
			case 'g':
				currentState = 82
			case 'W':
				currentState = 56
			case 'q':
				currentState = 81
			case 'f':
				currentState = 102
			case 'j':
				currentState = 84
			case 'E':
				currentState = 95
			case 'J':
				currentState = 59
			case 'm':
				currentState = 54
			case 'G':
				currentState = 100
			case 'b':
				currentState = 60
			case 'p':
				currentState = 70
			case 'y':
				currentState = 72
			case 'z':
				currentState = 64
			case 'K':
				currentState = 63
			case 'T':
				currentState = 74
			case 'V':
				currentState = 92
			case 'e':
				currentState = 105
			case 'h':
				currentState = 79
			case 's':
				currentState = 87
			case 'x':
				currentState = 93
			case 'D':
				currentState = 94
			case 'M':
				currentState = 90
			case 'A':
				currentState = 77
			case 'I':
				currentState = 62
			case 'O':
				currentState = 73
			case 'P':
				currentState = 65
			case 'c':
				currentState = 69
			case '_':
				currentState = 97
			case 'Z':
				currentState = 101
			case 'l':
				currentState = 106
			case 'o':
				currentState = 58
			case 'u':
				currentState = 55
			case 'w':
				currentState = 71
			case 'B':
				currentState = 88
			case 'F':
				currentState = 89
			case 'H':
				currentState = 98
			case 'L':
				currentState = 61
			case 'Q':
				currentState = 91
			case 'S':
				currentState = 78
			case 'U':
				currentState = 67
			case 'X':
				currentState = 68
			case 'i':
				currentState = 80
			case 'k':
				currentState = 57
			case 't':
				currentState = 76
			case 'v':
				currentState = 96
			case 'R':
				currentState = 103
			case 'C':
				currentState = 86
			case 'd':
				currentState = 66
			case 'n':
				currentState = 85
			default:
				break outer
			}

		case 72:
			switch c {
			case 'v':
				currentState = 96
			case 'y':
				currentState = 72
			case 'a':
				currentState = 104
			case 'n':
				currentState = 85
			case 's':
				currentState = 87
			case 'F':
				currentState = 89
			case 'H':
				currentState = 98
			case 'b':
				currentState = 60
			case 'i':
				currentState = 80
			case '_':
				currentState = 97
			case 'P':
				currentState = 65
			case 'f':
				currentState = 102
			case 't':
				currentState = 76
			case 'A':
				currentState = 77
			case 'D':
				currentState = 94
			case 'K':
				currentState = 63
			case 'z':
				currentState = 64
			case 'R':
				currentState = 103
			case 'X':
				currentState = 68
			case 'h':
				currentState = 79
			case 'L':
				currentState = 61
			case 'c':
				currentState = 69
			case 'x':
				currentState = 93
			case 'd':
				currentState = 66
			case 'V':
				currentState = 92
			case 'l':
				currentState = 106
			case 'm':
				currentState = 54
			case 'w':
				currentState = 71
			case 'O':
				currentState = 73
			case 'Y':
				currentState = 75
			case 'k':
				currentState = 57
			case 'J':
				currentState = 59
			case 'C':
				currentState = 86
			case 'E':
				currentState = 95
			case 'Z':
				currentState = 101
			case 'q':
				currentState = 81
			case 'u':
				currentState = 55
			case 'B':
				currentState = 88
			case 'I':
				currentState = 62
			case 'W':
				currentState = 56
			case 'M':
				currentState = 90
			case 'S':
				currentState = 78
			case 'g':
				currentState = 82
			case 'Q':
				currentState = 91
			case 'p':
				currentState = 70
			case 'r':
				currentState = 83
			case 'N':
				currentState = 99
			case 'T':
				currentState = 74
			case 'U':
				currentState = 67
			case 'o':
				currentState = 58
			case 'G':
				currentState = 100
			case 'e':
				currentState = 105
			case 'j':
				currentState = 84
			default:
				break outer
			}

		case 2:
			switch c {
			case 's':
				currentState = 87
			case 'f':
				currentState = 102
			case 'r':
				currentState = 83
			case 'Y':
				currentState = 75
			case 'a':
				currentState = 104
			case 'j':
				currentState = 84
			case 'B':
				currentState = 88
			case 'R':
				currentState = 103
			case 'h':
				currentState = 79
			case 'U':
				currentState = 67
			case 'X':
				currentState = 68
			case 'W':
				currentState = 56
			case 'L':
				currentState = 61
			case 'c':
				currentState = 69
			case 'e':
				currentState = 105
			case 'o':
				currentState = 58
			case 'q':
				currentState = 81
			case 'I':
				currentState = 62
			case 'V':
				currentState = 92
			case 'E':
				currentState = 95
			case 'P':
				currentState = 65
			case '_':
				currentState = 97
			case 'n':
				currentState = 85
			case 'v':
				currentState = 96
			case 'z':
				currentState = 64
			case 'K':
				currentState = 63
			case 'i':
				currentState = 80
			case 'l':
				currentState = 106
			case 'p':
				currentState = 70
			case 'u':
				currentState = 55
			case 'T':
				currentState = 74
			case 'd':
				currentState = 66
			case 'y':
				currentState = 72
			case 'D':
				currentState = 94
			case 'Z':
				currentState = 101
			case 'F':
				currentState = 89
			case 'x':
				currentState = 93
			case 't':
				currentState = 76
			case 'C':
				currentState = 86
			case 'k':
				currentState = 57
			case 'Q':
				currentState = 91
			case 'm':
				currentState = 54
			case 'w':
				currentState = 71
			case 'H':
				currentState = 98
			case 'N':
				currentState = 99
			case 'O':
				currentState = 73
			case 'g':
				currentState = 82
			case 'b':
				currentState = 60
			case 'J':
				currentState = 59
			case 'M':
				currentState = 90
			case 'S':
				currentState = 78
			case 'A':
				currentState = 77
			case 'G':
				currentState = 100
			default:
				break outer
			}

		case 71:
			switch c {
			case 'o':
				currentState = 58
			case 'H':
				currentState = 98
			case 'N':
				currentState = 99
			case 'R':
				currentState = 103
			case 'W':
				currentState = 56
			case 'i':
				currentState = 80
			case 'Q':
				currentState = 91
			case 'e':
				currentState = 105
			case 's':
				currentState = 87
			case 'y':
				currentState = 72
			case 'J':
				currentState = 59
			case 'P':
				currentState = 65
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case 'w':
				currentState = 71
			case 'D':
				currentState = 94
			case 'K':
				currentState = 63
			case 'L':
				currentState = 61
			case 'V':
				currentState = 92
			case 'k':
				currentState = 57
			case 'X':
				currentState = 68
			case 'u':
				currentState = 55
			case 'E':
				currentState = 95
			case 'F':
				currentState = 89
			case 'M':
				currentState = 90
			case 'Y':
				currentState = 75
			case 'b':
				currentState = 60
			case 'B':
				currentState = 88
			case 'O':
				currentState = 73
			case 'n':
				currentState = 85
			case 'r':
				currentState = 83
			case 'T':
				currentState = 74
			case 'U':
				currentState = 67
			case 'c':
				currentState = 69
			case 'd':
				currentState = 66
			case 'j':
				currentState = 84
			case 'v':
				currentState = 96
			case 'Z':
				currentState = 101
			case 'p':
				currentState = 70
			case 'q':
				currentState = 81
			case 'A':
				currentState = 77
			case 'C':
				currentState = 86
			case 'x':
				currentState = 93
			case 'z':
				currentState = 64
			case '_':
				currentState = 97
			case 'I':
				currentState = 62
			case 'S':
				currentState = 78
			case 'a':
				currentState = 104
			case 'G':
				currentState = 100
			case 'h':
				currentState = 79
			case 'l':
				currentState = 106
			case 'm':
				currentState = 54
			case 't':
				currentState = 76
			default:
				break outer
			}

		case 95:
			switch c {
			case 'x':
				currentState = 93
			case 'A':
				currentState = 77
			case 'T':
				currentState = 74
			case 'a':
				currentState = 104
			case 'i':
				currentState = 80
			case 'o':
				currentState = 58
			case 'q':
				currentState = 81
			case 'B':
				currentState = 88
			case 'd':
				currentState = 66
			case 'j':
				currentState = 84
			case '_':
				currentState = 97
			case 'O':
				currentState = 73
			case 'l':
				currentState = 106
			case 'p':
				currentState = 70
			case 'F':
				currentState = 89
			case 'E':
				currentState = 95
			case 'M':
				currentState = 90
			case 'U':
				currentState = 67
			case 'V':
				currentState = 92
			case 'h':
				currentState = 79
			case 'J':
				currentState = 59
			case 'G':
				currentState = 100
			case 'H':
				currentState = 98
			case 'N':
				currentState = 99
			case 'W':
				currentState = 56
			case 'c':
				currentState = 69
			case 'f':
				currentState = 102
			case 'z':
				currentState = 64
			case 'P':
				currentState = 65
			case 'Y':
				currentState = 75
			case 'b':
				currentState = 60
			case 'k':
				currentState = 57
			case 'm':
				currentState = 54
			case 't':
				currentState = 76
			case 'D':
				currentState = 94
			case 'n':
				currentState = 85
			case 'w':
				currentState = 71
			case 'L':
				currentState = 61
			case 'S':
				currentState = 78
			case 'v':
				currentState = 96
			case 'K':
				currentState = 63
			case 'Z':
				currentState = 101
			case 'e':
				currentState = 105
			case 'C':
				currentState = 86
			case 'u':
				currentState = 55
			case 'Q':
				currentState = 91
			case 'R':
				currentState = 103
			case 'X':
				currentState = 68
			case 'y':
				currentState = 72
			case 'I':
				currentState = 62
			case 'r':
				currentState = 83
			case 's':
				currentState = 87
			case 'g':
				currentState = 82
			default:
				break outer
			}

		case 21:
			switch c {
			case 'A':
				currentState = 77
			case 'H':
				currentState = 98
			case 'I':
				currentState = 62
			case 'O':
				currentState = 73
			case 'P':
				currentState = 65
			case 'w':
				currentState = 71
			case 'V':
				currentState = 92
			case 'm':
				currentState = 54
			case 'p':
				currentState = 70
			case 'v':
				currentState = 96
			case 'u':
				currentState = 55
			case 'R':
				currentState = 103
			case 'g':
				currentState = 82
			case 'z':
				currentState = 64
			case '_':
				currentState = 97
			case 'Z':
				currentState = 101
			case 'd':
				currentState = 66
			case 'j':
				currentState = 84
			case 'k':
				currentState = 57
			case 's':
				currentState = 87
			case 'x':
				currentState = 93
			case 'M':
				currentState = 90
			case 'l':
				currentState = 106
			case 'q':
				currentState = 81
			case 't':
				currentState = 76
			case 'E':
				currentState = 95
			case 'U':
				currentState = 67
			case 'n':
				currentState = 85
			case 'F':
				currentState = 89
			case 'G':
				currentState = 100
			case 'r':
				currentState = 83
			case 'D':
				currentState = 94
			case 'a':
				currentState = 104
			case 'Q':
				currentState = 91
			case 'T':
				currentState = 74
			case 'Y':
				currentState = 75
			case 'y':
				currentState = 72
			case 'K':
				currentState = 63
			case 'c':
				currentState = 69
			case 'X':
				currentState = 68
			case 'i':
				currentState = 80
			case 'o':
				currentState = 58
			case 'f':
				currentState = 102
			case 'C':
				currentState = 86
			case 'N':
				currentState = 99
			case 'B':
				currentState = 88
			case 'S':
				currentState = 78
			case 'W':
				currentState = 56
			case 'b':
				currentState = 60
			case 'h':
				currentState = 79
			case 'J':
				currentState = 59
			case 'L':
				currentState = 61
			case 'e':
				currentState = 105
			default:
				break outer
			}

		case 35:
			switch c {
			case 'd':
				currentState = 66
			case 'f':
				currentState = 102
			case 's':
				currentState = 87
			case 'M':
				currentState = 90
			case 'S':
				currentState = 78
			case 'p':
				currentState = 70
			case 'B':
				currentState = 88
			case 'L':
				currentState = 61
			case 'b':
				currentState = 60
			case 'u':
				currentState = 55
			case 'w':
				currentState = 71
			case 'g':
				currentState = 82
			case 'n':
				currentState = 85
			case 'q':
				currentState = 81
			case '_':
				currentState = 97
			case 'W':
				currentState = 56
			case 'Y':
				currentState = 75
			case 'h':
				currentState = 79
			case 'k':
				currentState = 57
			case 'T':
				currentState = 74
			case 'e':
				currentState = 105
			case 'E':
				currentState = 95
			case 'Q':
				currentState = 91
			case 'R':
				currentState = 103
			case 'c':
				currentState = 69
			case 't':
				currentState = 76
			case 'J':
				currentState = 59
			case 'N':
				currentState = 99
			case 'o':
				currentState = 58
			case 'H':
				currentState = 98
			case 'O':
				currentState = 73
			case 'i':
				currentState = 80
			case 'v':
				currentState = 96
			case 'D':
				currentState = 94
			case 'X':
				currentState = 68
			case 'j':
				currentState = 84
			case 'y':
				currentState = 72
			case 'z':
				currentState = 64
			case 'F':
				currentState = 89
			case 'V':
				currentState = 92
			case 'a':
				currentState = 104
			case 'm':
				currentState = 54
			case 'x':
				currentState = 93
			case 'C':
				currentState = 86
			case 'U':
				currentState = 67
			case 'Z':
				currentState = 101
			case 'K':
				currentState = 63
			case 'r':
				currentState = 83
			case 'P':
				currentState = 65
			case 'l':
				currentState = 106
			case 'A':
				currentState = 77
			case 'G':
				currentState = 100
			case 'I':
				currentState = 62
			default:
				break outer
			}

		case 52:
			switch c {
			case 'R':
				currentState = 103
			case 'v':
				currentState = 96
			case 'Z':
				currentState = 101
			case 'm':
				currentState = 54
			case 'w':
				currentState = 71
			case '_':
				currentState = 97
			case 'C':
				currentState = 86
			case 'D':
				currentState = 94
			case 'y':
				currentState = 72
			case 'T':
				currentState = 74
			case 'o':
				currentState = 58
			case 'V':
				currentState = 92
			case 'W':
				currentState = 56
			case 'e':
				currentState = 105
			case 'B':
				currentState = 88
			case 'H':
				currentState = 98
			case 'L':
				currentState = 61
			case 'Y':
				currentState = 75
			case 'f':
				currentState = 102
			case 'x':
				currentState = 93
			case 'h':
				currentState = 79
			case 'k':
				currentState = 57
			case 'A':
				currentState = 77
			case 'X':
				currentState = 68
			case 'P':
				currentState = 65
			case 'U':
				currentState = 67
			case 'O':
				currentState = 73
			case 'd':
				currentState = 66
			case 'i':
				currentState = 80
			case 'p':
				currentState = 70
			case 'E':
				currentState = 95
			case 'M':
				currentState = 90
			case 'S':
				currentState = 78
			case 'l':
				currentState = 106
			case 'Q':
				currentState = 91
			case 'a':
				currentState = 104
			case 'q':
				currentState = 81
			case 'u':
				currentState = 55
			case 'J':
				currentState = 59
			case 'N':
				currentState = 99
			case 'K':
				currentState = 63
			case 'b':
				currentState = 60
			case 's':
				currentState = 87
			case 'G':
				currentState = 100
			case 'j':
				currentState = 84
			case 'r':
				currentState = 83
			case 't':
				currentState = 76
			case 'z':
				currentState = 64
			case 'c':
				currentState = 69
			case 'g':
				currentState = 82
			case 'n':
				currentState = 85
			case 'F':
				currentState = 89
			case 'I':
				currentState = 62
			default:
				break outer
			}

		case 70:
			switch c {
			case 'f':
				currentState = 102
			case 'm':
				currentState = 54
			case 'p':
				currentState = 70
			case 'x':
				currentState = 93
			case 'J':
				currentState = 59
			case 'Q':
				currentState = 91
			case 'd':
				currentState = 66
			case 'K':
				currentState = 63
			case 'b':
				currentState = 60
			case 'a':
				currentState = 104
			case 'q':
				currentState = 81
			case 't':
				currentState = 76
			case 'H':
				currentState = 98
			case 'h':
				currentState = 79
			case 'B':
				currentState = 88
			case 'V':
				currentState = 92
			case 'c':
				currentState = 69
			case 'y':
				currentState = 72
			case 'P':
				currentState = 65
			case 'r':
				currentState = 83
			case 'l':
				currentState = 106
			case 'v':
				currentState = 96
			case 'I':
				currentState = 62
			case 'Z':
				currentState = 101
			case 'e':
				currentState = 105
			case 'i':
				currentState = 80
			case 'w':
				currentState = 71
			case 'R':
				currentState = 103
			case 'X':
				currentState = 68
			case 'Y':
				currentState = 75
			case 'n':
				currentState = 85
			case 's':
				currentState = 87
			case 'M':
				currentState = 90
			case 'N':
				currentState = 99
			case 'T':
				currentState = 74
			case 'W':
				currentState = 56
			case 'G':
				currentState = 100
			case 'U':
				currentState = 67
			case 'g':
				currentState = 82
			case 'A':
				currentState = 77
			case 'S':
				currentState = 78
			case 'o':
				currentState = 58
			case 'C':
				currentState = 86
			case 'L':
				currentState = 61
			case 'j':
				currentState = 84
			case 'u':
				currentState = 55
			case 'z':
				currentState = 64
			case '_':
				currentState = 97
			case 'D':
				currentState = 94
			case 'E':
				currentState = 95
			case 'F':
				currentState = 89
			case 'O':
				currentState = 73
			case 'k':
				currentState = 57
			default:
				break outer
			}

		case 84:
			switch c {
			case 'W':
				currentState = 56
			case 'd':
				currentState = 66
			case 'I':
				currentState = 62
			case 'r':
				currentState = 83
			case 'q':
				currentState = 81
			case 'b':
				currentState = 60
			case 'c':
				currentState = 69
			case 'e':
				currentState = 105
			case 'm':
				currentState = 54
			case 'w':
				currentState = 71
			case 'K':
				currentState = 63
			case 'Q':
				currentState = 91
			case 'F':
				currentState = 89
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case 'T':
				currentState = 74
			case 'J':
				currentState = 59
			case 'R':
				currentState = 103
			case 'x':
				currentState = 93
			case 'A':
				currentState = 77
			case 'k':
				currentState = 57
			case 'M':
				currentState = 90
			case 'U':
				currentState = 67
			case 'Y':
				currentState = 75
			case 'i':
				currentState = 80
			case 'z':
				currentState = 64
			case 'H':
				currentState = 98
			case 'h':
				currentState = 79
			case 'o':
				currentState = 58
			case 'v':
				currentState = 96
			case 'P':
				currentState = 65
			case 'D':
				currentState = 94
			case 't':
				currentState = 76
			case 'C':
				currentState = 86
			case 'G':
				currentState = 100
			case 'S':
				currentState = 78
			case 'V':
				currentState = 92
			case 'a':
				currentState = 104
			case 'l':
				currentState = 106
			case 'u':
				currentState = 55
			case 'B':
				currentState = 88
			case 'X':
				currentState = 68
			case 'p':
				currentState = 70
			case 'y':
				currentState = 72
			case 'E':
				currentState = 95
			case 'n':
				currentState = 85
			case 's':
				currentState = 87
			case '_':
				currentState = 97
			case 'Z':
				currentState = 101
			case 'O':
				currentState = 73
			case 'L':
				currentState = 61
			case 'j':
				currentState = 84
			case 'N':
				currentState = 99
			default:
				break outer
			}

		case 11:
			switch c {
			case 'y':
				currentState = 72
			case 'b':
				currentState = 60
			case 'g':
				currentState = 82
			case 'o':
				currentState = 58
			case 'T':
				currentState = 74
			case 'r':
				currentState = 83
			case 'n':
				currentState = 85
			case 'D':
				currentState = 94
			case 'E':
				currentState = 95
			case 'N':
				currentState = 99
			case 'O':
				currentState = 73
			case 'd':
				currentState = 66
			case 'l':
				currentState = 106
			case 'm':
				currentState = 54
			case 't':
				currentState = 76
			case 'H':
				currentState = 98
			case 'U':
				currentState = 67
			case 'X':
				currentState = 68
			case 'c':
				currentState = 69
			case 'e':
				currentState = 105
			case 'k':
				currentState = 57
			case 'u':
				currentState = 55
			case 'C':
				currentState = 86
			case 'I':
				currentState = 62
			case 'S':
				currentState = 78
			case 'W':
				currentState = 56
			case '_':
				currentState = 97
			case 'B':
				currentState = 88
			case 'h':
				currentState = 79
			case 'z':
				currentState = 64
			case 'K':
				currentState = 63
			case 'a':
				currentState = 104
			case 'A':
				currentState = 77
			case 'G':
				currentState = 100
			case 'j':
				currentState = 84
			case 'J':
				currentState = 59
			case 'L':
				currentState = 61
			case 'Q':
				currentState = 91
			case 'V':
				currentState = 92
			case 'Z':
				currentState = 101
			case 'i':
				currentState = 80
			case 'w':
				currentState = 71
			case 'M':
				currentState = 90
			case 'R':
				currentState = 103
			case 'Y':
				currentState = 75
			case 'v':
				currentState = 96
			case 'P':
				currentState = 65
			case 's':
				currentState = 87
			case 'F':
				currentState = 89
			case 'p':
				currentState = 70
			case 'q':
				currentState = 81
			case 'x':
				currentState = 93
			case 'f':
				currentState = 102
			default:
				break outer
			}

		case 32:
			switch c {
			case 'J':
				currentState = 59
			case 'T':
				currentState = 74
			case 'k':
				currentState = 57
			case 'z':
				currentState = 64
			case 'B':
				currentState = 88
			case 'X':
				currentState = 68
			case 'n':
				currentState = 85
			case 'r':
				currentState = 83
			case 'w':
				currentState = 71
			case 'C':
				currentState = 86
			case 'Z':
				currentState = 101
			case 'c':
				currentState = 69
			case 'j':
				currentState = 84
			case 'o':
				currentState = 58
			case 'A':
				currentState = 77
			case 'F':
				currentState = 89
			case 'I':
				currentState = 62
			case 'p':
				currentState = 70
			case 's':
				currentState = 87
			case 'E':
				currentState = 95
			case 'x':
				currentState = 93
			case 'O':
				currentState = 73
			case 'Q':
				currentState = 91
			case 'N':
				currentState = 99
			case 'l':
				currentState = 106
			case 'P':
				currentState = 65
			case 'a':
				currentState = 104
			case 'i':
				currentState = 80
			case 'L':
				currentState = 61
			case 'K':
				currentState = 63
			case 'R':
				currentState = 103
			case 'H':
				currentState = 98
			case 'U':
				currentState = 67
			case 'm':
				currentState = 54
			case 'q':
				currentState = 81
			case 'V':
				currentState = 92
			case 'y':
				currentState = 72
			case 'D':
				currentState = 94
			case 't':
				currentState = 76
			case 'g':
				currentState = 82
			case 'v':
				currentState = 96
			case 'd':
				currentState = 66
			case 'W':
				currentState = 56
			case 'Y':
				currentState = 75
			case 'h':
				currentState = 79
			case '_':
				currentState = 97
			case 'G':
				currentState = 100
			case 'S':
				currentState = 78
			case 'b':
				currentState = 60
			case 'e':
				currentState = 105
			case 'f':
				currentState = 102
			case 'u':
				currentState = 55
			case 'M':
				currentState = 90
			default:
				break outer
			}

		case 100:
			switch c {
			case 'b':
				currentState = 60
			case 'D':
				currentState = 94
			case 'H':
				currentState = 98
			case 'N':
				currentState = 99
			case 'c':
				currentState = 69
			case 'e':
				currentState = 105
			case 'A':
				currentState = 77
			case 'j':
				currentState = 84
			case 'p':
				currentState = 70
			case 'w':
				currentState = 71
			case 'B':
				currentState = 88
			case 'T':
				currentState = 74
			case 'a':
				currentState = 104
			case 'd':
				currentState = 66
			case 'J':
				currentState = 59
			case 'P':
				currentState = 65
			case 'Y':
				currentState = 75
			case 'y':
				currentState = 72
			case 'Q':
				currentState = 91
			case 'S':
				currentState = 78
			case 'V':
				currentState = 92
			case 'W':
				currentState = 56
			case 'm':
				currentState = 54
			case 'o':
				currentState = 58
			case 'v':
				currentState = 96
			case 'C':
				currentState = 86
			case 'E':
				currentState = 95
			case 'I':
				currentState = 62
			case 'X':
				currentState = 68
			case 'q':
				currentState = 81
			case 'r':
				currentState = 83
			case 't':
				currentState = 76
			case 'u':
				currentState = 55
			case 'R':
				currentState = 103
			case 'U':
				currentState = 67
			case 'L':
				currentState = 61
			case 'l':
				currentState = 106
			case 'n':
				currentState = 85
			case 's':
				currentState = 87
			case 'G':
				currentState = 100
			case 'Z':
				currentState = 101
			case 'z':
				currentState = 64
			case 'M':
				currentState = 90
			case 'O':
				currentState = 73
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case 'h':
				currentState = 79
			case 'K':
				currentState = 63
			case 'i':
				currentState = 80
			case 'k':
				currentState = 57
			case '_':
				currentState = 97
			case 'x':
				currentState = 93
			case 'F':
				currentState = 89
			default:
				break outer
			}

		case 0:
			switch c {
			case 'k':
				currentState = 15
			case 'q':
				currentState = 49
			case 'Y':
				currentState = 12
			case 'C':
				currentState = 53
			case 'S':
				currentState = 28
			case 'f':
				currentState = 2
			case 'o':
				currentState = 3
			case 'l':
				currentState = 8
			case 'x':
				currentState = 9
			case 'H':
				currentState = 10
			case 'J':
				currentState = 32
			case 'g':
				currentState = 14
			case 'h':
				currentState = 25
			case 'v':
				currentState = 30
			case 'z':
				currentState = 18
			case 'N':
				currentState = 26
			case 'a':
				currentState = 19
			case 'V':
				currentState = 1
			case 'K':
				currentState = 48
			case 'R':
				currentState = 29
			case 'T':
				currentState = 43
			case 'p':
				currentState = 52
			case 's':
				currentState = 44
			case 'G':
				currentState = 41
			case 'Q':
				currentState = 37
			case 'c':
				currentState = 7
			case 'r':
				currentState = 22
			case 'U':
				currentState = 33
			case 'b':
				currentState = 17
			case 'u':
				currentState = 5
			case 'y':
				currentState = 31
			case '_':
				currentState = 35
			case 'Z':
				currentState = 13
			case 'i':
				currentState = 34
			case 'm':
				currentState = 39
			case 't':
				currentState = 27
			case 'P':
				currentState = 42
			case 'j':
				currentState = 20
			case 'B':
				currentState = 23
			case 'd':
				currentState = 47
			case 'W':
				currentState = 11
			case 'n':
				currentState = 21
			case 'I':
				currentState = 36
			case 'D':
				currentState = 4
			case 'E':
				currentState = 46
			case 'X':
				currentState = 38
			case 'w':
				currentState = 45
			case 'A':
				currentState = 6
			case 'M':
				currentState = 51
			case 'e':
				currentState = 24
			case 'F':
				currentState = 50
			case 'L':
				currentState = 16
			case 'O':
				currentState = 40
			default:
				break outer
			}

		case 42:
			switch c {
			case 'V':
				currentState = 92
			case 'a':
				currentState = 104
			case 'd':
				currentState = 66
			case 'H':
				currentState = 98
			case 'K':
				currentState = 63
			case 'x':
				currentState = 93
			case 'y':
				currentState = 72
			case 'v':
				currentState = 96
			case 'I':
				currentState = 62
			case 'R':
				currentState = 103
			case 'i':
				currentState = 80
			case 'u':
				currentState = 55
			case 'r':
				currentState = 83
			case 'B':
				currentState = 88
			case 'D':
				currentState = 94
			case 'M':
				currentState = 90
			case 'e':
				currentState = 105
			case 'n':
				currentState = 85
			case 's':
				currentState = 87
			case 'J':
				currentState = 59
			case 'L':
				currentState = 61
			case 'Z':
				currentState = 101
			case 'm':
				currentState = 54
			case 'f':
				currentState = 102
			case 'l':
				currentState = 106
			case 'S':
				currentState = 78
			case 'p':
				currentState = 70
			case 't':
				currentState = 76
			case 'P':
				currentState = 65
			case 'g':
				currentState = 82
			case 'w':
				currentState = 71
			case '_':
				currentState = 97
			case 'T':
				currentState = 74
			case 'b':
				currentState = 60
			case 'j':
				currentState = 84
			case 'o':
				currentState = 58
			case 'E':
				currentState = 95
			case 'U':
				currentState = 67
			case 'W':
				currentState = 56
			case 'z':
				currentState = 64
			case 'Q':
				currentState = 91
			case 'h':
				currentState = 79
			case 'A':
				currentState = 77
			case 'N':
				currentState = 99
			case 'O':
				currentState = 73
			case 'X':
				currentState = 68
			case 'Y':
				currentState = 75
			case 'F':
				currentState = 89
			case 'c':
				currentState = 69
			case 'q':
				currentState = 81
			case 'C':
				currentState = 86
			case 'G':
				currentState = 100
			case 'k':
				currentState = 57
			default:
				break outer
			}

		case 61:
			switch c {
			case 'b':
				currentState = 60
			case 'r':
				currentState = 83
			case 'M':
				currentState = 90
			case 'R':
				currentState = 103
			case 'k':
				currentState = 57
			case 'P':
				currentState = 65
			case 'V':
				currentState = 92
			case 'a':
				currentState = 104
			case 'c':
				currentState = 69
			case 'l':
				currentState = 106
			case 'z':
				currentState = 64
			case 'L':
				currentState = 61
			case 'W':
				currentState = 56
			case 'E':
				currentState = 95
			case 'e':
				currentState = 105
			case 'm':
				currentState = 54
			case 't':
				currentState = 76
			case '_':
				currentState = 97
			case 'D':
				currentState = 94
			case 'F':
				currentState = 89
			case 'i':
				currentState = 80
			case 'o':
				currentState = 58
			case 'Y':
				currentState = 75
			case 'Z':
				currentState = 101
			case 'B':
				currentState = 88
			case 'G':
				currentState = 100
			case 'K':
				currentState = 63
			case 'O':
				currentState = 73
			case 'f':
				currentState = 102
			case 'h':
				currentState = 79
			case 'n':
				currentState = 85
			case 's':
				currentState = 87
			case 'U':
				currentState = 67
			case 'j':
				currentState = 84
			case 'N':
				currentState = 99
			case 'Q':
				currentState = 91
			case 'g':
				currentState = 82
			case 'A':
				currentState = 77
			case 'S':
				currentState = 78
			case 'u':
				currentState = 55
			case 'C':
				currentState = 86
			case 'H':
				currentState = 98
			case 'J':
				currentState = 59
			case 'p':
				currentState = 70
			case 'v':
				currentState = 96
			case 'X':
				currentState = 68
			case 'q':
				currentState = 81
			case 'x':
				currentState = 93
			case 'I':
				currentState = 62
			case 'd':
				currentState = 66
			case 'w':
				currentState = 71
			case 'y':
				currentState = 72
			case 'T':
				currentState = 74
			default:
				break outer
			}

		case 62:
			switch c {
			case 'H':
				currentState = 98
			case 'X':
				currentState = 68
			case 't':
				currentState = 76
			case 'O':
				currentState = 73
			case 'T':
				currentState = 74
			case 's':
				currentState = 87
			case 'U':
				currentState = 67
			case 'p':
				currentState = 70
			case 'C':
				currentState = 86
			case 'G':
				currentState = 100
			case 'c':
				currentState = 69
			case 'e':
				currentState = 105
			case 'l':
				currentState = 106
			case 'm':
				currentState = 54
			case 'v':
				currentState = 96
			case 'I':
				currentState = 62
			case 'h':
				currentState = 79
			case 'i':
				currentState = 80
			case 'j':
				currentState = 84
			case 'k':
				currentState = 57
			case 'V':
				currentState = 92
			case 'q':
				currentState = 81
			case 'b':
				currentState = 60
			case 'E':
				currentState = 95
			case 'K':
				currentState = 63
			case 'W':
				currentState = 56
			case 'Y':
				currentState = 75
			case 'g':
				currentState = 82
			case 'w':
				currentState = 71
			case 'D':
				currentState = 94
			case 'F':
				currentState = 89
			case 'n':
				currentState = 85
			case 'u':
				currentState = 55
			case 'z':
				currentState = 64
			case '_':
				currentState = 97
			case 'Z':
				currentState = 101
			case 'a':
				currentState = 104
			case 'd':
				currentState = 66
			case 'J':
				currentState = 59
			case 'P':
				currentState = 65
			case 'Q':
				currentState = 91
			case 'f':
				currentState = 102
			case 'x':
				currentState = 93
			case 'A':
				currentState = 77
			case 'B':
				currentState = 88
			case 'S':
				currentState = 78
			case 'o':
				currentState = 58
			case 'r':
				currentState = 83
			case 'M':
				currentState = 90
			case 'N':
				currentState = 99
			case 'R':
				currentState = 103
			case 'y':
				currentState = 72
			case 'L':
				currentState = 61
			default:
				break outer
			}

		case 66:
			switch c {
			case 'u':
				currentState = 55
			case 'M':
				currentState = 90
			case 'Z':
				currentState = 101
			case 'p':
				currentState = 70
			case 'y':
				currentState = 72
			case 'O':
				currentState = 73
			case 'Q':
				currentState = 91
			case 'R':
				currentState = 103
			case 'S':
				currentState = 78
			case 'W':
				currentState = 56
			case 'H':
				currentState = 98
			case 'N':
				currentState = 99
			case 'V':
				currentState = 92
			case 'Y':
				currentState = 75
			case 'm':
				currentState = 54
			case 'q':
				currentState = 81
			case 'C':
				currentState = 86
			case 'G':
				currentState = 100
			case 'I':
				currentState = 62
			case 'i':
				currentState = 80
			case 'o':
				currentState = 58
			case '_':
				currentState = 97
			case 'U':
				currentState = 67
			case 'e':
				currentState = 105
			case 'f':
				currentState = 102
			case 'T':
				currentState = 74
			case 'a':
				currentState = 104
			case 'A':
				currentState = 77
			case 'K':
				currentState = 63
			case 'P':
				currentState = 65
			case 'E':
				currentState = 95
			case 't':
				currentState = 76
			case 'D':
				currentState = 94
			case 'n':
				currentState = 85
			case 'r':
				currentState = 83
			case 'k':
				currentState = 57
			case 'w':
				currentState = 71
			case 'z':
				currentState = 64
			case 'J':
				currentState = 59
			case 'h':
				currentState = 79
			case 'v':
				currentState = 96
			case 'x':
				currentState = 93
			case 'L':
				currentState = 61
			case 'b':
				currentState = 60
			case 'g':
				currentState = 82
			case 'j':
				currentState = 84
			case 's':
				currentState = 87
			case 'B':
				currentState = 88
			case 'X':
				currentState = 68
			case 'd':
				currentState = 66
			case 'F':
				currentState = 89
			case 'c':
				currentState = 69
			case 'l':
				currentState = 106
			default:
				break outer
			}

		case 73:
			switch c {
			case 'F':
				currentState = 89
			case 'n':
				currentState = 85
			case 'p':
				currentState = 70
			case 'C':
				currentState = 86
			case 'J':
				currentState = 59
			case 'R':
				currentState = 103
			case 'S':
				currentState = 78
			case 'k':
				currentState = 57
			case 'B':
				currentState = 88
			case 'G':
				currentState = 100
			case 'Q':
				currentState = 91
			case 'U':
				currentState = 67
			case 'm':
				currentState = 54
			case 't':
				currentState = 76
			case 'T':
				currentState = 74
			case 'H':
				currentState = 98
			case 'M':
				currentState = 90
			case 'Y':
				currentState = 75
			case 'v':
				currentState = 96
			case 'K':
				currentState = 63
			case 'e':
				currentState = 105
			case 'o':
				currentState = 58
			case 'z':
				currentState = 64
			case 'A':
				currentState = 77
			case 'N':
				currentState = 99
			case 'Z':
				currentState = 101
			case 'f':
				currentState = 102
			case 'g':
				currentState = 82
			case 'r':
				currentState = 83
			case 'L':
				currentState = 61
			case 'W':
				currentState = 56
			case 'd':
				currentState = 66
			case 'i':
				currentState = 80
			case '_':
				currentState = 97
			case 'I':
				currentState = 62
			case 'E':
				currentState = 95
			case 'w':
				currentState = 71
			case 'x':
				currentState = 93
			case 'D':
				currentState = 94
			case 'X':
				currentState = 68
			case 'a':
				currentState = 104
			case 'b':
				currentState = 60
			case 'j':
				currentState = 84
			case 'u':
				currentState = 55
			case 'l':
				currentState = 106
			case 'y':
				currentState = 72
			case 's':
				currentState = 87
			case 'O':
				currentState = 73
			case 'P':
				currentState = 65
			case 'c':
				currentState = 69
			case 'h':
				currentState = 79
			case 'V':
				currentState = 92
			case 'q':
				currentState = 81
			default:
				break outer
			}

		case 76:
			switch c {
			case 'H':
				currentState = 98
			case 'l':
				currentState = 106
			case 'v':
				currentState = 96
			case 'C':
				currentState = 86
			case 'Q':
				currentState = 91
			case 'o':
				currentState = 58
			case 'q':
				currentState = 81
			case 'j':
				currentState = 84
			case 'z':
				currentState = 64
			case 'I':
				currentState = 62
			case 'J':
				currentState = 59
			case 'K':
				currentState = 63
			case 'W':
				currentState = 56
			case 'Y':
				currentState = 75
			case 'L':
				currentState = 61
			case 'R':
				currentState = 103
			case 'p':
				currentState = 70
			case 'T':
				currentState = 74
			case 'd':
				currentState = 66
			case 'f':
				currentState = 102
			case 'D':
				currentState = 94
			case 'E':
				currentState = 95
			case 'M':
				currentState = 90
			case 'V':
				currentState = 92
			case 'k':
				currentState = 57
			case 't':
				currentState = 76
			case 'A':
				currentState = 77
			case 'F':
				currentState = 89
			case 'O':
				currentState = 73
			case 'P':
				currentState = 65
			case '_':
				currentState = 97
			case 'N':
				currentState = 99
			case 'u':
				currentState = 55
			case 'y':
				currentState = 72
			case 'B':
				currentState = 88
			case 'h':
				currentState = 79
			case 'm':
				currentState = 54
			case 'S':
				currentState = 78
			case 'g':
				currentState = 82
			case 'x':
				currentState = 93
			case 'U':
				currentState = 67
			case 'i':
				currentState = 80
			case 'w':
				currentState = 71
			case 'X':
				currentState = 68
			case 'b':
				currentState = 60
			case 'e':
				currentState = 105
			case 's':
				currentState = 87
			case 'r':
				currentState = 83
			case 'G':
				currentState = 100
			case 'Z':
				currentState = 101
			case 'a':
				currentState = 104
			case 'c':
				currentState = 69
			case 'n':
				currentState = 85
			default:
				break outer
			}

		case 8:
			switch c {
			case 'j':
				currentState = 84
			case '_':
				currentState = 97
			case 'N':
				currentState = 99
			case 'R':
				currentState = 103
			case 'c':
				currentState = 69
			case 'i':
				currentState = 80
			case 'o':
				currentState = 58
			case 'I':
				currentState = 62
			case 'L':
				currentState = 61
			case 'W':
				currentState = 56
			case 'n':
				currentState = 85
			case 'k':
				currentState = 57
			case 'B':
				currentState = 88
			case 'T':
				currentState = 74
			case 'J':
				currentState = 59
			case 'X':
				currentState = 68
			case 'D':
				currentState = 94
			case 'f':
				currentState = 102
			case 't':
				currentState = 76
			case 'M':
				currentState = 90
			case 'V':
				currentState = 92
			case 'v':
				currentState = 96
			case 'x':
				currentState = 93
			case 'y':
				currentState = 72
			case 'O':
				currentState = 73
			case 'Q':
				currentState = 91
			case 'e':
				currentState = 105
			case 'g':
				currentState = 82
			case 'l':
				currentState = 106
			case 'A':
				currentState = 77
			case 'E':
				currentState = 95
			case 'G':
				currentState = 100
			case 'H':
				currentState = 98
			case 'a':
				currentState = 104
			case 'u':
				currentState = 55
			case 'b':
				currentState = 60
			case 'p':
				currentState = 70
			case 'q':
				currentState = 81
			case 'r':
				currentState = 83
			case 'Z':
				currentState = 101
			case 'z':
				currentState = 64
			case 'h':
				currentState = 79
			case 's':
				currentState = 87
			case 'F':
				currentState = 89
			case 'K':
				currentState = 63
			case 'S':
				currentState = 78
			case 'Y':
				currentState = 75
			case 'd':
				currentState = 66
			case 'C':
				currentState = 86
			case 'P':
				currentState = 65
			case 'U':
				currentState = 67
			case 'm':
				currentState = 54
			case 'w':
				currentState = 71
			default:
				break outer
			}

		case 86:
			switch c {
			case 'I':
				currentState = 62
			case 'Q':
				currentState = 91
			case 'f':
				currentState = 102
			case 'Y':
				currentState = 75
			case '_':
				currentState = 97
			case 'S':
				currentState = 78
			case 'X':
				currentState = 68
			case 'j':
				currentState = 84
			case 'p':
				currentState = 70
			case 'M':
				currentState = 90
			case 'U':
				currentState = 67
			case 'G':
				currentState = 100
			case 'P':
				currentState = 65
			case 'd':
				currentState = 66
			case 'D':
				currentState = 94
			case 'H':
				currentState = 98
			case 't':
				currentState = 76
			case 'v':
				currentState = 96
			case 'z':
				currentState = 64
			case 'K':
				currentState = 63
			case 'O':
				currentState = 73
			case 'm':
				currentState = 54
			case 's':
				currentState = 87
			case 'q':
				currentState = 81
			case 'A':
				currentState = 77
			case 'F':
				currentState = 89
			case 'N':
				currentState = 99
			case 'W':
				currentState = 56
			case 'B':
				currentState = 88
			case 'h':
				currentState = 79
			case 'C':
				currentState = 86
			case 'b':
				currentState = 60
			case 'Z':
				currentState = 101
			case 'c':
				currentState = 69
			case 'y':
				currentState = 72
			case 'L':
				currentState = 61
			case 'e':
				currentState = 105
			case 'T':
				currentState = 74
			case 'a':
				currentState = 104
			case 'g':
				currentState = 82
			case 'E':
				currentState = 95
			case 'J':
				currentState = 59
			case 'n':
				currentState = 85
			case 'o':
				currentState = 58
			case 'w':
				currentState = 71
			case 'V':
				currentState = 92
			case 'i':
				currentState = 80
			case 'l':
				currentState = 106
			case 'x':
				currentState = 93
			case 'R':
				currentState = 103
			case 'k':
				currentState = 57
			case 'r':
				currentState = 83
			case 'u':
				currentState = 55
			default:
				break outer
			}

		case 29:
			switch c {
			case 'Q':
				currentState = 91
			case 'T':
				currentState = 74
			case 'w':
				currentState = 71
			case '_':
				currentState = 97
			case 'A':
				currentState = 77
			case 'P':
				currentState = 65
			case 'Z':
				currentState = 101
			case 'h':
				currentState = 79
			case 'l':
				currentState = 106
			case 'R':
				currentState = 103
			case 'U':
				currentState = 67
			case 'W':
				currentState = 56
			case 'd':
				currentState = 66
			case 'f':
				currentState = 102
			case 'n':
				currentState = 85
			case 'Y':
				currentState = 75
			case 'k':
				currentState = 57
			case 'B':
				currentState = 88
			case 'K':
				currentState = 63
			case 'o':
				currentState = 58
			case 'x':
				currentState = 93
			case 'L':
				currentState = 61
			case 'b':
				currentState = 60
			case 'N':
				currentState = 99
			case 'X':
				currentState = 68
			case 'j':
				currentState = 84
			case 'G':
				currentState = 100
			case 'H':
				currentState = 98
			case 'J':
				currentState = 59
			case 'E':
				currentState = 95
			case 'V':
				currentState = 92
			case 'm':
				currentState = 54
			case 'y':
				currentState = 72
			case 'M':
				currentState = 90
			case 'O':
				currentState = 73
			case 'S':
				currentState = 78
			case 'a':
				currentState = 104
			case 'C':
				currentState = 86
			case 'c':
				currentState = 69
			case 'e':
				currentState = 105
			case 'q':
				currentState = 81
			case 'r':
				currentState = 83
			case 's':
				currentState = 87
			case 'z':
				currentState = 64
			case 'p':
				currentState = 70
			case 't':
				currentState = 76
			case 'F':
				currentState = 89
			case 'v':
				currentState = 96
			case 'D':
				currentState = 94
			case 'I':
				currentState = 62
			case 'g':
				currentState = 82
			case 'u':
				currentState = 55
			case 'i':
				currentState = 80
			default:
				break outer
			}

		case 92:
			switch c {
			case 'a':
				currentState = 104
			case 'e':
				currentState = 105
			case 's':
				currentState = 87
			case 'L':
				currentState = 61
			case 'z':
				currentState = 64
			case 'V':
				currentState = 92
			case 'd':
				currentState = 66
			case 'n':
				currentState = 85
			case 'T':
				currentState = 74
			case 'P':
				currentState = 65
			case 'r':
				currentState = 83
			case 'E':
				currentState = 95
			case 'H':
				currentState = 98
			case 'K':
				currentState = 63
			case 'l':
				currentState = 106
			case 't':
				currentState = 76
			case '_':
				currentState = 97
			case 'D':
				currentState = 94
			case 'G':
				currentState = 100
			case 'B':
				currentState = 88
			case 'Z':
				currentState = 101
			case 'j':
				currentState = 84
			case 'v':
				currentState = 96
			case 'C':
				currentState = 86
			case 'c':
				currentState = 69
			case 'o':
				currentState = 58
			case 'U':
				currentState = 67
			case 'h':
				currentState = 79
			case 'F':
				currentState = 89
			case 'i':
				currentState = 80
			case 'p':
				currentState = 70
			case 'u':
				currentState = 55
			case 'Q':
				currentState = 91
			case 'b':
				currentState = 60
			case 'k':
				currentState = 57
			case 'N':
				currentState = 99
			case 'Y':
				currentState = 75
			case 'f':
				currentState = 102
			case 'w':
				currentState = 71
			case 'x':
				currentState = 93
			case 'W':
				currentState = 56
			case 'y':
				currentState = 72
			case 'A':
				currentState = 77
			case 'J':
				currentState = 59
			case 'O':
				currentState = 73
			case 'R':
				currentState = 103
			case 'g':
				currentState = 82
			case 'M':
				currentState = 90
			case 'S':
				currentState = 78
			case 'X':
				currentState = 68
			case 'm':
				currentState = 54
			case 'q':
				currentState = 81
			case 'I':
				currentState = 62
			default:
				break outer
			}

		case 96:
			switch c {
			case 'h':
				currentState = 79
			case 'Y':
				currentState = 75
			case 'V':
				currentState = 92
			case 'W':
				currentState = 56
			case 'e':
				currentState = 105
			case 'k':
				currentState = 57
			case 'o':
				currentState = 58
			case 's':
				currentState = 87
			case 'A':
				currentState = 77
			case '_':
				currentState = 97
			case 'S':
				currentState = 78
			case 'z':
				currentState = 64
			case 'R':
				currentState = 103
			case 'd':
				currentState = 66
			case 't':
				currentState = 76
			case 'y':
				currentState = 72
			case 'Q':
				currentState = 91
			case 'b':
				currentState = 60
			case 'c':
				currentState = 69
			case 'P':
				currentState = 65
			case 'x':
				currentState = 93
			case 'Z':
				currentState = 101
			case 'K':
				currentState = 63
			case 'a':
				currentState = 104
			case 'j':
				currentState = 84
			case 'B':
				currentState = 88
			case 'J':
				currentState = 59
			case 'X':
				currentState = 68
			case 'm':
				currentState = 54
			case 'H':
				currentState = 98
			case 'G':
				currentState = 100
			case 'L':
				currentState = 61
			case 'O':
				currentState = 73
			case 'g':
				currentState = 82
			case 'D':
				currentState = 94
			case 'C':
				currentState = 86
			case 'M':
				currentState = 90
			case 'U':
				currentState = 67
			case 'f':
				currentState = 102
			case 'l':
				currentState = 106
			case 'q':
				currentState = 81
			case 'u':
				currentState = 55
			case 'F':
				currentState = 89
			case 'n':
				currentState = 85
			case 'w':
				currentState = 71
			case 'N':
				currentState = 99
			case 'I':
				currentState = 62
			case 'i':
				currentState = 80
			case 'p':
				currentState = 70
			case 'E':
				currentState = 95
			case 'r':
				currentState = 83
			case 'v':
				currentState = 96
			case 'T':
				currentState = 74
			default:
				break outer
			}

		case 105:
			switch c {
			case 'J':
				currentState = 59
			case 'M':
				currentState = 90
			case 'O':
				currentState = 73
			case 'k':
				currentState = 57
			case 'l':
				currentState = 106
			case '_':
				currentState = 97
			case 'H':
				currentState = 98
			case 'R':
				currentState = 103
			case 'h':
				currentState = 79
			case 'N':
				currentState = 99
			case 'v':
				currentState = 96
			case 'B':
				currentState = 88
			case 'L':
				currentState = 61
			case 'D':
				currentState = 94
			case 'S':
				currentState = 78
			case 'Z':
				currentState = 101
			case 'f':
				currentState = 102
			case 'j':
				currentState = 84
			case 's':
				currentState = 87
			case 't':
				currentState = 76
			case 'E':
				currentState = 95
			case 'P':
				currentState = 65
			case 'T':
				currentState = 74
			case 'i':
				currentState = 80
			case 'u':
				currentState = 55
			case 'w':
				currentState = 71
			case 'Y':
				currentState = 75
			case 'a':
				currentState = 104
			case 'q':
				currentState = 81
			case 'n':
				currentState = 85
			case 'z':
				currentState = 64
			case 'C':
				currentState = 86
			case 'K':
				currentState = 63
			case 'W':
				currentState = 56
			case 'd':
				currentState = 66
			case 'r':
				currentState = 83
			case 'x':
				currentState = 93
			case 'U':
				currentState = 67
			case 'c':
				currentState = 69
			case 'e':
				currentState = 105
			case 'I':
				currentState = 62
			case 'X':
				currentState = 68
			case 'g':
				currentState = 82
			case 'o':
				currentState = 58
			case 'F':
				currentState = 89
			case 'b':
				currentState = 60
			case 'm':
				currentState = 54
			case 'p':
				currentState = 70
			case 'Q':
				currentState = 91
			case 'y':
				currentState = 72
			case 'A':
				currentState = 77
			case 'G':
				currentState = 100
			case 'V':
				currentState = 92
			default:
				break outer
			}

		case 14:
			switch c {
			case 'M':
				currentState = 90
			case 'p':
				currentState = 70
			case '_':
				currentState = 97
			case 'x':
				currentState = 93
			case 'y':
				currentState = 72
			case 'S':
				currentState = 78
			case 'U':
				currentState = 67
			case 'c':
				currentState = 69
			case 'g':
				currentState = 82
			case 'h':
				currentState = 79
			case 'I':
				currentState = 62
			case 'T':
				currentState = 74
			case 'Z':
				currentState = 101
			case 'R':
				currentState = 103
			case 'l':
				currentState = 106
			case 'e':
				currentState = 105
			case 'i':
				currentState = 80
			case 'Y':
				currentState = 75
			case 'z':
				currentState = 64
			case 'a':
				currentState = 104
			case 't':
				currentState = 76
			case 'w':
				currentState = 71
			case 'E':
				currentState = 95
			case 'K':
				currentState = 63
			case 'L':
				currentState = 61
			case 'O':
				currentState = 73
			case 'o':
				currentState = 58
			case 'B':
				currentState = 88
			case 'Q':
				currentState = 91
			case 'H':
				currentState = 98
			case 'V':
				currentState = 92
			case 'm':
				currentState = 54
			case 'n':
				currentState = 85
			case 'q':
				currentState = 81
			case 'C':
				currentState = 86
			case 'N':
				currentState = 99
			case 'P':
				currentState = 65
			case 'j':
				currentState = 84
			case 'k':
				currentState = 57
			case 's':
				currentState = 87
			case 'A':
				currentState = 77
			case 'b':
				currentState = 60
			case 'd':
				currentState = 66
			case 'D':
				currentState = 94
			case 'W':
				currentState = 56
			case 'G':
				currentState = 100
			case 'f':
				currentState = 102
			case 'r':
				currentState = 83
			case 'u':
				currentState = 55
			case 'F':
				currentState = 89
			case 'J':
				currentState = 59
			case 'X':
				currentState = 68
			case 'v':
				currentState = 96
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

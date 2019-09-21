package reast

import "testing"

func TestString(t *testing.T) {
	ast := NewChoose(NewChoose(
		NewRepeat(NewSequence(NewPrimitive('c', 'c'),
			NewSequence(NewPrimitive('a', 'a'),
				NewPrimitive('t', 't')))),
		NewSequence(
			NewRepeat(NewPrimitive('a', 'a')),
			NewPrimitive('b', 'b'))),
		NewPrimitive('c', 'c'))

	expected := "(((cat)*|(a)*b)|c)"
	if ast.REString() != expected {
		t.Errorf("reast.REString() wrong. expected: %q, got=%q", expected, ast.REString())
	}
}

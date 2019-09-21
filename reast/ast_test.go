package reast

import "testing"

func TestString(t *testing.T) {
	ast := NewChoose(NewChoose(
		NewRepeat(NewSequence(NewPrimitive('c', 0),
			NewSequence(NewPrimitive('a', 0),
				NewPrimitive('t', 0)))),
		NewSequence(
			NewRepeat(NewPrimitive('a', 0)),
			NewPrimitive('b', 0))),
		NewPrimitive('c', 0))

	expected := "(((cat)*|(a)*b)|c)"
	if ast.REString() != expected {
		t.Errorf("reast.REString() wrong. expected: %q, got=%q", expected, ast.REString())
	}
}

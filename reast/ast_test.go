package reast

import "testing"

func TestString(t *testing.T) {
	ast := Choose{[]RegEx{
		&Repeat{&Sequence{[]RegEx{
			&Primitive{'c'},
			&Primitive{'a'},
			&Primitive{'t'},
		}}},
		&Sequence{[]RegEx{
			&Repeat{&Primitive{'a'}},
			&Primitive{'b'},
		}},
		&Primitive{'c'},
	}}

	expected := "((cat)*|(a)*b|c)"
	if ast.REString() != expected {
		t.Errorf("reast.REString() wrong. expected: %q, got=%q", expected, ast.REString())
	}
}

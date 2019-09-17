package ast

import "testing"

func TestToRex(t *testing.T) {
	ast := Choose{[]Node{
		&Repeat{&Concat{[]Node{
			&Literal{'c'},
			&Literal{'a'},
			&Literal{'t'},
		}}},
		&Concat{[]Node{
			&Repeat{&Literal{'a'}},
			&Literal{'b'},
		}},
		&Literal{'c'},
	}}

	expected := "((cat)*|(a)*b|c)"
	if ast.ToRegex() != expected {
		t.Errorf("ast.ToRegex() wrong. expected: %q, got=%q", expected, ast.ToRegex())
	}
}

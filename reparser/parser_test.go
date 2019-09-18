package reparser

import "testing"

func TestNormal(t *testing.T) {
	r := Parse("abc*((cat)*|dog)")
	expected := "ab(c)*((cat)*|dog)"
	if r.REString() != expected {
		t.Fatalf("reparser.Parse() expected %q. got %q", expected, r.REString())
	}
}

func TestCharset(t *testing.T) {
	r := Parse("[a-c1-4_]")
	expected := "(a|b|c|1|2|3|4|_)"
	if r.REString() != expected {
		t.Fatalf("reparser.Parse() expected %q. got %q", expected, r.REString())
	}
}

func TestEscCharset(t *testing.T) {
	r := Parse("[a-c\\]1-4_\\-]")
	expected := "(a|b|c|]|1|2|3|4|_|-)"
	if r.REString() != expected {
		t.Fatalf("reparser.Parse() expected %q. got %q", expected, r.REString())
	}
}
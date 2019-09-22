package reparser

import "testing"

func TestNormal(t *testing.T) {
	r, err := Parse("abc*((cat)*|dog)")
	if err != nil {
		t.Fatal(err)
	}
	expected := "ab(c)*((cat)*|dog)"
	if r.REString() != expected {
		t.Fatalf("reparser.Parse() expected %q. got %q", expected, r.REString())
	}
}

func TestCharset(t *testing.T) {
	r, err := Parse("[a-c1-4_]")
	if err != nil {
		t.Fatalf("reparser.Parse() failed: %s", err.Error())
	}
	expected := "((a-c|1-4)|_)"
	if r.REString() != expected {
		t.Fatalf("reparser.Parse() expected %q. got %q", expected, r.REString())
	}
}

func TestEscCharset(t *testing.T) {
	r, err := Parse("[a-c\\]1-4_\\-]")
	if err != nil {
		t.Fatalf("reparser.Parse() failed: %s", err.Error())
	}
	expected := "((((a-c|])|1-4)|_)|-)"
	if r.REString() != expected {
		t.Fatalf("reparser.Parse() expected %q. got %q", expected, r.REString())
	}
}

func TestExcludeCharset(t *testing.T) {
	r, err := Parse("[^a-c1-4_]")
	if err != nil {
		t.Fatalf("reparser.Parse() failed: %s", err.Error())
	}
	expected := "(((\x01-0|5-^)|`)|d-\U0010ffff)"
	if r.REString() != expected {
		t.Fatalf("reparser.Parse() expected %q. got %q", expected, r.REString())
	}
}


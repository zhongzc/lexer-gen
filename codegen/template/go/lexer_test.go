package _go

import (
	"testing"
)

func TestLexer(t *testing.T) {
	cs := NewStream(`
abdog
abcatcat
abccc
`)
	l := NewLexer(cs)
	for l.HasNext() {
		tk, err := l.NextToken()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(tk)
	}
}

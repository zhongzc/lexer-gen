package _tmpLexer

import (
	"testing"
)

func TestLexer(t *testing.T) {
	cs := NewStream(`
if (a_for_apple == 10000) {
	var b_for_ball = 10086;
	return b_for_banana;
} else {
	return 0;
}
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

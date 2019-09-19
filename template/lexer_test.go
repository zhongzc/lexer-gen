package template

import (
	"testing"
)

func TestLexer(t *testing.T) {
	cs := NewStream("baaab")
	am := &A{}
	err := am.RunGreedy(cs)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

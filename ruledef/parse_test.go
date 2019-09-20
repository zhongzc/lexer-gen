package ruledef

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	temp := `
IF if
CONST const
LP (
RP )
IDENT [A-Za-z_]
`
	res := Parse(strings.NewReader(temp))
	if len(res) != 5 {
		t.Errorf("Parse() length expected %d. got %d", 5, len(res))
	}
}

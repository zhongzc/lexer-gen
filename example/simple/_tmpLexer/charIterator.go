// generated by goGen.go - do not edit
package lexer

// character reader
type CharIterator interface {
	CurrentIndex() int
	NextChar() rune
	Peek() rune
	SetIndex(i int)
	SubString(l int, r int) []rune
	Loc(idx int) struct {
		Line int
		Col  int
	}
}

// Character reader implement
type CharStream struct {
	curIdx int
	runes  []rune
}

func NewStream(input string) *CharStream {
	return &CharStream{0, []rune(input)}
}

func (cs *CharStream) CurrentIndex() int {
	return cs.curIdx
}

func (cs *CharStream) NextChar() rune {
	if cs.curIdx >= len(cs.runes) {
		return 0
	}
	r := cs.runes[cs.curIdx]
	cs.curIdx++
	return r
}

func (cs *CharStream) Peek() rune {
	if cs.curIdx >= len(cs.runes) {
		return 0
	}
	return cs.runes[cs.curIdx]
}

func (cs *CharStream) SetIndex(i int) {
	cs.curIdx = i
}

func (cs *CharStream) SubString(from int, limit int) []rune {
	return cs.runes[from:limit]
}

func (cs *CharStream) Loc(idx int) struct {
	Line int
	Col  int
} {
	line := 1
	col := 1
	for i := 0; i <= idx; i++ {
		if cs.runes[i] != '\n' {
			col++
		} else {
			line++
			col = 0
		}
	}
	return struct {
		Line int
		Col  int
	}{Line: line, Col: col}
}
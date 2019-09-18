package fa

type Rule struct {
	From int
	By   rune
	To   int
}

func (r *Rule) Apply(from int, by rune) (next int, ok bool) {
	if r.From == from && r.By == by {
		return r.To, true
	} else {
		return -1, false
	}
}

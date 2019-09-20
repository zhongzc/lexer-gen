package fa

type StateSet map[int]bool

func NewSet(ints ...int) StateSet {
	var ss StateSet = make(map[int]bool)
	for _, n := range ints {
		ss[n] = true
	}
	return ss
}

func (ss StateSet) Equal(o StateSet) bool {
	if len(ss) != len(o) {
		return false
	}
	for k := range ss {
		if !o[k] {
			return false
		}
	}
	return true
}

func (ss StateSet) Add(o StateSet) {
	for k := range o {
		ss[k] = true
	}
}

func (ss StateSet) LE(o StateSet) bool {
	if len(ss) > len(o) {
		return false
	}
	for k := range ss {
		if !o[k] {
			return false
		}
	}
	return true
}

func (ss StateSet) Copy() (res StateSet) {
	res = make(map[int]bool)
	for k := range ss {
		res[k] = true
	}
	return
}

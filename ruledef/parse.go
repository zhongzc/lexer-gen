package ruledef

import (
	"bufio"
	"bytes"
	"io"
)

type NamedRE struct {
	Name string
	RE   string
}

func Parse(reader io.Reader) (res []*NamedRE) {
	rd := bufio.NewReader(reader)
	var err error
	var buf []byte
	for err == nil {
		buf, _, err = rd.ReadLine()
		buf = bytes.Trim(buf, " \n\t")
		if len(buf) == 0 {
			continue
		}

		bs := bytes.SplitN(buf, []byte{' '}, 2)
		res = append(res, &NamedRE{string(bs[0]), string(bs[1])})
	}
	return
}

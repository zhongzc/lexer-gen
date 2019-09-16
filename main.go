package main

import (
	"fmt"
	. "lexerGenerator/rgex"
)

func main() {
	l := Literal{Rune: '好'}
	c := Concat{Grounds: []Ground{l,
		CharacterSet{Ranges: []Range{{From: 'A', To: 'Z'}}},
		Repeat{Choose: Choose{Concats: []Concat{{Grounds: []Ground{l, CharacterSet{Ranges: []Range{{From: 'a', To: 'z'}}}}}}}}}}
	ch := Choose{Concats: []Concat{c, c}}
	fmt.Println(ch.ToRegex())
}

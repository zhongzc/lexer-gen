package main

import (
	"fmt"
	"github.com/zhongzc/lexerGen/example/scutlab1/sample/lexer"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	test1 := lexer.NewStream(string(bs))

	l := lexer.NewLexer(test1)
	idx := 1
	rmb := make(map[string]int)

	cnt := 0
	for l.HasNext() {
		tk, err := l.NextToken()
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if tk.Type == "t42" {
			continue
		}

		if tk.Type == "t36" || tk.Type == "t37" || tk.Type == "t38" {
			if tk.Type == "t38" {
				tk.Token = tk.Token[:len(tk.Token)-1]
			}
			if i, ok := rmb[string(tk.Token)]; ok {
				fmt.Printf("(%d, %d)\t", 36, i)
			} else {
				rmb[string(tk.Token)] = idx
				fmt.Printf("(%d, %d)\t", 36, idx)
				idx++
			}
		} else {
			i, _ := strconv.ParseInt(string(tk.Type)[1:], 10, 32)
			fmt.Printf("(%d, -)\t", i)
		}

		if cnt == 4 {
			fmt.Print("\n")
			cnt = 0
		} else {
			cnt++
		}

	}
}

//func (l *Lexer) NextToken() (t *Token, err error) {
//	if l.Chars.Peek() == 0 {
//		return nil, errors.New("reach EOF")
//	}
//	l.skipWhitespace()
//	idx := l.Chars.CurrentIndex()
//
//	for _, a := range l.automatas {
//		err = a.Automata.RunGreedy(l.Chars)
//		if err != nil {
//			if a.Name == "t42" && l.Chars.CurrentIndex()-idx > 1 {
//				loc := l.Chars.Loc(idx)
//				msg := fmt.Sprintf("\nincomplete comment. at line %d, col %d\n", loc.Line, loc.Col)
//				return nil, errors.New(msg)
//			} else if a.Name == "t38" && idx != l.Chars.CurrentIndex() {
//				loc := l.Chars.Loc(idx)
//				msg := fmt.Sprintf("\nincomplete characters. at line %d, col %d\n", loc.Line, loc.Col)
//				return nil, errors.New(msg)
//			}
//			l.Chars.SetIndex(idx)
//		} else {
//			if a.Name == "t37" && syntax.IsWordChar(l.Chars.Peek()) {
//				loc := l.Chars.Loc(l.Chars.CurrentIndex())
//				msg := fmt.Sprintf("\nwrong delimiter: %c. at line %d, col %d\n", l.Chars.Peek(),
//					loc.Line, loc.Col)
//				return nil, errors.New(msg)
//			}
//
//			var tv = l.Chars.SubString(idx, l.Chars.CurrentIndex())
//			idx = l.Chars.CurrentIndex()
//			t = &Token{TokenType(a.Name), TokenValue(tv)}
//			err = nil
//
//			l.skipWhitespace()
//			return
//		}
//	}
//
//	loc := l.Chars.Loc(idx)
//	msg := fmt.Sprintf("\nunexpected char: %c. at line %d, col %d\n", l.Chars.Peek(), loc.Line, loc.Col)
//	return nil, errors.New(msg)
//}

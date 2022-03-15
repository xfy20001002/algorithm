package main

//从stdin中读取某一行作为输入,并重新输出

import (
	"fmt"
)

func main() {
	var flag bool = true
	for flag {
		var s string
		n := ScanLine(&s)
		//fmt.Println(n)
		if n == 0 {
			break
		} else {
			fmt.Println(s)
		}
	}
}

func ScanLine(s *string) int {
	var c byte
	var n int
	for {
		_, err := fmt.Scanf("%c", &c)
		if nil != err {
			break
		}
		if c == '\n' {
			break
		} else {
			*s = *s + string(c)
			n++
		}
	}
	return n
}

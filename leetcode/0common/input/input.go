package input

import "fmt"

func ScanLine(s *string) int {
	var c byte
	var n int
	for {
		_, err := fmt.Scanf("%c", &c)
		if c == '\n' {
			break
		}
		if err != nil {
			break
		} else {
			n++
			*s = *s + string(c)
		}
		//fmt.Println("****", *s)
	}
	return n
}

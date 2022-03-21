package main

import (
	"fmt"
)

func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	//n为字符串长度
	src := "ABCDEFGHIJKXYZAB"
	target := "ABX"
	fmt.Println(findMinSubString(src, target))
}

func findMinSubString(src string, target string) string {
	var start int = 0
	var length int = 0
	var minStart int = 0
	var minLength int = len(src)
	var p1 int = 0                //前指针
	var p2 int = 0                //后指针
	records := make(map[byte]int) //record记录byte出现的次数
	//初始化
	if inTarget(target, src[0]) {
		records[src[0]] += 1
	}
	for p2 < len(src) {
		if isTrue(target, records) {
			start = p1
			length = p2 - p1 + 1
			if length < minLength {
				minLength = length
				minStart = start
			}
			p1++
			records[src[p1]] -= 1

		} else {
			if inTarget(target, src[p2]) {
				records[src[p2]] += 1
			}
			p2++
		}
	}
	return src[minStart : minStart+minLength]
}

func inTarget(target string, c byte) bool {
	for i := 0; i < len(target); i++ {
		if c == target[i] {
			return true
		}
	}
	return false
}

func isTrue(target string, m map[byte]int) bool {
	//判断是否所有target字符都在m中
	res := true
	for i := 0; i < len(target); i++ {
		if m[target[i]] < 1 {
			res = false
		}
	}
	return res
}

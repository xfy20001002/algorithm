package main

import "fmt"

func longestPalindrome(s string) string {
	//创建二维数组并初始化
	//babad
	n := len(s)
	flags := make([][]bool, n)
	for i := range flags {
		flags[i] = make([]bool, n)
		flags[i][i] = true
	}
	//babad
	//b a b a d
	//
	//逆序遍历
	var maxRes = 1
	var start = 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if flags[i][j] {
				if j != i {
					if i-1 >= 0 && j+1 < n && s[i-1] == s[j+1] {
						flags[i-1][j+1] = true
						if j-i+3 > maxRes {
							maxRes = j - i + 3
							start = i - 1
						}
					}
				}
			}
			if j == i {
				if i-1 >= 0 && j+1 < n && s[i-1] == s[j+1] {
					flags[i-1][j+1] = true
					if j-i+3 > maxRes {
						maxRes = j - i + 3
						start = i - 1
					}
				} else if j+1 < n && s[i] == s[j+1] {
					if j-i+2 > maxRes {
						maxRes = j - i + 2
						start = i
					}
				} else if i-1 >= 0 && s[j] == s[i-1] {
					if j-i+2 > maxRes {
						maxRes = j - i + 2
						start = i - 1
					}
				}
			}
		}
	}
	return s[start : start+maxRes]

}

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

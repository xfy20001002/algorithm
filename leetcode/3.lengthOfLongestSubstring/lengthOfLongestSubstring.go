package main

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstring1(s string) int {
	//solution1
	//傻瓜解法
	//思路：用1个数组记录以s[i]为首部的不重复字符串的最大长度
	//每次循环时，用1个哈希表存储之前遍历过的字符，如果遍历过的字符在哈希表中，则停止循环，并将哈希表中的值清空
	//时间复杂度：O(n^2)
	//空间复杂度：O(n)
	if len(s) == 0 {
		return 0
	}
	records := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		h := make(map[byte]bool)
		var j int
		for j = i; j < len(s); j++ {
			if _, ok := h[s[j]]; ok {
				records[i] = j - i
				break
			} else {
				h[s[j]] = true
			}
		}
		if j == len(s) {
			records[i] = j - i
		}
	}
	var res int
	for _, v := range records {
		res = max(res, v)
	}
	return res
}

func lengthOfLongestSubstring2(s string) int {
	//solution2
	//首字母且利用之前结果的算法过于复杂，边界条件较多，代码难写
	//abcac
	/*
		  0 1 2 3 4
		0 1 2 3 3 3
		1   1 2 3 3
		2     1 2 2
		3       1 2
		4         1
	*/
	//思路:滑动串口
	//利用l,r两个变量维护一个滑动窗口
	//利用一个哈希表维护滑动窗口中出现的变量
	//利用一个数组记录以i为起始字符的最大不重复子数组长度
	if len(s) == 0 {
		return 0
	}
	var l, r, res int
	l, r, res = 0, 0, 0
	existed := map[byte]int{}
	for r < len(s) {
		if v, ok := existed[s[r]]; ok {
			for l < v+1 {
				if l == v {
					existed[s[v]] = r
				} else {
					res = max(r-l, res)
					delete(existed, s[l])
				}
				l = l + 1
			}
		} else {
			existed[s[r]] = r
		}
		r = r + 1
	}
	res = max(res, r-l)
	return res
}

/*
func main() {
	fmt.Println(lengthOfLongestSubstring1("abcabcbb"))
	//res := lengthOfLongestSubstring2("abcac")
	res := lengthOfLongestSubstring2("abc")
	fmt.Println(res)
}
*/

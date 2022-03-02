package leetcode

//4种方法
//1.暴力法
//2.动态规划
//3.栈
//4.正逆序遍历

func LongestValidParentheses1(s string) int {

	//正逆序遍历
	left := 0
	right := 0
	length := 0
	maxLength := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		balance := left - right
		if balance < 0 {
			left = 0
			right = 0
		} else if balance == 0 {
			length = right * 2
			maxLength = max(maxLength, length)
		}
	}
	//逆序遍历
	left = 0
	right = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		balance := right - left
		if balance < 0 {
			left = 0
			right = 0
		} else if balance == 0 {
			length = right * 2
			maxLength = max(maxLength, length)
		}
	}
	return maxLength
}

func longestValidParentheses2(s string) int {
	//2.动态规划

	n := len(s)
	//dp[i]表示以i结尾的最长有效括号
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else {
			if i == 0 || i-dp[i-1]-1 < 0 {
				//")"
				//"())"
				dp[i] = 0
				continue
			}
			if s[i-dp[i-1]-1] == '(' {
				a := 0
				if i-dp[i-1]-2 > 0 {
					a = dp[i-dp[i-1]-2]
				}
				//合并前面的有效括号
				dp[i] = dp[i-1] + 2 + a
			}
		}
	}
	return max(dp)
}

func max(arr []int) int {
	res := 0
	for _, v := range arr {
		if res < v {
			res = v
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func longestValidParentheses3(s string) int {
	//3.栈
	stack := []int{-1}
	l := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				l = i - stack[len(stack)-1]
				maxLen = max(maxLen, l)
			}
		}
	}
	return maxLen
}

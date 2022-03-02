package leetcode

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

var res []string

func generateParenthesis(n int) []string {
	res = []string{}
	dfs("", 0, n)
	return res
}

func dfs(path string, pos, n int) {
	if len(path) == 2*n {
		if isValid(path) {
			res = append(res, path)
		}
		return
	}
	dfs(path+"(", pos+1, n)
	dfs(path+")", pos+1, n)
}

func isValid(path string) bool {
	left, right := 0, 0
	for i := 0; i < len(path); i++ {
		if path[i] == '(' {
			left = left + 1
		} else if path[i] == ')' {
			right = right + 1
		}
		if left < right {
			return false
		}
	}
	return left == right
}

package leetcode

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]


*/

func LetterCombinations(digits string) []string {
	maps := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	if digits == "" {
		return []string{}
	} else if len(digits) == 1 {
		return split(maps[digits[0]-'0'-2])
	}
	key := digits[0] - '0' - 2
	res := []string{}
	for i := 0; i < len(maps[key]); i++ {
		var subres []string
		subres = LetterCombinations(digits[1:])
		for j := 0; j < len(subres); j++ {
			s := string(maps[key][i]) + subres[j]
			res = append(res, s)
		}
	}
	return res
}
func split(s string) []string {
	var res []string
	for _, c := range s {
		res = append(res, string(c))
	}
	return res
}

package leetcode

func IsMatch(s string, p string) bool {
	n1 := len(s)
	n2 := len(p)
	dp := make([][]bool, n1+1)
	for i, _ := range dp {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true
	for j := 2; j < n2+1; j++ {
		dp[0][j] = (dp[0][j-2]) && (p[j-1] == '*')
	}
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if p[j] != '.' && p[j] != '*' {
				dp[i+1][j+1] = dp[i][j] && (p[j] == s[i])
			} else if p[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = dp[i+1][j-1] || dp[i+1][j] || ((p[j-1] == s[i] || p[j-1] == '.') && dp[i][j+1])
			}
		}
	}
	return dp[n1][n2]
}

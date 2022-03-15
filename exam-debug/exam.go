package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	//找到0的个数减去1的个数值为最大的区间
	//使得0的权重为1,1的权重为-1，遍历所有区间
	dp := make([][]int, n) //dp[i][j]表示数组nums[i,j]中0减一个数的值
	maxCount := 0
	var start, end int
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		if nums[i] == 0 {
			dp[i][i] = 1
			if maxCount < dp[i][i] {
				start = i
				end = i
				maxCount = dp[i][i]
			}
		} else {
			dp[i][i] = -1
		}
		for j := i + 1; j < n; j++ {
			if nums[j] == 0 {
				dp[i][j] = dp[i][j-1] + 1
				if maxCount < dp[i][j] {
					start = i
					end = j
					maxCount = dp[i][j]
				}
			} else {
				dp[i][j] = dp[i][j-1] - 1
			}
		}
	}
	var res = 0
	fmt.Println(start, end, maxCount)
	for i := 0; i < start; i++ {
		if nums[i] == 1 {
			res = res + 1
		}
	}
	for i := end + 1; i < n; i++ {
		if nums[i] == 1 {
			res = res + 1
		}
	}
	res = res + maxCount
	fmt.Println(res)
}

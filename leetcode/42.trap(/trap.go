package main

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
func trap1(height []int) int {
	//暴力解法
	//当前位置能装的雨水等于左右两端最大值高度的较小值减去当前位置的高度
	var ans int = 0
	for i := 1; i < len(height)-1; i++ {
		left, right := 0, 0
		for j := 0; j <= i; j++ {
			left = max(height[j], left)
		}
		for j := i; j < len(height); j++ {
			right = max(height[j], right)
		}
		ans = ans + min(left, right) - height[i]
	}
	return ans
}

func trap2(height []int) int {
	//利用left数组和right数组事先存储好左右两边的最大值
	//left[i]表示到0...i时数组中的最大值
	//right[i]表示i...len(heignt)-1时数组中的最大值
	n := len(height)
	lefts := make([]int, n)
	rights := make([]int, n)

	//填充left数组
	maxLeft := height[0]
	for i := 0; i < n; i++ {
		maxLeft = max(maxLeft, height[i])
		lefts[i] = maxLeft
	}
	//填充right数组
	maxRight := height[n-1]
	for i := n - 1; i >= 0; i-- {
		maxRight = max(maxRight, height[i])
		rights[i] = maxRight
	}
	//fmt.Println(rights)
	//fmt.Println(lefts)
	var ans int = 0
	//更新ans
	for i := 1; i < n-1; i++ {
		ans = ans + min(lefts[i], rights[i]) - height[i]
	}
	return ans
}

func trap3(height []int) int {

}
func trap4(height []int) int {

}

func main() {

}

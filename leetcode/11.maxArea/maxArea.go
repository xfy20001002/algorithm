package leetcode

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。
*/
func MaxArea(height []int) int {
	var l = 0
	var r = len(height) - 1
	res := area(height, l, r)
	for l < r {
		if height[l] < height[r] {
			l = l + 1
		} else {
			r = r - 1
		}
		res = max(res, area(height, l, r))
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
func area(height []int, l, r int) int {
	return (r - l) * min(height[l], height[r])
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

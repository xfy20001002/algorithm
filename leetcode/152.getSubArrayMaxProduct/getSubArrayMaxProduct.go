package main

/*
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个 32-位 整数。

子数组 是数组的连续子序列。



示例 1:

输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-product-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/

//得到子数组的最大连续乘积
func GetSubArrayMaxProduct(nums []int64) int64 {
	//2  3  -5   6     -1
	//2  6  -5   6     180  以i为结尾的子数组的乘积最大值
	//2  3  -30 -180   -1   以i为结尾的子数组的乘积最小值
	if len(nums) == 0 {
		return 0
	}
	maxProduct := nums[0]
	minProduct := nums[0]
	maxProductTemp := maxProduct
	minProductTemp := minProduct
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			//maxProductTemp, minProductTemp = minProductTemp, maxProductTemp
			maxProductTemp = max(minProductTemp*nums[i], nums[i])
			minProductTemp = min(maxProductTemp*nums[i], nums[i])
		} else {
			maxProductTemp = max(maxProductTemp*nums[i], nums[i])
			minProductTemp = min(minProductTemp*nums[i], nums[i])
		}
		maxProduct = max(maxProduct, maxProductTemp)
		minProduct = min(minProduct, minProductTemp)
		//maxProductTemp = max(nums[i], maxProductTemp*nums[i])
		//minProductTemp = min(nums[i], minProductTemp*nums[i])
		//maxProduct = max(maxProduct, maxProductTemp)
		//minProduct = min(minProduct, minProductTemp)
	}
	return maxProduct
}

func max(x, y int64) int64 {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int64) int64 {
	if x < y {
		return x
	} else {
		return y
	}
}

/*
func main() {

}
*/

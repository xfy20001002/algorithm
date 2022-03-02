package leetcode

import "sort"

/*
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。



示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func NextPermutation(nums []int) {
	n := len(nums)
	//思路 优先替换位置较后的数,替换后保证替换位置后的数为最小字典序(利用排序)
	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[j] > nums[i] {
				//fmt.Println(i,j)
				nums[i], nums[j] = nums[j], nums[i]
				//保证nums[i]...nums[n-1]为最小序 可以利用反转
				reverse(nums[i+1 : n])
				return
			}
		}
	}
	sort.Ints(nums)
}

func reverse(arr []int) {
	var i, j = 0, len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i = i + 1
		j = j - 1
	}
}

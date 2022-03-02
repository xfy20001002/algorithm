package leetcode

/*
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。



示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：

输入：nums = [1], target = 0
输出：-1


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func search(nums []int, target int) int {
	return searchHelper(nums, 0, len(nums)-1, target)
}

func searchHelper(nums []int, start, end, target int) int {
	//细节
	//旋转点在左 在右 正序 逆序
	//left right mid之间相重叠
	left := start
	right := end
	var mid = (left + right) / 2
	//fmt.Println(left,right,mid)
	if left > right {
		return -1
	}
	if nums[mid] == target {
		return mid
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	if nums[mid] > nums[left] && nums[mid] < nums[right] {
		//k为0的情况
		if target < nums[mid] {
			return searchHelper(nums, start, mid-1, target)
		} else {
			return searchHelper(nums, mid+1, end, target)
		}
	}
	if nums[left] < nums[mid] {
		//k在右边的情况
		if target > nums[mid] {
			//在mid右边找
			return searchHelper(nums, mid+1, end, target)
		} else {
			//可以在左边也能在右边
			if target > nums[right] {
				//在左边找
				return searchHelper(nums, start, mid-1, target)
			} else {
				//在右边找
				return searchHelper(nums, mid+1, end, target)
			}
		}
	} else {
		//k在左边的情况
		if target > nums[mid] {
			//在右边或左边找 可以直接用二分查找
			if target < nums[left] {
				//在右边找
				return searchHelper(nums, mid+1, end, target)
			} else {
				//在左边找
				return searchHelper(nums, start, mid-1, target)
			}
		} else {
			//在左边找
			return searchHelper(nums, start, mid-1, target)
		}
	}
	return -1
}

//不用递归解法

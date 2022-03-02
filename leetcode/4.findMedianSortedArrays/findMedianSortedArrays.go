package main

import (
	"math"
)

func Merge2Array(nums1 []int, nums2 []int) []int {
	var i1, i2, k int
	newArray := make([]int, len(nums1)+len(nums2))
	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] < nums2[i2] {
			newArray[k] = nums1[i1]
			i1 = i1 + 1
		} else {
			newArray[k] = nums2[i2]
			i2 = i2 + 1
		}
		k = k + 1
	}
	if i1 < len(nums1) {
		for i1 < len(nums1) {
			newArray[k] = nums1[i1]
			i1++
			k++
		}
	} else if i2 < len(nums2) {
		for i2 < len(nums2) {

			newArray[k] = nums2[i2]
			i2++
			k++
		}
	}
	return newArray
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//soluion1
	//合并两个数组，在新数组中求中位数
	newNums := Merge2Array(nums1, nums2)
	ls := len(nums1) + len(nums2)
	if ls == 0 {
		return float64(0)
	}
	var res float64
	mid := ls / 2
	if ls&1 == 1 {
		res = float64(newNums[mid])
	} else {
		res = (float64(newNums[mid]) + float64(newNums[mid-1])) / 2
	}
	return res
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	//2,3   4,5
	//2,3,6 4,5
	//遍历l/2+1次
	//维护l,r变量
	//循环结束后根据奇数偶数使用l,r得到想要的结果
	//注意点：需要考虑某一个数组完全超出的边界 如[1,2][3,4] 解决思路为添加初始哨兵，根据是否溢出边界来给哨兵变量赋值
	var n = 0
	var i, j = -1, -1
	var l, r = 0, 0
	m := len(nums1) + len(nums2)
	if m == 0 {
		return float64(0)
	}
	for ; n < m/2+1; n++ {
		v1, v2 := 1000001, 1000001
		if len(nums1) != i+1 {
			v1 = nums1[i+1]
		}
		if len(nums2) != j+1 {
			v2 = nums2[j+1]
		}
		if v1 < v2 {
			i = i + 1
			l = r
			r = v1
		} else {
			j = j + 1
			l = r
			r = v2
		}
	}
	var res float64
	if m&1 == 1 {
		res = float64(r)
	} else {
		res = (float64(l) + float64(r)) / 2
	}
	return res
}

func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	return float64(findKthNumber1(nums1, nums2, 0, len(nums1)-1, 0, len(nums2)-1, (l+1)/2)+findKthNumber1(nums1, nums2, 0, len(nums1)-1, 0, len(nums2)-1, (l+2)/2)) * 0.5
}

func findKthNumber1(nums1, nums2 []int, s1, e1, s2, e2, k int) int {
	//找到两个数组间第k小的树
	//思路
	//分别在nums1,nums2数组找到第k/2小的数,进行比较,小的则全部去除
	//同时需要进行边界分析
	//边界分析:
	//1.分情况讨论
	//2.维护一个变量使其始终为最小长度的数组
	if k == 1 {
		if s1 > e1 {
			return nums2[s2]
		} else if s2 > e2 {
			return nums1[s1]
		} else {
			return min(nums1[s1], nums2[s2])
		}
	}
	var n1, n2, res int
	if s1+k/2-1 > e1 {
		n2 = nums2[s2+k/2-1]
		if s1 <= e1 {
			n1 = nums1[e1]
		} else {
			n1 = math.MaxInt
		}
		if n1 < n2 {
			res = findKthNumber1(nums1, nums2, e1+1, e1, s2, e2, k-e1-1)
		} else {
			res = findKthNumber1(nums1, nums2, s1, e1, s2+k/2, e2, k-k/2)
		}
	} else if s2+k/2-1 > e2 {
		n1 = nums1[s1+k/2-1]
		if s2 <= e2 {
			n2 = nums2[e2]
		} else {
			n2 = math.MaxInt
		}
		if n1 < n2 {
			res = findKthNumber1(nums1, nums2, s1+k/2, e1, s2, e2, k-k/2)
		} else {
			res = findKthNumber1(nums1, nums2, s1, e1, e2+1, e2, k-e2-1)
		}
	} else {
		n1, n2 = nums1[s1+k/2-1], nums2[s2+k/2-1]
		if n1 < n2 {
			res = findKthNumber1(nums1, nums2, s1+k/2, e1, s2, e2, k-k/2)
		} else {
			res = findKthNumber1(nums1, nums2, s1, e1, s2+k/2, e2, k-k/2)
		}
	}
	return res
}
func findKthNumber2(nums1, nums2 []int, s1, e1, s2, e2, k int) int {
	len1 := e1 - s1 + 1
	len2 := e2 - s2 + 1
	if len1 > len2 {
		return findKthNumber2(nums2, nums1, s2, e2, s1, e1, k)
	}
	if len1 == 0 {
		return nums2[s2+k-1]
	}
	if k == 1 {
		return min(nums1[s1], nums2[s2])
	} else {
		var res int
		var i = min(s1+k/2-1, e1)
		var j = min(s2+k/2-1, e2)
		if nums1[i] < nums2[j] {
			res = findKthNumber2(nums1, nums2, i+1, e1, s2, e2, k-i-1+s1)
		} else {
			res = findKthNumber2(nums1, nums2, s1, e1, j+1, e2, k-j-1+s2)
		}
		return res
	}
}
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

/*
func findMedianSortedArrays4(nums1 []int, nums2 []int) float64 {

}
*/
/*
func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	findMedianSortedArrays2(nums1, nums2)
	//fmt.Println(math.MaxInt)
	//fmt.Println(findKthNumber2(nums1, nums2, 0, len(nums1)-1, 0, len(nums2)-1, 2))
	fmt.Println(findKthNumber2(nums1, nums2, 0, len(nums1)-1, 0, len(nums2)-1, 3))

}
*/

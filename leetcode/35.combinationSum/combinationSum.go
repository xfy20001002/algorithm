package main

import "fmt"

//组合总数
//找到所有和为target的切片

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	cur := make([]int, 0)
	combinationSumHelper(candidates, target, 0, 0, cur)

	return res
}

func combinationSumHelper(candidates []int, target, start, sum int, cur []int) {
	if sum == target {
		res = append(res, append([]int(nil), cur...)) //该写法是因为append的时候，会改变cur的值，所以需要拷贝一份
		//fmt.Println(cur)
		return
	}
	for i := start; i < len(candidates); i++ {
		if sum+candidates[i] <= target {
			cur = append(cur, candidates[i])
			combinationSumHelper(candidates, target, i, sum+candidates[i], cur)
			cur = cur[:len(cur)-1]
		}
	}
}

func main() {
	//test copy slice
	/*
		a := []int{1, 2, 3, 4, 5}
		b := append([]int(nil), a...)
		fmt.Println(b)
	*/
	combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(res)
}

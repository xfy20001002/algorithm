package main

import "fmt"

func Quick_sort(arr []int, start, end int) {
	if start >= end {
		return
	}
	demarcation := Partition(arr, start, end)
	Quick_sort(arr, start, demarcation-1)
	Quick_sort(arr, demarcation+1, end)
	//fmt.Println(partition(arr, start, end))
	//print(arr)
}

//3 1 6 5

func Partition(arr []int, start, end int) int {
	m := arr[start]
	for i := start + 1; i <= end; {
		if arr[i] > m {
			arr[i], arr[end] = arr[end], arr[i]
			end = end - 1
		} else {
			i = i + 1
		}
	}
	arr[start], arr[end] = arr[end], arr[start]
	return end
}

func Print(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}

/*
func main() {
	arr := []int{4, 1, 3, 5, 6, 4, 5}
	quick_sort(arr, 0, len(arr)-1)
	print(arr)
}
*/

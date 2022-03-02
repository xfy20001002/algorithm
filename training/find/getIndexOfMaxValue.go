package main

import "fmt"

func GetIndexOfMaxValue(arr []int) int {
	maxValue := arr[0]
	res := 0
	for i, v := range arr {
		if v > maxValue {
			maxValue = v
			res = i
		}
	}
	return res
}
func GetIndexOfMinValue(arr []int) int {
	minValue := arr[0]
	res := 0
	for i, v := range arr {
		//维护最小值及最小值索引
		if v < minValue {
			minValue = v
			res = i
		}
	}
	return res
}
func main() {
	var intArr = [...]int{3, -4, 93, 8, 12, 29}
	fmt.Println(GetIndexOfMaxValue(intArr[:]))
}

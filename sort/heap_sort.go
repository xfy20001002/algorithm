package main

import "fmt"

func Parent(i int) int {
	return i / 2
}
func Left(i int) int {
	return i * 2
}
func Right(i int) int {
	return i*2 + 1
}

func Max_heapify(arr []int, length int, i int) {
	left := Left(i)
	right := Right(i)
	var largest = i
	if left <= length && arr[left] > arr[largest] {
		largest = left
	}
	if right <= length && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		Max_heapify(arr, length, largest)
	}
}

func Build_max_heap(arr []int, length int) {
	//value:O 1 3 4
	//index:0 1 2 3
	for i := length / 2; i >= 1; i-- {
		Max_heapify(arr, length, i)
	}
}

func HeapSort(arr []int) {
	length := len(arr) - 1
	Build_max_heap(arr, length)
	for length > 0 {
		arr[length], arr[1] = arr[1], arr[length]
		length = length - 1
		Max_heapify(arr, length, 1)
	}
}

func main() {
	arr := []int{1000, 4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	Build_max_heap(arr, len(arr)-1)
	fmt.Println(arr)
	HeapSort(arr)
	fmt.Println(arr)
}

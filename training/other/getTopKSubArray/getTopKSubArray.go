package main

import (
	"container/heap"
	"fmt"
)

//find top k elements in an array

type intHeap []int

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	//fmt.Printf("Hello World!\n");
	arr := []int(12, 3, 4, 6, 5, 8, 3, 4, 6, 5)
	n := len(arr)
	k := 4
	h := intHeap{}
	heap.Init(&h)
	for i := 0; i < k; i++ {
		heap.Push(&h, arr[i])
	}
	for i := k; i < n; i++ {
		maxValue := heap.Top()
		if arr[i] < maxvalue {
			heap.Pop(&h)
			heap.Push(arr[i])
		}
	}
	//顺序输出堆的元素
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(&h)
	}

	fmt.Println(res)
}

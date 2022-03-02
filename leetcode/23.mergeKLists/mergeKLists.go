package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(arr []int) *ListNode {
	var head *ListNode
	var tail *ListNode
	for _, v := range arr {
		if head == nil {
			tail = &ListNode{Val: v}
			head = tail
		} else {
			tail.Next = &ListNode{Val: v}
			tail = tail.Next
		}
	}
	return head
}
func printList(head *ListNode) {
	p := head
	for p != nil {
		fmt.Printf("%d->", p.Val)
		p = p.Next
	}
}

//3种方法
//1.优先队列
//2.归并
//3.暴力
type Status struct {
	Val int
	Ptr *ListNode
}

type intHeap []Status

func (h intHeap) Len() int {
	return len(h)
}
func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h intHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(Status))
}

func (h *intHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

func mergeKLists1(lists []*ListNode) *ListNode {
	//time  O(n*log(k))
	//space O(n)
	//空间复杂度可以优化为O(k)
	var head *ListNode
	var tail *ListNode
	h := &intHeap{}
	heap.Init(h)
	for _, v := range lists {
		if v != nil {
			heap.Push(h, Status{Val: v.Val, Ptr: v})
		}
	}
	for len(*h) != 0 {
		v := heap.Pop(h).(Status)
		if tail == nil {
			tail = &ListNode{Val: v.Val}
			head = tail
		} else {
			tail.Next = &ListNode{Val: v.Val}
			tail = tail.Next
		}
		nPtr := v.Ptr.Next
		if nPtr != nil {
			heap.Push(h, Status{Val: nPtr.Val, Ptr: nPtr})
		}
	}
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//归并
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for interval := 1; interval < len(lists); interval = interval * 2 {
		for i := 0; i < len(lists)-interval; i = i + interval*2 {
			lists[i] = merge2Lists(lists[i], lists[i+interval])
		}
	}
	return lists[0]
}

func merge2Lists(list1, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		if list1 == nil {
			return list2
		} else {
			return list1
		}
	}
	var head ListNode
	tail, p1, p2 := &head, list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			tail.Next = p1
			p1 = p1.Next
			tail = tail.Next
		} else {
			tail.Next = p2
			p2 = p2.Next
			tail = tail.Next
		}
	}
	if p1 == nil {
		tail.Next = p2
	} else {
		tail.Next = p1
	}
	return head.Next
}

func main() {

}

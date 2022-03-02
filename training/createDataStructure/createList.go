package main

import "fmt"

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

/*
func main() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	head := createList(arr)
	printList(head)
}
*/

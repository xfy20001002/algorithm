package main

import (
	"fmt"
)

//两数相加

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	//solution1
	//先处理最短的一个链表，再处理较长的链表
	var head *ListNode
	var tail *ListNode
	carry := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		if head == nil {
			if sum > 9 {
				tail = &ListNode{Val: sum - 10}
				head = tail
				carry = 1
			} else {
				tail = &ListNode{Val: sum}
				head = tail
				carry = 0
			}
		} else {
			if sum > 9 {
				tail.Next = &ListNode{Val: sum - 10}
				tail = tail.Next

				carry = 1
			} else {
				tail.Next = &ListNode{Val: sum}
				tail = tail.Next

				carry = 0
			}
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil && l2 == nil && carry == 0 {
		return head
	} else {
		if l1 != nil {
			for l1 != nil {
				sum := l1.Val + carry
				if sum > 9 {
					tail.Next = &ListNode{Val: sum - 10}
					tail = tail.Next

					carry = 1
				} else {
					tail.Next = &ListNode{Val: sum}
					tail = tail.Next

					carry = 0
				}
				l1 = l1.Next
			}
			if carry == 1 {
				tail.Next = &ListNode{Val: 1}
				tail = tail.Next
			}
		} else if l2 != nil {
			for l2 != nil {
				sum := l2.Val + carry
				if sum > 9 {
					tail.Next = &ListNode{Val: sum - 10}
					tail = tail.Next

					carry = 1
				} else {
					tail.Next = &ListNode{Val: sum}
					tail = tail.Next

					carry = 0
				}
				l2 = l2.Next
			}
			if carry == 1 {
				tail.Next = &ListNode{Val: 1}
				tail = tail.Next
			}

		} else {
			tail.Next = &ListNode{Val: 1}
			tail = tail.Next
		}
		return head
	}
}

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	//一次处理完2个链表，最终只用判断是否有进位
	var head *ListNode
	var tail *ListNode
	var carry int = 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
		}
		if l2 != nil {
			n2 = l2.Val
		}
		/*
			sum ,carry = sum%10,carry/10
		*/
		sum := carry + n1 + n2
		if sum > 9 {
			head, tail = addElement(head, tail, sum-10)
			carry = 1
		} else {
			head, tail = addElement(head, tail, sum)
			carry = 0
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry == 1 {
		head, _ = addElement(head, tail, 1)
	}
	return head
}
func addElement(head, tail *ListNode, value int) (*ListNode, *ListNode) {
	if head == nil {
		head = &ListNode{Val: value}
		tail = head
	} else {
		tail.Next = &ListNode{Val: value}
		tail = tail.Next
	}
	return head, tail
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
	//ex1
	//2 4 3
	//5 6 4
	//ex2
	//2 4 6
	//2 4 4
	//ex3
	//2 4 6 9
	//2 5 4
	//ex4
	//[]
	//[]
	arr1 := []int{
		2, 4, 6, 9,
	}
	arr2 := []int{
		2, 5, 4,
	}
	l1 := createList(arr1)
	l2 := createList(arr2)
	res := AddTwoNumbers2(l1, l2)
	printList(res)
}
*/

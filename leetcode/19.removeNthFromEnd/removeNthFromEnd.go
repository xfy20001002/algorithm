package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	//让一个指针指向第三个节点
	var pre *ListNode = head
	var cur *ListNode = pre
	for i := 0; i < n; i++ {
		cur = cur.Next
	}
	if cur == nil {
		head = head.Next
		return head
	}
	for cur.Next != nil {
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = pre.Next.Next
	return head
}

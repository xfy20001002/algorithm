package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
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

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists merge two increment-sorted list
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return l1
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var root *ListNode
	if l1.Val < l2.Val {
		root = l1
		l1 = l1.Next
	} else {
		root = l2
		l2 = l2.Next
	}
	cur := root
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return root
}

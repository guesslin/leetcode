package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	one, two := head, head.Next
	for two != nil && two.Next != nil {
		one = one.Next
		two = two.Next.Next
	}
	two = one.Next
	one.Next = nil
	two = reverse(two)
	one = head
	for two != nil {
		if one.Val != two.Val {
			return false
		}
		one = one.Next
		two = two.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var rev *ListNode
	for head != nil {
		head.Next, rev, head = rev, head, head.Next
	}
	return rev
}

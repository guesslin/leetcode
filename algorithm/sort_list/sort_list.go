package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	l := 0
	cur := head
	for ; cur != nil; cur, l = cur.Next, l+1 {
	}
	if l <= 1 {
		return head
	}
	// break into two sub link list
	cur = head
	for i := 1; i < l/2; i++ {
		cur = cur.Next
	}
	cur, cur.Next = cur.Next, nil
	l1 := sortList(head)
	l2 := sortList(cur)
	return mergeTwoLists(l1, l2)
}

// mergeTwoLists merge two increment-sorted list
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
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

func printList(head *ListNode) {
	for ; head != nil; head = head.Next {
		fmt.Printf("%d", head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Printf("\n")
}

func main() {
	one := ListNode{Val: 1}
	two := ListNode{Val: 2}
	three := ListNode{Val: 3}
	four := ListNode{Val: 4}
	five := ListNode{Val: 5}
	one.Next = &three
	three.Next = &two
	two.Next = &five
	five.Next = &four
	head := &one
	printList(head)
	head = sortList(head)
	printList(head)
}

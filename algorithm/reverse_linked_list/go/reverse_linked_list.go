package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var rev *ListNode
	for head != nil {
		head.Next, rev, head = rev, head, head.Next
	}
	return rev
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
		if head != nil {
			fmt.Printf("-> ")
		}
	}
	fmt.Printf("\n")
}

func main() {
	head := ListNode{Val: 1}
	cur := &head
	for i := 2; i <= 5; i++ {
		cur.Next = &ListNode{Val: i}
		cur = cur.Next
	}
	printList(&head)
	cur = reverseList(&head)
	printList(cur)
}

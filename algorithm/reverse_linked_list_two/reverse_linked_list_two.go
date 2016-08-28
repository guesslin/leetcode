package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var dummy ListNode
	dummy.Next = head
	l := &dummy
	for i := 1; i < m; i++ {
		l = l.Next
	}
	r := l.Next
	var cur *ListNode
	for i := m; i < n; i++ {
		cur = r.Next
		r.Next = cur.Next
		cur.Next = l.Next
		l.Next = cur
	}
	return dummy.Next
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
	cur = reverseBetween(&head, 2, 4)
	printList(cur)
}

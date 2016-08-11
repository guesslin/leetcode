package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	for head != nil && head.Val == val {
		if head.Val == val {
			head = head.Next
		}
	}
	for cur := head; cur != nil; {
		if cur.Next != nil && cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func printLinkedList(head *ListNode) {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("%d ", cur.Val)
	}
	fmt.Printf("\n")
}

func main() {
	l := ListNode{1, nil}
	l.Next = &ListNode{2, nil}
	l.Next.Next = &ListNode{2, nil}
	l.Next.Next.Next = &ListNode{1, nil}
	printLinkedList(removeElements(&l, 2))

}

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var tmp ListNode
	tmp.Next = head
	remove := &tmp
	pToTail := head
	for i := 1; i <= n; i++ {
		pToTail = pToTail.Next
	}
	for pToTail != nil {
		pToTail = pToTail.Next
		remove = remove.Next
	}
	remove.Next = remove.Next.Next
	return tmp.Next
}

/*
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return head
}
*/

func printLinkedList(head *ListNode) {
	fmt.Printf("[ ")
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("%d ", cur.Val)
	}
	fmt.Printf("]\n")
}

func main() {
	l := ListNode{1, nil}
	l.Next = &ListNode{2, nil}
	//	l.Next.Next = &ListNode{3, nil}
	//	l.Next.Next.Next = &ListNode{4, nil}
	//	l.Next.Next.Next.Next = &ListNode{5, nil}
	printLinkedList(removeNthFromEnd(&l, 1))
}

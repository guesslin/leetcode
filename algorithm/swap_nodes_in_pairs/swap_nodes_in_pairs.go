package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// set two pointer to first and second element
	cur_f, cur_b := head, head.Next
	head = cur_b
	cur_f.Next = cur_b.Next
	cur_b.Next = cur_f
	for cur_f != nil && cur_b != nil {
		cur := cur_f
		cur_f = cur_f.Next
		if cur_f == nil {
			break
		}
		cur_b = cur_f.Next
		if cur_b == nil {
			break
		}
		cur_f.Next = cur_b.Next
		cur_b.Next = cur_f
		cur.Next = cur_b
	}
	return head
}

func printList(head *ListNode) {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("%d\t", cur.Val)
	}
	fmt.Printf("\n")
}

func main() {
	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	printList(&head)
	cur := swapPairs(&head)
	printList(cur)
}

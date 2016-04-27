package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for current, node := head, head.Next; node != nil; node = node.Next {
		if current.Val == node.Val {
			current.Next = node.Next
		} else {
			current = node
		}
	}
	return head
}

func main() {
	L4 := ListNode{3, nil}
	L3 := ListNode{3, &L4}
	L2 := ListNode{2, &L3}
	L1 := ListNode{2, &L2}
	r := deleteDuplicates(&L1)
	for n := r; n != nil; n = n.Next {
		fmt.Printf("%d\t", n.Val)
	}
	fmt.Println()
}

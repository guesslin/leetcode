package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// partition return a list generated by old list and partition by given x
// like Given 1->4->3->2->5->2 and x = 3, return 1->2->2->4->3->5
// all elements in list should keep original sequance.
func partition(head *ListNode, x int) *ListNode {
	var sHead, gHead ListNode
	small := &sHead
	greate := &gHead
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			small.Next = cur
			small = small.Next
		} else {
			greate.Next = cur
			greate = greate.Next
		}
	}
	greate.Next = nil
	small.Next = gHead.Next
	return sHead.Next
}

func printList(head *ListNode) {
	for ; head != nil; head = head.Next {
		fmt.Printf("%d->", head.Val)
	}
	fmt.Printf("NULL\n")
}

func main() {
	two := ListNode{Val: 2}
	one := ListNode{Val: 1}
	two.Next = &one
	printList(&two)
	head := partition(&two, 2)
	printList(head)
}

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := new(ListNode)
	var x, y, carry int
	temp := root
	for {
		x, y = 0, 0
		if l1 != nil {
			x, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			y, l2 = l2.Val, l2.Next
		}
		temp.Val = x + y + carry
		carry = temp.Val / 10
		temp.Val = temp.Val % 10
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		temp.Next = new(ListNode)
		temp = temp.Next
	}
	return root
}

func main() {
	L1 := new(ListNode) // 2->4->3->9
	L1.Next = new(ListNode)
	L1.Next.Next = new(ListNode)
	L1.Next.Next.Next = new(ListNode)
	L1.Val = 2
	L1.Next.Val = 4
	L1.Next.Next.Val = 3
	L1.Next.Next.Next.Val = 9

	L2 := new(ListNode) // 5->6->4
	L2.Next = new(ListNode)
	L2.Next.Next = new(ListNode)
	L2.Val = 5
	L2.Next.Val = 6
	L2.Next.Next.Val = 4

	L3 := new(ListNode) // 5
	L3.Val = 5
	L4 := new(ListNode) // 5
	L4.Val = 5

	result := addTwoNumbers(L1, L2)
	for node := result; node != nil; node = node.Next {
		fmt.Printf("%d\t", node.Val)
	}
	fmt.Printf("\n")
	result = addTwoNumbers(nil, L2)
	for node := result; node != nil; node = node.Next {
		fmt.Printf("%d\t", node.Val)
	}
	fmt.Printf("\n")
	result = addTwoNumbers(L3, L4)
	for node := result; node != nil; node = node.Next {
		fmt.Printf("%d\t", node.Val)
	}
	fmt.Printf("\n")
}

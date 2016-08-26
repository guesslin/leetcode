package main

import (
	"fmt"
)

/*
 * Given a list, rotate the list to the right by k places, where k is non-negative.
 * For example:
 *   Given 1->2->3->4->5->NULL and k = 2,
 *   return 4->5->1->2->3->NULL.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	l := 1
	cur := head
	// count the length of list
	for cur.Next != nil {
		l++
		cur = cur.Next
	}
	// after loop cur point to the last element in list
	last := cur
	cur = head
	k = l - (k % l)
	for i := 1; i < k; i++ {
		cur = cur.Next
	}
	last.Next = head
	head = cur.Next
	cur.Next = nil
	return head
}

func printList(head *ListNode) {
	for {
		fmt.Printf("%d", head.Val)
		head = head.Next
		if head != nil {
			fmt.Printf("->")
		} else {
			break
		}
	}
	fmt.Printf("\n")
}

func main() {
	five := ListNode{Val: 5}
	four := ListNode{Val: 4, Next: &five}
	three := ListNode{Val: 3, Next: &four}
	two := ListNode{Val: 2, Next: &three}
	one := ListNode{Val: 1, Next: &two}
	head := rotateRight(&one, 0)
	printList(head)
}

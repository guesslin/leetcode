package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

// @param head The linked list's head. Note that the head is guanranteed to be not null, so it contains at least one node.
func Constructor(head *ListNode) Solution {
	var s Solution
	s.head = head
	return s
}

// Returns a random node's value.
// Use Reservoir Sampling Algorithm as randomize algorithm
// https://en.wikipedia.org/wiki/Reservoir_sampling
// 可以改成取 k 個數字出來，只要修改判斷條件的 `if j = rand.Intn(i); j < k`
// 並且改為 res := make([]int, k)
//          res[j] = cursor.Val
func (this *Solution) GetRandom() int {
	cursor := this.head
	res, i := 0, 1
	for cursor != nil {
		if rand.Intn(i) == 0 {
			res = cursor.Val
		}
		i++
		cursor = cursor.Next
	}
	return res

}

func main() {
	head := ListNode{Val: 1, Next: nil}
	head.Next = &ListNode{Val: 2, Next: nil}
	head.Next.Next = &ListNode{Val: 3, Next: nil}
	s := Constructor(&head)
	fmt.Println(s.GetRandom())
	fmt.Println(s.GetRandom())
	fmt.Println(s.GetRandom())
}

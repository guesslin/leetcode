package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ld := depth(root.Left)
	rd := depth(root.Right)
	if ld == -1 || rd == -1 || abs(ld-rd) > 1 {
		return false
	}
	return true
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	ld := depth(node.Left)
	rd := depth(node.Right)
	if ld == -1 || rd == -1 || abs(ld-rd) > 1 {
		return -1
	}
	return max(ld, rd) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	n11 := TreeNode{11, nil, nil}
	n10 := TreeNode{10, nil, nil}
	n9 := TreeNode{9, nil, nil}
	n8 := TreeNode{8, nil, nil}
	n7 := TreeNode{7, nil, nil}
	n6 := TreeNode{6, nil, nil}
	n5 := TreeNode{5, nil, nil}
	n4 := TreeNode{4, nil, nil}
	n3 := TreeNode{3, nil, nil}
	n2 := TreeNode{2, nil, nil}
	n1 := TreeNode{1, nil, nil}
	n3.Right = &n7
	n3.Left = &n6
	n2.Right = &n5
	n2.Left = &n4
	n1.Right = &n3
	n1.Left = &n2
	fmt.Printf("Is Tree Balanced? Expect true, Get: %v\n", isBalanced(&n1))
	n7.Right = &n8
	n8.Right = &n9
	n9.Right = &n10
	fmt.Printf("Is Tree Balanced? Expect false, Get: %v\n", isBalanced(&n1))
	fmt.Printf("Is Tree Balanced? Expect true, Get: %v\n", isBalanced(&n11))
	fmt.Printf("Is Tree Balanced? Expect true, Get: %v\n", isBalanced(nil))
}

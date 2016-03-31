package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var s []*TreeNode
	var ans []int
	for {
		for root != nil {
			if root.Right != nil {
				s = append(s, root.Right)
			}
			ans = append(ans, root.Val)
			root = root.Left
		}
		if len(s) == 0 {
			break
		}
		root, s = s[len(s)-1], s[:len(s)-1]
	}
	return ans
}

func main() {
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
	fmt.Printf("PreOrder Traversal: %v\n", preorderTraversal(&n1))
}

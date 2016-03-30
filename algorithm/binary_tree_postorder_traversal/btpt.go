package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Peek(stack []*TreeNode) *TreeNode {
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return nil
}

// postorderTraversal traversal tree in right-left-middle sequence
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	s := make([]*TreeNode, 0)
	ans := make([]int, 0)
	for {
		for root != nil {
			if root.Right != nil {
				s = append(s, root.Right)
			}
			s = append(s, root)
			root = root.Left
		}
		root, s = s[len(s)-1], s[:len(s)-1]
		if root.Right != nil && Peek(s) == root.Right {
			s = s[:len(s)-1]
			s = append(s, root)
			root = root.Right
		} else {
			ans = append(ans, root.Val)
			root = nil
		}
		if len(s) == 0 {
			break
		}

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
	fmt.Println(postorderTraversal(&n1))
}

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

func print_s(stack []*TreeNode) {
	for _, s := range stack {
		fmt.Printf("%d ", s.Val)
	}
	fmt.Printf("\n")
}

// postorderTraversal traversal tree in right-left-middle sequence
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	var ans []int
	for {
		for root != nil {
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if root.Right != nil && Peek(stack) == root.Right {
			stack = stack[:len(stack)-1]
			stack = append(stack, root)
			root = root.Right
		} else {
			ans = append(ans, root.Val)
			root = nil
		}
		if len(stack) == 0 {
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
	fmt.Printf("PostOrder Traversal: %v\n", postorderTraversal(&n1))
}

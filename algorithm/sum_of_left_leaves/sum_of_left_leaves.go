package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	stack := make([]*TreeNode, 0)
	sum := 0
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left := cur.Left
		right := cur.Right
		if left != nil {
			if left.Left == nil && left.Right == nil {
				sum += left.Val
			} else {
				stack = append(stack, left)
			}
		}
		if right != nil {
			stack = append(stack, right)
		}
	}
	return sum
}

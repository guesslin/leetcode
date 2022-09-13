package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Runtime: 80 ms, faster than 97.98% of Go online submissions for Count Good Nodes in Binary Tree.
// Memory Usage: 10.4 MB, less than 82.98% of Go online submissions for Count Good Nodes in Binary Tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return visit(root, root.Val)
}

func visit(node *TreeNode, num int) int {
	l, r := 0, 0
	num = max(node.Val, num)
	if node.Left != nil {
		l = visit(node.Left, num)
	}
	if node.Right != nil {
		r = visit(node.Right, num)
	}
	if node.Val >= num {
		return l + r + 1
	}
	return l + r
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

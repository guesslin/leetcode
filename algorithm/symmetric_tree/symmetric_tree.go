package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	left := treeToOJInorder(root.Left)
	right := treeToOJInorder(root.Right)
	right = reverse(right)
	if strings.Compare(left, right) == 0 {
		if root.Left.Val != root.Right.Val {
			return false
		}
		return true
	}
	return false
}

func treeToOJInorder(node *TreeNode) string {
	if node == nil {
		return "#"
	}
	if node.Left == nil && node.Right == nil {
		return fmt.Sprintf("%d", node.Val)
	}
	return fmt.Sprintf("%s,%d,%s", treeToOJInorder(node.Left), node.Val, treeToOJInorder(node.Right))
}

func reverse(s string) string {
	r := strings.Split(s, ",")
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return strings.Join(r, ",")
}

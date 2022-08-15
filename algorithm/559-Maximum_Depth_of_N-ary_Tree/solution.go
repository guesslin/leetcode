package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	cur := 0
	for _, child := range root.Children {
		cur = max(cur, maxDepth(child))
	}
	return 1 + cur
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	nodes := make([]TreeNode, len(nums))
	for i := 0; i < len(nums); i++ {
		nodes[i].Val = nums[i]
	}
	lo, hi := 0, len(nums)-1
	mid := (lo + hi) / 2
	root := &nodes[mid]
	root.Left = buildTree(nodes, lo, mid-1)
	root.Right = buildTree(nodes, mid+1, hi)
	return root
}

func buildTree(nodes []TreeNode, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	mid := (lo + hi) / 2
	root := &nodes[mid]
	root.Left = buildTree(nodes, lo, mid-1)
	root.Right = buildTree(nodes, mid+1, hi)
	return root
}

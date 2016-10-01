package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	nums := make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		nums = append(nums, cur.Val)
	}
	return sortedArrayToBST(nums)
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

func buildTreeDFS(root *TreeNode, head *ListNode) *ListNode {
	if root.Right == nil && root.Left == nil {
		root.Val = head.Val
		return head.Next
	}
	if root.Left != nil {
		head = buildTreeDFS(root.Left, head)
	}
	root.Val = head.Val
	head = head.Next
	if root.Right != nil {
		head = buildTreeDFS(root.Right, head)
	}
	return head
}

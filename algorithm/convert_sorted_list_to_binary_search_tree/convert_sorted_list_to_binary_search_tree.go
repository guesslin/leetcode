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
	root := TreeNode{Val: 0}
	queue := make([]*TreeNode, 0)
	queue = append(queue, &root)
	l := head.Next
	var cur *TreeNode
	for l != nil {
		cur, queue = queue[0], queue[1:]
		if l != nil {
			cur.Left = &TreeNode{Val: 0}
			queue = append(queue, cur.Left)
			l = l.Next
		}
		if l != nil {
			cur.Right = &TreeNode{Val: 0}
			queue = append(queue, cur.Right)
			l = l.Next
		}
	}
	buildTreeDFS(&root, head)
	return &root
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

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for {
		level++
		level_len := len(queue)
		for i := 0; i < level_len; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if queue[i].Right == nil && queue[i].Left == nil {
				return level
			}
		}
		queue = queue[level_len:]
		if len(queue) == 0 {
			break
		}
	}
	return level
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	level := levelOrder(root)
	res := make([]int, 0, len(level))
	for i := 0; i < len(level); i++ {
		res = append(res, level[i][len(level[i])-1])
	}
	return res

}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for {
		level_len := len(queue)
		for i := 0; i < level_len; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		tmp := make([]int, 0)
		for i := 0; i < level_len; i++ {
			tmp = append(tmp, queue[i].Val)
		}
		res = append(res, tmp)
		queue = queue[level_len:]
		if len(queue) == 0 {
			break
		}
	}
	return res
}

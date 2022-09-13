package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		l := len(queue)

		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		average := calAverage(queue[:l])
		queue = queue[l:]
		res = append(res, average)
	}

	return res
}

func calAverage(in []*TreeNode) float64 {
	sum := int64(0)
	for _, node := range in {
		sum = sum + int64(node.Val)
	}
	return float64(sum) / float64(len(in))
}

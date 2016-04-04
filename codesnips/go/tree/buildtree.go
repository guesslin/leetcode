package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func toInt(s string) int {
	var result int
	for _, c := range s {
		if 48 <= c && c <= 57 {
			result = result*10 + int(c-48)
		} else {
			return 0
		}
	}
	return result
}

// BuildTree build *TreeNode by giving string sequence
// using queue to read
func BuildTree(nodes []string) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	n := new(TreeNode)
	n.Val = toInt(nodes[0])
	root := n
	var queue []*TreeNode
	for i := 1; i < len(nodes); i = i + 2 {
		if nodes[i] != "null" {
			m := new(TreeNode)
			m.Val = toInt(nodes[i])
			n.Left = m
			queue = append(queue, m)
		}
		if nodes[i+1] != "null" {
			m := new(TreeNode)
			m.Val = toInt(nodes[i+1])
			n.Right = m
			queue = append(queue, m)
		}
		n, queue = queue[0], queue[1:]
	}
	return root
}

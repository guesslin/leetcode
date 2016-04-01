package tree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree build *TreeNode by giving string sequence
// using queue to read
func BuildTree(nodes []string) *TreeNode {
	for i := 0; i < len(nodes); i++ {
		fmt.Println(nodes[i])
	}
	return nil
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Peek(stack []*TreeNode) *TreeNode {
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return nil
}

func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	var ans []int
	for len(ans) < k {
		for root != nil {
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]

		// root.Left will be nil, so put root.Val into ans
		ans = append(ans, root.Val)

		// check if root.Right has something, then keep going with it
		if root.Right != nil && Peek(stack) == root.Right {
			stack = stack[:len(stack)-1]
			root = root.Right
			continue
		} else {
			root = nil
		}
		if len(stack) == 0 {
			break
		}
	}
	return ans[k-1]
}

func main() {
}

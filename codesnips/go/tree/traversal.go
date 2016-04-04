package tree

func Peek(stack []*TreeNode) *TreeNode {
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return nil
}

// postorderTraversal traversal tree in right-left-middle sequence
func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	var ans []int
	for {
		for root != nil {
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if root.Right != nil && Peek(stack) == root.Right {
			stack = stack[:len(stack)-1]
			stack = append(stack, root)
			root = root.Right
		} else {
			ans = append(ans, root.Val)
			root = nil
		}
		if len(stack) == 0 {
			break
		}

	}
	return ans
}

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	var ans []int
	for {
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
	return ans
}

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var s []*TreeNode
	var ans []int
	for {
		for root != nil {
			if root.Right != nil {
				s = append(s, root.Right)
			}
			ans = append(ans, root.Val)
			root = root.Left
		}
		if len(s) == 0 {
			break
		}
		root, s = s[len(s)-1], s[:len(s)-1]
	}
	return ans
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	sumTree(root, 0)                    // build tree as range sum path
	return findSum(root, sum, []int{0}) // find range sum
}

func findSum(root *TreeNode, sum int, bases []int) int {
	if root == nil {
		return 0
	}
	res := 0
	for _, base := range bases {
		if root.Val == sum+base {
			res++
		}
	}
	bases = append(bases, root.Val)
	right := findSum(root.Right, sum, bases)
	left := findSum(root.Left, sum, bases)
	bases = bases[:len(bases)-1]
	return right + left + res
}

func sumTree(root *TreeNode, base int) {
	if root == nil {
		return
	}
	root.Val += base
	sumTree(root.Left, root.Val)
	sumTree(root.Right, root.Val)
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var leftDeep, rightDeep int

	if root == nil {
		return 0
	}

	leftDeep = maxDepth(root.Left)
	rightDeep = maxDepth(root.Right)
	if leftDeep > rightDeep {
		return leftDeep + 1
	}
	return rightDeep + 1
}

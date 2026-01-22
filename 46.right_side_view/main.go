package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	levelMap := make(map[int]struct{})
	result := make([]int, 0)

	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		depth += 1
		if _, ok := levelMap[depth]; !ok {
			result = append(result, root.Val)
			levelMap[depth] = struct{}{}
		}
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}
	dfs(root, 0)
	return result
}

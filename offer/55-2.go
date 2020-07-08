/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := true
	// dfs用于获取每个节点的深度
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *bool) int {
	if root == nil {
		return 0
	}
	leftDep := dfs(root.Left, res)
	rightDep := dfs(root.Right, res)
	if getAbs(leftDep - rightDep) > 1 {
		*res = false
	}
	return getMax(leftDep, rightDep) + 1
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
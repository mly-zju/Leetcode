/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	// 深度：dfs或者层次遍历
	if root == nil {
		return 0
	}
	return 1 + getMax(maxDepth(root.Left), maxDepth(root.Right))
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
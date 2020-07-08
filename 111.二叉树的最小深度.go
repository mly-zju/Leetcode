/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (40.16%)
 * Likes:    216
 * Dislikes: 0
 * Total Accepted:    53.8K
 * Total Submissions: 129.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 * 
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 返回它的最小深度  2.
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return getMinLevel(root, 0)
}

func getMinLevel(root *TreeNode, res int) int {
	if root == nil {
		return res
	} else if root.Left == nil && root.Right == nil {
		return res + 1
	} else if root.Left == nil {
		return getMinLevel(root.Right, res + 1)
	} else if root.Right == nil {
		return getMinLevel(root.Left, res + 1)
	}

	return getMin(getMinLevel(root.Left, res + 1), getMinLevel(root.Right, res + 1))
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end


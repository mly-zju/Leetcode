/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 *
 * https://leetcode-cn.com/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (46.12%)
 * Likes:    148
 * Dislikes: 0
 * Total Accepted:    12.5K
 * Total Submissions: 27K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。
 *
 * 示例 :
 * 给定二叉树
 *
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \
 * ⁠     4   5
 *
 *
 * 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
 *
 * 注意：两结点之间的路径长度是以它们之间边的数目表示。
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
var max int

func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	getLevel(root)
	return max
}

func getLevel(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftLevel := getLevel(root.Left)
	rightLevel := getLevel(root.Right)
	// 计算max
	if max < leftLevel+rightLevel {
		max = leftLevel + rightLevel
	}
	return 1 + getMax(leftLevel, rightLevel)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end


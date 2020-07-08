/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
 *
 * https://leetcode-cn.com/problems/longest-univalue-path/description/
 *
 * algorithms
 * Easy (36.85%)
 * Likes:    156
 * Dislikes: 0
 * Total Accepted:    7.6K
 * Total Submissions: 20.1K
 * Testcase Example:  '[5,4,5,1,1,5]'
 *
 * 给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。
 *
 * 注意：两个节点之间的路径长度由它们之间的边数表示。
 *
 * 示例 1:
 *
 * 输入:
 *
 *
 * ⁠             5
 * ⁠            / \
 * ⁠           4   5
 * ⁠          / \   \
 * ⁠         1   1   5
 *
 *
 * 输出:
 *
 *
 * 2
 *
 *
 * 示例 2:
 *
 * 输入:
 *
 *
 * ⁠             1
 * ⁠            / \
 * ⁠           4   5
 * ⁠          / \   \
 * ⁠         4   4   5
 *
 *
 * 输出:
 *
 *
 * 2
 *
 *
 * 注意: 给定的二叉树不超过10000个结点。 树的高度不超过1000。
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
var maxLen int

func longestUnivaluePath(root *TreeNode) int {
	maxLen = 0
	if root == nil {
		return maxLen
	}

	dfs(root)
	return maxLen
}

// 返回以root为根的最长同值单边节点长度，以及具体的同值val
func dfs(root *TreeNode) (prevSum, prevVal int) {
	if root == nil {
		return 0, 0
	}

	// 计算以当前节点为根的最长同值
	curMax := 0
	if root.Left != nil {
		leftSum, leftVal := dfs(root.Left)
		if leftVal == root.Val {
			curMax += leftSum + 1
			// 同时计算一下prevSum，因为只能取一条边的
			if leftSum + 1 > prevSum {
				prevSum = leftSum + 1
			}
		}
	}
	if root.Right != nil {
		rightSum, rightVal := dfs(root.Right)
		if rightVal == root.Val {
			curMax += rightSum + 1
			// 同时计算一下prevSum，因为只能取一条边的
			if rightSum + 1 > prevSum {
				prevSum = rightSum + 1
			}
		}
	}

	if curMax > maxLen {
		maxLen = curMax
	}
	return prevSum, root.Val
}

// @lc code=end


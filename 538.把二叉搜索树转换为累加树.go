/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 *
 * https://leetcode-cn.com/problems/convert-bst-to-greater-tree/description/
 *
 * algorithms
 * Easy (56.88%)
 * Likes:    160
 * Dislikes: 0
 * Total Accepted:    10.7K
 * Total Submissions: 18.6K
 * Testcase Example:  '[5,2,13]'
 *
 * 给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater
 * Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
 *
 * 例如：
 *
 *
 * 输入: 二叉搜索树:
 * ⁠             5
 * ⁠           /   \
 * ⁠          2     13
 *
 * 输出: 转换为累加树:
 * ⁠            18
 * ⁠           /   \
 * ⁠         20     13
 *
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
var sum int
var prev int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	prev = 0
	revInorder(root)
	return root
}

func revInorder(root *TreeNode) {
	if root == nil {
		return
	}
	// 先右边中序遍历
	revInorder(root.Right)

	// 暂存当前值
	tmp := root.Val
	// 对比前驱是否大于当前
	if prev > root.Val {
		root.Val += sum
	}
	sum += tmp
	prev = tmp
	// 然后左边遍历
	revInorder(root.Left)
}

// @lc code=end


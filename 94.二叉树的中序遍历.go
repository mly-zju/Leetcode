/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (67.95%)
 * Likes:    406
 * Dislikes: 0
 * Total Accepted:    110K
 * Total Submissions: 156.4K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
 * 
 * 示例:
 * 
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 * 
 * 输出: [1,3,2]
 * 
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	// 1. 递归
	// res = dfs(root, res)
	// return res

	// 2. 迭代
	cache := []*TreeNode{}
	for root != nil || len(cache) != 0 {
		for root != nil {
			cache = append(cache, root)
			root = root.Left
		}
		newLen := len(cache)
		root = cache[newLen-1]
		cache = cache[0:newLen-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func dfs(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = dfs(root.Left, res)
	res = append(res, root.Val)
	res = dfs(root.Right, res)
	return res
}
// @lc code=end


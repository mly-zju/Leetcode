/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (61.75%)
 * Likes:    208
 * Dislikes: 0
 * Total Accepted:    71.6K
 * Total Submissions: 111.3K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 前序 遍历。
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
 * 输出: [1,2,3]
 * 
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
func preorderTraversal(root *TreeNode) []int {
	// 1. 递归
	// res := []int{}
	// if root == nil {
	// 	return res
	// }

	// res = dfs(root, res)
	// return res

	// 2. 迭代
	res := []int{}
	if root == nil {
		return res
	}
	// 迭代肯定要基于栈，不用问
	cache := []*TreeNode{}
	for root != nil || len(cache) != 0 {
		for root != nil {
			res = append(res, root.Val)
			cache = append(cache, root)
			root = root.Left
		}
		newLen := len(cache)
		root = cache[newLen - 1].Right
		cache = cache[0:newLen - 1]
	}
	return res
}

func dfs(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = dfs(root.Left, res)
	res = dfs(root.Right, res)
	return res
}
// @lc code=end


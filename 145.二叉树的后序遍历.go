/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (67.54%)
 * Likes:    230
 * Dislikes: 0
 * Total Accepted:    53.2K
 * Total Submissions: 75.7K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 后序 遍历。
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
 * 输出: [3,2,1]
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
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	// 1. 递归
	// res = dfs(root, res)
	// return res

	// 2. 迭代
	seenMap := make(map[*TreeNode]bool)
	cache := []*TreeNode{}
	for root != nil || len(cache) != 0 {
		for root != nil && !seenMap[root] {
			cache = append(cache, root)
			root = root.Left
		}
		length := len(cache)
		root = cache[length-1]
		cache = cache[0:length-1]

		if !seenMap[root] {
			seenMap[root] = true
			cache = append(cache, root)
			root = root.Right
		} else {
			res = append(res, root.Val)
			// 将root置为nil，等待下一轮从cache获取
			root = nil
		}
	}
	return res
}

func dfs(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = dfs(root.Left, res)
	res = dfs(root.Right, res)
	res = append(res, root.Val)
	return res
}
// @lc code=end


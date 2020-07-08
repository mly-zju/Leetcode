import "strconv"

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (59.98%)
 * Likes:    166
 * Dislikes: 0
 * Total Accepted:    18.1K
 * Total Submissions: 30K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 输入:
 *
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 *
 * 输出: ["1->2->5", "1->3"]
 *
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	cache := []int{}
	res = dfs(root, cache, res)
	return res
}

func dfs(root *TreeNode, cache []int, res []string) []string {
	if root == nil {
		return res
	} 

	cache = append(cache, root.Val)
	if root.Left == nil && root.Right == nil {
		res = append(res, getString(cache))
		return res
	} else {
		if root.Left != nil {
			res = dfs(root.Left, cache, res)
		}
		if root.Right != nil {
			res = dfs(root.Right, cache, res)
		}
		return res
	}
}

func getString(cache []int) string {
	res := ""
	length := len(cache)
	for i := 0; i < length; i++ {
		if i != length - 1 {
			res += strconv.Itoa(cache[i]) + "->"
		} else {
			res += strconv.Itoa(cache[i])
		}
	}
	return res
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (70.94%)
 * Likes:    467
 * Dislikes: 0
 * Total Accepted:    130.5K
 * Total Submissions: 180.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 * 
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 返回它的最大深度 3 。
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
func maxDepth(root *TreeNode) int {
	// 1. 递归
	// res := 0
	// return getMaxLevel(root, res)

	// 2. 迭代：可以基于层次遍历，直到某一层全部为nil，则就是层数
	res := 0
	if root == nil {
		return res
	}

	cache := []*TreeNode{root}
	for len(cache) != 0 {
		levelLen := len(cache)
		for i := 0; i < levelLen; i++ {
			if cache[i].Left != nil {
				cache = append(cache, cache[i].Left)
			}
			if cache[i].Right!= nil {
				cache = append(cache, cache[i].Right)
			}
		}
		cache = cache[levelLen:]
		res++
	}
	return res
}

func getMaxLevel(root *TreeNode, res int) int {
	if root == nil {
		return res
	}

	res++
	leftLevel := getMaxLevel(root.Left, 0)
	rightLevel := getMaxLevel(root.Right, 0)
	if leftLevel > rightLevel {
		return res + leftLevel
	} else {
		return res + rightLevel
	}
}
// @lc code=end


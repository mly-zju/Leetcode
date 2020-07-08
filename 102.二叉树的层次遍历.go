/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (58.78%)
 * Likes:    305
 * Dislikes: 0
 * Total Accepted:    56.5K
 * Total Submissions: 94.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 * 例如:
 * 给定二叉树: [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其层次遍历结果：
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
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
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	cache := []*TreeNode{}
	if root == nil {
		return res
	}
	cache = append(cache, root)
	for len(cache) > 0 {
		levelLen := len(cache)
		levelArr := []int{}
		for i := 0; i < levelLen; i++ {
			levelArr = append(levelArr, cache[i].Val)
			if cache[i].Left != nil {
				cache = append(cache, cache[i].Left)
			}
			if cache[i].Right != nil {
				cache = append(cache, cache[i].Right)
			}
		}
		res = append(res, levelArr)
		cache = cache[levelLen:]
	}
	return res
}

// @lc code=end


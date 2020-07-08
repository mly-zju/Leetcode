/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
 *
 * https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/description/
 *
 * algorithms
 * Easy (60.95%)
 * Likes:    81
 * Dislikes: 0
 * Total Accepted:    9.9K
 * Total Submissions: 16.1K
 * Testcase Example:  '[3,9,20,15,7]'
 *
 * 给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.
 *
 * 示例 1:
 *
 * 输入:
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 输出: [3, 14.5, 11]
 * 解释:
 * 第0层的平均值是 3,  第1层是 14.5, 第2层是 11. 因此返回 [3, 14.5, 11].
 *
 *
 * 注意：
 *
 *
 * 节点值的范围在32位有符号整数范围内。
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
func averageOfLevels(root *TreeNode) []float64 {
	cache := []*TreeNode{}
	res := []float64{}
	if root == nil {
		return res
	}

	cache = append(cache, root)
	for len(cache) > 0 {
		levelLen := len(cache)
		levelSum := 0
		for i := 0; i < levelLen; i++ {
			levelSum += cache[i].Val
			if cache[i].Left != nil {
				cache = append(cache, cache[i].Left)
			}
			if cache[i].Right != nil {
				cache = append(cache, cache[i].Right)
			}
		}
		levelAvg := float64(levelSum) / float64(levelLen)
		res = append(res, levelAvg)
		cache = cache[levelLen:]
	}

	return res
}

// @lc code=end


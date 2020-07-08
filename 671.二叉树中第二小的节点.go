/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
 *
 * https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree/description/
 *
 * algorithms
 * Easy (44.28%)
 * Likes:    53
 * Dislikes: 0
 * Total Accepted:    5.5K
 * Total Submissions: 12.5K
 * Testcase Example:  '[2,2,5,null,null,5,7]'
 *
 * 给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或
 * 0。如果一个节点有两个子节点的话，那么这个节点的值不大于它的子节点的值。
 *
 * 给出这样的一个二叉树，你需要输出所有节点中的第二小的值。如果第二小的值不存在的话，输出 -1 。
 *
 * 示例 1:
 *
 *
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 2   5
 * ⁠    / \
 * ⁠   5   7
 *
 * 输出: 5
 * 说明: 最小的值是 2 ，第二小的值是 5 。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 2   2
 *
 * 输出: -1
 * 说明: 最小的值是 2, 但是不存在第二小的值。
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
func findSecondMinimumValue(root *TreeNode) int {
	// 层次遍历, 上层的肯定是最小的
	if root == nil {
		return -1
	}

	// ansArr第一个值肯定是root
	ansArr := []int{}
	cache := []*TreeNode{root}

	for len(cache) > 0 {
		levelLen := len(cache)
		for i := 0; i < levelLen; i++ {
			ansLen := len(ansArr)
			if ansLen == 0 || (ansLen == 1 && cache[i].Val != ansArr[0]) {
				// 第一个值肯定是root.Val, 第二个值必须不等于第一个值
				ansArr = append(ansArr, cache[i].Val)
			} else if ansLen > 1 && ansArr[1] > cache[i].Val && ansArr[0] != cache[i].Val {
				// 否则需要对比第二个值的大小
				ansArr[1] = cache[i].Val
			}

			// 有条件的推送cache
			if ansLen < 2 || cache[i].Val <= ansArr[1] {
				if cache[i].Left != nil {
					cache = append(cache, cache[i].Left)
					cache = append(cache, cache[i].Right)
				}
			}
		}
		cache = cache[levelLen:]
	}

	if len(ansArr) < 2 {
		return -1
	}
	return ansArr[1]
}

// @lc code=end


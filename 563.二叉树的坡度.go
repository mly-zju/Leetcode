/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
 *
 * https://leetcode-cn.com/problems/binary-tree-tilt/description/
 *
 * algorithms
 * Easy (49.92%)
 * Likes:    41
 * Dislikes: 0
 * Total Accepted:    6K
 * Total Submissions: 11.9K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个二叉树，计算整个树的坡度。
 *
 * 一个树的节点的坡度定义即为，该节点左子树的结点之和和右子树结点之和的差的绝对值。空结点的的坡度是0。
 *
 * 整个树的坡度就是其所有节点的坡度之和。
 *
 * 示例:
 *
 *
 * 输入:
 * ⁠        1
 * ⁠      /   \
 * ⁠     2     3
 * 输出: 1
 * 解释:
 * 结点的坡度 2 : 0
 * 结点的坡度 3 : 0
 * 结点的坡度 1 : |2-3| = 1
 * 树的坡度 : 0 + 0 + 1 = 1
 *
 *
 * 注意:
 *
 *
 * 任何子树的结点的和不会超过32位整数的范围。
 * 坡度的值不会超过32位整数的范围。
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
var tilt int

func findTilt(root *TreeNode) int {
	tilt = 0
	getSum(root)
	return tilt
}

func getSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum := getSum(root.Left)
	rightSum := getSum(root.Right)
	tilt += getAbs(leftSum - rightSum)
	return leftSum + rightSum + root.Val
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end


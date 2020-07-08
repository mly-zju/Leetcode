/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
 *
 * https://leetcode-cn.com/problems/count-complete-tree-nodes/description/
 *
 * algorithms
 * Medium (59.50%)
 * Likes:    94
 * Dislikes: 0
 * Total Accepted:    10.2K
 * Total Submissions: 16.3K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * 给出一个完全二叉树，求出该树的节点个数。
 *
 * 说明：
 *
 *
 * 完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第
 * h 层，则该层包含 1~ 2^h 个节点。
 *
 * 示例:
 *
 * 输入:
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠/ \  /
 * 4  5 6
 *
 * 输出: 6
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

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	// 判断左右子树高度
	leftLevel := getLevel(root.Left, 0)
	rightLevel := getLevel(root.Right, 0)
	// 如果二者相等，说明最后一层已经到了右子树，则加上左子树的值之后，计算右子树
	if leftLevel == rightLevel {
		res = pow(2, leftLevel) + countNodes(root.Right)
	} else {
		res = pow(2, rightLevel) + countNodes(root.Left)
	}
	return res
}

func getLevel(root *TreeNode, prev int) int {
	if root == nil {
		return prev
	}
	// 将本层计入
	prev += 1

	left := getLevel(root.Left, prev)
	right := getLevel(root.Right, prev)
	if left > right {
		return left
	}
	return right
}

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	res := 1
	for n > 0 {
		res *= x
		n--
	}
	return res
}

// @lc code=end


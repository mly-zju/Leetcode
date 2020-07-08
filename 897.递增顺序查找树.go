/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序查找树
 *
 * https://leetcode-cn.com/problems/increasing-order-search-tree/description/
 *
 * algorithms
 * Easy (61.48%)
 * Likes:    39
 * Dislikes: 0
 * Total Accepted:    5.3K
 * Total Submissions: 8.5K
 * Testcase Example:  '[5,3,6,2,4,null,8,1,null,null,null,7,9]'
 *
 * 给定一个树，按中序遍历重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点。
 *
 *
 *
 * 示例 ：
 *
 * 输入：[5,3,6,2,4,null,8,1,null,null,null,7,9]
 *
 * ⁠      5
 * ⁠     / \
 * ⁠   3    6
 * ⁠  / \    \
 * ⁠ 2   4    8
 * /        / \
 * 1        7   9
 *
 * 输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
 *
 * ⁠1
 * \
 * 2
 * \
 * 3
 * \
 * 4
 * \
 * 5
 * \
 * 6
 * \
 * 7
 * \
 * 8
 * \
 * ⁠                9
 *
 *
 *
 * 提示：
 *
 *
 * 给定树中的结点数介于 1 和 100 之间。
 * 每个结点都有一个从 0 到 1000 范围内的唯一整数值。
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
var newRoot *TreeNode
func increasingBST(root *TreeNode) *TreeNode {
	// 1. 迭代
	// 中序遍历 + 前驱节点重建
	// cache := []*TreeNode{}
	// var prevNode *TreeNode
	// var newRoot *TreeNode
	// for root != nil || len(cache) != 0 {
	// 	for root != nil {
	// 		cache = append(cache, root)
	// 		root = root.Left
	// 	}
	// 	newLen := len(cache)
	// 	root = cache[newLen-1]
	// 	if newRoot == nil {
	// 		newRoot = root
	// 	}
	// 	// 每个节点丢弃左子树
	// 	root.Left = nil
	// 	if prevNode == nil {
	// 		prevNode = root
	// 	} else {
	// 		prevNode.Right = root
	// 		prevNode = root
	// 	}
	// 	root = root.Right
	// 	cache = cache[0:newLen-1]
	// }
	// return newRoot

	// 2. 递归
	newRoot = nil
	var prevNode *TreeNode
	dfs(root, prevNode)
	return newRoot
}

func dfs(root, prevNode *TreeNode) *TreeNode {
	if root == nil {
		return prevNode
	}

	prevNode = dfs(root.Left, prevNode)

	if prevNode == nil {
		prevNode = root
		newRoot = prevNode
	} else {
		prevNode.Left = nil
		prevNode.Right = root
	}

	prevNode = root
	prevNode = dfs(root.Right, prevNode)
	return prevNode
}

// @lc code=end


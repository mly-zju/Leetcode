/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
 *
 * https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (54.20%)
 * Likes:    90
 * Dislikes: 0
 * Total Accepted:    10.2K
 * Total Submissions: 18.2K
 * Testcase Example:  '[1,null,3,2]'
 *
 * 给定一个所有节点为非负值的二叉搜索树，求树中任意两节点的差的绝对值的最小值。
 * 
 * 示例 :
 * 
 * 
 * 输入:
 * 
 * ⁠  1
 * ⁠   \
 * ⁠    3
 * ⁠   /
 * ⁠  2
 * 
 * 输出:
 * 1
 * 
 * 解释:
 * 最小绝对差为1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
 * 
 * 
 * 注意: 树中至少有2个节点。
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
var minVal int
func getMinimumDifference(root *TreeNode) int {
	// 中序遍历的是排序值，绝对值肯定在相邻元素中间

	// -1代表未初始化
	// 1. 递归法
	// minVal = -1
	// dfs(root, nil)
	// return minVal

	// 2. 迭代法
	minVal = -1
	var prevNode *TreeNode
	cache := []*TreeNode{}
	for root != nil || len(cache) > 0 {
		for root != nil {
			cache = append(cache, root)
			root = root.Left
		}
		length := len(cache)
		root = cache[length-1]
		cache = cache[0:length-1]

		if prevNode == nil {
			prevNode = root
		} else if minVal == -1 || getAbs(prevNode.Val - root.Val) < minVal {
			minVal = getAbs(prevNode.Val - root.Val)
		}
		prevNode = root
		root = root.Right
	}
	return minVal
}

func dfs(root, prevNode *TreeNode) *TreeNode {
	if root == nil {
		return prevNode
	}

	prevNode = dfs(root.Left, prevNode)
	// 更新minVal的条件：prevNode必须不为nil，且(minVal未初始化 或者 比minVal小)

	if prevNode != nil && (minVal == -1 || getAbs(prevNode.Val - root.Val) < minVal) {
		minVal = getAbs(prevNode.Val - root.Val)
	}
	prevNode = root
	prevNode = dfs(root.Right, prevNode)
	return prevNode
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
// @lc code=end


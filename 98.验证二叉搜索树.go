/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (27.62%)
 * Likes:    318
 * Dislikes: 0
 * Total Accepted:    48.8K
 * Total Submissions: 175.7K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 *
 * 假设一个二叉搜索树具有如下特征：
 *
 *
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 *
 * 示例 1:
 *
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
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
// var validFlag bool
func isValidBST(root *TreeNode) bool {
	// 中序遍历+前驱节点，只要保证总是小于前驱即可
	// 迭代或者递归都行
	// 递归的话，需要修改全局变量了

	// 1. 递归
	// validFlag = true
	// dfs(root, nil)
	// return validFlag

	// 2. 迭代
	cache := []*TreeNode{}
	var prevNode *TreeNode
	for root != nil || len(cache) != 0 {
		for root != nil {
			cache = append(cache, root)
			root = root.Left
		}
		newLen := len(cache)
		root = cache[newLen - 1]
		if prevNode != nil && prevNode.Val >= root.Val {
			return false
		}
		prevNode = root
		cache = cache[0:newLen-1]
		root = root.Right
	}
	return true
}

/*
func dfs(root, prevNode *TreeNode) *TreeNode {
	if root == nil {
		return prevNode
	}

	prevNode = dfs(root.Left, prevNode)
	if prevNode != nil && prevNode.Val >= root.Val {
		validFlag = false
	}
	prevNode = root
	prevNode = dfs(root.Right, prevNode)

	return prevNode
}
*/

// @lc code=end


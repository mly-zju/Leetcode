/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (64.19%)
 * Likes:    125
 * Dislikes: 0
 * Total Accepted:    17K
 * Total Submissions: 26.1K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	length := len(inorder)
	if length == 0 {
		return nil
	}

	mid := 0
	// 寻找根节点位置
	for index, val := range inorder {
		if val == postorder[length-1] {
			mid = index
			break
		}
	}

	return &TreeNode{
		Val: inorder[mid],
		Left: buildTree(inorder[0:mid], postorder[0:mid]),
		Right: buildTree(inorder[mid+1:], postorder[mid:length-1]),
	}
}

// @lc code=end


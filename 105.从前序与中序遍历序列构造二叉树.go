/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (61.51%)
 * Likes:    270
 * Dislikes: 0
 * Total Accepted:    29.9K
 * Total Submissions: 47.9K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 性质：前序节点的第一个元素决定了根节点；中序节点的中间点决定了左右子树的划分；
	length := len(inorder)
	if length == 0 {
		return nil
	}

	// 寻找根节点，对于pre来说，第一个就是, 寻找中序的二分节点
	mid := 0
	for index, val := range inorder {
		if val == preorder[0] {
			mid = index
			break
		}
	}

	return &TreeNode{
		Val:   inorder[mid],
		Left:  buildTree(preorder[1:mid+1], inorder[0:mid]),
		Right: buildTree(preorder[mid+1:], inorder[mid+1:]),
	}
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 *
 * https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (65.03%)
 * Likes:    202
 * Dislikes: 0
 * Total Accepted:    17.6K
 * Total Submissions: 26.6K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * 给定一个二叉树，原地将它展开为链表。
 *
 * 例如，给定二叉树
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   5
 * ⁠/ \   \
 * 3   4   6
 *
 * 将其展开为：
 *
 * 1
 * ⁠\
 * ⁠ 2
 * ⁠  \
 * ⁠   3
 * ⁠    \
 * ⁠     4
 * ⁠      \
 * ⁠       5
 * ⁠        \
 * ⁠         6
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
func flatten(root *TreeNode) {
	// 要用原地算法。从题目看，顺序是前序的顺序，但是不能直接用前序遍历，会破坏链接结构，可以用右子树优先的后序遍历

	// 1. 递归
	// var prevNode *TreeNode
	// dfs(root, prevNode)

	// 2. 迭代
	cache := []*TreeNode{}
	visitMap := map[*TreeNode]bool{}
	var prevNode *TreeNode
	for root != nil || len(cache) != 0 {
		for visitMap[root] != true && root != nil {
			cache = append(cache, root)
			root = root.Right
		}
		newLen := len(cache)
		root = cache[newLen - 1]
		cache = cache[0:newLen - 1]
		if visitMap[root] {
			// 如果已经访问过，才访问自己本身
			root.Left = nil
			root.Right = prevNode
			prevNode = root
			root = nil
		} else {
			visitMap[root] = true
			// left已经先访问了，因此可以修改为nil没关系
			cache = append(cache, root)
			root = root.Left
		}
	}
}

func dfs(root, prevNode *TreeNode) *TreeNode {
	if root == nil {
		return prevNode
	}
	// 右子树优先的后序
	prevNode = dfs(root.Right, prevNode)
	prevNode = dfs(root.Left, prevNode)
	if prevNode != nil {
		root.Right = prevNode
		root.Left = nil
	}
	prevNode = root
	return prevNode
}

// @lc code=end


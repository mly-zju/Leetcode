/*
 * @lc app=leetcode.cn id=572 lang=golang
 *
 * [572] 另一个树的子树
 *
 * https://leetcode-cn.com/problems/subtree-of-another-tree/description/
 *
 * algorithms
 * Easy (41.74%)
 * Likes:    124
 * Dislikes: 0
 * Total Accepted:    9.5K
 * Total Submissions: 22.7K
 * Testcase Example:  '[3,4,5,1,2]\n[4,1,2]'
 *
 * 给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s
 * 也可以看做它自身的一棵子树。
 *
 * 示例 1:
 * 给定的树 s:
 *
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \
 * ⁠1   2
 *
 *
 * 给定的树 t：
 *
 *
 * ⁠  4
 * ⁠ / \
 * ⁠1   2
 *
 *
 * 返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。
 *
 * 示例 2:
 * 给定的树 s：
 *
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \
 * ⁠1   2
 * ⁠   /
 * ⁠  0
 *
 *
 * 给定的树 t：
 *
 *
 * ⁠  4
 * ⁠ / \
 * ⁠1   2
 *
 *
 * 返回 false。
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
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		if t == nil {
			return true
		}
		return false
	}

	return isSame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s, t *TreeNode) bool {
	if s == nil && t == nil {
		// 双方都空，则相等
		return true
	} else if s == nil || t == nil {
		// 否则如果只有一方为空，则不相等
		return false
	}

	// 如果双方都不为空
	if s.Val != t.Val {
		return false
	}
	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}

// @lc code=end


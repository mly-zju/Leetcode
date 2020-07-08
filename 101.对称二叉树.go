/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (48.53%)
 * Likes:    631
 * Dislikes: 0
 * Total Accepted:    99.6K
 * Total Submissions: 198.3K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 * 
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 * 
 * 
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 * 
 * 
 * 说明:
 * 
 * 如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 1. 递归
	// return dfs(root.Left, root.Right)

	// 2. bfs迭代法
	cache := []*TreeNode{root}
	for len(cache) != 0 {
		// 每一层都先直接判别是否对称相等
		levelLen := len(cache)
		l, r := 0, levelLen - 1
		for l < r {
			if cache[l] == nil && cache[r] == nil {
				l++
				r--
			} else if cache[l] == nil || cache[r] == nil {
				return false
			} else if cache[l].Val == cache[r].Val {
				l++
				r--
			} else {
				return false
			}
		}

		for i := 0; i < levelLen; i++ {
			if cache[i] != nil {
				cache = append(cache, cache[i].Left)
				cache = append(cache, cache[i].Right)
			}
		}
		cache = cache[levelLen:]
	}
	return true
}

func dfs(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return dfs(t1.Left, t2.Right) && dfs(t1.Right, t2.Left)
}
// @lc code=end


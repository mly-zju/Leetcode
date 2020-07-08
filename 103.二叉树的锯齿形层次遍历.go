/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (51.82%)
 * Likes:    112
 * Dislikes: 0
 * Total Accepted:    22.5K
 * Total Submissions: 42.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回锯齿形层次遍历如下：
 *
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
 * ]
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	cache := []*TreeNode{}
	res := [][]int{}
	if root == nil {
		return res
	}

	cache = append(cache, root)
	flag := 0
	for len(cache) != 0 {
		// 记录每一层的长度
		levelLen := len(cache)
		levelArr := []int{}

		for i := 0; i < levelLen; i++ {
			levelArr = append(levelArr, cache[i].Val)
			if cache[i].Left != nil {
				cache = append(cache, cache[i].Left)
			}
			if cache[i].Right != nil {
				cache = append(cache, cache[i].Right)
			}
		}

		// 根据flag判断迭代顺序, 如果为1需要执行反转
		if flag % 2 == 1 {
			l, r := 0, levelLen - 1
			for l < r {
				levelArr[l], levelArr[r] = levelArr[r], levelArr[l]
				l++
				r--
			}
		}
		res = append(res, levelArr)

		// 去掉父层级
		cache = cache[levelLen:]
		flag++
	}

	return res
}

// @lc code=end


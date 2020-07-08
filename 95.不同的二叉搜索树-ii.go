/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (59.22%)
 * Likes:    215
 * Dislikes: 0
 * Total Accepted:    13.7K
 * Total Submissions: 22.8K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * 解释:
 * 以上的输出对应以下 5 种不同结构的二叉搜索树：
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
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
func generateTrees(n int) []*TreeNode {
	return dfs(1, n)
}

func dfs(l, r int) []*TreeNode {
	res := []*TreeNode{}
	if l > r {
		return res
	} else if l == r {
		res = append(res, &TreeNode{
			Val: l,
		})
		return res
	}

	// 根节点存在各种可能性
	for mid := l; mid <= r; mid++ {
		leftTrees, rightTrees := dfs(l, mid-1), dfs(mid+1, r)
		leftLen, rightLen := len(leftTrees), len(rightTrees)
		// 开始执行组合, 需要防止外层为空导致无法获取
		if leftLen == 0 && rightLen == 0 {
			return res
		} else if leftLen == 0 {
			for i := 0; i < rightLen; i++ {
				res = append(res, &TreeNode{
					Val: mid,
					Right: rightTrees[i],
				})
			}
		} else if rightLen == 0 {
			for i := 0; i < leftLen; i++ {
				res = append(res, &TreeNode{
					Val: mid,
					Left: leftTrees[i],
				})
			}
		} else {
			for i := 0; i < leftLen; i++ {
				for j := 0; j < rightLen; j++ {
					res = append(res, &TreeNode{
						Val: mid,
						Left: leftTrees[i],
						Right: rightTrees[j],
					})
				}
			}
		}
	}

	return res
}

// @lc code=end


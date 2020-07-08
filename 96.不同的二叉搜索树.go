/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (62.28%)
 * Likes:    302
 * Dislikes: 0
 * Total Accepted:    20.6K
 * Total Submissions: 32.6K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
 *
 * 示例:
 *
 * 输入: 3
 * 输出: 5
 * 解释:
 * 给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 */

// @lc code=start
func numTrees(n int) int {
	return dfs(1, n)
}

func dfs(start, end int) int {
	if start > end {
		return 0
	} else if start == end {
		return 1
	}

	res := 0
	for mid := start; mid <= end; mid++ {
		leftNum := dfs(start, mid-1)
		rightNum := dfs(mid+1, end)
		if leftNum > 0 && rightNum > 0 {
			res += leftNum * rightNum
		} else {
			res += leftNum + rightNum
		}
	}
	return res
}

// @lc code=end


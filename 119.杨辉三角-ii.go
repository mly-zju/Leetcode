/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (57.95%)
 * Likes:    88
 * Dislikes: 0
 * Total Accepted:    27.1K
 * Total Submissions: 46.3K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 3
 * 输出: [1,3,3,1]
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(k) 空间复杂度吗？
 *
 */

// @lc code=start
func getRow(rowIndex int) []int {
	res := []int{1}
	tmp := 1
	for i := 1; i <= rowIndex; i++ {
		for j := 1; j < i; j++ {
			prev := tmp
			// 每一轮缓存当前元素，下一轮用
			tmp = res[j]
			res[j] = prev + tmp
		}
		res = append(res, 1)
	}
	return res
}

// @lc code=end


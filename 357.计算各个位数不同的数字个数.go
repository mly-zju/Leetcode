/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 计算各个位数不同的数字个数
 *
 * https://leetcode-cn.com/problems/count-numbers-with-unique-digits/description/
 *
 * algorithms
 * Medium (46.80%)
 * Likes:    42
 * Dislikes: 0
 * Total Accepted:    5.3K
 * Total Submissions: 10.9K
 * Testcase Example:  '2'
 *
 * 给定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10^n 。
 *
 * 示例:
 *
 * 输入: 2
 * 输出: 91
 * 解释: 答案应为除去 11,22,33,44,55,66,77,88,99 外，在 [0,100) 区间内的所有数字。
 *
 *
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	// dp[i-1]等于dp[i]情况下前导位为0的情况
	dp := []int{}
}

// @lc code=end


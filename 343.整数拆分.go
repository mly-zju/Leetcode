/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 *
 * https://leetcode-cn.com/problems/integer-break/description/
 *
 * algorithms
 * Medium (52.48%)
 * Likes:    130
 * Dislikes: 0
 * Total Accepted:    13.5K
 * Total Submissions: 25.4K
 * Testcase Example:  '2'
 *
 * 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
 *
 * 示例 1:
 *
 * 输入: 2
 * 输出: 1
 * 解释: 2 = 1 + 1, 1 × 1 = 1。
 *
 * 示例 2:
 *
 * 输入: 10
 * 输出: 36
 * 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
 *
 * 说明: 你可以假设 n 不小于 2 且不大于 58。
 *
 */

// @lc code=start
func integerBreak(n int) int {
	// dp := []int{}
	// for i := 0; i <= n; i++ {
	// 	if i < 2 {
	// 		dp = append(dp, i)
	// 	} else {
	// 		tmpMax := 0
	// 		for j := 1; j < i; j++ {
	// 			tmpNum := getMax(dp[i-j]*j, (i-j)*j)
	// 			if tmpNum > tmpMax {
	// 				tmpMax = tmpNum
	// 			}
	// 		}
	// 		dp = append(dp, tmpMax)
	// 	}
	// }
	// return dp[n]

	// dp[i]代表正整数i可以得到的最大乘积和
	dp := []int{0}
	for i := 1; i <= n; i++ {
		tmpMax := 0
		for j := 1; j < i; j++ {
			thisMax := getMax(dp[i-j]*j, (i-j)*j)
			if thisMax > tmpMax {
				tmpMax = thisMax
			}
		}
		dp = append(dp, tmpMax)
	}
	return dp[n]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end


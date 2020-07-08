package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 *
 * https://leetcode-cn.com/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (50.70%)
 * Likes:    254
 * Dislikes: 0
 * Total Accepted:    28.4K
 * Total Submissions: 55.1K
 * Testcase Example:  '12'
 *
 * 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
 *
 * 示例 1:
 *
 * 输入: n = 12
 * 输出: 3
 * 解释: 12 = 4 + 4 + 4.
 *
 * 示例 2:
 *
 * 输入: n = 13
 * 输出: 2
 * 解释: 13 = 4 + 9.
 *
 */

// @lc code=start
func numSquares(n int) int {
	// dp[i]代表组成i的完全平方数最小个数，则对于n, 遍历找到最小的dp[i - j * j] + 1
	if n == 1 {
		return 1
	}

	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		tmpMin := i
		for j := 1; j*j <= i; j++ {
			tmp := dp[i-j*j] + 1
			if tmp < tmpMin {
				tmpMin = tmp
			}
		}
		dp = append(dp, tmpMin)
	}

	return dp[n]
}

func getMax(n int) int {
	return int(math.Floor(math.Sqrt(float64(n))))
}

// func main() {
// 	x := 12
// 	fmt.Println(numSquares(x))
// }

// @lc code=end

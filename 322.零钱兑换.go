/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (35.50%)
 * Likes:    294
 * Dislikes: 0
 * Total Accepted:    32.3K
 * Total Submissions: 88.6K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 *
 * 示例 1:
 *
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3
 * 解释: 11 = 5 + 5 + 1
 *
 * 示例 2:
 *
 * 输入: coins = [2], amount = 3
 * 输出: -1
 *
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 *
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	length := len(coins)
	if length == 0 {
		return -1
	}

	sort.Ints(coins)
	// dp[i][j]是coins[0~i]个凑j的最小银币个数
	// dp[i][j] = getMin(dp[i-1][j], dp[i][j-coins[i]] + 1)
	dp := make([]int, amount + 1)
	// 出了dp[0], 其他初始化最大值
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 0; i < length; i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = getMin(dp[j], dp[j-coins[i]] + 1)
		}
	}
	if dp[amount] == amount + 1 {
		return -1
	}
	return dp[amount]
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end


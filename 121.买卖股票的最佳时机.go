/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (50.73%)
 * Likes:    588
 * Dislikes: 0
 * Total Accepted:    88.3K
 * Total Submissions: 172.7K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 *
 * 如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
 *
 * 注意你不能在买入股票前卖出股票。
 *
 * 示例 1:
 *
 * 输入: [7,1,5,3,6,4]
 * 输出: 5
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
 *
 *
 * 示例 2:
 *
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 *
 */

// @lc code=start
func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}

	// 基于状态机的动态规划：
	// sale<-hold/sale, hold<-free/hold, free只能有初始态一次
	var dpSale, dpHold, dpFree int
	for i := 0; i < length; i++ {
		// 无论何时，dpFree只能有一次初始态，那就是0
		dpFree = 0
		if i == 0 {
			dpHold = -prices[i]
			dpSale = 0
		} else {
			ndpHold := getMax(dpFree-prices[i], dpHold)
			ndpSale := getMax(dpSale, dpHold+prices[i])

			dpHold = ndpHold
			dpSale = ndpSale
		}
	}

	return getMax(dpFree, dpSale)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end


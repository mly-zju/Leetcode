func maxProfit(prices []int) int {
	// 每天卖的最大值等于该天减去之前的最小值
	max := 0
	for i, length := 1, len(prices); i < length; i++ {
		for j := 0; j < i; j++ {
			tmp := prices[i] - prices[j]
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}
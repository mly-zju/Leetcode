func numWays(n int) int {
	const MAX_NUM = 1e9 + 7
	dp := []int{}
	for i := 0; i <= n; i++ {
		if i <= 1 {
			dp = append(dp, 1)
		} else {
			dp = append(dp, (dp[i - 1] + dp[i - 2]) % MAX_NUM)
		}
	}
	return dp[n]
}
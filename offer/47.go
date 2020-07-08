func maxValue(grid [][]int) int {
	rNum := len(grid)
	if rNum == 0 {
		return 0
	}
	cNum := len(grid[0])
	if cNum == 0 {
		return 0
	}

	// 由于只和上一层和本层有关系，所以可以一维数组
	// dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + valud[i][j]
	dp := make([]int, cNum)
	// 开始查看
	for i := 0; i < rNum; i++ {
		for j := 0; j < cNum; j++ {
			if i == 0 && j == 0 {
				dp[j] = grid[i][j]
			} else if i == 0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				dp[j] = getMax(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[cNum - 1]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
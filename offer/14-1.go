func cuttingRope(n int) int {
	// dp[i]代表i米长时候的最大乘积
	dp := []int{0}
	for i := 1; i <= n; i++ {
		if i == 1 || i == 2 {
			dp = append(dp, 1)
		} else {
			tmpMax := 0
			for j := 1; j <= (i + 1) / 2; j++ {
				a := dp[j] * dp[i-j]
				b := dp[j] * (i-j)
				c := j * dp[i-j]
				d := j * (i-j)
				newMax := getMax(getMax(a, b), getMax(c, d))
				tmpMax = getMax(newMax, tmpMax)
			}
			dp = append(dp, tmpMax)
		}
	}
	return dp[n]
}

func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func translateNum(num int) int {
	// 动态规划，类似跳台阶问题
	dp := []int{}
	str := strconv.Itoa(num)
	length := len(str)
	for i := 0; i < length; i++ {
		if i == 0 {
			dp = append(dp, 1)
		} else {
			prevNum := int(str[i - 1] - '0')
			curNum := int(str[i] - '0')
			twoNum := prevNum * 10 + curNum
			if twoNum <= 25 && prevNum != 0 {
				if i == 1 {
					dp = append(dp, dp[i - 1] + 1)
				} else {
					dp = append(dp, dp[i - 1] + dp[i - 2])
				}
			} else {
				dp = append(dp, dp[i - 1])
			}
		}
	}
	return dp[length - 1]
}
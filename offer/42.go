func maxSubArray(nums []int) int {
	// dp[i]代表以i结尾的数组中最大和
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp := make([]int, length)
	max := nums[0]
	for i := 0; i < length; i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else {
			if dp[i - 1] > 0 {
				dp[i] = dp[i-1] + nums[i]
			} else {
				dp[i] = nums[i]
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
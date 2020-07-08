package main

/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 *
 * https://leetcode-cn.com/problems/partition-equal-subset-sum/description/
 *
 * algorithms
 * Medium (44.38%)
 * Likes:    166
 * Dislikes: 0
 * Total Accepted:    13.6K
 * Total Submissions: 31.1K
 * Testcase Example:  '[1,5,11,5]'
 *
 * 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
 *
 * 注意:
 *
 *
 * 每个数组中的元素不会超过 100
 * 数组的大小不会超过 200
 *
 *
 * 示例 1:
 *
 * 输入: [1, 5, 11, 5]
 *
 * 输出: true
 *
 * 解释: 数组可以分割成 [1, 5, 5] 和 [11].
 *
 *
 *
 *
 * 示例 2:
 *
 * 输入: [1, 2, 3, 5]
 *
 * 输出: false
 *
 * 解释: 数组不能分割成两个元素和相等的子集.
 *
 *
 *
 *
 */

// @lc code=start
func canPartition(nums []int) bool {
	// 0-1背包问题，求最大值能否正好等于
	length := len(nums)
	if length == 0 {
		return false
	}

	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2

	// dp[i][j]代表资源从0~i且背包容量为j时候的最大值，
	// dp[i][j] = getMax(dp[i-1][j], dp[i-1][j-room[i]] + value[i])
	// 这里room[i]和value[i]都等于nums[i]本身
	dp := make([]int, sum+1)

	for i := 0; i < length; i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] = getMax(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[sum] == sum
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 *
 * https://leetcode-cn.com/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (34.67%)
 * Likes:    152
 * Dislikes: 0
 * Total Accepted:    15.4K
 * Total Submissions: 43.6K
 * Testcase Example:  '[2,3,2]'
 *
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
 *
 * 示例 1:
 *
 * 输入: [2,3,2]
 * 输出: 3
 * 解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
 *
 *
 * 示例 2:
 *
 * 输入: [1,2,3,1]
 * 输出: 4
 * 解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 */

// @lc code=start
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	}

	// 可以搞成两个动态规划：nums[1:], nums[0:length - 1]
	res1 := _rob(nums[1:])
	res2 := _rob(nums[0 : len(nums)-1])
	if res1 > res2 {
		return res1
	}
	return res2
}

func _rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	dp := []int{nums[0]}
	for i := 1; i < length; i++ {
		prev := dp[i-1]
		cur := nums[i]
		if i-2 >= 0 {
			cur += dp[i-2]
		}

		if cur > prev {
			dp = append(dp, cur)
		} else {
			dp = append(dp, prev)
		}
	}

	return dp[length-1]
}

// @lc code=end


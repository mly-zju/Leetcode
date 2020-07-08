/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (47.49%)
 * Likes:    1300
 * Dislikes: 0
 * Total Accepted:    111.1K
 * Total Submissions: 232.4K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 示例:
 *
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 *
 *
 * 进阶:
 *
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 *
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// max := 0
	// sum := []int{}
	// for index, val := range nums {
	// 	curSum := 0
	// 	if index == 0 {
	// 		sum = append(sum, val)
	// 		max = val
	// 	} else {
	// 		if val > val+sum[index-1] {
	// 			curSum = val
	// 		} else {
	// 			curSum = val + sum[index-1]
	// 		}
	// 		sum = append(sum, curSum)

	// 		if curSum > max {
	// 			max = curSum
	// 		}
	// 	}
	// }
	// return max

	length := len(nums)
	if length == 0 {
		return 0
	}

	// dp[i]代表以i结尾的子树组最大和,
	dp := []int{}
	max := nums[0]
	for i := 0; i < length; i++ {
		if i == 0 {
			dp = append(dp, nums[i])
		} else {
			if dp[i-1] > 0 {
				dp = append(dp, dp[i-1]+nums[i])
			} else {
				dp = append(dp, nums[i])
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

// @lc code=end


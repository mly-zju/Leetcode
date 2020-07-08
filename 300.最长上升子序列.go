/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 *
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (43.28%)
 * Likes:    361
 * Dislikes: 0
 * Total Accepted:    39K
 * Total Submissions: 89.7K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 *
 * 示例:
 *
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 *
 * 说明:
 *
 *
 * 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
 * 你算法的时间复杂度应该为 O(n^2) 。
 *
 *
 * 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
 *
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	// 1. O(n^2)的dp算法
	length := len(nums)
	if length == 0 {
		return 0
	}

	// dp[i]代表以i结尾时候的最大升序子序列长度
	dp := []int{}
	res := 0
	for i := 0; i < length; i++ {
		maxLen := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > maxLen {
				maxLen = dp[j] + 1
			}
		}
		if maxLen > res {
			res = maxLen
		}
		dp = append(dp, maxLen)
	}

	return res
}

// @lc code=end


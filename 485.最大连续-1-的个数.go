/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 *
 * https://leetcode-cn.com/problems/max-consecutive-ones/description/
 *
 * algorithms
 * Easy (54.59%)
 * Likes:    66
 * Dislikes: 0
 * Total Accepted:    20.3K
 * Total Submissions: 37.1K
 * Testcase Example:  '[1,0,1,1,0,1]'
 *
 * 给定一个二进制数组， 计算其中最大连续1的个数。
 *
 * 示例 1:
 *
 *
 * 输入: [1,1,0,1,1,1]
 * 输出: 3
 * 解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
 *
 *
 * 注意：
 *
 *
 * 输入的数组只包含 0 和1。
 * 输入数组的长度是正整数，且不超过 10,000。
 *
 *
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	// 1. 遍历+状态计数器
	// length := len(nums)
	// if length == 0 {
	// 	return 0
	// }

	// curNum := 0
	// max := 0
	// for _, val := range nums {
	// 	if val == 1 {
	// 		curNum++
	// 	} else if val == 0 {
	// 		if curNum > max {
	// 			max = curNum
	// 		}
	// 		curNum = 0
	// 	}
	// }
	// // 还要看最后一个元素
	// if curNum > max {
	// 	max = curNum
	// }
	// return max

	// 2. 动态规划
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp := []int{}
	max := 0
	for i := 0; i < length; i++ {
		if i == 0 {
			if nums[i] == 1 {
				dp = append(dp, 1)
				max = 1
			} else {
				dp = append(dp, 0)
			}
		} else {
			if nums[i] == 0 {
				dp = append(dp, 0)
			} else {
				cur := dp[i-1] + 1
				dp = append(dp, cur)
				if cur > max {
					max = cur
				}
			}
		}
	}
	return max
}

// @lc code=end


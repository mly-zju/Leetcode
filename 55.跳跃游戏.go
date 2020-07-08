/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 *
 * https://leetcode-cn.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (36.17%)
 * Likes:    374
 * Dislikes: 0
 * Total Accepted:    44.9K
 * Total Submissions: 121.5K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 *
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 *
 * 判断你是否能够到达最后一个位置。
 *
 * 示例 1:
 *
 * 输入: [2,3,1,1,4]
 * 输出: true
 * 解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
 *
 *
 * 示例 2:
 *
 * 输入: [3,2,1,0,4]
 * 输出: false
 * 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
 *
 *
 */

// @lc code=start
func canJump(nums []int) bool {
	// 动态规划，通过dp[i]如果能达到最后，是一个好坐标，否则不是,
	// 则判断dp[i]是否好坐标，要看dp[i + 1]...dp[i + val]里面是否包含好坐标
	length := len(nums)
	dp := []bool{}
	// for i := 0; i < length; i++ {
	// 	if i == 0 {
	// 		dp = append(dp, true)
	// 	} else {
	// 		// val需要倒着取
	// 		val := nums[length-1-i]
	// 		flag := false
	// 		for j := 1; j <= val; j++ {
	// 			if i-j >= 0 && dp[i-j] {
	// 				flag = true
	// 				break
	// 			}
	// 		}
	// 		dp = append(dp, flag)
	// 	}
	// }
	// return dp[len(dp)-1]

	// dp[i]代表位置i能否实现
	for i := 0; i < length; i++ {
		if i == 0 {
			dp = append(dp, true)
		} else {
			ans := false
			for j := 0; j < i; j++ {
				if dp[j] && j+nums[j] >= i {
					ans = true
					break
				}
			}
			dp = append(dp, ans)
		}
	}

	return dp[length-1]
}

// @lc code=end


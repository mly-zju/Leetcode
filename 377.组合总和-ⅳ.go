/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 *
 * https://leetcode-cn.com/problems/combination-sum-iv/description/
 *
 * algorithms
 * Medium (39.59%)
 * Likes:    93
 * Dislikes: 0
 * Total Accepted:    5.7K
 * Total Submissions: 14.2K
 * Testcase Example:  '[1,2,3]\n4'
 *
 * 给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。
 *
 * 示例:
 *
 *
 * nums = [1, 2, 3]
 * target = 4
 *
 * 所有可能的组合为：
 * (1, 1, 1, 1)
 * (1, 1, 2)
 * (1, 2, 1)
 * (1, 3)
 * (2, 1, 1)
 * (2, 2)
 * (3, 1)
 *
 * 请注意，顺序不同的序列被视作不同的组合。
 *
 * 因此输出为 7。
 *
 *
 * 进阶：
 * 如果给定的数组中含有负数会怎么样？
 * 问题会产生什么变化？
 * 我们需要在题目中添加什么限制来允许负数的出现？
 *
 * 致谢：
 * 特别感谢 @pbrother 添加此问题并创建所有测试用例。
 *
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if target <= 0 {
		return 0
	}

	dp := []int{1}
	for i := 1; i <= target; i++ {
		tmpSum := 0
		for j := 0; j < length; j++ {
			nTarget := i - nums[j]
			if nTarget >= 0 {
				tmpSum += dp[nTarget]
			}
		}
		dp = append(dp, tmpSum)
	}

	return dp[target]
}

// @lc code=end


import "sort"

/*
 * @lc app=leetcode.cn id=561 lang=golang
 *
 * [561] 数组拆分 I
 *
 * https://leetcode-cn.com/problems/array-partition-i/description/
 *
 * algorithms
 * Easy (67.36%)
 * Likes:    121
 * Dislikes: 0
 * Total Accepted:    21.6K
 * Total Submissions: 31.8K
 * Testcase Example:  '[1,4,3,2]'
 *
 * 给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到
 * n 的 min(ai, bi) 总和最大。
 *
 * 示例 1:
 *
 *
 * 输入: [1,4,3,2]
 *
 * 输出: 4
 * 解释: n 等于 2, 最大总和为 4 = min(1, 2) + min(3, 4).
 *
 *
 * 提示:
 *
 *
 * n 是正整数,范围在 [1, 10000].
 * 数组中的元素范围在 [-10000, 10000].
 *
 *
 */

// @lc code=start
func arrayPairSum(nums []int) int {
	// 贪心算法，肯定要找尽可能大的，所以从大到小排序，取奇数index的相加即可
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	sum := 0
	for i, length := 1, len(nums); i < length; i += 2 {
		sum += nums[i]
	}
	return sum
}

// @lc code=end


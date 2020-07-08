import "sort"

/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 *
 * https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/description/
 *
 * algorithms
 * Easy (33.02%)
 * Likes:    160
 * Dislikes: 0
 * Total Accepted:    10.7K
 * Total Submissions: 32.3K
 * Testcase Example:  '[2,6,4,8,10,9,15]'
 *
 * 给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
 *
 * 你找到的子数组应是最短的，请输出它的长度。
 *
 * 示例 1:
 *
 *
 * 输入: [2, 6, 4, 8, 10, 9, 15]
 * 输出: 5
 * 解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
 *
 *
 * 说明 :
 *
 *
 * 输入的数组长度范围在 [1, 10,000]。
 * 输入的数组可能包含重复元素 ，所以升序的意思是<=。
 *
 *
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	// 1. 方法1： 排序，然后对比最开始不匹配和最后不匹配的index，计算个数
	sortNums := make([]int, len(nums))
	copy(sortNums, nums)
	sort.Ints(sortNums)
	start, end := -1, -1
	for i, length := 0, len(nums); i < length; i++ {
		if nums[i] != sortNums[i] {
			if start == -1 {
				start = i
			} else {
				end = i
			}
		}
	}
	if start == -1 {
		return 0
	}
	return end - start + 1
}

// @lc code=end


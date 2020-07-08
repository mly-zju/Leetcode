package main

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (39.75%)
 * Likes:    151
 * Dislikes: 0
 * Total Accepted:    20.5K
 * Total Submissions: 51K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
 *
 * 示例:
 *
 * 输入: s = 7, nums = [2,3,1,2,4,3]
 * 输出: 2
 * 解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
 *
 *
 * 进阶:
 *
 * 如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 *
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	// 滑块法
	l, r := 0, 0
	minLen := length + 1
	sum := nums[0]
	for l <= r && r < length {
		if sum < s {
			r++
			if r < length {
				sum += nums[r]
			}
		} else {
			// 先计算当前长度
			tmpLen := r - l + 1
			if tmpLen < minLen {
				minLen = tmpLen
			}
			sum -= nums[l]
			l++
		}
	}
	if minLen == length+1 {
		return 0
	}
	return minLen
}

// func main() {
// 	s := 15
// 	nums := []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}
// 	fmt.Println(minSubArrayLen(s, nums))
// }

// @lc code=end

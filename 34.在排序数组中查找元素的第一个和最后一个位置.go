/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (37.69%)
 * Likes:    251
 * Dislikes: 0
 * Total Accepted:    45.6K
 * Total Submissions: 119.3K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 如果数组中不存在目标值，返回 [-1, -1]。
 *
 * 示例 1:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: [3,4]
 *
 * 示例 2:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: [-1,-1]
 *
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	// 不仅仅是寻找到就截止，而是要寻找上界、下界
	length := len(nums)
	if length == 0 {
		return []int{-1, -1}
	}
	// l, r := 0, length-1
	low, high := binarySearch(nums, target, false), binarySearch(nums, target, true)
	return []int{low, high}
}

func binarySearch(nums []int, target int, isHigh bool) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	l, r := 0, length - 1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			// 处理上下界
			if isHigh {
				if mid + 1 < length && nums[mid+1] == target {
					// 如果上界存在，继续寻找
					l = mid + 1
				} else {
					return mid
				}
			} else {
				if mid - 1 >= 0 && nums[mid-1] == target {
					r = mid - 1
				} else {
					return mid
				}
			}
		}
	}
	return -1
}

// @lc code=end


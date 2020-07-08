/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (36.10%)
 * Likes:    418
 * Dislikes: 0
 * Total Accepted:    55.9K
 * Total Submissions: 154.4K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
 *
 * 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 *
 * 你可以假设数组中不存在重复的元素。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 示例 1:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 0
 * 输出: 4
 *
 *
 * 示例 2:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 3
 * 输出: -1
 *
 */

// @lc code=start
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	l, r := 0, length - 1
	return _search(nums, l, r, target)
}

func _search(nums []int, l, r, target int) int {
	if l > r {
		return -1
	} else if l == r {
		if nums[l] == target {
			return l
		}
		return -1
	}

	mid := (l + r) / 2
	if nums[mid] == target {
		return mid
	}

	// 总是优先寻找有序的一边
	if l < mid - 1 && nums[l] <= nums[mid-1] {
		res := _search(nums, l, mid - 1, target)
		if res != -1 {
			return res
		}
		return _search(nums, mid + 1, r, target)
	} else {
		res := _search(nums, mid + 1, r, target)
		if res != -1 {
			return res
		}
		return _search(nums, l, mid - 1, target)
	}
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (34.26%)
 * Likes:    72
 * Dislikes: 0
 * Total Accepted:    13.3K
 * Total Submissions: 38.6K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
 *
 * 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
 *
 * 示例 1:
 *
 * 输入: nums = [2,5,6,0,0,1,2], target = 0
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: nums = [2,5,6,0,0,1,2], target = 3
 * 输出: false
 *
 * 进阶:
 *
 *
 * 这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
 * 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 *
 *
 */

// @lc code=start
func search(nums []int, target int) bool {
	// 还是一样的，始终会有一边是有序数组
	length := len(nums)
	if length == 0 {
		return false
	}
	l, r := 0, length-1
	return _search(nums, l, r, target)
}

func _search(nums []int, l, r, target int) bool {
	if l > r {
		return false
	}

	mid := (l + r) / 2
	if nums[mid] == target {
		return true
	}

	// 查看哪边是有序的
	if l <= mid-1 && nums[l] <= nums[mid-1] {
		return _search(nums, l, mid-1, target) || _search(nums, mid+1, r, target)
	} else if mid+1 <= r && nums[mid+1] <= nums[r] {
		return _search(nums, mid+1, r, target) || _search(nums, l, mid-1, target)
	}

	return false
}

// @lc code=end


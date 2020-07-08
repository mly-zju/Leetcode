package main

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 *
 * https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (49.56%)
 * Likes:    105
 * Dislikes: 0
 * Total Accepted:    23K
 * Total Submissions: 46.1K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] )。
 *
 * 请找出其中最小的元素。
 *
 * 你可以假设数组中不存在重复元素。
 *
 * 示例 1:
 *
 * 输入: [3,4,5,1,2]
 * 输出: 1
 *
 * 示例 2:
 *
 * 输入: [4,5,6,7,0,1,2]
 * 输出: 0
 *
 */

// @lc code=start
func findMin(nums []int) int {
	// 其实就是找旋转点，因为最小值一定藏在旋转点所在区间
	length := len(nums)
	if length == 0 {
		return 0
	}
	l, r := 0, length-1
	// 排除有序序列
	if nums[r] > nums[l] {
		return nums[0]
	}

	var mid int
	for l < r {
		mid = (l + r) / 2
		// 总是寻找非有序的一段，因为最小点一定在非有序的一边
		if nums[mid] > nums[r] {
			// 如果mid大于r，那么mid+1~r一定非有序
			l = mid + 1
		} else if nums[mid] < nums[r] {
			// 如果mid小于r，那么l~mid一定非有序(要包含mid，因为mid可能正好是最小点)
			r = mid
		} else {
			// 如果相等，那么不确定是否有序，但是至少可以缩小范围
			r--
		}
	}
	return nums[l]
}

// func main() {
// 	x := []int{1, 2}
// 	fmt.Println(findMin(x))
// }

// @lc code=end

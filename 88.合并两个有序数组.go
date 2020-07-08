/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (45.32%)
 * Likes:    424
 * Dislikes: 0
 * Total Accepted:    110.4K
 * Total Submissions: 236.4K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
 * 
 * 说明:
 * 
 * 
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 * 
 * 
 * 示例:
 * 
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 * 
 * 输出: [1,2,2,3,5,6]
 * 
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	// 双指针合并法：拉链法，类似快慢指针，但是由于在两个数组
	// 可以两个都从头开始，配上移位函数。更好的办法是从尾部开始, 这里也变三指针了，还有一个修改基准指针
	l1, l2, cursor := m - 1, n - 1, m + n - 1
	for l1 >= 0 && l2 >= 0 {
		if nums1[l1] >= nums2[l2] {
			nums1[cursor] = nums1[l1]
			l1--
			cursor--
		} else {
			nums1[cursor] = nums2[l2]
			l2--
			cursor--
		}
	}
	// 继续迭代
	for l1 >= 0 {
		nums1[cursor] = nums1[l1]
		l1--
		cursor--
	}
	for l2 >= 0 {
		nums1[cursor] = nums2[l2]
		l2--
		cursor--
	}
}
// @lc code=end


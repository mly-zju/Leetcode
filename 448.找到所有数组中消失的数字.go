/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 *
 * https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (53.03%)
 * Likes:    283
 * Dislikes: 0
 * Total Accepted:    26.7K
 * Total Submissions: 47.4K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * 给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
 * 
 * 找到所有在 [1, n] 范围之间没有出现在数组中的数字。
 * 
 * 您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
 * 
 * 示例:
 * 
 * 
 * 输入:
 * [4,3,2,7,8,2,3,1]
 * 
 * 输出:
 * [5,6]
 * 
 * 
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	// 最简单当然是哈希统计，或者排序统计数组。
	// 但是要求没有额外空间，那么可以做标识，出现的，相应索引的值改为负数，这样检测没编程负数的索引即可。
	length := len(nums)
	for i := 0; i < length; i++ {
		setIndex := getAbs(nums[i]) - 1
		if nums[setIndex] > 0 {
			nums[setIndex] = -nums[setIndex]
		}
	}

	res := []int{}
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
// @lc code=end


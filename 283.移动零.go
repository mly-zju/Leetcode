/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (56.62%)
 * Likes:    498
 * Dislikes: 0
 * Total Accepted:    108.7K
 * Total Submissions: 182.8K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 * 
 * 示例:
 * 
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 * 
 * 说明:
 * 
 * 
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 * 
 * 
 */

// @lc code=start
func moveZeroes(nums []int)  {
	// 27题目的类似版本，快慢指针之后把最后剩下的全部置为0
	length := len(nums)
	if length == 0 {
		return
	}

	s, f := 0, 0
	for f < length {
		for f < length && nums[f] == 0 {
			f++
		}
		if f < length {
			nums[s] = nums[f]
			s++
			f++
		}
	}
	for s < length {
		nums[s] = 0
		s++
	}
}
// @lc code=end


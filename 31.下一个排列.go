import "sort"

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 *
 * https://leetcode-cn.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (32.06%)
 * Likes:    304
 * Dislikes: 0
 * Total Accepted:    28.5K
 * Total Submissions: 88.7K
 * Testcase Example:  '[1,2,3]'
 *
 * 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
 *
 * 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
 *
 * 必须原地修改，只允许使用额外常数空间。
 *
 * 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 *
 */

// @lc code=start
func nextPermutation(nums []int) {
	// 从后向前，找第一个递减的索引
	length := len(nums)
	index := -1
	for i := length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			index = i
			break
		}
	}
	if index == -1 {
		// 如果全部递增，则直接反转
		l, r := 0, length-1
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	} else {
		// 否则寻找index后面大于它的最小的数字，交换并将index后面元素升序排列
		minIndex := index + 1
		for i := index + 1; i < length; i++ {
			if nums[i] > nums[index] && nums[i] < nums[minIndex] {
				minIndex = i
			}
		}
		nums[index], nums[minIndex] = nums[minIndex], nums[index]
		sort.Ints(nums[index+1:])
	}
}

// @lc code=end


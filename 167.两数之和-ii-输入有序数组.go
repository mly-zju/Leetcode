/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 *
 * https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/description/
 *
 * algorithms
 * Easy (49.96%)
 * Likes:    243
 * Dislikes: 0
 * Total Accepted:    67.9K
 * Total Submissions: 129.1K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
 * 
 * 函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
 * 
 * 说明:
 * 
 * 
 * 返回的下标值（index1 和 index2）不是从零开始的。
 * 你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
 * 
 * 
 * 示例:
 * 
 * 输入: numbers = [2, 7, 11, 15], target = 9
 * 输出: [1,2]
 * 解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
 * 
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	// 1. 基于统计hash来搞
	// statMap := make(map[int]int)
	// // 先遍历一遍，执行统计
	// for index, val := range numbers {
	// 	statMap[val] = index + 1
	// }
	// res := []int{}
	// for index, val := range numbers {
	// 	if statMap[target - val] > 0 {
	// 		res = append(res, index + 1)
	// 		res = append(res, statMap[target - val])
	// 		break
	// 	}
	// }
	// return res

	// 2. 基于首尾指针：前提是已经排序
	length := len(numbers)
	l, r := 0, length - 1
	res := []int{}
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			res = append(res, []int{l+1, r+1}...)
			break
		}
	}
	return res
}
// @lc code=end


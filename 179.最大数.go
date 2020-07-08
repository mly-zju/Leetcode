package main

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 *
 * https://leetcode-cn.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (33.65%)
 * Likes:    173
 * Dislikes: 0
 * Total Accepted:    15.5K
 * Total Submissions: 45K
 * Testcase Example:  '[10,2]'
 *
 * 给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
 *
 * 示例 1:
 *
 * 输入: [10,2]
 * 输出: 210
 *
 * 示例 2:
 *
 * 输入: [3,30,34,5,9]
 * 输出: 9534330
 *
 * 说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
 *
 */

// @lc code=start
func largestNumber(nums []int) string {
	// 数字按位对比大小，超出的截取，继续进行循环对比
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		astr, bstr := strconv.Itoa(a), strconv.Itoa(b)
		// 组合使得字符串长度相等
		aS, bS := astr+bstr, bstr+astr
		return aS > bS
	})

	res := ""
	for _, num := range nums {
		res += strconv.Itoa(num)
	}

	// 还需要对0特殊处理
	resNum, _ := strconv.Atoi(res)
	if resNum == 0 {
		return "0"
	}
	return res
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
// 	fmt.Println(largestNumber(a))
// }

// @lc code=end

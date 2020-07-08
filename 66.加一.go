/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 *
 * https://leetcode-cn.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (40.89%)
 * Likes:    427
 * Dislikes: 0
 * Total Accepted:    119.6K
 * Total Submissions: 277.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 * 
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 * 
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 * 
 * 示例 1:
 * 
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 * 
 * 
 * 示例 2:
 * 
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 * 
 * 
 */

// @lc code=start
func plusOne(digits []int) []int {
	length := len(digits)
	res := []int{}
	if length == 0 {
		return res
	}

	// 进位标识
	flag := 0
	for i := length - 1; i >= 0; i-- {
		tmpNum := 0
		if i == length - 1 {
			tmpNum = digits[i] + 1	
		} else {
			tmpNum = digits[i] + flag
			flag = 0
		}

		if tmpNum >= 10 {
			flag = 1
			tmpNum = tmpNum % 10
		}
		res = append(res, tmpNum)
	}
	// 最后处理进位
	if flag > 0 {
		res = append(res, flag)
	}
	lenr := len(res)
	l, r := 0, lenr - 1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
// @lc code=end


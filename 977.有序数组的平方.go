/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 *
 * https://leetcode-cn.com/problems/squares-of-a-sorted-array/description/
 *
 * algorithms
 * Easy (70.68%)
 * Likes:    74
 * Dislikes: 0
 * Total Accepted:    29.8K
 * Total Submissions: 42.1K
 * Testcase Example:  '[-4,-1,0,3,10]'
 *
 * 给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：[-4,-1,0,3,10]
 * 输出：[0,1,9,16,100]
 * 
 * 
 * 示例 2：
 * 
 * 输入：[-7,-3,2,3,11]
 * 输出：[4,9,9,49,121]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= A.length <= 10000
 * -10000 <= A[i] <= 10000
 * A 已按非递减顺序排序。
 * 
 * 
 */

// @lc code=start
func sortedSquares(A []int) []int {
	// 首尾指针构建一个非递增新数组，然后反转即可。
	// 或者也可以一开始寻找奇数偶数交界出，两边扩散指针
	length := len(A)
	res := []int{}
	if length == 0 {
		return res
	}
	l, r := 0, length - 1
	for l < r {
		if getAbs(A[l]) > getAbs(A[r]) {
			res = append(res, A[l] * A[l])
			l++
		} else {
			res = append(res, A[r] * A[r])
			r--
		}
	}
	// 最后相等以后还要再把相等元素插入
	res = append(res, A[r] * A[r])

	// 执行反转
	l, r = 0, length - 1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
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


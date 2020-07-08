/*
 * @lc app=leetcode.cn id=896 lang=golang
 *
 * [896] 单调数列
 *
 * https://leetcode-cn.com/problems/monotonic-array/description/
 *
 * algorithms
 * Easy (49.23%)
 * Likes:    37
 * Dislikes: 0
 * Total Accepted:    9.1K
 * Total Submissions: 18.4K
 * Testcase Example:  '[1,2,2,3]'
 *
 * 如果数组是单调递增或单调递减的，那么它是单调的。
 *
 * 如果对于所有 i <= j，A[i] <= A[j]，那么数组 A 是单调递增的。 如果对于所有 i <= j，A[i]> = A[j]，那么数组 A
 * 是单调递减的。
 *
 * 当给定的数组 A 是单调数组时返回 true，否则返回 false。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：[1,2,2,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 * 输入：[6,5,4,4]
 * 输出：true
 *
 *
 * 示例 3：
 *
 * 输入：[1,3,2]
 * 输出：false
 *
 *
 * 示例 4：
 *
 * 输入：[1,2,4,5]
 * 输出：true
 *
 *
 * 示例 5：
 *
 * 输入：[1,1,1]
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= A.length <= 50000
 * -100000 <= A[i] <= 100000
 *
 *
 */

// @lc code=start
func isMonotonic(A []int) bool {
	// 1. 判断是否单调递增或者单调递减
	// incr, desc := true, true
	// for i, length := 1, len(A); i < length; i++ {
	// 	if A[i] < A[i-1] {
	// 		// 如果后一个小于前一个，就不是递增
	// 		incr = false
	// 	} else if A[i] > A[i-1] {
	// 		// 如果后一个大于前一个，就不是递减
	// 		desc = false
	// 	}
	// 	if !incr && !desc {
	// 		return false
	// 	}
	// }
	// return incr || desc

	// 2. 基于数据流或状态机，后一个减前一个是否始终>=0或者<=0
	flag := 0 // 0无状态，1递增，-1递减
	for i, length := 1, len(A); i < length; i++ {
		if A[i] > A[i-1] {
			if flag == 0 {
				flag = 1
			} else if flag == -1 {
				return false
			}
		} else if A[i] < A[i-1] {
			if flag == 0 {
				flag = -1
			} else if flag == 1 {
				return false
			}
		}
	}
	return true
}

// @lc code=end


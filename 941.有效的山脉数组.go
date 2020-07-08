/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 *
 * https://leetcode-cn.com/problems/valid-mountain-array/description/
 *
 * algorithms
 * Easy (34.07%)
 * Likes:    28
 * Dislikes: 0
 * Total Accepted:    5.1K
 * Total Submissions: 15K
 * Testcase Example:  '[2,1]'
 *
 * 给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。
 *
 * 让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：
 *
 *
 * A.length >= 3
 * 在 0 < i < A.length - 1 条件下，存在 i 使得：
 *
 * A[0] < A[1] < ... A[i-1] < A[i]
 * A[i] > A[i+1] > ... > A[B.length - 1]
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：[2,1]
 * 输出：false
 *
 *
 * 示例 2：
 *
 * 输入：[3,5,5]
 * 输出：false
 *
 *
 * 示例 3：
 *
 * 输入：[0,3,2,1]
 * 输出：true
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= A.length <= 10000
 * 0 <= A[i] <= 10000
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
func validMountainArray(A []int) bool {
	length := len(A)
	if length < 3 {
		return false
	}

	flag := 0
	// 状态必须从0->1->-1
	for i := 1; i < length; i++ {
		if A[i] == A[i-1] {
			// 不能相等，必须严格递增或者递减
			return false
		} else if A[i] > A[i-1] {
			// 此时之前flag状态不能为-1, 因为不可以-1 -> 1
			if flag == -1 {
				return false
			}
			flag = 1
		} else if A[i] < A[i-1] {
			// 此时之前flag的状态不能为0，因为不可以0->-1
			if flag == 0 {
				return false
			}
			flag = -1
		}
	}
	// flag最后的状态必须为-1
	return flag == -1
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 *
 * https://leetcode-cn.com/problems/sort-array-by-parity-ii/description/
 *
 * algorithms
 * Easy (65.25%)
 * Likes:    64
 * Dislikes: 0
 * Total Accepted:    17.7K
 * Total Submissions: 26.9K
 * Testcase Example:  '[4,2,5,7]'
 *
 * 给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。
 *
 * 对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。
 *
 * 你可以返回任何满足上述条件的数组作为答案。
 *
 *
 *
 * 示例：
 *
 * 输入：[4,2,5,7]
 * 输出：[4,5,2,7]
 * 解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= A.length <= 20000
 * A.length % 2 == 0
 * 0 <= A[i] <= 1000
 *
 *
 *
 *
 */

// @lc code=start
func sortArrayByParityII(A []int) []int {
	ei, oi := 0, 1
	length := len(A)
	for ei < length && oi < length {
		// 先找偶数位置非偶数的
		for ei < length && A[ei]%2 == 0 {
			ei += 2
		}
		// 再找奇数位置非奇数的
		for oi < length && A[oi]%2 == 1 {
			oi += 2
		}
		if ei < length && oi < length {
			A[ei], A[oi] = A[oi], A[ei]
			ei += 2
			oi += 2
		}
	}
	return A
}

// @lc code=end


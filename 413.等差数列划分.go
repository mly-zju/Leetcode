/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 *
 * https://leetcode-cn.com/problems/arithmetic-slices/description/
 *
 * algorithms
 * Medium (57.13%)
 * Likes:    74
 * Dislikes: 0
 * Total Accepted:    7.7K
 * Total Submissions: 13.2K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 如果一个数列至少有三个元素，并且任意两个相邻元素之差相同，则称该数列为等差数列。
 *
 * 例如，以下数列为等差数列:
 *
 *
 * 1, 3, 5, 7, 9
 * 7, 7, 7, 7
 * 3, -1, -5, -9
 *
 * 以下数列不是等差数列。
 *
 *
 * 1, 1, 2, 5, 7
 *
 *
 *
 * 数组 A 包含 N 个数，且索引从0开始。数组 A 的一个子数组划分为数组 (P, Q)，P 与 Q 是整数且满足 0<=P<Q<N 。
 *
 * 如果满足以下条件，则称子数组(P, Q)为等差数组：
 *
 * 元素 A[P], A[p + 1], ..., A[Q - 1], A[Q] 是等差的。并且 P + 1 < Q 。
 *
 * 函数要返回数组 A 中所有为等差数组的子数组个数。
 *
 *
 *
 * 示例:
 *
 *
 * A = [1, 2, 3, 4]
 *
 * 返回: 3, A 中有三个子等差数组: [1, 2, 3], [2, 3, 4] 以及自身 [1, 2, 3, 4]。
 *
 *
 */

// @lc code=start
func numberOfArithmeticSlices(A []int) int {
	// 从条件来看，子等差数组长度位3个以上
	// dp[n]代表以第n位结尾的等差数列长度
	length := len(A)
	if length < 3 {
		return 0
	}

	dp := []int{1, 2}
	sum := 0
	for i := 2; i < length; i++ {
		if dp[i-1] == 1 {
			// 如果前者是1，后者一定是2
			dp = append(dp, 2)
		} else {
			if A[i]-A[i-1] == A[i-1]-A[i-2] {
				dp = append(dp, dp[i-1]+1)
			} else {
				// 每次断掉的时候，说明找到了区域最大连续等差，则找出前者的排列组合
				// 注意每次断掉以后变2，而不是变1
				dp = append(dp, 2)
				prevLen := dp[i-1]
				for j := 1; j <= prevLen-3+1; j++ {
					sum += j
				}
			}
		}
	}

	// 最后再执行一次, 执行前需要查看是否满足断掉的条件
	if dp[length-1] >= 3 {
		prevLen := dp[length-1]
		for j := 1; j <= prevLen-3+1; j++ {
			sum += j
		}
	}

	return sum
}

// @lc code=end


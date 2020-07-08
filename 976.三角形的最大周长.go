import "sort"

/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长
 *
 * https://leetcode-cn.com/problems/largest-perimeter-triangle/description/
 *
 * algorithms
 * Easy (54.28%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    7.9K
 * Total Submissions: 14.4K
 * Testcase Example:  '[2,1,2]'
 *
 * 给定由一些正数（代表长度）组成的数组 A，返回由其中三个长度组成的、面积不为零的三角形的最大周长。
 *
 * 如果不能形成任何面积不为零的三角形，返回 0。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：[2,1,2]
 * 输出：5
 *
 *
 * 示例 2：
 *
 * 输入：[1,2,1]
 * 输出：0
 *
 *
 * 示例 3：
 *
 * 输入：[3,2,3,4]
 * 输出：10
 *
 *
 * 示例 4：
 *
 * 输入：[3,6,2,3]
 * 输出：8
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= A.length <= 10000
 * 1 <= A[i] <= 10^6
 *
 *
 */

// @lc code=start
func largestPerimeter(A []int) int {
	// 排序后验证是否三角形
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})
	length := len(A)
	if length < 3 {
		return 0
	}

	a, b, c := 0, 1, 2
	flag := false
	for c < length {
		if A[a] >= A[b]+A[c] || A[a]-A[c] >= A[b] {
			// 如果两边之和小于等于第三边，证明a太长
			// 同样, 如果两边之差大于等于第三边，也是证明有一个边太长
			// 这个时候，只能均匀的移动，不能跳着移动，因为会更加加大差距(在有序的情况下)
			a++
			b++
			c++
		} else {
			flag = true
			break
		}
	}
	if flag {
		return A[a] + A[b] + A[c]
	}
	return 0
}

// @lc code=end


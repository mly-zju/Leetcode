import "sort"

/*
 * @lc app=leetcode.cn id=870 lang=golang
 *
 * [870] 优势洗牌
 *
 * https://leetcode-cn.com/problems/advantage-shuffle/description/
 *
 * algorithms
 * Medium (34.65%)
 * Likes:    40
 * Dislikes: 0
 * Total Accepted:    3.8K
 * Total Submissions: 11.2K
 * Testcase Example:  '[2,7,11,15]\n[1,10,4,11]'
 *
 * 给定两个大小相等的数组 A 和 B，A 相对于 B 的优势可以用满足 A[i] > B[i] 的索引 i 的数目来描述。
 *
 * 返回 A 的任意排列，使其相对于 B 的优势最大化。
 *
 *
 *
 * 示例 1：
 *
 * 输入：A = [2,7,11,15], B = [1,10,4,11]
 * 输出：[2,11,7,15]
 *
 *
 * 示例 2：
 *
 * 输入：A = [12,24,8,32], B = [13,25,32,11]
 * 输出：[24,32,8,12]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= A.length = B.length <= 10000
 * 0 <= A[i] <= 10^9
 * 0 <= B[i] <= 10^9
 *
 *
 */

// @lc code=start
func advantageCount(A []int, B []int) []int {
	length := len(A)
	sort.Ints(A)
	res := []int{}
	for i := 0; i < length; i++ {
		// 找到最小限度能满足的数字，没有就推入最小值
		ansIndex := -1
		lenA := len(A)
		// 这里不再通过标记法，而是使用过的直接删除，降低时间复杂度
		for j := lenA - 1; j >= 0; j-- {
			if A[j] > B[i] {
				ansIndex = j
			}
		}

		if ansIndex != -1 {
			res = append(res, A[ansIndex])
			A = append(A[0:ansIndex], A[ansIndex+1:]...)
		} else {
			res = append(res, A[0])
			A = A[1:]
		}
	}

	return res
}

// @lc code=end


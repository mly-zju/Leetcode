/*
 * @lc app=leetcode.cn id=1051 lang=golang
 *
 * [1051] 高度检查器
 *
 * https://leetcode-cn.com/problems/height-checker/description/
 *
 * algorithms
 * Easy (71.50%)
 * Likes:    30
 * Dislikes: 0
 * Total Accepted:    7.6K
 * Total Submissions: 10.6K
 * Testcase Example:  '[1,1,4,2,1,3]'
 *
 * 学校在拍年度纪念照时，一般要求学生按照 非递减 的高度顺序排列。
 *
 * 请你返回至少有多少个学生没有站在正确位置数量。该人数指的是：能让所有学生以 非递减 高度排列的必要移动人数。
 *
 *
 *
 * 示例：
 *
 * 输入：[1,1,4,2,1,3]
 * 输出：3
 * 解释：
 * 高度为 4、3 和最后一个 1 的学生，没有站在正确的位置。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= heights.length <= 100
 * 1 <= heights[i] <= 100
 *
 *
 */

// @lc code=start
func heightChecker(heights []int) int {
	// 1. 排序后比较
	// length := len(heights)
	// heights2 := make([]int, length)
	// copy(heights2, heights)
	// sort.Ints(heights2)
	// res := 0
	// for i := 0; i < length; i++ {
	// 	if heights[i] != heights2[i] {
	// 		res++
	// 	}
	// }
	// return res

	// 2. 基于计数排序
	statArr := [101]int{}
	for _, val := range heights {
		statArr[val]++
	}

	cursor := 0
	res := 0
	for i := 1; i < 101; i++ {
		for j := 0; j < statArr[i]; j++ {
			if heights[cursor] != i {
				res++
			}
			cursor++
		}
	}
	return res
}

// @lc code=end


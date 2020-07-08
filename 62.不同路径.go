/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 *
 * https://leetcode-cn.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (56.29%)
 * Likes:    360
 * Dislikes: 0
 * Total Accepted:    56.1K
 * Total Submissions: 97.3K
 * Testcase Example:  '3\n2'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 *
 * 问总共有多少条不同的路径？
 *
 *
 *
 * 例如，上图是一个7 x 3 的网格。有多少可能的路径？
 *
 * 说明：m 和 n 的值均不超过 100。
 *
 * 示例 1:
 *
 * 输入: m = 3, n = 2
 * 输出: 3
 * 解释:
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向右 -> 向下
 * 2. 向右 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向右
 *
 *
 * 示例 2:
 *
 * 输入: m = 7, n = 3
 * 输出: 28
 *
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	// 1. 动态规划
	// res := [][]int{}
	// for i := 0; i < m; i++ {
	// 	lineArr := []int{}
	// 	for j := 0; j < n; j++ {
	// 		if i == 0 || j == 0 {
	// 			lineArr = append(lineArr, 1)
	// 		} else {
	// 			lineArr = append(lineArr, res[i-1][j]+lineArr[j-1])
	// 		}
	// 	}
	// 	res = append(res, lineArr)
	// }

	// return res[m-1][n-1]

	// 2. 回溯法
	res := 0
	return _run(0, 0, m, n, res)
}

func _run(curM, curN, m, n, res int) int {
	if curM == m - 1 && curN == n - 1 {
		res++
		return res
	} else {
		// 1. 如果可以向右的话
		if curM < m - 1 {
			res = _run(curM+1, curN, m, n, res)
		}

		// 2. 如果可以向下的话
		if curN < n - 1 {
			res = _run(curM, curN+1, m, n, res)
		}
		return res
	}
}

// @lc code=end


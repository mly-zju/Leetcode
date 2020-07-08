/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode-cn.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (62.33%)
 * Likes:    338
 * Dislikes: 0
 * Total Accepted:    46K
 * Total Submissions: 72.5K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 * 示例:
 *
 * 输入:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * 输出: 7
 * 解释: 因为路径 1→3→1→1→1 的总和最小。
 *
 *
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	rLen := len(grid)
	if rLen == 0 {
		return 0
	}
	cLen := len(grid[0])
	if cLen == 0 {
		return 0
	}

	// 和63题相比，res里面存放的就不是路径数量，而是最小路径和
	res := [][]int{}
	for i := 0; i < rLen; i++ {
		lineArr := []int{}
		for j := 0; j < cLen; j++ {
			if i == 0 && j == 0 {
				lineArr = append(lineArr, grid[i][j])
			} else if i == 0 {
				lineArr = append(lineArr, lineArr[j-1]+grid[i][j])
			} else if j == 0 {
				lineArr = append(lineArr, res[i-1][j]+grid[i][j])
			} else {
				lineArr = append(lineArr, getMin(res[i-1][j], lineArr[j-1])+grid[i][j])
			}
		}
		res = append(res, lineArr)
	}

	return res[rLen-1][cLen-1]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end


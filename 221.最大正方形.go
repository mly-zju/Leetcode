/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 *
 * https://leetcode-cn.com/problems/maximal-square/description/
 *
 * algorithms
 * Medium (38.20%)
 * Likes:    182
 * Dislikes: 0
 * Total Accepted:    16.4K
 * Total Submissions: 42.4K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
 *
 * 示例:
 *
 * 输入:
 *
 * 1 0 1 0 0
 * 1 0 1 1 1
 * 1 1 1 1 1
 * 1 0 0 1 0
 *
 * 输出: 4
 *
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	// dp[i][j]代表以第i行j列结尾的最大正方形边长
	rNum := len(matrix)
	if rNum == 0 {
		return rNum
	}
	cNum := len(matrix[0])
	if cNum == 0 {
		return cNum
	}

	dp := [][]int{}
	max := 0
	for i := 0; i < rNum; i++ {
		lineArr := []int{}
		for j := 0; j < cNum; j++ {
			if matrix[i][j] == '0' {
				lineArr = append(lineArr, 0)
			} else {
				prevI, prevJ := i-1, j-1
				// 如果前驱元素已经小于界限，当前最大1
				if prevI < 0 || prevJ < 0 {
					lineArr = append(lineArr, 1)
					if max < 1 {
						max = 1
					}
				} else {
					// tmpMax := 1 + checkMax(matrix, i, j, dp[prevI][prevJ])
					tmpMax := 1 + getMin(dp[prevI][prevJ], getMin(dp[i-1][j], lineArr[j-1]))
					lineArr = append(lineArr, tmpMax)
					if tmpMax > max {
						max = tmpMax
					}
				}
			}
		}
		dp = append(dp, lineArr)
	}

	// max存放的是边长
	return max * max
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func checkMax(matrix [][]byte, i, j, limit int) int {
	if limit == 0 {
		return 0
	}

	rMax, cMax := 0, 0
	for k := 1; k <= limit; k++ {
		if matrix[i][j-k] == '1' {
			rMax++
		} else {
			break
		}
	}

	for k := 1; k <= limit; k++ {
		if matrix[i-k][j] == '1' {
			cMax++
		} else {
			break
		}
	}

	if rMax < cMax {
		return rMax
	}
	return cMax
}

// @lc code=end


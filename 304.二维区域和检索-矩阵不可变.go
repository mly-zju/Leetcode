/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 *
 * https://leetcode-cn.com/problems/range-sum-query-2d-immutable/description/
 *
 * algorithms
 * Medium (40.76%)
 * Likes:    52
 * Dislikes: 0
 * Total Accepted:    4.3K
 * Total Submissions: 10.4K
 * Testcase Example:  '["NumMatrix","sumRegion","sumRegion","sumRegion"]\n' +
  '[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]'
 *
 * 给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2)。
 *
 *
 * 上图子矩阵左上角 (row1, col1) = (2, 1) ，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。
 *
 * 示例:
 *
 * 给定 matrix = [
 * ⁠ [3, 0, 1, 4, 2],
 * ⁠ [5, 6, 3, 2, 1],
 * ⁠ [1, 2, 0, 1, 5],
 * ⁠ [4, 1, 0, 1, 7],
 * ⁠ [1, 0, 3, 0, 5]
 * ]
 *
 * sumRegion(2, 1, 4, 3) -> 8
 * sumRegion(1, 1, 2, 2) -> 11
 * sumRegion(1, 2, 2, 4) -> 12
 *
 *
 * 说明:
 *
 *
 * 你可以假设矩阵不可变。
 * 会多次调用 sumRegion 方法。
 * 你可以假设 row1 ≤ row2 且 col1 ≤ col2。
 *
 *
*/

// @lc code=start
type NumMatrix struct {
	// 基于行缓存来做
	sumLine [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	numMatrix := NumMatrix{
		sumLine: [][]int{},
	}

	// 初始化，填充sumLine缓存
	rNum := len(matrix)
	if rNum == 0 {
		return numMatrix
	}
	cNum := len(matrix[0])
	if cNum == 0 {
		return numMatrix
	}

	for i := 0; i < rNum; i++ {
		lineCache := []int{}
		tmp := 0
		for j := 0; j < cNum; j++ {
			tmp += matrix[i][j]
			lineCache = append(lineCache, tmp)
		}
		numMatrix.sumLine = append(numMatrix.sumLine, lineCache)
	}

	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	rStart, rEnd, cStart, cEnd := row1, row2, col1, col2

	sum := 0
	for i := rStart; i <= rEnd; i++ {
		if cStart-1 < 0 {
			sum += this.sumLine[i][cEnd]
		} else {
			sum += this.sumLine[i][cEnd] - this.sumLine[i][cStart-1]
		}
	}

	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end


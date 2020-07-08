package main

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix-ii/description/
 *
 * algorithms
 * Medium (37.73%)
 * Likes:    175
 * Dislikes: 0
 * Total Accepted:    32K
 * Total Submissions: 84.6K
 * Testcase Example:  '[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n' +
  '5'
 *
 * 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
 *
 *
 * 每行的元素从左到右升序排列。
 * 每列的元素从上到下升序排列。
 *
 *
 * 示例:
 *
 * 现有矩阵 matrix 如下：
 *
 * [
 * ⁠ [1,   4,  7, 11, 15],
 * ⁠ [2,   5,  8, 12, 19],
 * ⁠ [3,   6,  9, 16, 22],
 * ⁠ [10, 13, 14, 17, 24],
 * ⁠ [18, 21, 23, 26, 30]
 * ]
 *
 *
 * 给定 target = 5，返回 true。
 *
 * 给定 target = 20，返回 false。
 *
*/

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	// 左下角开始遍历
	rNum := len(matrix)
	if rNum == 0 {
		return false
	}
	cNum := len(matrix[0])
	if cNum == 0 {
		return false
	}

	r, c := rNum-1, 0
	flag := false
	for r >= 0 && c < cNum {
		if matrix[r][c] > target {
			// 太大了，向上走一点
			r--
		} else if matrix[r][c] < target {
			c++
		} else {
			flag = true
			break
		}
	}
	return flag
}

// func main() {
// 	target := 5
// 	nums := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
// 	fmt.Println(searchMatrix(nums, target))
// }

// @lc code=end

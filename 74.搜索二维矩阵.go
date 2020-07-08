package main

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (35.91%)
 * Likes:    102
 * Dislikes: 0
 * Total Accepted:    22.7K
 * Total Submissions: 62.3K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 *
 *
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 *
 *
 * 示例 1:
 *
 * 输入:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * 输出: false
 *
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	rNum := len(matrix)
	if rNum == 0 {
		return false
	}
	cNum := len(matrix[0])
	if cNum == 0 {
		return false
	}

	// 1. 一个非n*n的有序矩阵二分查找, 可以先通过二分锁定行，再通过二分锁定列
	// // 先确定行
	// rl, rr := 0, rNum-1
	// rm := 0
	// flag := false
	// for rl <= rr {
	// 	rm = (rl + rr) / 2
	// 	if matrix[rm][0] > target {
	// 		// 如果当前太大
	// 		rr = rm - 1
	// 	} else if matrix[rm][cNum-1] < target {
	// 		// 如果当前太小
	// 		rl = rm + 1
	// 	} else {
	// 		// 否则代表可能在于此行
	// 		flag = true
	// 		break
	// 	}
	// }

	// if flag == false {
	// 	return false
	// }

	// cl, cr := 0, cNum-1
	// cm := 0
	// flag = false
	// // 开始对此行执行标准的二分
	// for cl <= cr {
	// 	cm = (cl + cr) / 2
	// 	if matrix[rm][cm] == target {
	// 		flag = true
	// 		break
	// 	} else if matrix[rm][cm] > target {
	// 		cr = cm - 1
	// 	} else {
	// 		cl = cm + 1
	// 	}
	// }

	// return flag

	// 2. 还是一次二分查找，只不过取值的时候需要折算成l和r来取
	l, r := 0, rNum*cNum-1
	mid := 0
	flag := false
	for l <= r {
		mid = (l + r) / 2
		midPos := getIndex(mid, rNum, cNum)
		midVal := matrix[midPos[0]][midPos[1]]
		if midVal > target {
			r = mid - 1
		} else if midVal < target {
			l = mid + 1
		} else {
			flag = true
			break
		}
	}
	return flag
}

func getIndex(val, rNum, cNum int) []int {
	c := val % cNum
	r := val / cNum
	return []int{r, c}
}

// func main() {
// 	x := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
// 	// x := [][]int{{1, 1}}
// 	fmt.Println(searchMatrix(x, 13))
// }

// @lc code=end

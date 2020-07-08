import "math"

/*
 * @lc app=leetcode.cn id=861 lang=golang
 *
 * [861] 翻转矩阵后的得分
 *
 * https://leetcode-cn.com/problems/score-after-flipping-matrix/description/
 *
 * algorithms
 * Medium (71.04%)
 * Likes:    48
 * Dislikes: 0
 * Total Accepted:    3.5K
 * Total Submissions: 4.8K
 * Testcase Example:  '[[0,0,1,1],[1,0,1,0],[1,1,0,0]]'
 *
 * 有一个二维矩阵 A 其中每个元素的值为 0 或 1 。
 *
 * 移动是指选择任一行或列，并转换该行或列中的每一个值：将所有 0 都更改为 1，将所有 1 都更改为 0。
 *
 * 在做出任意次数的移动后，将该矩阵的每一行都按照二进制数来解释，矩阵的得分就是这些数字的总和。
 *
 * 返回尽可能高的分数。
 *
 *
 *
 *
 *
 *
 * 示例：
 *
 * 输入：[[0,0,1,1],[1,0,1,0],[1,1,0,0]]
 * 输出：39
 * 解释：
 * 转换为 [[1,1,1,1],[1,0,0,1],[1,1,1,1]]
 * 0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= A.length <= 20
 * 1 <= A[0].length <= 20
 * A[i][j] 是 0 或 1
 *
 *
 */

// @lc code=start
func matrixScore(A [][]int) int {
	// 贪心：先保证每行第一个为1，然后之后每列如果0比1多，执行反转
	rNum := len(A)
	cNum := len(A[0])

	for r := 0; r < rNum; r++ {
		if A[r][0] == 0 {
			// 执行反转
			for c := 0; c < cNum; c++ {
				A[r][c] = getReverse(A[r][c])
			}
		}
	}

	// 计算列的二分之一
	cMax := int(math.Ceil(float64(rNum) / 2))
	// 遍历从1开始的每列
	for c := 1; c < cNum; c++ {
		c1Sum := 0
		for r := 0; r < rNum; r++ {
			if A[r][c] == 1 {
				c1Sum++
			}
		}
		if c1Sum < cMax {
			// 执行反转
			for r := 0; r < rNum; r++ {
				A[r][c] = getReverse(A[r][c])
			}
		}
	}

	// 计算每行的二进制并求和
	res := 0
	for r := 0; r < rNum; r++ {
		res += getSum(A[r])
	}
	return res
}

func getReverse(a int) int {
	if a == 1 {
		return 0
	}
	return 1
}

func getSum(arr []int) int {
	res := 0
	for _, val := range arr {
		res = res*2 + val
	}
	return res
}

// @lc code=end


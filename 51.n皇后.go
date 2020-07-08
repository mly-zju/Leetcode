/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 *
 * https://leetcode-cn.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (68.49%)
 * Likes:    358
 * Dislikes: 0
 * Total Accepted:    31.3K
 * Total Submissions: 45.7K
 * Testcase Example:  '4'
 *
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 * 
 * 
 * 
 * 上图为 8 皇后问题的一种解法。
 * 
 * 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
 * 
 * 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 * 
 * 示例:
 * 
 * 输入: 4
 * 输出: [
 * ⁠[".Q..",  // 解法 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 * 
 * ⁠["..Q.",  // 解法 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * 解释: 4 皇后问题存在两个不同的解法。
 * 
 * 
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	res := [][]string{}
	// 初始化棋盘map
	seenMap := [][]bool{}
	for i := 0; i < n; i++ {
		lineArr := make([]bool, n)
		seenMap = append(seenMap, lineArr)
	}

	return dfs(seenMap, 0, res)
}

func dfs(seenMap [][]bool, lineIndex int, res [][]string) [][]string {
	n := len(seenMap)
	if lineIndex == n {
		// 如果已经到了最后一行，遍历map，推入答案
		tmpAns := []string{}
		for i := 0; i < n; i++ {
			tmpStr := ""
			for j := 0; j < n; j++ {
				if seenMap[i][j] {
					tmpStr += "Q"
				} else {
					tmpStr += "."
				}
			}
			tmpAns = append(tmpAns, tmpStr)
		}

		res = append(res, tmpAns)
	} else {
		// 搜索该行每一列，尝试寻找答案
		for i := 0; i < n; i++ {
			if isOk(seenMap, lineIndex, i) {
				seenMap[lineIndex][i] = true
				res = dfs(seenMap, lineIndex + 1, res)
				seenMap[lineIndex][i] = false
			}
		}
	}
	return res
}

func isOk(seenMap [][]bool, r, c int) bool {
	n := len(seenMap)
	
	// 先检查所有同列是否有Q
	for i := 0; i < n; i++ {
		if seenMap[i][c] {
			return false
		}
	}

	// 再检查四处对角线
	// 1. 左上角
	i, j := r-1, c-1
	for i >= 0 && j >= 0 {
		if seenMap[i][j] {
			return false
		}
		i--
		j--
	}	

	// 2. 右下角
	i, j = r + 1, c + 1
	for i < n && j < n {
		if seenMap[i][j] {
			return false
		}
		i++
		j++
	}

	// 3. 右上角
	i, j = r - 1, c + 1
	for i >= 0 && j < n {
		if seenMap[i][j] {
			return false
		}
		i--
		j++
	}

	// 4. 左下角
	i, j = r + 1, c - 1
	for i < n && j >= 0 {
		if seenMap[i][j] {
			return false
		}
		i++
		j--
	}

	return true
}




// @lc code=end


/*
 * @lc app=leetcode.cn id=1138 lang=golang
 *
 * [1138] 字母板上的路径
 *
 * https://leetcode-cn.com/problems/alphabet-board-path/description/
 *
 * algorithms
 * Medium (33.60%)
 * Likes:    5
 * Dislikes: 0
 * Total Accepted:    1.5K
 * Total Submissions: 4.4K
 * Testcase Example:  '"leet"'
 *
 * 我们从一块字母板上的位置 (0, 0) 出发，该坐标对应的字符为 board[0][0]。
 *
 * 在本题里，字母板为board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"].
 *
 * 我们可以按下面的指令规则行动：
 *
 *
 * 如果方格存在，'U' 意味着将我们的位置上移一行；
 * 如果方格存在，'D' 意味着将我们的位置下移一行；
 * 如果方格存在，'L' 意味着将我们的位置左移一列；
 * 如果方格存在，'R' 意味着将我们的位置右移一列；
 * '!' 会把在我们当前位置 (r, c) 的字符 board[r][c] 添加到答案中。
 *
 *
 * 返回指令序列，用最小的行动次数让答案和目标 target 相同。你可以返回任何达成目标的路径。
 *
 *
 *
 * 示例 1：
 *
 * 输入：target = "leet"
 * 输出："DDR!UURRR!!DDD!"
 *
 *
 * 示例 2：
 *
 * 输入：target = "code"
 * 输出："RR!DDRR!UUL!R!"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= target.length <= 100
 * target 仅含有小写英文字母。
 *
 *
 */

// @lc code=start
func alphabetBoardPath(target string) string {
	// 贪心算法，每次都找下一个满足目标且距离当前坐标最近的位置, 位置计算函数使用曼哈顿距离
	board := []string{"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"}
	r, c := 0, 0
	res := ""
	// 统计每个字母的坐标
	rboardMap := make(map[byte]int)
	cboardMap := make(map[byte]int)
	for i, rLen := 0, len(board); i < rLen; i++ {
		for j, cLen := 0, len(board[i]); j < cLen; j++ {
			rboardMap[board[i][j]] = i
			cboardMap[board[i][j]] = j
		}
	}

	// 遍历target，找到每个字母的目标坐标，然后基于步长尝试加指令
	stepMap := map[string]int{
		"U": -1, "D": 1, "L": -1, "R": 1,
	}
	for i, length := 0, len(target); i < length; i++ {
		r1, c1 := rboardMap[target[i]], cboardMap[target[i]]
		var rDirect, cDirect string
		if r1 > r {
			rDirect = "D"
		} else {
			rDirect = "U"
		}
		if c1 > c {
			cDirect = "R"
		} else {
			cDirect = "L"
		}

		for r != r1 || c != c1 {
			// 优先column
			if c != c1 && c+stepMap[cDirect] < len(board[r]) {
				c += stepMap[cDirect]
				res += cDirect
			} else {
				r += stepMap[rDirect]
				res += rDirect
			}
		}
		res += "!"
	}
	return res
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (71.83%)
 * Likes:    685
 * Dislikes: 0
 * Total Accepted:    64.4K
 * Total Submissions: 88.5K
 * Testcase Example:  '3'
 *
 * 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
 *
 * 例如，给出 n = 3，生成结果为：
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
 *
 */

// @lc code=start
var res []string

func generateParenthesis(n int) []string {
	// 回溯法
	res = []string{}
	if n == 0 {
		return res
	}
	leftNum, rightNum := n, n
	total := n * 2
	buf := make([]byte, total)
	_search(total, leftNum, rightNum, buf)
	return res
}

func _search(total, leftNum, rightNum int, buf []byte) {
	if leftNum == 0 && rightNum == 0 {
		res = append(res, string(buf))
	} else {
		curIndex := total - leftNum - rightNum
		// 先看能否填充左括号, 只要leftNum>0,永远可以
		if leftNum > 0 {
			buf[curIndex] = '('
			_search(total, leftNum-1, rightNum, buf)
		}
		// 再看能否填充右括号，只要leftNum < rightNum, 都可以
		if rightNum > leftNum {
			buf[curIndex] = ')'
			_search(total, leftNum, rightNum-1, buf)
		}
	}
}

// @lc code=end


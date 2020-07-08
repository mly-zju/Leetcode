import "strings"

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (50.93%)
 * Likes:    534
 * Dislikes: 0
 * Total Accepted:    67.9K
 * Total Submissions: 130.8K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */

// @lc code=start
var res []string

func letterCombinations(digits string) []string {
	// 回溯算法

	res = []string{}
	// 数字对应字母map
	charMap := [][]string{
		{},
		{},
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}
	cNum := len(digits)
	if cNum == 0 {
		return res
	}
	buf := make([]string, cNum)
	_search(digits, buf, 0, charMap)
	return res
}

func _search(digits string, buf []string, step int, charMap [][]string) {
	if len(buf) == step {
		res = append(res, strings.Join(buf, ""))
	} else {
		tmpNum := int(digits[step] - '0')
		tryList := charMap[tmpNum]
		for _, val := range tryList {
			buf[step] = val
			_search(digits, buf, step+1, charMap)
		}
	}
}

// @lc code=end


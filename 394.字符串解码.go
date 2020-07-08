/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 *
 * https://leetcode-cn.com/problems/decode-string/description/
 *
 * algorithms
 * Medium (46.02%)
 * Likes:    143
 * Dislikes: 0
 * Total Accepted:    9.8K
 * Total Submissions: 21.1K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * 给定一个经过编码的字符串，返回它解码后的字符串。
 *
 * 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
 *
 * 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
 *
 * 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
 *
 * 示例:
 *
 *
 * s = "3[a]2[bc]", 返回 "aaabcbc".
 * s = "3[a2[c]]", 返回 "accaccacc".
 * s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
 *
 *
 */

// @lc code=start
type Ele struct {
	multi int
	str   string
}

func decodeString(s string) string {
	stack := []Ele{}
	strCache := ""
	multiCache := 0
	for _, ch := range s {
		if isNumber(ch) {
			// 如果是数字，增加multiCache
			multiCache = multiCache*10 + int(ch-'0')
		} else if isAlpha(ch) {
			// 如果是字母，增加strCache
			strCache += string(ch)
		} else if ch == '[' {
			// 每次遇到开括号，入栈
			stack = append(stack, Ele{
				multi: multiCache,
				str:   strCache,
			})
			multiCache = 0
			strCache = ""
		} else if ch == ']' {
			// 每次遇到闭括号，当前cache和前一个合并
			length := len(stack)
			myEle := stack[length-1]
			stack = stack[0 : length-1]
			tmpStr := ""
			for myEle.multi > 0 {
				tmpStr += strCache
				myEle.multi--
			}
			strCache = myEle.str + tmpStr
		}
	}
	return strCache
}

func isNumber(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isAlpha(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// @lc code=end


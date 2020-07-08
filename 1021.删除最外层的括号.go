/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 *
 * https://leetcode-cn.com/problems/remove-outermost-parentheses/description/
 *
 * algorithms
 * Easy (74.91%)
 * Likes:    89
 * Dislikes: 0
 * Total Accepted:    21.1K
 * Total Submissions: 27.6K
 * Testcase Example:  '"(()())(())"'
 *
 * 有效括号字符串为空 ("")、"(" + A + ")" 或 A + B，其中 A 和 B 都是有效的括号字符串，+
 * 代表字符串的连接。例如，""，"()"，"(())()" 和 "(()(()))" 都是有效的括号字符串。
 * 
 * 如果有效字符串 S 非空，且不存在将其拆分为 S = A+B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。
 * 
 * 给出一个非空有效字符串 S，考虑将其进行原语化分解，使得：S = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。
 * 
 * 对 S 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 S 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入："(()())(())"
 * 输出："()()()"
 * 解释：
 * 输入字符串为 "(()())(())"，原语化分解得到 "(()())" + "(())"，
 * 删除每个部分中的最外层括号后得到 "()()" + "()" = "()()()"。
 * 
 * 示例 2：
 * 
 * 输入："(()())(())(()(()))"
 * 输出："()()()()(())"
 * 解释：
 * 输入字符串为 "(()())(())(()(()))"，原语化分解得到 "(()())" + "(())" + "(()(()))"，
 * 删除每隔部分中的最外层括号后得到 "()()" + "()" + "()(())" = "()()()()(())"。
 * 
 * 
 * 示例 3：
 * 
 * 输入："()()"
 * 输出：""
 * 解释：
 * 输入字符串为 "()()"，原语化分解得到 "()" + "()"，
 * 删除每个部分中的最外层括号后得到 "" + "" = ""。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * S.length <= 10000
 * S[i] 为 "(" 或 ")"
 * S 是一个有效括号字符串
 * 
 * 
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	// 每次栈变空的时候，把期间的缓存去头去尾，添加到答案即可
	vLen := 0
	length := len(S)
	cache := []byte{}
	res := ""

	for i := 0; i < length; i++ {
		cache = append(cache, S[i])
		if S[i] == '(' {
			vLen++
		} else {
			vLen--
		}
		if vLen == 0 {
			// 每当vLen为0说明自洽了，把期间的cache掐头去尾添加到res
			if len(cache) > 0 {
				res += string(cache[1:len(cache) - 1])
				cache = cache[0:0]
			}
		}
	}
	return res
}
// @lc code=end


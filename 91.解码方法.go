/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 *
 * https://leetcode-cn.com/problems/decode-ways/description/
 *
 * algorithms
 * Medium (21.81%)
 * Likes:    230
 * Dislikes: 0
 * Total Accepted:    24.1K
 * Total Submissions: 107.9K
 * Testcase Example:  '"12"'
 *
 * 一条包含字母 A-Z 的消息通过以下方式进行了编码：
 *
 * 'A' -> 1
 * 'B' -> 2
 * ...
 * 'Z' -> 26
 *
 *
 * 给定一个只包含数字的非空字符串，请计算解码方法的总数。
 *
 * 示例 1:
 *
 * 输入: "12"
 * 输出: 2
 * 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
 *
 *
 * 示例 2:
 *
 * 输入: "226"
 * 输出: 3
 * 解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
 *
 *
 */

// @lc code=start
func numDecodings(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	// 如果第一位为0， 则直接返回0
	if s[0] == '0' {
		return 0
	}

	dp := []int{}
	for i := 0; i < length; i++ {
		if i == 0 {
			dp = append(dp, 1)
		} else {
			tmp := 0
			// 如果当前一位数合法
			if s[i] != '0' {
				tmp += dp[i-1]
			}
			// 如果当前两位数合法
			if isDoubleValid(s[i-1], s[i]) {
				if i-2 >= 0 {
					tmp += dp[i-2]
				} else {
					tmp++
				}
			}
			dp = append(dp, tmp)
		}
	}

	return dp[length-1]
}

func isDoubleValid(a, b byte) bool {
	if a == '1' {
		return true
	} else if a == '2' && b <= '6' {
		return true
	}
	return false
}

// @lc code=end


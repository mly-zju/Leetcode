/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文字符串 Ⅱ
 *
 * https://leetcode-cn.com/problems/valid-palindrome-ii/description/
 *
 * algorithms
 * Easy (33.07%)
 * Likes:    89
 * Dislikes: 0
 * Total Accepted:    7.9K
 * Total Submissions: 23.7K
 * Testcase Example:  '"aba"'
 *
 * 给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
 *
 * 示例 1:
 *
 *
 * 输入: "aba"
 * 输出: True
 *
 *
 * 示例 2:
 *
 *
 * 输入: "abca"
 * 输出: True
 * 解释: 你可以删除c字符。
 *
 *
 * 注意:
 *
 *
 * 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
 *
 *
 */

// @lc code=start
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		// 如果发现不一致，尝试跳过一个继续判断
		if s[l] != s[r] {
			return isPr(s, l+1, r) || isPr(s, l, r-1)
		}
		l++
		r--
	}
	return true
}

func isPr(s string, start, end int) bool {
	l, r := start, end
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// @lc code=end


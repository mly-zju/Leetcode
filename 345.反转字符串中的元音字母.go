/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 *
 * https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (47.93%)
 * Likes:    76
 * Dislikes: 0
 * Total Accepted:    24.8K
 * Total Submissions: 50.4K
 * Testcase Example:  '"hello"'
 *
 * 编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
 *
 * 示例 1:
 *
 * 输入: "hello"
 * 输出: "holle"
 *
 *
 * 示例 2:
 *
 * 输入: "leetcode"
 * 输出: "leotcede"
 *
 * 说明:
 * 元音字母不包含字母"y"。
 *
 */

// @lc code=start
func reverseVowels(s string) string {
	checkTurn := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	length := len(s)
	l, r := 0, length-1
	bs := []byte(s)
	for l < r {
		for l < r && checkTurn[bs[l]] != true {
			l++
		}
		for l < r && checkTurn[bs[r]] != true {
			r--
		}
		// 交换
		if l < r {
			bs[l], bs[r] = bs[r], bs[l]
			l++
			r--
		}
	}
	return string(bs)
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 *
 * https://leetcode-cn.com/problems/letter-case-permutation/description/
 *
 * algorithms
 * Easy (59.15%)
 * Likes:    120
 * Dislikes: 0
 * Total Accepted:    11.9K
 * Total Submissions: 19.2K
 * Testcase Example:  '"a1b2"'
 *
 * 给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。
 *
 *
 * 示例:
 * 输入: S = "a1b2"
 * 输出: ["a1b2", "a1B2", "A1b2", "A1B2"]
 *
 * 输入: S = "3z4"
 * 输出: ["3z4", "3Z4"]
 *
 * 输入: S = "12345"
 * 输出: ["12345"]
 *
 *
 * 注意：
 *
 *
 * S 的长度不超过12。
 * S 仅由数字和字母组成。
 *
 *
 */

// @lc code=start
var ans []string

func letterCasePermutation(S string) []string {
	ans = []string{}
	// 将字符串转为byte数组
	buf := []byte(S)
	_search(0, buf)
	return ans
}

func _search(start int, buf []byte) {
	length := len(buf)
	if length == start {
		ans = append(ans, string(buf))
	} else {
		// 每一步根据是否是字母，可能有不止一种探索方式
		tmp := buf[start]
		if tmp >= 'a' && tmp <= 'z' {
			buf[start] = tmp - 32
			_search(start+1, buf)
		} else if tmp >= 'A' && tmp <= 'Z' {
			buf[start] = tmp + 32
			_search(start+1, buf)
		}

		// 什么都不变也是一种探索
		buf[start] = tmp
		_search(start+1, buf)
	}
}

// @lc code=end


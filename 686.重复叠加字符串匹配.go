import "strings"

/*
 * @lc app=leetcode.cn id=686 lang=golang
 *
 * [686] 重复叠加字符串匹配
 *
 * https://leetcode-cn.com/problems/repeated-string-match/description/
 *
 * algorithms
 * Easy (31.97%)
 * Likes:    51
 * Dislikes: 0
 * Total Accepted:    4.9K
 * Total Submissions: 15.4K
 * Testcase Example:  '"abcd"\n"cdabcdab"'
 *
 * 给定两个字符串 A 和 B, 寻找重复叠加字符串A的最小次数，使得字符串B成为叠加后的字符串A的子串，如果不存在则返回 -1。
 *
 * 举个例子，A = "abcd"，B = "cdabcdab"。
 *
 * 答案为 3， 因为 A 重复叠加三遍后为 “abcdabcdabcd”，此时 B 是其子串；A 重复叠加两遍后为"abcdabcd"，B
 * 并不是其子串。
 *
 * 注意:
 *
 * A 与 B 字符串的长度在1和10000区间范围内。
 *
 */

// @lc code=start
func repeatedStringMatch(A string, B string) int {
	i := 1
	C := A
	// 如果C小于B，肯定要叠加
	for len(C) < len(B) {
		C += A
		i++
	}

	// 检测是否子串
	if strings.Index(C, B) != -1 {
		return i
	}
	// 否则需要再加一次叠加来判断。理由就是B的尾部需要能完整遍历A, 因此仅仅比B长还不行，必须完整的长出至少一个A的length。
	C += A
	i++
	if strings.Index(C, B) != -1 {
		return i
	}

	return -1
}

// @lc code=end


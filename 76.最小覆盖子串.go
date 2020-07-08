/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode-cn.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (35.56%)
 * Likes:    384
 * Dislikes: 0
 * Total Accepted:    29.9K
 * Total Submissions: 84K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
 *
 * 示例：
 *
 * 输入: S = "ADOBECODEBANC", T = "ABC"
 * 输出: "BANC"
 *
 * 说明：
 *
 *
 * 如果 S 中不存这样的子串，则返回空字符串 ""。
 * 如果 S 中存在这样的子串，我们保证它是唯一的答案。
 *
 *
 */

// @lc code=start
package main

func minWindow(s string, t string) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	// 统计t的词频
	statMap := make(map[string]int)
	for _, ch := range t {
		statMap[string(ch)]++
	}

	// 滑动
	l, r := 0, 0
	curMap := make(map[string]int)
	curMap[string(s[0])]++
	minStr := ""
	for l < length && r < length {
		if isOk(curMap, statMap) {
			if minStr == "" || r-l+1 < len(minStr) {
				minStr = s[l : r+1]
			}
			curMap[string(s[l])]--
			l++
		} else {
			r++
			if r < length {
				curMap[string(s[r])]++
			}
		}
	}
	return minStr
}

func isOk(curMap, statMap map[string]int) bool {
	for key, target := range statMap {
		if curMap[key] < target {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
// }

// @lc code=end

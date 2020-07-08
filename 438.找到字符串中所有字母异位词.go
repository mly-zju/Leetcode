/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 *
 * https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Medium (41.96%)
 * Likes:    236
 * Dislikes: 0
 * Total Accepted:    20.1K
 * Total Submissions: 48K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
 * 
 * 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
 * 
 * 说明：
 * 
 * 
 * 字母异位词指字母相同，但排列不同的字符串。
 * 不考虑答案输出的顺序。
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入:
 * s: "cbaebabacd" p: "abc"
 * 
 * 输出:
 * [0, 6]
 * 
 * 解释:
 * 起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
 * 起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入:
 * s: "abab" p: "ab"
 * 
 * 输出:
 * [0, 1, 2]
 * 
 * 解释:
 * 起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
 * 起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
 * 起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
 * 
 * 
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	res := []int{}
	lens, lenp := len(s), len(p)
	if lens < lenp {
		return res
	}
	
	targetMap, curMap := make(map[byte]int), make(map[byte]int)
	// 初始化targetMap
	for _, ch := range p {
		targetMap[byte(ch)]++
	}

	for i := 0; i < lens; i++ {
		if i < lenp {
			curMap[s[i]]++
		} else {
			curMap[s[i-lenp]]--
			curMap[s[i]]++
		}

		if isOk(targetMap, curMap) {
			res = append(res, i-lenp+1)
		}
	}
	return res
}

func isOk(targetMap, curMap map[byte]int) bool {
	for key, target := range targetMap {
		if curMap[key] != target {
			return false
		}
	}
	return true
}
// @lc code=end


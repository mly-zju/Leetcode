import "strings"

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 *
 * https://leetcode-cn.com/problems/word-break/description/
 *
 * algorithms
 * Medium (42.73%)
 * Likes:    245
 * Dislikes: 0
 * Total Accepted:    26.1K
 * Total Submissions: 61.1K
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
 *
 * 说明：
 *
 *
 * 拆分时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 *
 *
 * 示例 1：
 *
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
 *
 *
 * 示例 2：
 *
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
 * 注意你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 *
 *
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// dp[i]代表在i结尾的字符串是否可以被匹配。则dp[i+wordLen] = (dp[i:wordLen] == word 且dp[i-1]为true)
	length := len(s)
	// 下标也是从1开始，0代表空字符串
	dp := make([]bool, length + 1)
	dp[0] = true

	for i := 1; i <= length; i++ {
		flag := false
		for _, word := range wordDict {
			wLen := len(word)
			if i - wLen >= 0 && dp[i-wLen] && word == s[i-wLen:i] {
				flag = true
				break
			}
		}
		dp[i] = flag
	}

	return dp[length]
}

// @lc code=end


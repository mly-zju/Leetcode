/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 *
 * https://leetcode-cn.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (26.89%)
 * Likes:    1116
 * Dislikes: 0
 * Total Accepted:    68.8K
 * Total Submissions: 251.4K
 * Testcase Example:  '"aa"\n"a"'
 *
 * 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
 * 
 * '.' 匹配任意单个字符
 * '*' 匹配零个或多个前面的那一个元素
 * 
 * 
 * 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
 * 
 * 说明:
 * 
 * 
 * s 可能为空，且只包含从 a-z 的小写字母。
 * p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * s = "aa"
 * p = "a"
 * 输出: false
 * 解释: "a" 无法匹配 "aa" 整个字符串。
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * s = "aa"
 * p = "a*"
 * 输出: true
 * 解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
 * 
 * 
 * 示例 3:
 * 
 * 输入:
 * s = "ab"
 * p = ".*"
 * 输出: true
 * 解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
 * 
 * 
 * 示例 4:
 * 
 * 输入:
 * s = "aab"
 * p = "c*a*b"
 * 输出: true
 * 解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
 * 
 * 
 * 示例 5:
 * 
 * 输入:
 * s = "mississippi"
 * p = "mis*is*p*."
 * 输出: false
 * 
 */

// @lc code=start
func isMatch(s string, p string) bool {
	// 难点在于*的匹配，需要将*和前一个x看做一个整体x*
	lens, lenp := len(s), len(p)

	// dp[i][j]需要处理边界条件
	dp := [][]bool{}
	for i := 0; i <= lens; i++ {
		dp = append(dp, make([]bool, lenp+1))
	}
	
	// 执行初始化
	dp[0][0] = true
	// s不为0 p为0的情况肯定false，不用管
	// p为0s不为0的情况，需要看p是不是可以自抵消
	for j := 2; j <= lenp; j++ {
		// 自抵消条件是本身是*, 且上上一个也是true
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= lens; i++ {
		for j := 1; j <= lenp; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if j >= 2 && p[j-1] == '*' {
				// 如果当前不相等，看看是不是星号
				// 星号的情况下，c*看做一个整体，采纳一次，多次，或者直接丢弃
				// 采纳至少一次的前提是前一个和当前要相等
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					// 如果只匹配一次，那么*等于忽略了，需要dp[i][j-1]可以匹配
					onceRes := dp[i][j-1]
					// 如果匹配多次，那么需要s[i-1]本身就等于s[i-2], 又由于星号可重复的缘故，需要dp[i-1][j]匹配
					multiRes := dp[i-1][j]
					dp[i][j] = onceRes || multiRes
				}
				
				// 然后继续尝试丢弃，因为及时前一个s和当前p相等，也可能不满足采纳的条件，需要看丢弃
				dp[i][j] = dp[i][j] || dp[i][j-2]
			}
		}
	}

	return dp[lens][lenp]
}
// @lc code=end


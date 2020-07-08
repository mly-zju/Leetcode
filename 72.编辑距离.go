/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 *
 * https://leetcode-cn.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (56.64%)
 * Likes:    772
 * Dislikes: 0
 * Total Accepted:    50.8K
 * Total Submissions: 85.6K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
 * 
 * 你可以对一个单词进行如下三种操作：
 * 
 * 
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 * 
 * 
 * 示例 2：
 * 
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 * 
 * 
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	// 处理边界条件
	if len1 == 0 {
		return len2
	} else if len2 == 0 {
		return len1
	}

	dp := [][]int{}
	// 初始化, 为了处理方便，下标从1开始, 0代表空
	for i := 0; i <= len1; i++ {
		dp = append(dp, make([]int, len2+1))
	}
	// 初始化0:x和x:0, 即某个l长的要转为kong，那改动的自然是l长
	for i := 0; i <= len1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 1. 改
				changeNum := dp[i-1][j-1] + 1
				// 2. 给word1增
				addNum := dp[i][j-1] + 1
				// 3. 给word1减
				delNum := dp[i-1][j] + 1

				dp[i][j] = getMin(changeNum, getMin(addNum, delNum))
			}
		}
	}

	return dp[len1][len2]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (34.82%)
 * Likes:    884
 * Dislikes: 0
 * Total Accepted:    188.6K
 * Total Submissions: 521.2K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 * 
 * 如果不存在公共前缀，返回空字符串 ""。
 * 
 * 示例 1:
 * 
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 * 
 * 
 * 示例 2:
 * 
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 * 
 * 
 * 说明:
 * 
 * 所有输入只包含小写字母 a-z 。
 * 
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	// 最简单的就是两两对比一遍，O(n^2)
	// 这里用更简单方法：按位比较，遇到不全等的就返回
	length := len(strs)
	if length == 0 {
		return ""
	}

	// 寻找字符串最短长度
	minLen := len(strs[0])
	for i := 1; i < length; i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

	sameCount := 0
	for i := 0; i < minLen; i++ {
		sameFlag := true
		for j := 1; j < length; j++ {
			if strs[j][i] != strs[0][i] {
				sameFlag = false
				break
			}
		}

		if sameFlag {
			sameCount++
		} else {
			break
		}
	}

	return strs[0][0:sameCount]
}

// @lc code=end


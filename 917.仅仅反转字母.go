/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 *
 * https://leetcode-cn.com/problems/reverse-only-letters/description/
 *
 * algorithms
 * Easy (50.26%)
 * Likes:    28
 * Dislikes: 0
 * Total Accepted:    5.6K
 * Total Submissions: 11.1K
 * Testcase Example:  '"ab-cd"'
 *
 * 给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入："ab-cd"
 * 输出："dc-ba"
 *
 *
 * 示例 2：
 *
 * 输入："a-bC-dEf-ghIj"
 * 输出："j-Ih-gfE-dCba"
 *
 *
 * 示例 3：
 *
 * 输入："Test1ng-Leet=code-Q!"
 * 输出："Qedo1ct-eeLg=ntse-T!"
 *
 *
 *
 *
 * 提示：
 *
 *
 * S.length <= 100
 * 33 <= S[i].ASCIIcode <= 122
 * S 中不包含 \ or "
 *
 *
 */

// @lc code=start
func reverseOnlyLetters(S string) string {
	sArr := []byte(S)
	l, r := 0, len(sArr)-1
	for l < r {
		for !isAlpha(sArr[l]) && l < r {
			l++
		}
		for !isAlpha(sArr[r]) && l < r {
			r--
		}
		if l < r {
			sArr[l], sArr[r] = sArr[r], sArr[l]
			l++
			r--
		}
	}
	return string(sArr)
}

func isAlpha(ch byte) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=1071 lang=golang
 *
 * [1071] 字符串的最大公因子
 *
 * https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/description/
 *
 * algorithms
 * Easy (46.75%)
 * Likes:    17
 * Dislikes: 0
 * Total Accepted:    2.5K
 * Total Submissions: 5.4K
 * Testcase Example:  '"ABCABC"\n"ABC"'
 *
 * 对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。
 *
 * 返回字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。
 *
 *
 *
 * 示例 1：
 *
 * 输入：str1 = "ABCABC", str2 = "ABC"
 * 输出："ABC"
 *
 *
 * 示例 2：
 *
 * 输入：str1 = "ABABAB", str2 = "ABAB"
 * 输出："AB"
 *
 *
 * 示例 3：
 *
 * 输入：str1 = "LEET", str2 = "CODE"
 * 输出：""
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= str1.length <= 1000
 * 1 <= str2.length <= 1000
 * str1[i] 和 str2[i] 为大写英文字母
 *
 *
 */
package main

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {
	// 1. 模拟辗转相除
	// modStr := getMod(str1, str2)
	// if modStr == "" {
	// 	return str2
	// } else if modStr == "a" {
	// 	return ""
	// } else {
	// 	return gcdOfStrings(str2, modStr)
	// }

	// 2. 解法2：先判断是否有解，如果有，找两者字符串长度的最大公因子即可
	if str1+str2 != str2+str1 {
		return ""
	} else {
		maxLen := gcd(len(str1), len(str2))
		return str1[0:maxLen]
	}
}

func gcd(a, b int) int {
	if a == b {
		return a
	} else if a > b {
		return gcd(b, a-b)
	} else {
		return gcd(a, b-a)
	}
}

func getMod(str1, str2 string) string {
	len1, len2 := len(str1), len(str2)
	if len1 < len2 {
		return str1
	}

	sub := ""
	for i := 0; i < len1; i += len2 {
		sub = str1[i:len2]
		if sub != str2 {
			if i == 0 {
				// 如果i为0，说明整体都不匹配, 返回小写a代表整体不匹配
				return "a"
			} else {
				return str1[i:]
			}
		}
	}
	return ""
}

// func main() {
// 	str1, str2 := "LEET", "CODE"
// 	gcd := gcdOfStrings(str1, str2)
// 	fmt.Println(gcd)
// }

// @lc code=end

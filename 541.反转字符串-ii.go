/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 *
 * https://leetcode-cn.com/problems/reverse-string-ii/description/
 *
 * algorithms
 * Easy (48.92%)
 * Likes:    46
 * Dislikes: 0
 * Total Accepted:    7.7K
 * Total Submissions: 15.6K
 * Testcase Example:  '"abcdefg"\n2'
 *
 * 给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。如果剩余少于 k
 * 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。
 *
 * 示例:
 *
 *
 * 输入: s = "abcdefg", k = 2
 * 输出: "bacdfeg"
 *
 *
 * 要求:
 *
 *
 * 该字符串只包含小写的英文字母。
 * 给定字符串的长度和 k 在[1, 10000]范围内。
 *
 *
 */
package main

import (
	"strings"
)

// @lc code=start
func reverseStr(s string, k int) string {
	strArr := strings.Split(s, "")
	gap := 2 * k
	for i, length := 0, len(s); i < length; i++ {
		if i%gap != 0 {
			continue
		}

		if length-i < k {
			// 判断剩余是否少于k, 如果是则全部反转
			reversePart(strArr, i, length-1)
		} else {
			// 否则反转k个
			reversePart(strArr, i, i+k-1)
		}
	}
	return strings.Join(strArr, "")
}

func reversePart(s []string, start int, end int) {
	i := 0
	for true {
		if start+i >= end-i {
			return
		} else {
			s[start+i], s[end-i] = s[end-i], s[start+i]
			i++
		}
	}
}

// func main() {
// 	s := "abcdefg"
// 	k := 2
// 	s2 := reverseStr(s, k)
// 	fmt.Println(s2)
// }

// @lc code=end

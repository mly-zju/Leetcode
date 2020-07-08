/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (67.66%)
 * Likes:    123
 * Dislikes: 0
 * Total Accepted:    30.7K
 * Total Submissions: 45.1K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
 *
 * 示例 1:
 *
 *
 * 输入: "Let's take LeetCode contest"
 * 输出: "s'teL ekat edoCteeL tsetnoc"
 *
 *
 * 注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
 *
 */

// @lc code=start
package main

func reverseWords(s string) string {
	// 1. 遍历
	cArr := []byte(s)
	start, end := 0, 0
	for i, length := 0, len(cArr); i < length; i++ {
		if cArr[i] == ' ' {
			end = i - 1
			reversePart(cArr, start, end)
			start = i + 1
		} else if i == length-1 {
			end = i
			reversePart(cArr, start, end)
		}
	}
	return string(cArr)
}

func reversePart(cArr []byte, start, end int) {
	if start >= end {
		return
	}

	i := 0
	for start+i < end-i {
		cArr[start+i], cArr[end-i] = cArr[end-i], cArr[start+i]
		i++
	}
}

// func main() {
// 	b := reverseWords("Let's take Leetcode contest")
// 	fmt.Println(b)
// }

// @lc code=end

/*
 * @lc app=leetcode.cn id=696 lang=golang
 *
 * [696] 计数二进制子串
 *
 * https://leetcode-cn.com/problems/count-binary-substrings/description/
 *
 * algorithms
 * Easy (47.01%)
 * Likes:    100
 * Dislikes: 0
 * Total Accepted:    5.1K
 * Total Submissions: 10.7K
 * Testcase Example:  '"00110"'
 *
 * 给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。
 *
 * 重复出现的子串要计算它们出现的次数。
 *
 * 示例 1 :
 *
 *
 * 输入: "00110011"
 * 输出: 6
 * 解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。
 *
 * 请注意，一些重复出现的子串要计算它们出现的次数。
 *
 * 另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
 *
 *
 * 示例 2 :
 *
 *
 * 输入: "10101"
 * 输出: 4
 * 解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
 *
 *
 * 注意：
 *
 *
 * s.length 在1到50,000之间。
 * s 只包含“0”或“1”字符。
 *
 *
 */
package main

// @lc code=start
func countBinarySubstrings(s string) int {
	// 1. 暴力破解(行不通，时间超时了)
	// sArr := []byte(s)
	// length := len(sArr)
	// count := 0
	// for gap := 2; gap <= length; gap += 2 {
	// 	for i := 0; i <= length-gap; i++ {
	// 		if isNeed(sArr, i, gap) {
	// 			count++
	// 		}
	// 	}
	// }
	// return count

	// prev, cur, count := 0, 0, 0
	// sArr := []byte(s)
	// prevC := sArr[0]
	// // 扫描，找出连续数字的分界点, 不等，或者到了最后一个元素
	// for _, b := range sArr {
	// 	if b == prevC {
	// 		cur++
	// 	} else {
	// 		count += getMin(prev, cur)
	// 		prevC = b
	// 		prev = cur
	// 		cur = 1
	// 	}
	// }

	// // 最后结束时候要再比较一次
	// count += getMin(prev, cur)
	// return count

	// 3. 基于统计数组
	sArr := []int{}
	prevC := rune(s[0])
	count := 0
	for _, ch := range s {
		if ch == prevC {
			count++
		} else {
			sArr = append(sArr, count)
			count = 1
			prevC = ch
		}
	}
	// 结束时候要再统计一次
	sArr = append(sArr, count)

	res := 0
	for i, length := 0, len(sArr); i < length-1; i++ {
		res += getMin(sArr[i], sArr[i+1])
	}
	return res
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func isNeed(sArr []byte, i, gap int) bool {
	mid := gap/2 + i
	var l, r byte
	if sArr[i] == '0' {
		l = '0'
		r = '1'
	} else {
		l = '1'
		r = '0'
	}

	for j := i; j < mid; j++ {
		if (sArr[j] == l) && (sArr[j-i+mid] == r) {
			continue
		} else {
			return false
		}
	}
	return true
}

// func main() {
// 	s := "10101"
// 	r := countBinarySubstrings(s)
// 	fmt.Println(r)
// }

// @lc code=end

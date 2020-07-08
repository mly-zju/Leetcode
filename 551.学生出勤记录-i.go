package main

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=551 lang=golang
 *
 * [551] 学生出勤记录 I
 *
 * https://leetcode-cn.com/problems/student-attendance-record-i/description/
 *
 * algorithms
 * Easy (48.85%)
 * Likes:    19
 * Dislikes: 0
 * Total Accepted:    7.8K
 * Total Submissions: 15.9K
 * Testcase Example:  '"PPALLP"'
 *
 * 给定一个字符串来代表一个学生的出勤记录，这个记录仅包含以下三个字符：
 *
 *
 * 'A' : Absent，缺勤
 * 'L' : Late，迟到
 * 'P' : Present，到场
 *
 *
 * 如果一个学生的出勤记录中不超过一个'A'(缺勤)并且不超过两个连续的'L'(迟到),那么这个学生会被奖赏。
 *
 * 你需要根据这个学生的出勤记录判断他是否会被奖赏。
 *
 * 示例 1:
 *
 * 输入: "PPALLP"
 * 输出: True
 *
 *
 * 示例 2:
 *
 * 输入: "PPALLL"
 * 输出: False
 *
 *
 */

// @lc code=start
func checkRecord(s string) bool {
	// 记录前一个是否是l
	// lPrev := false
	// aNum, lNum, lMax := 0, 0, 0

	// for _, bt := range s {
	// 	if bt == 'L' {
	// 		// 如果是L，则给当前lNum++, lPrev设置true，代表前一个是L
	// 		lNum++
	// 		lPrev = true
	// 	} else {
	// 		// 如果不是L, 如果之前一个是L，则看当前最长是多少
	// 		if lPrev {
	// 			if lNum > lMax {
	// 				lMax = lNum
	// 			}
	// 		}
	// 		if bt == 'A' {
	// 			aNum++
	// 		}
	// 		lPrev = false
	// 		lNum = 0
	// 	}
	// }

	// // 最后结束时候再看看L的连续长度
	// if lPrev && (lNum > lMax) {
	// 	lMax = lNum
	// }

	// return (aNum <= 1) && (lMax <= 2)

	// 2. 方法2：统计A个数，并证明LLL是否是子串
	aNum := 0
	for i, length := 0, len(s); i < length; i++ {
		if s[i] == 'A' {
			aNum++
			if aNum > 1 {
				return false
			}
		}
	}

	if strings.Index(s, "LLL") == -1 {
		return true
	}
	return false
}

// @lc code=end

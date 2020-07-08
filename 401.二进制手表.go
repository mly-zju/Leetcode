package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 *
 * https://leetcode-cn.com/problems/binary-watch/description/
 *
 * algorithms
 * Easy (49.27%)
 * Likes:    97
 * Dislikes: 0
 * Total Accepted:    9.1K
 * Total Submissions: 17.8K
 * Testcase Example:  '0'
 *
 * 二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。
 *
 * 每个 LED 代表一个 0 或 1，最低位在右侧。
 *
 *
 *
 * 例如，上面的二进制手表读取 “3:25”。
 *
 * 给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。
 *
 * 案例:
 *
 *
 * 输入: n = 1
 * 返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16",
 * "0:32"]
 *
 *
 *
 * 注意事项:
 *
 *
 * 输出的顺序没有要求。
 * 小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
 * 分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。
 *
 *
 */

// @lc code=start
var totalArr []string

func readBinaryWatch(num int) []string {
	// 回溯法，类似八皇后问题，探索所有可以放置的可能性
	// 主要设计：探索圆满终止条件、当前一步探索不满足，返回上一步条件

	// 全局变量存放答案
	totalArr = []string{}
	// 探索路径设计为一个10位数组，点亮的部分为1，否则为0
	buf := make([]int, 10)
	_search(0, buf, num)
	return totalArr
}

func _search(start int, buf []int, num int) {
	if start == 10 {
		// 如果已经亮了num个值，push答案
		if getUsed(buf) == num {
			totalArr = append(totalArr, getStr(buf))
		}
	} else {
		// 每一位都有放或者不放两种选择
		for i := 0; i <= 1; i++ {
			buf[start] = i
			if isOk(start, buf, num) {
				_search(start+1, buf, num)
			}
			// 每一次假设完，要还原为初始状态
			buf[start] = 0
		}
	}
}

func isOk(start int, buf []int, num int) bool {
	// 如果当前小时级别大于11，或者分钟级别大于59，则不ok
	hour, minite := 0, 0
	for i := 0; i < 10; i++ {
		if i < 4 {
			hour = hour*2 + buf[i]
		} else {
			minite = minite*2 + buf[i]
		}
	}
	if hour > 11 || minite > 59 {
		return false
	}

	return true
}

func getUsed(buf []int) int {
	res := 0
	for _, val := range buf {
		if val == 1 {
			res += 1
		}
	}
	return res
}

func getStr(buf []int) string {
	hour := 0
	minite := 0
	for i := 0; i < 10; i++ {
		if i < 4 {
			hour = hour*2 + buf[i]
		} else {
			minite = minite*2 + buf[i]
		}
	}

	res := strconv.Itoa(hour) + ":"

	if minite < 10 {
		res = res + "0" + strconv.Itoa(minite)
	} else {
		res = res + strconv.Itoa(minite)
	}
	return res
}

// func main() {
// 	fmt.Println(readBinaryWatch(1))
// }

// @lc code=end

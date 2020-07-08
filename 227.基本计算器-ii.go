/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 *
 * https://leetcode-cn.com/problems/basic-calculator-ii/description/
 *
 * algorithms
 * Medium (34.79%)
 * Likes:    113
 * Dislikes: 0
 * Total Accepted:    12.6K
 * Total Submissions: 36.1K
 * Testcase Example:  '"3+2*2"'
 *
 * 实现一个基本的计算器来计算一个简单的字符串表达式的值。
 *
 * 字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
 *
 * 示例 1:
 *
 * 输入: "3+2*2"
 * 输出: 7
 *
 *
 * 示例 2:
 *
 * 输入: " 3/2 "
 * 输出: 1
 *
 * 示例 3:
 *
 * 输入: " 3+5 / 2 "
 * 输出: 5
 *
 *
 * 说明：
 *
 *
 * 你可以假设所给定的表达式都是有效的。
 * 请不要使用内置的库函数 eval。
 *
 *
 */

// @lc code=start
func calculate(s string) int {
	cache := []int{}
	op := ""
	cacheStr := ""

	for i, length := 0, len(s); i < length; i++ {
		// 忽略空字符串
		if isNumber(s[i]) {
			cacheStr += string(s[i])
		} else if isOp(s[i]) {
			// 如果遇到了运算符，则停止下来，执行处理数字的运算
			curNum, _ := strconv.Atoi(cacheStr)
			cacheStr = ""

			cache = dealNum(cache, curNum, op)

			op = string(s[i])
		}
	}
	
	// 最后还要再处理一次
	curNum, _ := strconv.Atoi(cacheStr)
	cache = dealNum(cache, curNum, op)

	res := 0
	for _, num := range cache {
		res += num
	}
	return res
}

func dealNum(cache []int, curNum int, op string) []int {
	cacheLen := len(cache)
	if op == "+" || op == "" {
		cache = append(cache, curNum)
	} else if op == "-" {
		cache = append(cache, -curNum)
	} else if op == "*" {
		cache[cacheLen-1] *= curNum
	} else {
		cache[cacheLen-1] /= curNum
	}
	return cache
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func isOp(c byte) bool {
	return c == '+' || c == '-' || c == '/' || c == '*'
}


// @lc code=end

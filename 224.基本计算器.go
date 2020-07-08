/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 *
 * https://leetcode-cn.com/problems/basic-calculator/description/
 *
 * algorithms
 * Hard (37.05%)
 * Likes:    157
 * Dislikes: 0
 * Total Accepted:    11.2K
 * Total Submissions: 30.1K
 * Testcase Example:  '"1 + 1"'
 *
 * 实现一个基本的计算器来计算一个简单的字符串表达式的值。
 *
 * 字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。
 *
 * 示例 1:
 *
 * 输入: "1 + 1"
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: " 2-1 + 2 "
 * 输出: 3
 *
 * 示例 3:
 *
 * 输入: "(1+(4+5+2)-3)+(6+8)"
 * 输出: 23
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
package main

import (
	"strconv"
)

func calculate(s string) int {
	x, _ := _calculate(s, 0)
	return x
}

// 返回1：四则运算结果，返回2：括号结束的s索引
func _calculate(s string, i int) (int, int) {
	stack := []int{}
	cacheStr := ""
	op := ""

	length := len(s)
	for i < length {
		if isNumber(s[i]) {
			cacheStr += string(s[i])
			i++
		} else if isOp(s[i]) {
			stack = dealNum(stack, cacheStr, op)
			cacheStr = ""
			op = string(s[i])
			i++
		} else if s[i] == ')' {
			// 遇到结尾括号，说明需要返回当前结果
			stack = dealNum(stack, cacheStr, op)
			res := 0
			for _, num := range stack {
				res += num
			}
			return res, i
		} else if s[i] == '(' {
			innerVal, endIndex := _calculate(s, i+1)
			// 将innerVal转为cacheStr对应的字符串
			cacheStr = strconv.Itoa(innerVal)
			i = endIndex + 1
		} else {
			// 否则是空格，忽略
			i++
		}
	}

	// 最后再执行一次
	stack = dealNum(stack, cacheStr, op)
	res := 0
	for _, num := range stack {
		res += num
	}
	return res, i
}

// 本题只需要处理加减，但是加上乘除也很简单，这里给出了例子
func dealNum(stack []int, cacheStr string, op string) []int {
	curNum, _ := strconv.Atoi(cacheStr)
	sLen := len(stack)
	if op == "+" || op == "" {
		stack = append(stack, curNum)
	} else if op == "-" {
		stack = append(stack, -curNum)
	} else if op == "*" {
		stack[sLen-1] *= curNum
	} else {
		stack[sLen-1] /= curNum
	}
	return stack
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func isOp(c byte) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

// func main() {
// 23 is right
// fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
// 2 is right
// fmt.Println(calculate("1+2*3 - 10/(1+1)"))
// }

// @lc code=end

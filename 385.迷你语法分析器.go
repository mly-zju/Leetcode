import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=385 lang=golang
 *
 * [385] 迷你语法分析器
 *
 * https://leetcode-cn.com/problems/mini-parser/description/
 *
 * algorithms
 * Medium (40.38%)
 * Likes:    12
 * Dislikes: 0
 * Total Accepted:    1.1K
 * Total Submissions: 2.8K
 * Testcase Example:  '"324"'
 *
 * 给定一个用字符串表示的整数的嵌套列表，实现一个解析它的语法分析器。
 *
 * 列表中的每个元素只可能是整数或整数嵌套列表
 *
 * 提示：你可以假定这些字符串都是格式良好的：
 *
 *
 * 字符串非空
 * 字符串不包含空格
 * 字符串只包含数字0-9, [, - ,, ]
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 给定 s = "324",
 *
 * 你应该返回一个 NestedInteger 对象，其中只包含整数值 324。
 *
 *
 *
 *
 * 示例 2：
 *
 *
 * 给定 s = "[123,[456,[789]]]",
 *
 * 返回一个 NestedInteger 对象包含一个有两个元素的嵌套列表：
 *
 * 1. 一个 integer 包含值 123
 * 2. 一个包含两个元素的嵌套列表：
 * ⁠   i.  一个 integer 包含值 456
 * ⁠   ii. 一个包含一个元素的嵌套列表
 * ⁠        a. 一个 integer 包含值 789
 *
 *
 *
 *
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	length := len(s)
	if length == 0 {
		return nil
	}

	stack := []NestedInteger{}
	cache := ""
	for i := 0; i < length; i++ {
		if s[i] == '[' {
			// 每当遇到开始括号，要新建
			stack = append(stack, NestedInteger{})
		} else if s[i] == ',' || s[i] == ']' {
			// 遇到逗号或者结尾括号，解析cache
			sLen := len(stack)
			if cache != "" {
				number, _ := strconv.Atoi(cache)
				cache = ""
				tmp := NestedInteger{}
				tmp.SetInteger(number)
				stack[sLen - 1].Add(tmp)
			}

			// 同时如果是结尾括号，需要加到上上一个元素中(如果存在的话), 同时把最新元素去掉
			if s[i] == ']' && sLen >= 2 {
				stack[sLen - 2].Add(stack[sLen - 1])
				stack = stack[0:sLen - 1]
			}
		} else {
			cache += string(s[i])
		}
	}

	// 如果cache不为空，说明最外层没有括号
	if cache != "" {
		number, _ := strconv.Atoi(cache)
		tmp := &NestedInteger{}
		tmp.SetInteger(number)
		return tmp
	}

	return &stack[0]
}

// @lc code=end


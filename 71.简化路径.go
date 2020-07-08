package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 *
 * https://leetcode-cn.com/problems/simplify-path/description/
 *
 * algorithms
 * Medium (37.59%)
 * Likes:    75
 * Dislikes: 0
 * Total Accepted:    17.4K
 * Total Submissions: 46.1K
 * Testcase Example:  '"/home/"'
 *
 * 以 Unix 风格给出一个文件的绝对路径，你需要简化它。或者换句话说，将其转换为规范路径。
 *
 * 在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..）
 * 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。更多信息请参阅：Linux / Unix中的绝对路径 vs 相对路径
 *
 * 请注意，返回的规范路径必须始终以斜杠 / 开头，并且两个目录名之间必须只有一个斜杠 /。最后一个目录名（如果存在）不能以 /
 * 结尾。此外，规范路径必须是表示绝对路径的最短字符串。
 *
 *
 *
 * 示例 1：
 *
 * 输入："/home/"
 * 输出："/home"
 * 解释：注意，最后一个目录名后面没有斜杠。
 *
 *
 * 示例 2：
 *
 * 输入："/../"
 * 输出："/"
 * 解释：从根目录向上一级是不可行的，因为根是你可以到达的最高级。
 *
 *
 * 示例 3：
 *
 * 输入："/home//foo/"
 * 输出："/home/foo"
 * 解释：在规范路径中，多个连续斜杠需要用一个斜杠替换。
 *
 *
 * 示例 4：
 *
 * 输入："/a/./b/../../c/"
 * 输出："/c"
 *
 *
 * 示例 5：
 *
 * 输入："/a/../../b/../c//.//"
 * 输出："/c"
 *
 *
 * 示例 6：
 *
 * 输入："/a//b////c/d//././/.."
 * 输出："/a/b/c"
 *
 */

// @lc code=start
func simplifyPath(path string) string {
	// strArr := strings.Split(path, "/")
	// res := []string{}
	// for _, val := range strArr {
	// 	if val == ".." {
	// 		length := len(res)
	// 		if length > 0 {
	// 			res = res[0 : len(res)-1]
	// 		}
	// 	} else if val == "." {
	// 		continue
	// 	} else if val != "" {
	// 		res = append(res, val)
	// 	}
	// }
	// resStr := ""
	// for _, val := range res {
	// 	resStr = resStr + "/" + val
	// }
	// if resStr == "" {
	// 	resStr = "/"
	// }
	// return resStr

	// 基于栈，以'/'为检测标识，然后对内容执行压栈/出栈/空白操作，整体结束之后，从栈中剩余元素拼接字符串
	stack := []string{}
	cache := ""
	for i, length := 0, len(path); i < length; i++ {
		if path[i] == '/' {
			// sLen := len(stack)
			// if sLen == 0 && cache != "" && cache != "." {
			// 	stack = append(stack, cache)
			// } else {
			// 	if cache == "" || cache == "." {
			// 	} else if cache == ".." {
			// 		stack = stack[0 : sLen-1]
			// 	} else {
			// 		stack = append(stack, cache)
			// 	}
			// }

			stack = dealCache(stack, cache)

			// 清空cache
			cache = ""
		} else {
			cache += string(path[i])
		}
	}

	// 最后还要再处理一下
	stack = dealCache(stack, cache)

	// 组装
	res := ""
	for _, str := range stack {
		res += "/" + str
	}
	if res == "" {
		return "/"
	}
	return res
}

func dealCache(stack []string, cache string) []string {
	sLen := len(stack)
	if cache == "" || cache == "." {
		// cache为空，do nothing
	} else if cache == ".." {
		if sLen > 0 {
			stack = stack[0 : sLen-1]
		}
	} else {
		// 普通字符的情况下，还要处理一下"."的情况
		if sLen > 0 && stack[sLen-1] == "." {
			stack[sLen-1] = cache
		} else {
			stack = append(stack, cache)
		}
	}
	return stack
}

// func main() {
// 	s := "/a//b////c/d//././/.."
// 	fmt.Println(simplifyPath(s))
// }

// @lc code=end

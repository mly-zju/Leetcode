import "bytes"

/*
 * @lc app=leetcode.cn id=709 lang=golang
 *
 * [709] 转换成小写字母
 *
 * https://leetcode-cn.com/problems/to-lower-case/description/
 *
 * algorithms
 * Easy (74.38%)
 * Likes:    97
 * Dislikes: 0
 * Total Accepted:    29.5K
 * Total Submissions: 39.4K
 * Testcase Example:  '"Hello"'
 *
 * 实现函数 ToLowerCase()，该函数接收一个字符串参数 str，并将该字符串中的大写字母转换成小写字母，之后返回新的字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: "Hello"
 * 输出: "hello"
 *
 * 示例 2：
 *
 *
 * 输入: "here"
 * 输出: "here"
 *
 * 示例 3：
 *
 *
 * 输入: "LOVELY"
 * 输出: "lovely"
 *
 *
 */

// @lc code=start
func toLowerCase(str string) string {
	// sArr := []byte(str)
	// for i, length := 0, len(sArr); i < length; i++ {
	// 	if sArr[i] >= 'A' && sArr[i] <= 'Z' {
	// 		sArr[i] = sArr[i] + 32
	// 	}
	// }
	// return string(sArr)
	sArr := bytes.ToLower([]byte(str))
	return string(sArr)
}

// @lc code=end


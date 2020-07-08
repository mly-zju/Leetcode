package main

/*
 * @lc app=leetcode.cn id=467 lang=golang
 *
 * [467] 环绕字符串中唯一的子字符串
 *
 * https://leetcode-cn.com/problems/unique-substrings-in-wraparound-string/description/
 *
 * algorithms
 * Medium (35.53%)
 * Likes:    38
 * Dislikes: 0
 * Total Accepted:    1.5K
 * Total Submissions: 3.9K
 * Testcase Example:  '"a"'
 *
 * 把字符串 s 看作是“abcdefghijklmnopqrstuvwxyz”的无限环绕字符串，所以 s
 * 看起来是这样的："...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd....".
 *
 * 现在我们有了另一个字符串 p 。你需要的是找出 s 中有多少个唯一的 p 的非空子串，尤其是当你的输入是字符串 p ，你需要输出字符串 s 中 p
 * 的不同的非空子串的数目。
 *
 * 注意: p 仅由小写的英文字母组成，p 的大小可能超过 10000。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: "a"
 * 输出: 1
 * 解释: 字符串 S 中只有一个"a"子字符。
 *
 *
 *
 *
 * 示例 2:
 *
 *
 * 输入: "cac"
 * 输出: 2
 * 解释: 字符串 S 中的字符串“cac”只有两个子串“a”、“c”。.
 *
 *
 *
 *
 * 示例 3:
 *
 *
 * 输入: "zab"
 * 输出: 6
 * 解释: 在字符串 S 中有六个子串“z”、“a”、“b”、“za”、“ab”、“zab”。.
 *
 *
 *
 *
 */

// @lc code=start
func findSubstringInWraproundString(p string) int {
	// 记录以每个字符结尾的最长长度, 最大26，超出26的就会有重复了
	maxAbc := map[byte]int{}
	// dp记录以第i位字符结尾的最长连续字符串长度
	dp := []int{}

	length := len(p)
	if length == 0 || length == 1 {
		return length
	}

	for i := 0; i < length; i++ {
		if i == 0 {
			dp = append(dp, 1)
		} else {
			if p[i]-p[i-1] == 1 || (p[i-1] == 'z' && p[i] == 'a') {
				dp = append(dp, dp[i-1]+1)
			} else {
				dp = append(dp, 1)
			}
		}

		if dp[i] > maxAbc[p[i]] {
			maxAbc[p[i]] = dp[i]
		}
	}

	res := 0
	for _, val := range maxAbc {
		res += val
	}
	return res
}

// func main() {
// 	a := "zab"
// 	fmt.Println(findSubstringInWraproundString(a))
// }

// @lc code=end

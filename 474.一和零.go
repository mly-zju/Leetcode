package main

/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 *
 * https://leetcode-cn.com/problems/ones-and-zeroes/description/
 *
 * algorithms
 * Medium (44.86%)
 * Likes:    99
 * Dislikes: 0
 * Total Accepted:    4.6K
 * Total Submissions: 9.5K
 * Testcase Example:  '["10","0001","111001","1","0"]\n5\n3'
 *
 * 在计算机界中，我们总是追求用有限的资源获取最大的收益。
 *
 * 现在，假设你分别支配着 m 个 0 和 n 个 1。另外，还有一个仅包含 0 和 1 字符串的数组。
 *
 * 你的任务是使用给定的 m 个 0 和 n 个 1 ，找到能拼出存在于数组中的字符串的最大数量。每个 0 和 1 至多被使用一次。
 *
 * 注意:
 *
 *
 * 给定 0 和 1 的数量都不会超过 100。
 * 给定字符串数组的长度不会超过 600。
 *
 *
 * 示例 1:
 *
 *
 * 输入: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
 * 输出: 4
 *
 * 解释: 总共 4 个字符串可以通过 5 个 0 和 3 个 1 拼出，即 "10","0001","1","0" 。
 *
 *
 * 示例 2:
 *
 *
 * 输入: Array = {"10", "0", "1"}, m = 1, n = 1
 * 输出: 2
 *
 * 解释: 你可以拼出 "10"，但之后就没有剩余数字了。更好的选择是拼出 "0" 和 "1" 。
 *
 *
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	// 多维背包问题
	length := len(strs)
	if length == 0 {
		return 0
	}

	// dp[i][j][k]代表选取0-i个商品时候, 0的容量未j，1的容量为k时候能装物品的最大值
	// dp[i][j][k] = getMax(dp[i-1][j][k], dp[i-1][j-num0[i]][k-num1[i]] + 1)
	dp := [][]int{}
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}

	for i := 0; i < length; i++ {
		num0, num1 := getNum(strs[i])
		for j := m; j >= num0; j-- {
			for k := n; k >= num1; k-- {
				dp[j][k] = getMax(dp[j][k], dp[j-num0][k-num1]+1)
			}
		}
	}
	return dp[m][n]
}

func getNum(a string) (int, int) {
	zeroNum, oneNum := 0, 0
	for i, length := 0, len(a); i < length; i++ {
		if a[i] == '0' {
			zeroNum++
		} else {
			oneNum++
		}
	}

	return zeroNum, oneNum
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	strs := []string{"10", "0", "1"}
// 	m, n := 1, 1
// 	fmt.Println(findMaxForm(strs, m, n))
// }

// @lc code=end

package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=1025 lang=golang
 *
 * [1025] 除数博弈
 *
 * https://leetcode-cn.com/problems/divisor-game/description/
 *
 * algorithms
 * Easy (68.39%)
 * Likes:    49
 * Dislikes: 0
 * Total Accepted:    9.7K
 * Total Submissions: 14.2K
 * Testcase Example:  '2'
 *
 * 爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。
 *
 * 最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：
 *
 *
 * 选出任一 x，满足 0 < x < N 且 N % x == 0 。
 * 用 N - x 替换黑板上的数字 N 。
 *
 *
 * 如果玩家无法执行这些操作，就会输掉游戏。
 *
 * 只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 false。假设两个玩家都以最佳状态参与游戏。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：2
 * 输出：true
 * 解释：爱丽丝选择 1，鲍勃无法进行操作。
 *
 *
 * 示例 2：
 *
 * 输入：3
 * 输出：false
 * 解释：爱丽丝选择 1，鲍勃也选择 1，然后爱丽丝无法进行操作。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= N <= 1000
 *
 *
 */

// @lc code=start
func divisorGame(N int) bool {
	// // 设置dp[n]为当数字为n的时候是否可以赢
	// dp := [1001]bool{}
	// // 1的时候先手肯定赢不了
	// dp[1] = false
	// // 2的时候先后肯定赢
	// dp[2] = true
	// // 接下来就是看先手能否将对手推到必输的数字x(即dp(x)为false)
	// for i := 3; i <= N; i++ {
	// 	// 遍历i的公因数，看看减去之后是否必输。公因数必然小于开根号。
	// 	dp[i] = false
	// 	for j, endNum := 1, int(math.Sqrt(float64(i))); j < endNum; j++ {
	// 		if i%j == 0 && dp[i-j] == false {
	// 			dp[i] = true
	// 			break
	// 		}
	// 	}
	// }
	// return dp[N]

	dp := []bool{false}
	for i := 1; i <= N; i++ {
		if i == 1 {
			dp = append(dp, false)
		} else if i == 2 {
			dp = append(dp, true)
		} else {
			ans := false
			// 遍历，查看能否找到一个满足条件的值x, 使得N%x== 0 且dp[N-x] = false
			for x := 1; x <= int(math.Sqrt(float64(i))); x++ {
				if i%x == 0 && dp[i-x] == false {
					ans = true
					break
				}
			}
			dp = append(dp, ans)
		}
	}

	return dp[N]
}

// func main() {
// 	N := 4
// 	fmt.Println(divisorGame(N))
// }

// @lc code=end

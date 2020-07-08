/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode-cn.com/problems/triangle/description/
 *
 * algorithms
 * Medium (61.72%)
 * Likes:    261
 * Dislikes: 0
 * Total Accepted:    31.1K
 * Total Submissions: 49.6K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
 *
 * 例如，给定三角形：
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 *
 * 说明：
 *
 * 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
 *
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	// 计算有多少行
	rNum := len(triangle)
	if rNum == 0 {
		return 0
	}

	// 正常思路是需要使用O(n^2), 但是可以简化为O(n), 复用每一行的空间, 外加一个临时变量，替换前临时存一下
	dp := []int{}
	tmp := 0
	for i := 0; i < rNum; i++ {
		// 先初始化
		dp = append(dp, 0)
	}

	// 每一层都要遍历每个节点，代表如果走当前节点的话，最小值是多少
	for i := 0; i < rNum; i++ {
		for j := 0; j < i + 1; j++ {
			// 头和尾需要特殊处理
			if j == 0 {
				tmp = dp[j]
				dp[j] = triangle[i][j] + dp[j]
			} else if j == i {
				// 本来应该dp[j-1], 但是dp[j-1]已经被修改，需要走到tmp
				dp[j] = triangle[i][j] + tmp
			} else {
				newTmp := dp[j]
				dp[j] = triangle[i][j] + getMin(dp[j], tmp)
				tmp = newTmp
			}
		}
	}

	// 遍历最后一层找到最小值
	min := dp[0]
	for i, length := 1, len(dp); i < length; i++ {
		if dp[i] < min {
			min = dp[i]
		}
	}
	return min
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end


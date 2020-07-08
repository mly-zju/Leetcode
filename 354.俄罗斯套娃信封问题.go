/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 *
 * https://leetcode-cn.com/problems/russian-doll-envelopes/description/
 *
 * algorithms
 * Hard (34.48%)
 * Likes:    116
 * Dislikes: 0
 * Total Accepted:    8.3K
 * Total Submissions: 24K
 * Testcase Example:  '[[5,4],[6,4],[6,7],[2,3]]'
 *
 * 给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h)
 * 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
 * 
 * 请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
 * 
 * 说明:
 * 不允许旋转信封。
 * 
 * 示例:
 * 
 * 输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
 * 输出: 3 
 * 解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
 * 
 * 
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	// 先执行排序：按照w升序且h降序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	// 然后针对h，寻找LIS
	return findLIS(envelopes)
}

func findLIS(envelopes [][]int) int {
	length := len(envelopes)
	if length == 0 {
		return 0
	}
	// dp[i]代表以i结尾的最大LIS
	dp := make([]int, length)
	
	totalMax := 1
	for i := 0; i < length; i++ {
		if i == 0 {
			dp[i] = 1
		} else {
			tmpMax := 1
			for j := 0; j < i; j++ {
				if envelopes[j][1] < envelopes[i][1] && tmpMax < dp[j] + 1 {
					tmpMax = dp[j] + 1
				}
			}
			dp[i] = tmpMax
		}
		if dp[i] > totalMax {
			totalMax = dp[i]
		}
	}
	return totalMax
}
// @lc code=end


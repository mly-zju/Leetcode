/*
 * @lc app=leetcode.cn id=813 lang=golang
 *
 * [813] 最大平均值和的分组
 *
 * https://leetcode-cn.com/problems/largest-sum-of-averages/description/
 *
 * algorithms
 * Medium (43.59%)
 * Likes:    52
 * Dislikes: 0
 * Total Accepted:    1.5K
 * Total Submissions: 3.3K
 * Testcase Example:  '[9,1,2,3,9]\n3'
 *
 * 我们将给定的数组 A 分成 K 个相邻的非空子数组 ，我们的分数由每个子数组内的平均值的总和构成。计算我们所能得到的最大分数是多少。
 *
 * 注意我们必须使用 A 数组中的每一个数进行分组，并且分数不一定需要是整数。
 *
 *
 * 示例:
 * 输入:
 * A = [9,1,2,3,9]
 * K = 3
 * 输出: 20
 * 解释:
 * A 的最优分组是[9], [1, 2, 3], [9]. 得到的分数是 9 + (1 + 2 + 3) / 3 + 9 = 20.
 * 我们也可以把 A 分成[9, 1], [2], [3, 9].
 * 这样的分组得到的分数为 5 + 2 + 6 = 13, 但不是最大值.
 *
 *
 * 说明:
 *
 *
 * 1 <= A.length <= 100.
 * 1 <= A[i] <= 10000.
 * 1 <= K <= A.length.
 * 答案误差在 10^-6 内被视为是正确的。
 *
 *
 */

// @lc code=start
func largestSumOfAverages(A []int, K int) float64 {
	// 区间动态规划，dp[i][j]代表0-i数组范围内划分j组的最大值
	length := len(A)
	if length < K {
		return 0
	}

	dp := [][]float64{}
	// 初始化
	for i := 0; i < length; i++ {
		lineArr := []float64{}
		for j := 0; j <= K; j++ {
			lineArr = append(lineArr, 0)
		}
		dp = append(dp, lineArr)
	}

	for i := 0; i < length; i++ {
		for j := 1; j <= K && j <= i+1; j++ {
			if j == 1 {
				// 如果只有1组，直接取平均值
				dp[i][j] = getSum(A, 0, i) / float64(i+1)
			} else if j == i+1 {
				// 如果组数和元素数相等，直接取和
				dp[i][j] = getSum(A, 0, i)
			} else {
				tmpMax := float64(0)
				// 从0开始，前面有j-1组，每组至少1个元素，所以最后一组最早也要从j-1开始, x代表之前一组结束位置的索引
				for x := j - 2; x < i; x++ {
					// 这里只看取一组的情况，则都是dp[x][j-1], 遍历x找最大值, 其中x最小也应该是j-2, 最大i - 1
					tmp := dp[x][j-1] + getSum(A, x+1, i)/float64(i-x)
					if tmp > tmpMax {
						tmpMax = tmp
					}
				}
				dp[i][j] = tmpMax
			}
		}
	}

	return dp[length-1][K]
}

func getSum(A []int, i, j int) float64 {
	res := 0
	for index := i; index <= j; index++ {
		res += A[index]
	}
	return float64(res)
}

// @lc code=end


import "sort"

/*
 * @lc app=leetcode.cn id=368 lang=golang
 *
 * [368] 最大整除子集
 *
 * https://leetcode-cn.com/problems/largest-divisible-subset/description/
 *
 * algorithms
 * Medium (35.46%)
 * Likes:    55
 * Dislikes: 0
 * Total Accepted:    3.8K
 * Total Submissions: 10.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给出一个由无重复的正整数组成的集合，找出其中最大的整除子集，子集中任意一对 (Si，Sj) 都要满足：Si % Sj = 0 或 Sj % Si =
 * 0。
 *
 * 如果有多个目标子集，返回其中任何一个均可。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: [1,2] (当然, [1,3] 也正确)
 *
 *
 * 示例 2:
 *
 * 输入: [1,2,4,8]
 * 输出: [1,2,4,8]
 *
 *
 */

// @lc code=start
func largestDivisibleSubset(nums []int) []int {
	// 先排序
	sort.Ints(nums)

	// 对每个num[i]来说，只要num[i]%num[j]==0, 那么对于dp[j]下面的所有元素，一定也满足该条件



	// // dp+前驱prev来做回溯, dp代表第i个位置时候最大连续值
	// length := len(nums)
	// if length == 0 {
	// 	return []int{}
	// }
	// dp, prev := []int{}, []int{}

	// maxIndex := 0
	// max := 0
	// for i := 0; i < length; i++ {
	// 	if i == 0 {
	// 		dp = append(dp, 1)
	// 		prev = append(prev, -1)
	// 		max = 1
	// 		maxIndex = 0
	// 	} else {
	// 		tmpMax := 1
	// 		tmpPrevIndex := -1
	// 		for j := 0; j < i; j++ {
	// 			if nums[i]%nums[j] == 0 && dp[j]+1 > tmpMax {
	// 				tmpMax = dp[j] + 1
	// 				tmpPrevIndex = j
	// 			}
	// 		}
	// 		dp = append(dp, tmpMax)
	// 		prev = append(prev, tmpPrevIndex)
	// 		if tmpMax > max {
	// 			max = tmpMax
	// 			maxIndex = i
	// 		}
	// 	}
	// }

	// res := []int{}
	// // 开始用前驱数组来看答案
	// for maxIndex != -1 {
	// 	res = append(res, nums[maxIndex])
	// 	maxIndex = prev[maxIndex]
	// }

	// return res
}

// @lc code=end


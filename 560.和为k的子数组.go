/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (44.20%)
 * Likes:    261
 * Dislikes: 0
 * Total Accepted:    20.8K
 * Total Submissions: 46.8K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
 * 
 * 示例 1 :
 * 
 * 
 * 输入:nums = [1,1,1], k = 2
 * 输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
 * 
 * 
 * 说明 :
 * 
 * 
 * 数组的长度为 [1, 20,000]。
 * 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
 * 
 * 
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	// 使用前缀和数组，因为有负数所以无法用滑动窗口
	length := len(nums)
	if length == 0 {
		return 0
	}
	cacheNums := []int{}
	sum := 0
	// 计算前缀和
	for _, num := range nums {
		sum += num
		cacheNums = append(cacheNums, sum)
	}

	res := 0
	// sum[i:j] = cacheNums[j] - cacheNums[i-1]
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if i == 0 {
				if cacheNums[j] == k {
					res++
				}
			} else if cacheNums[j] - cacheNums[i-1] == k {
				res++
			}
		}
	}
	return res
}
// @lc code=end


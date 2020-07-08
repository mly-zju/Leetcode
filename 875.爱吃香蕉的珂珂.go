/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 *
 * https://leetcode-cn.com/problems/koko-eating-bananas/description/
 *
 * algorithms
 * Medium (40.71%)
 * Likes:    50
 * Dislikes: 0
 * Total Accepted:    8.9K
 * Total Submissions: 21.9K
 * Testcase Example:  '[3,6,7,11]\n8'
 *
 * 珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。
 * 
 * 珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K
 * 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。  
 * 
 * 珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
 * 
 * 返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入: piles = [3,6,7,11], H = 8
 * 输出: 4
 * 
 * 
 * 示例 2：
 * 
 * 输入: piles = [30,11,23,4,20], H = 5
 * 输出: 30
 * 
 * 
 * 示例 3：
 * 
 * 输入: piles = [30,11,23,4,20], H = 6
 * 输出: 23
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= piles.length <= 10^4
 * piles.length <= H <= 10^9
 * 1 <= piles[i] <= 10^9
 * 
 * 
 */

// @lc code=start
func minEatingSpeed(piles []int, H int) int {
	// 需要求最小的k，可以二分查找降低复杂度
	length := len(piles)
	if length == 0 {
		return 0
	}
	kmin, kmax := 1, 1
	// 寻找kmax
	for _, val := range piles {
		if val > kmax {
			kmax = val
		}
	}

	// 开始二分查找, 这个其实就是寻找下界的二分法
	var kmid int
	for kmin < kmax {
		kmid = (kmin + kmax) / 2
		eatResult := getEatResult(piles, H, kmid)
		if eatResult {
			kmax = kmid
		} else {
			kmin = kmid + 1
		}
	}
	return kmin
}

// getEatResult 返回是否可以吃完
func getEatResult(piles []int, H, kmid int) bool {
	// 这里不是贪心法，因为不是求规定时间内最多能吃多少，而是吃完要多久，怎么吃都是一样的
	totalTime := 0
	for _, val := range piles {
		// 直接计算需要吃几次
		tmp := int(math.Ceil(float64(val) / float64(kmid)))
		totalTime += tmp
	}
	return totalTime <= H
}
// @lc code=end


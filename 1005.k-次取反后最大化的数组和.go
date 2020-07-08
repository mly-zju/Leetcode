/*
 * @lc app=leetcode.cn id=1005 lang=golang
 *
 * [1005] K 次取反后最大化的数组和
 *
 * https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations/description/
 *
 * algorithms
 * Easy (47.79%)
 * Likes:    20
 * Dislikes: 0
 * Total Accepted:    4.1K
 * Total Submissions: 8.3K
 * Testcase Example:  '[4,2,3]\n1'
 *
 * 给定一个整数数组 A，我们只能用以下方法修改该数组：我们选择某个个索引 i 并将 A[i] 替换为 -A[i]，然后总共重复这个过程 K
 * 次。（我们可以多次选择同一个索引 i。）
 *
 * 以这种方式修改数组后，返回数组可能的最大和。
 *
 *
 *
 * 示例 1：
 *
 * 输入：A = [4,2,3], K = 1
 * 输出：5
 * 解释：选择索引 (1,) ，然后 A 变为 [4,-2,3]。
 *
 *
 * 示例 2：
 *
 * 输入：A = [3,-1,0,2], K = 3
 * 输出：6
 * 解释：选择索引 (1, 2, 2) ，然后 A 变为 [3,1,0,2]。
 *
 *
 * 示例 3：
 *
 * 输入：A = [2,-3,-1,5,-4], K = 2
 * 输出：13
 * 解释：选择索引 (1, 4) ，然后 A 变为 [2,3,-1,5,4]。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= A.length <= 10000
 * 1 <= K <= 10000
 * -100 <= A[i] <= 100
 *
 *
 */

// @lc code=start
func largestSumAfterKNegations(A []int, K int) int {
	// 贪心算法：优先选择最小的负数。如果没有负数，优先选择0.如果没有0，优先选择最小正数。
	// 总之就是选择最小值。
	// 1. 其实可以基于堆来做，这里先用普通数组. 复杂度O(K*len(A))
	// for i := 0; i < K; i++ {
	// 	index := getTarget(A)
	// 	A[index] = -A[index]
	// }
	// sum := 0
	// for _, val := range A {
	// 	sum += val
	// }
	// return sum

	// 2. 基于统计数组来做，-100 <= A[i] <= 100，用201长度的数组记录出现频率。复杂度O(K*201)
	number := [201]int{}
	for _, val := range A {
		number[val+100]++
	}

	cursor := 0
	for i := 0; i < K; i++ {
		// 找最小的非0值
		for number[cursor] == 0 {
			cursor++
		}
		number[cursor]--
		// 对应数是 -(cursor - 100), 再加100偏移量，即200-cursor
		number[200-cursor]++
		// 如果是正数，那么下一轮最小值一定是对应的负数
		if cursor > 100 {
			cursor = 200 - cursor
		}
	}

	sum := 0
	for index, val := range number {
		if val != 0 {
			sum += (index - 100) * val
		}
	}
	return sum

	// 3. 基于堆? todo..
}

func getTarget(A []int) int {
	minNum := A[0]
	minIndex := 0
	for index, val := range A {
		if minNum > val {
			minNum = val
			minIndex = index
		}
	}
	return minIndex
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=659 lang=golang
 *
 * [659] 分割数组为连续子序列
 *
 * https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/description/
 *
 * algorithms
 * Medium (41.66%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    1.4K
 * Total Submissions: 3.6K
 * Testcase Example:  '[1,2,3,3,4,5]'
 *
 * 输入一个按升序排序的整数数组（可能包含重复数字），你需要将它们分割成几个子序列，其中每个子序列至少包含三个连续整数。返回你是否能做出这样的分割？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: [1,2,3,3,4,5]
 * 输出: True
 * 解释:
 * 你可以分割出这样两个连续子序列 :
 * 1, 2, 3
 * 3, 4, 5
 *
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入: [1,2,3,3,4,4,5,5]
 * 输出: True
 * 解释:
 * 你可以分割出这样两个连续子序列 :
 * 1, 2, 3, 4, 5
 * 3, 4, 5
 *
 *
 *
 *
 * 示例 3：
 *
 *
 * 输入: [1,2,3,4,4,5]
 * 输出: False
 *
 *
 *
 *
 * 提示：
 *
 *
 * 输入的数组长度范围为 [1, 10000]
 *
 *
 *
 *
 */

// @lc code=start
func isPossible(nums []int) bool {
	// 贪心算法，能加就加
	statMap := map[int]int{}
	endMap := map[int]int{}
	flag := true

	for _, val := range nums {
		statMap[val]++
	}

	for _, val := range nums {
		if statMap[val] == 0 {
			continue
		}

		if endMap[val-1] > 0 {
			statMap[val]--
			endMap[val-1]--
			endMap[val]++
		} else if statMap[val+1] > 0 && statMap[val+2] > 0 {
			endMap[val+2]++
			statMap[val]--
			statMap[val+1]--
			statMap[val+2]--
		} else {
			flag = false
			break
		}

	}
	return flag
}

// @lc code=end


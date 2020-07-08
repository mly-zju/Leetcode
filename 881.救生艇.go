import "sort"

/*
 * @lc app=leetcode.cn id=881 lang=golang
 *
 * [881] 救生艇
 *
 * https://leetcode-cn.com/problems/boats-to-save-people/description/
 *
 * algorithms
 * Medium (37.86%)
 * Likes:    31
 * Dislikes: 0
 * Total Accepted:    5.1K
 * Total Submissions: 13.2K
 * Testcase Example:  '[1,2]\n3'
 *
 * 第 i 个人的体重为 people[i]，每艘船可以承载的最大重量为 limit。
 *
 * 每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。
 *
 * 返回载到每一个人所需的最小船数。(保证每个人都能被船载)。
 *
 *
 *
 * 示例 1：
 *
 * 输入：people = [1,2], limit = 3
 * 输出：1
 * 解释：1 艘船载 (1, 2)
 *
 *
 * 示例 2：
 *
 * 输入：people = [3,2,2,1], limit = 3
 * 输出：3
 * 解释：3 艘船分别载 (1, 2), (2) 和 (3)
 *
 *
 * 示例 3：
 *
 * 输入：people = [3,5,3,4], limit = 5
 * 输出：4
 * 解释：4 艘船分别载 (3), (3), (4), (5)
 *
 * 提示：
 *
 *
 * 1 <= people.length <= 50000
 * 1 <= people[i] <= limit <= 30000
 *
 *
 */

// @lc code=start
func numRescueBoats(people []int, limit int) int {
	// 贪心：尽可能占用重量，并尽可能填满2人
	// 所以people排序，首尾指针，先装两个大的，第二个如果装不下，再装小的
	sort.Ints(people)
	l, r := 0, len(people)-1

	res := 0
	for l <= r {
		// 先装大的
		tmp := people[r]
		r--
		if l <= r {
			// 再尝试装大的，装不下才装小的
			if limit-tmp >= people[r] {
				r--
			} else if limit-tmp >= people[l] {
				l++
			}
		}
		res++
	}

	return res
}

// @lc code=end


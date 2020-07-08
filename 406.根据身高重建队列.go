import "sort"

/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 *
 * https://leetcode-cn.com/problems/queue-reconstruction-by-height/description/
 *
 * algorithms
 * Medium (62.34%)
 * Likes:    179
 * Dislikes: 0
 * Total Accepted:    11.5K
 * Total Submissions: 18.9K
 * Testcase Example:  '[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]'
 *
 * 假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。
 * 编写一个算法来重建这个队列。
 *
 * 注意：
 * 总人数少于1100人。
 *
 * 示例
 *
 *
 * 输入:
 * [[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
 *
 * 输出:
 * [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
 *
 *
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	// 从大到小排序后插入：理由是小的插在大的前面不会影响大的
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] > people[j][0] {
			return true
		} else if people[i][0] < people[j][0] {
			return false
		}

		// 如果第一位无法比较，比较第二位，升序
		return people[i][1] < people[j][1]
	})

	res := [][]int{}
	// 执行交换
	for _, p := range people {
		res = append(res, p)
		newLen := len(res)
		start := p[1]
		for i := newLen - 1; i > start; i-- {
			res[i] = res[i-1]
		}

		res[start] = p
	}

	return res
}

// @lc code=end


import "sort"

/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 *
 * https://leetcode-cn.com/problems/non-overlapping-intervals/description/
 *
 * algorithms
 * Medium (43.16%)
 * Likes:    74
 * Dislikes: 0
 * Total Accepted:    6K
 * Total Submissions: 13.8K
 * Testcase Example:  '[[1,2]]'
 *
 * 给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
 *
 * 注意:
 *
 *
 * 可以认为区间的终点总是大于它的起点。
 * 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
 *
 *
 * 示例 1:
 *
 *
 * 输入: [ [1,2], [2,3], [3,4], [1,3] ]
 *
 * 输出: 1
 *
 * 解释: 移除 [1,3] 后，剩下的区间没有重叠。
 *
 *
 * 示例 2:
 *
 *
 * 输入: [ [1,2], [1,2], [1,2] ]
 *
 * 输出: 2
 *
 * 解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
 *
 *
 * 示例 3:
 *
 *
 * 输入: [ [1,2], [2,3] ]
 *
 * 输出: 0
 *
 * 解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
 *
 *
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	// 其实就是找重叠的区间数量，和452射气球问题很类似
	length := len(intervals)
	if length == 0 {
		return 0
	}

	// 按照end从小到大排序
	sort.Slice(intervals, func (i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	res := 0
	prevEnd := intervals[0][1]
	for i := 1; i < length; i++ {
		if intervals[i][0] < prevEnd {
			// 如果重叠，就需要移除
			res++
		} else {
			// 不重叠，就需要更新最新的prevEnd
			prevEnd = intervals[i][1]
		}
	}
	return res
}

// @lc code=end


import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (38.43%)
 * Likes:    216
 * Dislikes: 0
 * Total Accepted:    38K
 * Total Submissions: 97.2K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 *
 * 示例 1:
 *
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2:
 *
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	// 先排序，再合并, 排序按照起始边界大小排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 执行合并
	res := [][]int{}
	resLen := 0
	for _, arr := range intervals {
		if resLen == 0 {
			res = append(res, arr)
			resLen++
		} else {
			tmpArr := res[resLen-1]
			// 如果end大于新元素的start，需要合并并更新
			if tmpArr[1] >= arr[0] {
				if arr[1] > tmpArr[1] {
					tmpArr[1] = arr[1]
				}
				res[resLen-1] = tmpArr

            // 否则不用更新区间，直接推入
			} else {
				res = append(res, arr)
				resLen++
			}
		}
	}
	return res
}

// @lc code=end


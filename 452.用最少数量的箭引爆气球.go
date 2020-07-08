import "sort"

/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 *
 * https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/description/
 *
 * algorithms
 * Medium (46.87%)
 * Likes:    74
 * Dislikes: 0
 * Total Accepted:    6.7K
 * Total Submissions: 14.1K
 * Testcase Example:  '[[10,16],[2,8],[1,6],[7,12]]'
 *
 *
 * 在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。由于它是水平的，所以y坐标并不重要，因此只要知道开始和结束的x坐标就足够了。开始坐标总是小于结束坐标。平面内最多存在10^4个气球。
 *
 * 一支弓箭可以沿着x轴从不同点完全垂直地射出。在坐标x处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart
 * ≤ x ≤ xend，则该气球会被引爆。可以射出的弓箭的数量没有限制。
 * 弓箭一旦被射出之后，可以无限地前进。我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。
 *
 * Example:
 *
 *
 * 输入:
 * [[10,16], [2,8], [1,6], [7,12]]
 *
 * 输出:
 * 2
 *
 * 解释:
 * 对于该样例，我们可以在x = 6（射爆[2,8],[1,6]两个气球）和 x = 11（射爆另外两个气球）。
 *
 *
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	// 这个和435区间调度很类似，435要找的是需要移除的数量，而本题目则是找到移除之后，无重叠的数量，二者互补
	length := len(points)
	if length == 0 {
		return 0
	}

	// 以结束坐标进行排序，顺序扫描，有重叠的就可以一箭射爆
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	prevEnd := points[0][1]
	res := 1
	for i := 1; i < length; i++ {
		if points[i][0] > prevEnd {
			// 如果不重叠，则箭数需要+1，且更新最新结束值
			res++
			prevEnd = points[i][1]
		}
	}
	return res
}

// @lc code=end


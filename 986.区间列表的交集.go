/*
 * @lc app=leetcode.cn id=986 lang=golang
 *
 * [986] 区间列表的交集
 *
 * https://leetcode-cn.com/problems/interval-list-intersections/description/
 *
 * algorithms
 * Medium (61.20%)
 * Likes:    37
 * Dislikes: 0
 * Total Accepted:    4.7K
 * Total Submissions: 7.7K
 * Testcase Example:  '[[0,2],[5,10],[13,23],[24,25]]\n[[1,5],[8,12],[15,24],[25,26]]'
 *
 * 给定两个由一些闭区间组成的列表，每个区间列表都是成对不相交的，并且已经排序。
 * 
 * 返回这两个区间列表的交集。
 * 
 * （形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <=
 * b。两个闭区间的交集是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3]。）
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 
 * 输入：A = [[0,2],[5,10],[13,23],[24,25]], B = [[1,5],[8,12],[15,24],[25,26]]
 * 输出：[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
 * 注意：输入和所需的输出都是区间对象组成的列表，而不是数组或列表。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= A.length < 1000
 * 0 <= B.length < 1000
 * 0 <= A[i].start, A[i].end, B[i].start, B[i].end < 10^9
 * 
 * 
 */

// @lc code=start
func intervalIntersection(A [][]int, B [][]int) [][]int {
	// 区间求交集，拉链法来搞一下
	lena, lenb := len(A), len(B)
	al, bl := 0, 0

	res := [][]int{}
	for al < lena && bl < lenb {
		// 情况1：a在前面,
		if A[al][0] <= B[bl][0] {
			// 要想有交集，b的start要小于a的end
			if A[al][1] >= B[bl][0] {
				// 此时又分为2中情况：一种是b完全被a遮盖，另一种非完全遮盖
				if B[bl][1] <= A[al][1] {
					res = append(res, []int{B[bl][0], B[bl][1]})
					// 完全遮盖，则b移动
					bl++
				} else {
					res = append(res, []int{B[bl][0], A[al][1]})
					// 非完全遮盖，则a移动
					al++
				}
            // 如果没有交集，a移动
			} else {
				al++
			}
        // 情况2：b在前面
		} else {
			// 有交集的条件
			if B[bl][1] >= A[al][0] {
				// 区分是否完全包裹
				if A[al][1] <= B[bl][1] {
					res = append(res, []int{A[al][0], A[al][1]})
					al++
				} else {
					res = append(res, []int{A[al][0], B[bl][1]})
					bl++
				}
			} else {
				bl++
			}
		}
	}

	return res
}
// @lc code=end


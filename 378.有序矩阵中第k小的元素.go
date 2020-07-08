package main

/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第K小的元素
 *
 * https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/description/
 *
 * algorithms
 * Medium (56.67%)
 * Likes:    111
 * Dislikes: 0
 * Total Accepted:    10.2K
 * Total Submissions: 18K
 * Testcase Example:  '[[1,5,9],[10,11,13],[12,13,15]]\n8'
 *
 * 给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第k小的元素。
 * 请注意，它是排序后的第k小元素，而不是第k个元素。
 *
 * 示例:
 *
 *
 * matrix = [
 * ⁠  [ 1,  5,  9],
 * ⁠  [10, 11, 13],
 * ⁠  [12, 13, 15]
 * ],
 * k = 8,
 *
 * 返回 13。
 *
 *
 * 说明:
 * 你可以假设 k 的值永远是有效的, 1 ≤ k ≤ n^2 。
 *
 */

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
	// 最简单的是基于小顶堆的优先队列。
	// 不过这里使用二分法玩一下

	// 1. 基于value的二分法
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	return binarySearch(matrix, l, r, k)
}

func binarySearch(matrix [][]int, l, r, k int) int {
	length := len(matrix)
	var mid int
	for l < r {
		mid = (l + r) / 2
		// count统计矩阵有多少个元素小于等于mid
		count := 0
		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				if matrix[i][j] <= mid {
					count++
					if count >= k {
						break
					}
				}
			}
		}

		if count >= k {
			// 猜的有可能太大，或者正好，因此r不设置mid-1而是mid
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// func main() {
// 	// m := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
// 	m := [][]int{{1, 2}, {1, 3}}
// 	fmt.Println(kthSmallest(m, 3))
// }

// @lc code=end

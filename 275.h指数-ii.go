/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H指数 II
 *
 * https://leetcode-cn.com/problems/h-index-ii/description/
 *
 * algorithms
 * Medium (37.88%)
 * Likes:    22
 * Dislikes: 0
 * Total Accepted:    4.4K
 * Total Submissions: 11.2K
 * Testcase Example:  '[0,1,3,5,6]'
 *
 * 给定一位研究者论文被引用次数的数组（被引用次数是非负整数），数组已经按照升序排列。编写一个方法，计算出研究者的 h 指数。
 *
 * h 指数的定义: “h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的 （N 篇论文中）至多有 h
 * 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）"
 *
 *
 *
 * 示例:
 *
 * 输入: citations = [0,1,3,5,6]
 * 输出: 3
 * 解释: 给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 0, 1, 3, 5, 6 次。
 * 由于研究者有 3 篇论文每篇至少被引用了 3 次，其余两篇论文每篇被引用不多于 3 次，所以她的 h 指数是 3。
 *
 *
 *
 * 说明:
 *
 * 如果 h 有多有种可能的值 ，h 指数是其中最大的那个。
 *
 *
 *
 * 进阶：
 *
 *
 * 这是 H指数 的延伸题目，本题中的 citations 数组是保证有序的。
 * 你可以优化你的算法到对数时间复杂度吗？
 *
 *
 */

// @lc code=start
func hIndex(citations []int) int {
	// 已经排序了，首先O(n)的方法很简单

	// 1. O(n)
	// max := 0
	// length := len(citations)
	// for i := length - 1; i >= 0; i-- {
	// 	tmp := getMin(citations[i], length-i)
	// 	if tmp > max {
	// 		max = tmp
	// 	}
	// }
	// return max

	// 2. O(logn)
	// 原理：理想情况是找到val==length - index的，如果val > length - index, 那么就向左走；如果val < length - index, 那么就向右走
	length := len(citations)
	l, r := 0, length-1
	mid := 0
	res := 0
	for l <= r {
		mid = (l + r) / 2
		if length-mid == citations[mid] {
			// 如果找到了相等的那么一定是最大的
			res = citations[mid]
			break
		} else if length-mid > citations[mid] {
			tmp := getMin(length-mid, citations[mid])
			if tmp > res {
				res = tmp
			}
			l = mid + 1
		} else {
			tmp := getMin(length-mid, citations[mid])
			if tmp > res {
				res = tmp
			}
			r = mid - 1
		}
	}
	return res
}

func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end


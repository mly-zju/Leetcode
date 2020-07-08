/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (57.74%)
 * Likes:    869
 * Dislikes: 0
 * Total Accepted:    95.1K
 * Total Submissions: 161.4K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 *
 *
 *
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 *
 *
 * 示例:
 *
 * 输入: [1,8,6,2,5,4,8,3,7]
 * 输出: 49
 *
 */

// @lc code=start
func maxArea(height []int) int {
	// 1. 暴力法
	// length := len(height)
	// max := 0
	// for i := 0; i < length-1; i++ {
	// 	for j := i + 1; j < length; j++ {
	// 		tmp := getMin(height[i], height[j]) * getAbs(j-i)
	// 		if tmp > max {
	// 			max = tmp
	// 		}
	// 	}
	// }
	// return max

	// 2. 首尾双指针，类似贪心算法，尝试朝着更大的方向移动
	l, r := 0, len(height)-1
	max := 0
	for l < r {
		tmp := getMin(height[l], height[r]) * getAbs(r-l)
		if tmp > max {
			max = tmp
		}
		// 总是移动较短的线段，因为移动较长的没有任何收益
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end


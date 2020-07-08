import "sort"

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (41.68%)
 * Likes:    281
 * Dislikes: 0
 * Total Accepted:    51.2K
 * Total Submissions: 122.2K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 * 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
 *
 * 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 *
 *
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	// 类似15题，排序+左边移位+右边首尾指针，不同的是需要记录全局最小值
	sort.Ints(nums)
	res := -1

	length := len(nums)
	for i := 0; i < length - 2; i++ {
		// 如果当前和前一个相等，则不用重复计算了
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		// 开始首尾指针寻找
		l, r := i + 1, length - 1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			// 每轮都更新一下res
			tmpGap := getAbs(target - tmp)
			if res == -1 {
				res = tmp
			} else if getAbs(target - res) > tmpGap {
				res = tmp
			}

			if tmp > target {
				r--
			} else if tmp < target {
				l++
			} else {
				// 如果遇到相等，直接返回
				return tmp
			}

		}
	}
	return res
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end


import "sort"

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (35.89%)
 * Likes:    302
 * Dislikes: 0
 * Total Accepted:    37.5K
 * Total Submissions: 104K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
 * + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 *
 * 注意：
 *
 * 答案中不可以包含重复的四元组。
 *
 * 示例：
 *
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 *
 * 满足要求的四元组集合为：
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	// 先排序
	sort.Ints(nums)
	length := len(nums)
	res := [][]int{}
	for i := 0; i < length-3; i++ {
		// 前一个元素需要去重
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			// 第二个元素也去重
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 首尾指针开始计算
			l, r := j+1, length-1
			sum := 0
			for l < r {
				sum = nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					// l/r跳过重复的
					lNum, rNum := nums[l], nums[r]
					for l < length && nums[l] == lNum {
						l++
					}
					for r >= 0 && nums[r] == rNum {
						r--
					}
				}
			}
		}
	}

	return res
}

// @lc code=end


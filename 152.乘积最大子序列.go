/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (35.45%)
 * Likes:    316
 * Dislikes: 0
 * Total Accepted:    27.1K
 * Total Submissions: 75K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。
 *
 * 示例 1:
 *
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 */

// @lc code=start
func maxProduct(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	// 需要两个dp，一个存放最小，一个存放最大
	dpMax, dpMin := []int{nums[0]}, []int{nums[0]}
	max := nums[0]

	for i := 1; i < length; i++ {
		if nums[i] == 0 {
			dpMax = append(dpMax, 0)
			dpMin = append(dpMin, 0)
		} else if nums[i] > 0 {
			// 先推最大值
			if dpMax[i-1] > 0 {
				dpMax = append(dpMax, nums[i]*dpMax[i-1])
			} else {
				dpMax = append(dpMax, nums[i])
			}

			// 再推最小值
			if dpMin[i-1] <= 0 {
				dpMin = append(dpMin, nums[i]*dpMin[i-1])
			} else {
				dpMin = append(dpMin, nums[i])
			}
		} else if nums[i] < 0 {
			// 先推最大值
			if dpMin[i-1] <= 0 {
				dpMax = append(dpMax, nums[i]*dpMin[i-1])
			} else {
				dpMax = append(dpMax, nums[i])
			}

			// 再推最小值
			if dpMax[i-1] > 0 {
				dpMin = append(dpMin, nums[i]*dpMax[i-1])
			} else {
				dpMin = append(dpMin, nums[i])
			}
		}
	}

	for _, val := range dpMax {
		if val > max {
			max = val
		}
	}
	return max

	// for i := 1; i < length; i++ {
	// 	if nums[i] == 0 {
	// 		dpMax = append(dpMax, 0)
	// 		dpMin = append(dpMin, 0)
	// 		if max < 0 {
	// 			max = 0
	// 		}
	// 	} else if nums[i] > 0 {
	// 		// 大于0的，dpMax插入cur * dpMax[i - 1](前提是i-1大于0), dpMin也要插入cur * dpMin[i - 1](前提是i-1小于0)
	// 		tmpMax := nums[i]
	// 		if dpMax[i-1] > 0 {
	// 			tmpMax = nums[i] * dpMax[i-1]
	// 		}
	// 		dpMax = append(dpMax, tmpMax)

	// 		tmpMin := nums[i]
	// 		if dpMin[i-1] < 0 {
	// 			tmpMin = nums[i] * dpMin[i-1]
	// 		}
	// 		dpMin = append(dpMin, tmpMin)

	// 		if tmpMax > max {
	// 			max = tmpMax
	// 		}
	// 	} else {
	// 		// 小于0的，dpMax插入cur * dpMin[i - 1](前提是i-1小于0), dpMin插入cur * dpMax[i - 1](前提i-1大于0)
	// 		tmpMax := nums[i]
	// 		if dpMin[i-1] < 0 {
	// 			tmpMax = dpMin[i-1] * nums[i]
	// 		}
	// 		dpMax = append(dpMax, tmpMax)

	// 		tmpMin := nums[i]
	// 		if dpMax[i-1] > 0 {
	// 			tmpMin = dpMax[i-1] * nums[i]
	// 		}
	// 		dpMin = append(dpMin, tmpMin)

	// 		if tmpMax > max {
	// 			max = tmpMax
	// 		}
	// 	}
	// }

	return max
}

// @lc code=end


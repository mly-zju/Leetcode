/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 *
 * https://leetcode-cn.com/problems/wiggle-subsequence/description/
 *
 * algorithms
 * Medium (39.30%)
 * Likes:    106
 * Dislikes: 0
 * Total Accepted:    8K
 * Total Submissions: 19.9K
 * Testcase Example:  '[1,7,4,9,2,5]'
 *
 * 如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为摆动序列。第一个差（如果存在的话）可能是正数或负数。少于两个元素的序列也是摆动序列。
 *
 * 例如， [1,7,4,9,2,5] 是一个摆动序列，因为差值 (6,-3,5,-7,3) 是正负交替出现的。相反, [1,4,7,2,5] 和
 * [1,7,4,5,5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
 *
 * 给定一个整数序列，返回作为摆动序列的最长子序列的长度。 通过从原始序列中删除一些（也可以不删除）元素来获得子序列，剩下的元素保持其原始顺序。
 *
 * 示例 1:
 *
 * 输入: [1,7,4,9,2,5]
 * 输出: 6
 * 解释: 整个序列均为摆动序列。
 *
 *
 * 示例 2:
 *
 * 输入: [1,17,5,10,13,15,10,5,16,8]
 * 输出: 7
 * 解释: 这个序列包含几个长度为 7 摆动序列，其中一个可为[1,17,10,13,10,16,8]。
 *
 * 示例 3:
 *
 * 输入: [1,2,3,4,5,6,7,8,9]
 * 输出: 2
 *
 * 进阶:
 * 你能否用 O(n) 时间复杂度完成此题?
 *
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	// 贪心算法，寻找peak
	length := len(nums)
	if length < 2 {
		return length
	}

	// 贪心算法思路: 让上升的继续上升直到达到最大值，这样后面能组成下降序列的概率更大；下降序列也一样
	res := 1
	prev := nums[0]
	// stat 0待定，1上升，-1下降
	stat := 0
	for i := 1; i < length; i++ {
		if stat == 0 {
			if nums[i] > prev {
				prev = nums[i]
				stat = 1
				res++
			} else if nums[i] < prev {
				prev = nums[i]
				stat = -1
				res++
			}
		} else if stat == 1 {
			if nums[i] < prev {
				prev = nums[i]
				stat = -1
				res++
			} else if nums[i] > prev {
				// 如果本来是上升，但是又遇到更大的，直接替换，但是res不增加
				prev = nums[i]
			}
		} else if stat == -1 {
			if nums[i] > prev {
				prev = nums[i]
				stat = 1
				res++
			} else if nums[i] < prev {
				prev = nums[i]
			}
		}
	}

	return res

	// isUp 0待定，1目前上行，-1目前下行
	// isUp := 0
	// res := 1
	// for i := 1; i < length; i++ {
	// 	if isUp == 0 && nums[i] != nums[i-1] {
	// 		res++
	// 		if nums[i] > nums[i-1] {
	// 			// 当前开始上行
	// 			isUp = 1
	// 		} else {
	// 			// 当前开始下行
	// 			isUp = -1
	// 		}
	// 	} else if isUp == 1 && nums[i] < nums[i-1] {
	// 		// 如果上行过程中第一次遇到下行，加到摆动序列，并设置isUp为下行
	// 		isUp = -1
	// 		res++
	// 	} else if isUp == -1 && nums[i] > nums[i-1] {
	// 		isUp = 1
	// 		res++
	// 	}
	// }
	// return res

	// 2. 带状态的dp
	// dpUp[i]代表以i为最后一个上升元素时候的最大长度，dpDown[i]代表以i为最后一个下降元素时候的最大长度
	// dpUp, dpDown := []int{}, []int{}
	// for i := 0; i < length; i++ {
	// 	tmpMaxUp := 1
	// 	tmpMaxDown := 1
	// 	for j := 0; j < i; j++ {
	// 		// 先求dpUp, 需要从dpDown序列中找
	// 		if nums[i] > nums[j] && dpDown[j]+1 > tmpMaxUp {
	// 			tmpMaxUp = dpDown[j] + 1
	// 		}

	// 		if nums[i] < nums[j] && dpUp[j]+1 > tmpMaxDown {
	// 			tmpMaxDown = dpUp[j] + 1
	// 		}
	// 	}

	// 	dpUp = append(dpUp, tmpMaxUp)
	// 	dpDown = append(dpDown, tmpMaxDown)
	// }

	// max := 0
	// for _, val := range dpUp {
	// 	if val > max {
	// 		max = val
	// 	}
	// }
	// for _, val := range dpDown {
	// 	if val > max {
	// 		max = val
	// 	}
	// }
	// return max
}

// @lc code=end


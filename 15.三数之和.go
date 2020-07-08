import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (24.02%)
 * Likes:    1472
 * Dislikes: 0
 * Total Accepted:    111.2K
 * Total Submissions: 457.4K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 * 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	// 1. 统计hash+排序+二重循环，可以去重
	// sort.Ints(nums)
	// length := len(nums)
	// statMap := make(map[int]int)
	// for _, val := range nums {
	// 	statMap[val]++
	// }

	// // 二重循环, hash去重
	// filterMap := make(map[string]bool)
	// res := [][]int{}
	// for i := 0; i < length-1; i++ {
	// 	for j := i + 1; j < length; j++ {
	// 		target := 0 - nums[i] - nums[j]
	// 		key := getKey(nums[i], nums[j], target)
	// 		// 先将前两个元素减去，防止后续重复出现
	// 		statMap[nums[i]]--
	// 		statMap[nums[j]]--
	// 		// target必须大于nums[j]，否则可能存在重复的风险, 然后再看hash
	// 		if target >= nums[j] && !filterMap[key] && statMap[target] > 0 {
	// 			filterMap[key] = true
	// 			res = append(res, []int{nums[i], nums[j], target})
	// 		}
	// 		statMap[nums[i]]++
	// 		statMap[nums[j]]++
	// 	}
	// }
	// return res

	// 2. 排序+移动最左指针(可以和前一个比较去重) + 右边双指针的方法
	// 先去重
	sort.Ints(nums)
	res := [][]int{}
	for i, length := 0, len(nums); i < length-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			// 如果非0，且等于前一个元素，那么结果一定是重复的，跳过
			continue
		}
		l, r := i+1, length-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				// 如果大于0，r变小一点
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				lNum, rNum := nums[l], nums[r]
				// 跳过重复的元素
				for l < length && nums[l] == lNum {
					l++
				}
				for r >= 0 && nums[r] == rNum {
					r--
				}
			}
		}
	}
	return res
}

func getKey(a, b, c int) string {
	return strconv.Itoa(a) + "-" + strconv.Itoa(b) + "-" + strconv.Itoa(c)
}

// @lc code=end


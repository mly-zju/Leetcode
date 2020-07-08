/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (54.27%)
 * Likes:    213
 * Dislikes: 0
 * Total Accepted:    35.9K
 * Total Submissions: 63.4K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列，返回所有不重复的全排列。
 *
 * 示例:
 *
 * 输入: [1,1,2]
 * 输出:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 */

// @lc code=start
var res [][]int

func permuteUnique(nums []int) [][]int {
	res = [][]int{}
	// 标记某位数字是否已经使用
	numFlagArr := make([]int, len(nums))
	ans := []int{}
	round := 0
	_search(nums, ans, numFlagArr, round)
	return res
}

func _search(nums, ans, numFlagArr []int, round int) {
	if len(nums) == round {
		// 在结束的地方去重
		// flag := true
		// for _, arr := range res {
		// 	if isSame(arr, ans) {
		// 		flag = false
		// 		break
		// 	}
		// }

		// if flag {
		newEle := make([]int, len(ans))
		copy(newEle, ans)
		res = append(res, newEle)
		// }
	} else {
		// 遍历前，进行重复性检测，发现当前和全局答案中前面部分有一样的，直接不再向下执行
		for _, arr := range res {
			if isSame(ans, arr) {
				return
			}
		}

		for i, length := 0, len(nums); i < length; i++ {
			if numFlagArr[i] != 1 {
				numFlagArr[i] = 1
				oldLen := len(ans)
				ans = append(ans, nums[i])
				_search(nums, ans, numFlagArr, round+1)
				// 回溯结束后复位
				numFlagArr[i] = 0
				ans = ans[0:oldLen]
			}
		}
	}
}

func isSame(a, b []int) bool {
	for i, length := 0, len(a); i < length; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// @lc code=end


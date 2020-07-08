import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (56.93%)
 * Likes:    197
 * Dislikes: 0
 * Total Accepted:    37.1K
 * Total Submissions: 62.2K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用一次。
 *
 * 说明：
 *
 *
 * 所有数字（包括目标数）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 所求解集为:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 所求解集为:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 */

// @lc code=start
var res [][]int

func combinationSum2(candidates []int, target int) [][]int {
	// 清空全局
	res = [][]int{}
	buf := []int{}
	// seenMap记录答案，用于去重
	seenMap := map[string]bool{}
	// 排序sort，保证按照索引走，得到组合基本都是排序的
	sort.Ints(candidates)
	dfs(candidates, target, 0, buf, seenMap)
	return res
}

func dfs(candidates []int, target, index int, buf []int, seenMap map[string]bool) {
	if target == 0 {
		// 如果已经成功减少到0，执行推入答案操作，不过在这里记得去重
		key := getStr(buf)
		if !seenMap[key] {
			seenMap[key] = true
			newBuf := make([]int, len(buf))
			copy(newBuf, buf)
			res = append(res, newBuf)
		}
	} else {
		// 如果已经超出限制，不再尝试
		if index >= len(candidates) {
			return
		}

		// 1. 尝试不推入, 直接进入下一环节
		dfs(candidates, target, index + 1, buf, seenMap)

		// 2. 尝试推入并进入下一环节
		if candidates[index] <= target {
			buf = append(buf, candidates[index])
			dfs(candidates, target - candidates[index], index + 1, buf, seenMap)
			// 返回的时候推出buf
			buf = buf[0:len(buf) - 1]
		}
	}
}

func getStr(buf []int) string {
	res := ""
	for _, num := range buf {
		res += "-" + strconv.Itoa(num)
	}
	return res
}


// @lc code=end


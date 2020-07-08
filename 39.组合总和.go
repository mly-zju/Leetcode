package main

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (66.92%)
 * Likes:    510
 * Dislikes: 0
 * Total Accepted:    59.5K
 * Total Submissions: 87.4K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的数字可以无限制重复被选取。
 *
 * 说明：
 *
 *
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [2,3,6,7], target = 7,
 * 所求解集为:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,3,5], target = 8,
 * 所求解集为:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 */

// @lc code=start
var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}

	ans := []int{}
	// 为了防止重复插入，需要有一个index变量，每次都从index开始，可选index自身和大于index的
	index := 0
	_search(candidates, target, index, ans)
	return res
}

func _search(candidates []int, target, index int, ans []int) {
	if target == 0 {
		newEle := make([]int, len(ans))
		copy(newEle, ans)
		res = append(res, newEle)
	} else {
		for i, val := range candidates {
			if i >= index && isOk(target, val) {
				ans := append(ans, val)
				_search(candidates, target-val, i, ans)
				// 回溯结束pop一下
				ans = ans[0 : len(ans)-1]
			}
		}
	}
}

func isOk(target, val int) bool {
	return target >= val
}

// func main() {
// 	candidates := []int{7, 3, 2}
// 	target := 18
// 	fmt.Println(combinationSum(candidates, target))
// }

// @lc code=end

package main

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (71.82%)
 * Likes:    510
 * Dislikes: 0
 * Total Accepted:    73K
 * Total Submissions: 99K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个没有重复数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */

// @lc code=start
var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	// 标记某位数字是否已经使用
	numFlagArr := make([]int, len(nums))
	ans := []int{}
	round := 0
	_search(nums, ans, numFlagArr, round)
	return res
}

func _search(nums, ans, numFlagArr []int, round int) {
	if round == len(numFlagArr) {
		// 如果已经到达轮次限制，则停止
		nArr := make([]int, len(ans))
		copy(nArr, ans)
		res = append(res, nArr)
	} else {
		// 这个每一轮都要选择，不存在不选的情况，只不过存在选哪个的情况，因此每轮都选择自己所有可选的
		for i, length := 0, len(numFlagArr); i < length; i++ {
			if numFlagArr[i] != 1 {
				numFlagArr[i] = 1
				// 执行推入, 进入下一轮
				ans = append(ans, nums[i])
				_search(nums, ans, numFlagArr, round+1)
				// 回退
				numFlagArr[i] = 0
				ans = ans[0 : len(ans)-1]
			}
		}
	}
}

// func main() {
// 	a := []int{5, 4, 6, 2}
// 	b := permute(a)
// 	fmt.Println(b)
// }

// @lc code=end

package main

/*
 * @lc app=leetcode.cn id=456 lang=golang
 *
 * [456] 132模式
 *
 * https://leetcode-cn.com/problems/132-pattern/description/
 *
 * algorithms
 * Medium (24.14%)
 * Likes:    89
 * Dislikes: 0
 * Total Accepted:    4.2K
 * Total Submissions: 16.9K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个整数序列：a1, a2, ..., an，一个132模式的子序列 ai, aj, ak 被定义为：当 i < j < k 时，ai < ak <
 * aj。设计一个算法，当给定有 n 个数字的序列时，验证这个序列中是否含有132模式的子序列。
 *
 * 注意：n 的值小于15000。
 *
 * 示例1:
 *
 *
 * 输入: [1, 2, 3, 4]
 *
 * 输出: False
 *
 * 解释: 序列中不存在132模式的子序列。
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3, 1, 4, 2]
 *
 * 输出: True
 *
 * 解释: 序列中有 1 个132模式的子序列： [1, 4, 2].
 *
 *
 * 示例 3:
 *
 *
 * 输入: [-1, 3, 2, 0]
 *
 * 输出: True
 *
 * 解释: 序列中有 3 个132模式的的子序列: [-1, 3, 2], [-1, 3, 0] 和 [-1, 2, 0].
 *
 *
 */

// @lc code=start
type Stack struct {
	list []int
}

func (this *Stack) push(n int) {
	this.list = append(this.list, n)
}

func (this *Stack) isEmpty() bool {
	return len(this.list) == 0
}

func (this *Stack) pop() int {
	length := len(this.list)
	res := this.list[length-1]
	this.list = this.list[0 : length-1]
	return res
}

func (this *Stack) top() int {
	return this.list[len(this.list)-1]
}

func (this *Stack) rtop() int {
	return this.list[len(this.list)-2]
}

func (this *Stack) len() int {
	return len(this.list)
}

func find132pattern(nums []int) bool {
	// 下一个更大值的类似问题，这次找的是下一个次大值
	// 本质上依然是暴力求解，只不过每一轮都排除了不必要的值，减少了复杂度
	stack := Stack{}
	length := len(nums)

	flag := false
	secondMax := 0

	for i := length - 1; i >= 0; i-- {
		if flag && nums[i] < stack.top() && nums[i] < secondMax {
			return true
		}
		for !stack.isEmpty() && nums[i] > stack.top() {
			flag = true
			// secondMax最终记录的是最后一个出栈的元素，因为其最大
			secondMax = stack.pop()
		}
		stack.push(nums[i])
	}
	return false
}

// func main() {
// 	// a := []int{3, 5, 0, 3, 4}
// 	// a := []int{1, 2, 3, 4}
// 	// a := []int{-2, 1, 1}
// 	fmt.Println(find132pattern(a))
// }

// @lc code=end

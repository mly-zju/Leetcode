/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 *
 * https://leetcode-cn.com/problems/next-greater-element-ii/description/
 *
 * algorithms
 * Medium (49.05%)
 * Likes:    63
 * Dislikes: 0
 * Total Accepted:    5.2K
 * Total Submissions: 10.4K
 * Testcase Example:  '[1,2,1]'
 *
 * 给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x
 * 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。
 *
 * 示例 1:
 *
 *
 * 输入: [1,2,1]
 * 输出: [2,-1,2]
 * 解释: 第一个 1 的下一个更大的数是 2；
 * 数字 2 找不到下一个更大的数；
 * 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
 *
 *
 * 注意: 输入数组的长度不会超过 10000。
 *
 */

// @lc code=start
type Mystack []int

func (this *Mystack) push(a int) {
	(*this) = append((*this), a)
}

func (this *Mystack) isEmpty() bool {
	return len((*this)) == 0
}

func (this *Mystack) pop() int {
	length := len(*this)
	tmp := (*this)[length-1]
	(*this) = (*this)[0 : length-1]
	return tmp
}

func (this *Mystack) top() int {
	length := len(*this)
	tmp := (*this)[length-1]
	return tmp
}

func nextGreaterElements(nums []int) []int {
	stack := Mystack{}
	ansArr := []int{}
	// 先全量推入一轮, 因为设定是循环数组
	length := len(nums)
	for i := length - 1; i >= 0; i-- {
		stack.push(nums[i])
	}

	// 开始执行单调栈
	for i := length - 1; i >= 0; i-- {
		for !stack.isEmpty() && nums[i] >= stack.top() {
			stack.pop()
		}

		answer := -1
		if !stack.isEmpty() {
			answer = stack.top()
		}
		ansArr = append(ansArr, answer)
		stack.push(nums[i])
	}

	// 数组反转
	l, r := 0, len(ansArr)-1
	for l < r {
		ansArr[l], ansArr[r] = ansArr[r], ansArr[l]
		l++
		r--
	}
	return ansArr
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 *
 * https://leetcode-cn.com/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (55.38%)
 * Likes:    175
 * Dislikes: 0
 * Total Accepted:    18.4K
 * Total Submissions: 33.1K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 * 根据每日 气温 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。
 *
 * 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4,
 * 2, 1, 1, 0, 0]。
 *
 * 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
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

func dailyTemperatures(T []int) []int {
	var indexStack, valStack Mystack
	length := len(T)
	res := []int{}
	if length == 0 {
		return res
	}

	for i := length - 1; i >= 0; i-- {
		for !valStack.isEmpty() && T[i] >= valStack.top() {
			valStack.pop()
			indexStack.pop()
		}

		// 如果为空，说明之后都不会升高了
		if indexStack.isEmpty() {
			res = append(res, 0)
		} else {
			res = append(res, indexStack.top() - i)
		}

		valStack.push(T[i])
		indexStack.push(i)
	}

	// 执行反转res
	l, r := 0, length - 1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}

// @lc code=end


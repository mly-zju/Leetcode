package main

/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉K位数字
 *
 * https://leetcode-cn.com/problems/remove-k-digits/description/
 *
 * algorithms
 * Medium (26.37%)
 * Likes:    121
 * Dislikes: 0
 * Total Accepted:    6.7K
 * Total Submissions: 25.4K
 * Testcase Example:  '"1432219"\n3'
 *
 * 给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
 *
 * 注意:
 *
 *
 * num 的长度小于 10002 且 ≥ k。
 * num 不会包含任何前导零。
 *
 *
 * 示例 1 :
 *
 *
 * 输入: num = "1432219", k = 3
 * 输出: "1219"
 * 解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
 *
 *
 * 示例 2 :
 *
 *
 * 输入: num = "10200", k = 1
 * 输出: "200"
 * 解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
 *
 *
 * 示例 3 :
 *
 *
 * 输入: num = "10", k = 2
 * 输出: "0"
 * 解释: 从原数字移除所有的数字，剩余为空就是0。
 *
 *
 */

// @lc code=start
type Stack struct {
	list []rune
}

func (this *Stack) push(n rune) {
	this.list = append(this.list, n)
}

func (this *Stack) isEmpty() bool {
	return len(this.list) == 0
}

func (this *Stack) pop() rune {
	length := len(this.list)
	res := this.list[length-1]
	this.list = this.list[0 : length-1]
	return res
}

func (this *Stack) top() rune {
	return this.list[len(this.list)-1]
}

func removeKdigits(num string, k int) string {
	stack := Stack{}
	length := len(num)
	if k >= length {
		return "0"
	}
	// 维护单调递增栈
	for _, ch := range num {
		for !stack.isEmpty() && ch < stack.top() && k > 0 {
			stack.pop()
			k--
		}
		stack.push(ch)
	}

	// 如果k不为0，需要继续截取，由于单调递增，去掉最后一部分即可
	for k > 0 {
		stack.pop()
		k--
	}
	// 将栈底的0去掉
	resArr := stack.list
	for len(resArr) > 0 && resArr[0] == '0' {
		resArr = resArr[1:]
	}

	if len(resArr) == 0 {
		return "0"
	}
	return string(resArr)
}

// func main() {
// 	s := removeKdigits("1111111", 3)
// 	fmt.Println(s)
// }

// @lc code=end

/*
 * @lc app=leetcode.cn id=856 lang=golang
 *
 * [856] 括号的分数
 *
 * https://leetcode-cn.com/problems/score-of-parentheses/description/
 *
 * algorithms
 * Medium (54.41%)
 * Likes:    64
 * Dislikes: 0
 * Total Accepted:    3K
 * Total Submissions: 5.5K
 * Testcase Example:  '"()"'
 *
 * 给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：
 *
 *
 * () 得 1 分。
 * AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
 * (A) 得 2 * A 分，其中 A 是平衡括号字符串。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入： "()"
 * 输出： 1
 *
 *
 * 示例 2：
 *
 * 输入： "(())"
 * 输出： 2
 *
 *
 * 示例 3：
 *
 * 输入： "()()"
 * 输出： 2
 *
 *
 * 示例 4：
 *
 * 输入： "(()(()))"
 * 输出： 6
 *
 *
 *
 *
 * 提示：
 *
 *
 * S 是平衡括号字符串，且只含有 ( 和 ) 。
 * 2 <= S.length <= 50
 *
 *
 */

// @lc code=start
type Node struct {
	base   int
}

type Mystack []Node

func (this *Mystack) push(a Node) {
	(*this) = append((*this), a)
}

func (this *Mystack) isEmpty() bool {
	return len((*this)) == 0
}

func (this *Mystack) pop() Node {
	length := len(*this)
	tmp := (*this)[length-1]
	(*this) = (*this)[0 : length-1]
	return tmp
}

func (this *Mystack) top() Node {
	length := len(*this)
	tmp := (*this)[length-1]
	return tmp
}

func scoreOfParentheses(S string) int {
	// 1. 括号匹配问题，遇到(入栈，遇到)出栈
	var nodeStack Mystack
	length := len(S)
	res := 0
	if length == 0 {
		return res
	}

	for i := 0; i < length; i++ {
		if S[i] == '(' {
			// 遇到开括号就推入
			nodeStack.push(Node{
				base: 0,
			})
		} else {
			// 处理闭括号
			prevNode := nodeStack.pop()
			// 计算当前节点值
			curVal := 0
			if prevNode.base == 0 {
				curVal = 1
			} else {
				curVal = prevNode.base * 2
			}

			// 接着看有没有父节点
			if nodeStack.isEmpty() {
				// 没有父节点，直接累加到全局
				res += curVal
			} else {
				// 有父节点，累加到父节点
				fatherNode := nodeStack.pop()
				fatherNode.base += curVal
				nodeStack.push(fatherNode)
			}
		}
	}
	return res

	// 2. 基于递归
	// index := 0
	// res, _ := dfs(S, index)
	// return res
}


// @lc code=end


package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=636 lang=golang
 *
 * [636] 函数的独占时间
 *
 * https://leetcode-cn.com/problems/exclusive-time-of-functions/description/
 *
 * algorithms
 * Medium (48.01%)
 * Likes:    27
 * Dislikes: 0
 * Total Accepted:    1.2K
 * Total Submissions: 2.6K
 * Testcase Example:  '2\n["0:start:0","1:start:2","1:end:5","0:end:6"]'
 *
 * 给出一个非抢占单线程CPU的 n 个函数运行日志，找到函数的独占时间。
 *
 * 每个函数都有一个唯一的 Id，从 0 到 n-1，函数可能会递归调用或者被其他函数调用。
 *
 * 日志是具有以下格式的字符串：function_id：start_or_end：timestamp。例如："0:start:0" 表示函数 0 从 0
 * 时刻开始运行。"0:end:0" 表示函数 0 在 0 时刻结束。
 *
 * 函数的独占时间定义是在该方法中花费的时间，调用其他函数花费的时间不算该函数的独占时间。你需要根据函数的 Id 有序地返回每个函数的独占时间。
 *
 * 示例 1:
 *
 * 输入:
 * n = 2
 * logs =
 * ["0:start:0",
 * ⁠"1:start:2",
 * ⁠"1:end:5",
 * ⁠"0:end:6"]
 * 输出:[3, 4]
 * 说明：
 * 函数 0 在时刻 0 开始，在执行了  2个时间单位结束于时刻 1。
 * 现在函数 0 调用函数 1，函数 1 在时刻 2 开始，执行 4 个时间单位后结束于时刻 5。
 * 函数 0 再次在时刻 6 开始执行，并在时刻 6 结束运行，从而执行了 1 个时间单位。
 * 所以函数 0 总共的执行了 2 +1 =3 个时间单位，函数 1 总共执行了 4 个时间单位。
 *
 *
 * 说明：
 *
 *
 * 输入的日志会根据时间戳排序，而不是根据日志Id排序。
 * 你的输出会根据函数Id排序，也就意味着你的输出数组中序号为 0 的元素相当于函数 0 的执行时间。
 * 两个函数不会在同时开始或结束。
 * 函数允许被递归调用，直到运行结束。
 * 1 <= n <= 100
 *
 *
 */

// @lc code=start
type Log struct {
	id      int
	logType string
	time    int // 记录日志打点时间
	gap     int // 记录中间由于其他子任务导致的空白时长
}

type Mystack []Log

func (this *Mystack) push(a Log) {
	(*this) = append((*this), a)
}

func (this *Mystack) isEmpty() bool {
	return len((*this)) == 0
}

func (this *Mystack) pop() Log {
	length := len(*this)
	tmp := (*this)[length-1]
	(*this) = (*this)[0 : length-1]
	return tmp
}

func (this *Mystack) top() Log {
	length := len(*this)
	tmp := (*this)[length-1]
	return tmp
}

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	stack := Mystack{}
	for _, logStr := range logs {
		newLog := resolveLog(logStr)
		if newLog.logType == "start" {
			// 遇到start，就需要新入栈日志
			stack.push(newLog)
		} else {
			// 如果是end，计算自己的时长：end-start-gap
			myOld := stack.pop()
			spendTime := newLog.time - myOld.time + 1 - myOld.gap
			// 这里+而不是直接设置，因为可能同一个函数会分段&嵌套执行，这种时候是累加的需要
			res[myOld.id] += spendTime

			if !stack.isEmpty() {
				// 如果stack不空，说明有父元素，同时要把自己的使用时长累加到父进程的gap时间中
				myPrev := stack.pop()
				// 上一层gap需要累加，累加值是这一层的开始到结束，不减去这一层的gap
				myPrev.gap += newLog.time - myOld.time + 1
				stack.push(myPrev)
			}
		}
	}

	return res
}

func resolveLog(str string) Log {
	strArr := strings.Split(str, ":")
	id, _ := strconv.Atoi(strArr[0])
	time, _ := strconv.Atoi(strArr[2])
	return Log{
		id:      id,
		logType: strArr[1],
		time:    time,
	}
}

// func main() {
// 	a := []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}
// 	n := 1
// 	// 8
// 	fmt.Println(exclusiveTime(n, a))
// }

// @lc code=end

package main

/*
 * @lc app=leetcode.cn id=767 lang=golang
 *
 * [767] 重构字符串
 *
 * https://leetcode-cn.com/problems/reorganize-string/description/
 *
 * algorithms
 * Medium (35.67%)
 * Likes:    49
 * Dislikes: 0
 * Total Accepted:    3.3K
 * Total Submissions: 8.7K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。
 *
 * 若可行，输出任意可行的结果。若不可行，返回空字符串。
 *
 * 示例 1:
 *
 *
 * 输入: S = "aab"
 * 输出: "aba"
 *
 *
 * 示例 2:
 *
 *
 * 输入: S = "aaab"
 * 输出: ""
 *
 *
 * 注意:
 *
 *
 * S 只包含小写字母并且长度在[1, 500]区间内。
 *
 *
 */

// @lc code=start
type Ele struct {
	title rune
	num   int
}

type Mystack struct {
	list []Ele
}

func (this *Mystack) ShiftUp(index int) {
	topIndex := (index - 1) / 2
	if topIndex >= 0 && this.list[index].num > this.list[topIndex].num {
		this.list[index], this.list[topIndex] = this.list[topIndex], this.list[index]
		this.ShiftUp(topIndex)
	}
}

func (this *Mystack) ShiftDown(index int) {
	leftIndex, rightIndex := index*2+1, index*2+2
	largeIndex := index
	length := len(this.list)

	if leftIndex < length && this.list[leftIndex].num > this.list[largeIndex].num {
		largeIndex = leftIndex
	}
	if rightIndex < length && this.list[rightIndex].num > this.list[largeIndex].num {
		largeIndex = rightIndex
	}

	if largeIndex != index {
		this.list[largeIndex], this.list[index] = this.list[index], this.list[largeIndex]
		this.ShiftDown(largeIndex)
	}
}

func (this *Mystack) Queue(ele Ele) {
	this.list = append(this.list, ele)
	newLen := len(this.list)
	this.ShiftUp(newLen - 1)
}

func (this *Mystack) Dequeue() Ele {
	oldLen := len(this.list)
	topEle := this.list[0]
	this.list[0] = this.list[oldLen-1]
	this.list = this.list[0 : oldLen-1]
	this.ShiftDown(0)
	return topEle
}

func (this *Mystack) Len() int {
	return len(this.list)
}

func reorganizeString(S string) string {
	// 统计词频，只要最大的不大于N/2就可以，因此思路就是每次都吐出2个最多和第二多的，然后重新统计词频(统计后的肯定也满足最大的不大于N/2)
	// 所以可以基于优先队列来做, 最大堆
	// 排序有讲究，为了防止词频相同情况下无法排序, 还要按照字母顺序排一下, 或者能保证推入和推出顺序一样也可以

	// 先统计词频
	N := (len(S) + 1) / 2
	statMap := map[rune]int{}
	for _, val := range S {
		statMap[val]++
		// 一旦有一个字母频次大于N, 证明无法重构
		if statMap[val] > N {
			return ""
		}
	}

	// 组建优先队列
	mystack := Mystack{}
	for title, num := range statMap {
		mystack.Queue(Ele{
			title: title,
			num:   num,
		})
	}

	res := ""
	// 循环输出
	for mystack.Len() > 0 {
		first := mystack.Dequeue()
		res = res + string(first.title)
		first.num--

		var second Ele
		if mystack.Len() > 0 {
			second = mystack.Dequeue()
			res = res + string(second.title)
			second.num--
		}

		if first.num > 0 {
			mystack.Queue(first)
		}
		if second.num > 0 {
			mystack.Queue(second)
		}
	}

	return res
}

// func main() {
// 	fmt.Println(reorganizeString("abbabbaaab"))
// }

// @lc code=end

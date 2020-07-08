/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 *
 * https://leetcode-cn.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (58.30%)
 * Likes:    197
 * Dislikes: 0
 * Total Accepted:    26.4K
 * Total Submissions: 44.8K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
 *
 * 示例 1:
 *
 * 输入: nums = [1,1,1,2,2,3], k = 2
 * 输出: [1,2]
 *
 *
 * 示例 2:
 *
 * 输入: nums = [1], k = 1
 * 输出: [1]
 *
 * 说明：
 *
 *
 * 你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
 * 你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
 *
 *
 */

// @lc code=start
type StatInfo struct {
	num   int
	times int
}

type MyStack struct {
	list []StatInfo
	k    int
}

func (this *MyStack) shiftUp(index int) {
	topIndex := (index - 1) / 2
	if topIndex >= 0 && this.list[index].times < this.list[topIndex].times {
		this.list[index], this.list[topIndex] = this.list[topIndex], this.list[index]
		this.shiftUp(topIndex)
	}
}

func (this *MyStack) shiftDown(index int) {
	leftIndex, rightIndex := index*2+1, index*2+2
	minIndex := index
	length := len(this.list)
	if leftIndex < length && this.list[leftIndex].times < this.list[minIndex].times {
		minIndex = leftIndex
	}
	if rightIndex < length && this.list[rightIndex].times < this.list[minIndex].times {
		minIndex = rightIndex
	}

	if minIndex != index {
		this.list[index], this.list[minIndex] = this.list[minIndex], this.list[index]
		this.shiftDown(minIndex)
	}
}

func (this *MyStack) queue(a StatInfo) {
	this.list = append(this.list, a)
	newLen := len(this.list)
	this.shiftUp(newLen - 1)
	if newLen > this.k {
		this.dequeue()
	}
}

func (this *MyStack) dequeue() StatInfo {
	length := len(this.list)
	tmp := this.list[0]
	this.list[0] = this.list[length-1]
	this.list = this.list[0 : length-1]
	this.shiftDown(0)
	return tmp
}

func topKFrequent(nums []int, k int) []int {
	// 第一种解法：统计词频后，执行排序, 复杂度O(nlogn)
	// type statInfo struct {
	// 	num   int
	// 	times int
	// }
	// statMap := map[int]int{}
	// for _, val := range nums {
	// 	statMap[val]++
	// }
	// statArr := []statInfo{}
	// for key, val := range statMap {
	// 	statArr = append(statArr, statInfo{
	// 		num:   key,
	// 		times: val,
	// 	})
	// }
	// sort.Slice(statArr, func(i, j int) bool {
	// 	return statArr[i].times > statArr[j].times
	// })

	// res := []int{}
	// for i := 0; i < k; i++ {
	// 	res = append(res, statArr[i].num)
	// }
	// return res

	// 第二种解法：统计词频后，执行优先队列, 复杂度O(nlogk)
	statMap := map[int]int{}
	for _, val := range nums {
		statMap[val]++
	}

	stack := MyStack{
		k: k,
	}
	for num, times := range statMap {
		stack.queue(StatInfo{
			num:   num,
			times: times,
		})
	}

	res := []int{}
	for len(stack.list) > 0 {
		res = append(res, stack.dequeue().num)
	}
	return res
}

// @lc code=end


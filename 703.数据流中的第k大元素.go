/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第K大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/description/
 *
 * algorithms
 * Easy (40.20%)
 * Likes:    82
 * Dislikes: 0
 * Total Accepted:    9.8K
 * Total Submissions: 23.9K
 * Testcase Example:  '["KthLargest","add","add","add","add","add"]\n' +
  '[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]'
 *
 * 设计一个找到数据流中第K大元素的类（class）。注意是排序后的第K大元素，不是第K个不同的元素。
 *
 * 你的 KthLargest 类需要一个同时接收整数 k 和整数数组nums 的构造器，它包含数据流中的初始元素。每次调用
 * KthLargest.add，返回当前数据流中第K大的元素。
 *
 * 示例:
 *
 *
 * int k = 3;
 * int[] arr = [4,5,8,2];
 * KthLargest kthLargest = new KthLargest(3, arr);
 * kthLargest.add(3);   // returns 4
 * kthLargest.add(5);   // returns 5
 * kthLargest.add(10);  // returns 5
 * kthLargest.add(9);   // returns 8
 * kthLargest.add(4);   // returns 8
 *
 *
 * 说明:
 * 你可以假设 nums 的长度≥ k-1 且k ≥ 1。
 *
*/

// @lc code=start
package main

type KthLargest struct {
	list     []int
	maxLimit int
}

func (this *KthLargest) shiftUp(index int) {
	topIndex := (index - 1) / 2
	if topIndex >= 0 && this.list[topIndex] > this.list[index] {
		this.list[topIndex], this.list[index] = this.list[index], this.list[topIndex]
		this.shiftUp(topIndex)
	}
}

func (this *KthLargest) shiftDown(index int) {
	leftIndex, rightIndex := index*2+1, index*2+2
	minIndex := index
	length := len(this.list)
	if leftIndex < length && this.list[leftIndex] < this.list[minIndex] {
		minIndex = leftIndex
	}
	if rightIndex < length && this.list[rightIndex] < this.list[minIndex] {
		minIndex = rightIndex
	}

	if index != minIndex {
		this.list[index], this.list[minIndex] = this.list[minIndex], this.list[index]
		this.shiftDown(minIndex)
	}
}

func (this *KthLargest) queue(a int) {
	this.list = append(this.list, a)
	newLen := len(this.list)
	this.shiftUp(newLen - 1)
}

func (this *KthLargest) dequeue() {
	newLen := len(this.list)
	this.list[0] = this.list[newLen-1]
	this.list = this.list[0 : newLen-1]
	this.shiftDown(0)
}

func Constructor(k int, nums []int) KthLargest {
	res := KthLargest{
		list:     []int{},
		maxLimit: k,
	}
	for _, val := range nums {
		res.Add(val)
	}
	return res
}

func (this *KthLargest) Add(val int) int {
	this.queue(val)
	if len(this.list) > this.maxLimit {
		this.dequeue()
	}
	return this.list[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// func main() {
// 	kthLargest := Constructor(3, []int{4, 5, 8, 2})
// 	fmt.Println(kthLargest.Add(3))
// 	fmt.Println(kthLargest.Add(5))
// 	fmt.Println(kthLargest.Add(10))
// 	fmt.Println(kthLargest.Add(9))
// 	fmt.Println(kthLargest.Add(4))
// }

// @lc code=end

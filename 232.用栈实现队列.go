/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 *
 * https://leetcode-cn.com/problems/implement-queue-using-stacks/description/
 *
 * algorithms
 * Easy (61.15%)
 * Likes:    140
 * Dislikes: 0
 * Total Accepted:    33K
 * Total Submissions: 52.3K
 * Testcase Example:  '["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * 使用栈实现队列的下列操作：
 * 
 * 
 * push(x) -- 将一个元素放入队列的尾部。
 * pop() -- 从队列首部移除元素。
 * peek() -- 返回队列首部的元素。
 * empty() -- 返回队列是否为空。
 * 
 * 
 * 示例:
 * 
 * MyQueue queue = new MyQueue();
 * 
 * queue.push(1);
 * queue.push(2);  
 * queue.peek();  // 返回 1
 * queue.pop();   // 返回 1
 * queue.empty(); // 返回 false
 * 
 * 说明:
 * 
 * 
 * 你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty
 * 操作是合法的。
 * 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
 * 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
 * 
 * 
 */

// @lc code=start
type MyStack struct {
	elements []int	
}

func (this *MyStack) Push(x int) {
	this.elements = append(this.elements, x)
}

func (this *MyStack) IsEmpty() bool {
	return len(this.elements) == 0
}

func (this *MyStack) Size() int {
	return len(this.elements)
}

func (this *MyStack) Pop() int {
	length := this.Size()
	ele := this.elements[length - 1]
	this.elements = this.elements[0:length - 1]
	return ele
}

func (this *MyStack) Top() int {
	length := this.Size()
	ele := this.elements[length - 1]
	return ele
}


type MyQueue struct {
	cache1 MyStack
	cache2 MyStack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.cache1.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.cache2.IsEmpty() {
		for !this.cache1.IsEmpty() {
			this.cache2.Push(this.cache1.Pop())
		}
	}
	return this.cache2.Pop()
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.cache2.IsEmpty() {
		for !this.cache1.IsEmpty() {
			this.cache2.Push(this.cache1.Pop())
		}
	}
	return this.cache2.Top()
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.cache1.IsEmpty() && this.cache2.IsEmpty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end


// 两个栈，一个正常，另一个使用单调递减栈即可。只不过这样，push的复杂度不一定是O(1)
type MinStack struct {
	list1 []int
	list2 []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	this.list1 = append(this.list1, x)
	// 单调递减栈
	len2 := len(this.list2)
	if len2 == 0 {
		this.list2 = append(this.list2, x)
		return
	}

	topEle := this.list2[len2 - 1]
	if topEle > x {
		this.list2 = append(this.list2, x)
	} else {
		// 否则，重复推入一次栈顶元素
		this.list2 = append(this.list2, topEle)
	}
}


func (this *MinStack) Pop()  {
	this.list1 = this.list1[0:len(this.list1) - 1]
	this.list2 = this.list2[0:len(this.list2) - 1]
}


func (this *MinStack) Top() int {
	return this.list1[len(this.list1) - 1]
}


func (this *MinStack) Min() int {
	return this.list2[len(this.list2) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
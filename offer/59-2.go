type MaxQueue struct {
	// list1存放全部数据
	list1 []int
	// list2则是单调递减栈
	list2 []int
}


func Constructor() MaxQueue {
	return MaxQueue{list1: []int{}, list2: []int{}}
}


func (this *MaxQueue) Max_value() int {
	if len(this.list2) == 0 {
		return -1
	}
	return this.list2[0]
}


func (this *MaxQueue) Push_back(value int)  {
	this.list1 = append(this.list1, value)

	for len2 := len(this.list2); len2 > 0 && this.list2[len2 - 1] < value; len2-- {
		this.list2 = this.list2[0:len2 - 1]
	}
	this.list2 = append(this.list2, value)
}


func (this *MaxQueue) Pop_front() int {
	if len(this.list1) == 0 {
		return -1
	}

	popValue := this.list1[0]
	this.list1 = this.list1[1:]
	if popValue == this.list2[0] {
		this.list2 = this.list2[1:]
	}
	return popValue
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
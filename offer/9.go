type CQueue struct {
	list1 []int
	list2 []int
}


func Constructor() CQueue {
	return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
	this.list1 = append(this.list1, value)
}


func (this *CQueue) DeleteHead() int {
	// 先看list2是否空
	len2 := len(this.list2)
	if len2 > 0 {
		head := this.list2[len2-1]
		this.list2 = this.list2[0:len2-1]
		return head
	}
	
	// 否则看list1是否为空
	len1 := len(this.list1)
	if len1 == 0 {
		return -1
	}

	for i := len1 - 1; i >= 0; i-- {
		this.list2 = append(this.list2, this.list1[i])
	}
	this.list1 = []int{}
	return this.DeleteHead()
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
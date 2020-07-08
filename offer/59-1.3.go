type MaxQueue struct {
	list1 []int
	list2 []int
}

func (this *MaxQueue) Queue(val int) {
	this.list1 = append(this.list1, val)
	// 类似单调栈，需要将小于自己的全部推出
	len2 := len(this.list2)
	for len2 > 0 && this.list2[len2-1] < val {
		this.list2 = this.list2[0:len2-1]
		len2--
	}
	this.list2 = append(this.list2, val)
}

func (this *MaxQueue) Dequeue() {
	val := this.list1[0]
	this.list1 = this.list1[1:]
	if val == this.list2[0] {
		this.list2 = this.list2[1:]
	}
}


func (this *MaxQueue) GetMax() int {
	return this.list2[0]
}


func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	length := len(nums)
	if length < k || length == 0 {
		return res
	}

	maxQueue := MaxQueue{}
	for i := 0; i < length; i++ {
		if i < k {
			maxQueue.Queue(nums[i])
			if i == k - 1 {
				res = append(res, maxQueue.GetMax())
			}
		} else {
			maxQueue.Dequeue()
			maxQueue.Queue(nums[i])
			res = append(res, maxQueue.GetMax())
		}
	}
	return res
}
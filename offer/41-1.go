// 第一种方法：基于二分查找+二分查找插入的方法
type MedianFinder struct {
	list []int
	len int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{list: []int{}}
}


func (this *MedianFinder) AddNum(num int)  {
	// 先二分查找再二分插入
	if this.len == 0 {
		this.list = append(this.list, num)
		this.len++
		return
	}

	l, r := 0, this.len - 1
	var mid int
	for l < r {
		mid = (l + r) / 2
		if this.list[mid] > num {
			r = mid - 1
		} else if this.list[mid] < num {
			l = mid + 1
		} else {
			l = mid
			break
		}
	}

	// 执行插入操作
	insertIndex := l
	if this.list[l] < num {
		insertIndex = l + 1
	}
	this.list = append(this.list, 0)
	this.len++
	for i := this.len - 1; i > insertIndex; i-- {
		this.list[i] = this.list[i-1]
	}
	this.list[insertIndex] = num
}


func (this *MedianFinder) FindMedian() float64 {
	if this.len % 2 == 1 {
		return float64(this.list[this.len / 2])
	}
	return (float64(this.list[(this.len - 1) / 2]) + float64(this.list[this.len / 2])) / 2
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
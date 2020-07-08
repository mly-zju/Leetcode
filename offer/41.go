// 重点在于在流中每次快速找到中间值。可以基于搜索二叉树(或者二分插入，查找比二叉树快)，或者基于优先队列(容量永远为当前所有元素的(count+1)/2)

// 最小堆优先队列
type Heap struct {
	list []int
}

func (this *Heap) Len() int {
	return len(this.list)
}

func (this *Heap) ShiftUp(index int) {
	pIndex := (index - 1) / 2
	if pIndex >= 0 && this.list[index] < this.list[pIndex] {
		this.list[index], this.list[pIndex] = this.list[pIndex], this.list[index]
		this.ShiftUp(pIndex)
	}
} 

func (this *Heap) ShiftDown(index int) {
	leftIndex, rightIndex := index * 2 + 1, index * 2 + 2
	minIndex := index
	length := this.Len()

	if leftIndex < length && this.list[minIndex] > this.list[leftIndex] {
		minIndex = leftIndex
	}
	if rightIndex < length && this.list[minIndex] > this.list[rightIndex] {
		minIndex = rightIndex
	}

	if minIndex != index {
		this.list[index], this.list[minIndex] = this.list[minIndex], this.list[index]
		this.ShiftDown(minIndex)
	}
}

func (this *Heap) Queue(ele int) {
	this.list = append(this.list, ele)
	this.ShiftUp(len(this.list) - 1)
}

func (this *Heap) Dequeue() int {
	length := this.Len()
	dropEle := this.list[0]
	this.list[0] = this.list[length - 1]
	this.list = this.list[0:length - 1]
	this.ShiftDown(0)
	return dropEle
}

func (this *Heap) Top() int {
	return this.list[0]
}

type MedianFinder struct {
	heap *Heap
	count int
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		heap: &Heap{},
	}
}


func (this *MedianFinder) AddNum(num int)  {
	this.heap.Queue(num)
	this.count++
	// 计算当前堆的容量上限
	var limit int
	if this.count % 2 == 0 {
		limit = this.count / 2 + 1
	} else {
		limit = this.count / 2 + 2
	}
	// 去除多于limit的元素
	if this.heap.Len() > limit {
		this.heap.Dequeue()
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.heap.Len() == 1 {
		return float64(this.heap.Top())
	}
	ele1 := this.heap.Dequeue()
	ele2 := this.heap.Dequeue()
	// 再推回去
	this.heap.Queue(ele1)
	this.heap.Queue(ele2)
	if this.count % 2 == 0 {
		return (float64(ele1) + float64(ele2)) / 2
	}
	return float64(ele2)
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
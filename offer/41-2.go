// 第二种方法：基于两个堆，左边最大堆右边最小堆
/*
假想我们现在有两个容器 A, B 这两个容器将我们的整体数据分成两部分,且 A 中的数据都小于 B 中的数据,并且 A 中的最后一个数据是 A 里面最大的
B 的第一个数据是 B 中最小的.
好了 我们有了上面的条件,那我们怎么找到中位数呢?
当整体数目为奇数时,中间的那个数就是所求.当整体数目为偶数时,中间两个数的和再除以 2 ,就能得到结果
但这和我们上面的两个容器有什么关系呢?
我们只要将上面的两个容器的数据数目只差保持在 1 之内即可,也就是说 A 和 B 将整体数据以中位数划分开来了
*/


// 优先队列
type Heap struct {
	list []int
	isBig bool // 是否最大堆，true是最大堆，否则是最小堆
}

func (this *Heap) Len() int {
	return len(this.list)
}

func (this *Heap) compare(i, j int) bool {
	if this.isBig {
		return this.list[i] > this.list[j]
	}
	return this.list[i] < this.list[j]
}

func (this *Heap) ShiftUp(index int) {
	pIndex := (index - 1) / 2
	if pIndex >= 0 && this.compare(index, pIndex) {
		this.list[index], this.list[pIndex] = this.list[pIndex], this.list[index]
		this.ShiftUp(pIndex)
	}
} 

func (this *Heap) ShiftDown(index int) {
	leftIndex, rightIndex := index * 2 + 1, index * 2 + 2
	sIndex := index
	length := this.Len()

	if leftIndex < length && this.compare(leftIndex, sIndex) {
		sIndex = leftIndex
	}
	if rightIndex < length && this.compare(rightIndex, sIndex) {
		sIndex = rightIndex
	}

	if sIndex != index {
		this.list[index], this.list[sIndex] = this.list[sIndex], this.list[index]
		this.ShiftDown(sIndex)
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
	BigHeap Heap
	SmallHeap Heap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		BigHeap: Heap{
			list: []int{},
			isBig: true,
		},
		SmallHeap: Heap{
			list: []int{},
			isBig: false,
		},
	}
}


func (this *MedianFinder) AddNum(num int)  {
	bigLen, smallLen := this.BigHeap.Len(), this.SmallHeap.Len()
	// small的新元素总是从big来，保证small的所有都是大于big
	// 同样，big的新元素总是从small来，保证big所有都小于small
	if bigLen == smallLen {
		this.BigHeap.Queue(num)
		biggest := this.BigHeap.Dequeue()
		this.SmallHeap.Queue(biggest)
	} else {
		this.SmallHeap.Queue(num)
		smallest := this.SmallHeap.Dequeue()
		this.BigHeap.Queue(smallest)
	}
}


func (this *MedianFinder) FindMedian() float64 {
	bigLen, smallLen := this.BigHeap.Len(), this.SmallHeap.Len()
	if (bigLen + smallLen) % 2 == 0 {
		ele1, ele2 := this.BigHeap.Dequeue(), this.SmallHeap.Dequeue()
		// 再填充回去
		this.BigHeap.Queue(ele1)
		this.SmallHeap.Queue(ele2)
		return float64(ele1 + ele2) / 2
	} else {
		ele := this.SmallHeap.Dequeue()
		this.SmallHeap.Queue(ele)
		return float64(ele)
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
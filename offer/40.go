// 最大堆优先队列
type Heap struct {
	list []int
}

func (this *Heap) Len() int {
	return len(this.list)
}

func (this *Heap) ShiftUp(index int) {
	pIndex := (index - 1) / 2
	if pIndex >= 0 && this.list[index] > this.list[pIndex] {
		this.list[index], this.list[pIndex] = this.list[pIndex], this.list[index]
		this.ShiftUp(pIndex)
	}
} 

func (this *Heap) ShiftDown(index int) {
	leftIndex, rightIndex := index * 2 + 1, index * 2 + 2
	largetIndex := index
	length := this.Len()

	if leftIndex < length && this.list[largetIndex] < this.list[leftIndex] {
		largetIndex = leftIndex
	}
	if rightIndex < length && this.list[largetIndex] < this.list[rightIndex] {
		largetIndex = rightIndex
	}

	if largetIndex != index {
		this.list[index], this.list[largetIndex] = this.list[largetIndex], this.list[index]
		this.ShiftDown(largetIndex)
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

func getLeastNumbers(arr []int, k int) []int {
	myHeap := &Heap{}
	for _, num := range arr {
		myHeap.Queue(num)
		if myHeap.Len() > k {
			myHeap.Dequeue()
		}
	}
	return myHeap.list
}
/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 *
 * https://leetcode-cn.com/problems/task-scheduler/description/
 *
 * algorithms
 * Medium (45.27%)
 * Likes:    131
 * Dislikes: 0
 * Total Accepted:    8.3K
 * Total Submissions: 18.2K
 * Testcase Example:  '["A","A","A","B","B","B"]\n2'
 *
 * 给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26
 * 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU
 * 在任何一个单位时间内都可以执行一个任务，或者在待命状态。
 * 
 * 然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
 * 
 * 你需要计算完成所有任务所需要的最短时间。
 * 
 * 示例 1：
 * 
 * 
 * 输入: tasks = ["A","A","A","B","B","B"], n = 2
 * 输出: 8
 * 执行顺序: A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
 * 
 * 
 * 注：
 * 
 * 
 * 任务的总个数为 [1, 10000]。
 * n 的取值范围为 [0, 100]。
 * 
 * 
 */

// @lc code=start
// 最大堆
type MyHeap struct {
	list []int
}

func (h *MyHeap) ShiftDown(index int) {
	leftIndex, rightIndex := index * 2 + 1, index * 2 + 2
	maxIndex := index
	length := len(h.list)

	if leftIndex < length && h.list[leftIndex] > h.list[maxIndex] {
		maxIndex = leftIndex
	}
	if rightIndex < length && h.list[rightIndex] > h.list[maxIndex] {
		maxIndex = rightIndex
	}

	if maxIndex != index {
		h.list[maxIndex], h.list[index] = h.list[index], h.list[maxIndex]
		h.ShiftDown(maxIndex)
	}
}

func (h *MyHeap) ShiftUp(index int) {
	topIndex := (index - 1) / 2
	if topIndex < 0 {
		return
	}

	if h.list[index] > h.list[topIndex] {
		h.list[index], h.list[topIndex] = h.list[topIndex], h.list[index]
		h.ShiftUp(topIndex)
	}
}

func (h *MyHeap) Enqueue(val int) {
	oldLen := len(h.list)
	h.list = append(h.list, val)
	h.ShiftUp(oldLen)
}

func (h *MyHeap) Dequeue() int {
	res := h.list[0]
	length := len(h.list)
	h.list[0] = h.list[length - 1]
	h.list = h.list[0:length-1]
	h.ShiftDown(0)
	return res
}

func (h *MyHeap) IsEmpty() bool {
	return len(h.list) == 0
}

func leastInterval(tasks []byte, n int) int {
	// 先构造统计数组
	statMap := make(map[byte]int)
	for _, by := range tasks {
		statMap[by]++
	}
	// 构造优先队列
	var myHeap MyHeap
	for _, num := range statMap {
		myHeap.Enqueue(num)
	}

	// 开始执行任务调度
	res := 0
	for !myHeap.IsEmpty() {
		cache := []int{}
		roundCount := n + 1
		for roundCount > 0 {
			if !myHeap.IsEmpty() {
				// 如果有任务，执行调度
				tmp := myHeap.Dequeue()
				tmp--
				if tmp > 0 {
					cache = append(cache, tmp)
				}
				res++
			} else if len(cache) != 0 {
				// 没有可执行任务，执行空转
				res++
			} else {
				// 否则如果任务完全为空，停止执行
				break
			}
			roundCount--
		}
		// 调度过后，任务恢复可执行
		for _, num := range cache {
			myHeap.Enqueue(num)
		}
	}

	return res
}
// @lc code=end


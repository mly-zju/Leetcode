/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的K对数字
 *
 * https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums/description/
 *
 * algorithms
 * Medium (35.62%)
 * Likes:    44
 * Dislikes: 0
 * Total Accepted:    3.5K
 * Total Submissions: 9K
 * Testcase Example:  '[1,7,11]\n[2,4,6]\n3'
 *
 * 给定两个以升序排列的整形数组 nums1 和 nums2, 以及一个整数 k。
 *
 * 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2。
 *
 * 找到和最小的 k 对数字 (u1,v1), (u2,v2) ... (uk,vk)。
 *
 * 示例 1:
 *
 * 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
 * 输出: [1,2],[1,4],[1,6]
 * 解释: 返回序列中的前 3 对数：
 * ⁠    [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
 *
 *
 * 示例 2:
 *
 * 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
 * 输出: [1,1],[1,1]
 * 解释: 返回序列中的前 2 对数：
 * [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
 *
 *
 * 示例 3:
 *
 * 输入: nums1 = [1,2], nums2 = [3], k = 3
 * 输出: [1,3],[2,3]
 * 解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
 *
 *
 */

// @lc code=start
type Pair struct {
	a int
	b int
}

type Mystack struct {
	list []Pair
}

func (this *Mystack) Sum(index int) int {
	return this.list[index].a + this.list[index].b
}

func (this *Mystack) ShiftUp(index int) {
	topIndex := (index - 1) / 2
	if topIndex >= 0 && this.Sum(index) > this.Sum(topIndex) {
		this.list[index], this.list[topIndex] = this.list[topIndex], this.list[index]
		this.ShiftUp(topIndex)
	}
}

func (this *Mystack) ShiftDown(index int) {
	lefti, righti := index*2+1, index*2+2
	largei := index
	length := len(this.list)
	if lefti < length && this.Sum(largei) < this.Sum(lefti) {
		largei = lefti
	}
	if righti < length && this.Sum(largei) < this.Sum(righti) {
		largei = righti
	}

	if largei != index {
		this.list[index], this.list[largei] = this.list[largei], this.list[index]
		this.ShiftDown(largei)
	}
}

func (this *Mystack) Queue(ele Pair) {
	this.list = append(this.list, ele)
	newLen := len(this.list)
	this.ShiftUp(newLen - 1)
}

func (this *Mystack) Dequeue() Pair {
	var tmp Pair
	length := len(this.list)
	if length == 0 {
		return tmp
	}

	tmp = this.list[0]
	this.list[0] = this.list[length-1]
	this.list = this.list[0 : length-1]
	this.ShiftDown(0)
	return tmp
}

func (this *Mystack) len() int {
	return len(this.list)
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// 1. 基于堆来做
	stack := Mystack{}
	res := [][]int{}
	for i, length1 := 0, len(nums1); i < length1; i++ {
		for j, length2 := 0, len(nums2); j < length2; j++ {
			stack.Queue(Pair{
				a: nums1[i],
				b: nums2[j],
			})

			for stack.len() > k {
				stack.Dequeue()
			}
		}
	}

	for stack.len() > 0 {
		pair := stack.Dequeue()
		res = append(res, []int{pair.a, pair.b})
	}
	// 反转
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}

// @lc code=end


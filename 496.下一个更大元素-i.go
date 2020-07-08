/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 *
 * https://leetcode-cn.com/problems/next-greater-element-i/description/
 *
 * algorithms
 * Easy (61.33%)
 * Likes:    117
 * Dislikes: 0
 * Total Accepted:    14.7K
 * Total Submissions: 23.6K
 * Testcase Example:  '[4,1,2]\n[1,3,4,2]'
 *
 * 给定两个没有重复元素的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2
 * 中的下一个比其大的值。
 *
 * nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出-1。
 *
 * 示例 1:
 *
 *
 * 输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
 * 输出: [-1,3,-1]
 * 解释:
 * ⁠   对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
 * ⁠   对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
 * ⁠   对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
 *
 * 示例 2:
 *
 *
 * 输入: nums1 = [2,4], nums2 = [1,2,3,4].
 * 输出: [3,-1]
 * 解释:
 * 对于num1中的数字2，第二个数组中的下一个较大数字是3。
 * ⁠   对于num1中的数字4，第二个数组中没有下一个更大的数字，因此输出 -1。
 *
 *
 * 注意:
 *
 *
 * nums1和nums2中所有元素是唯一的。
 * nums1和nums2 的数组大小都不超过1000。
 *
 *
 */
// @lc code=start
type Mystack []int

func (this *Mystack) push(a int) {
	(*this) = append((*this), a)
}

func (this *Mystack) isEmpty() bool {
	return len((*this)) == 0
}

func (this *Mystack) pop() int {
	length := len(*this)
	tmp := (*this)[length-1]
	(*this) = (*this)[0 : length-1]
	return tmp
}

func (this *Mystack) top() int {
	length := len(*this)
	tmp := (*this)[length-1]
	return tmp
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 1. 暴力法
	// len1, len2 := len(nums1), len(nums2)
	// res := []int{}
	// for i := 0; i < len1; i++ {
	// 	start := -1
	// 	hasFlag := false
	// 	for j := 0; j < len2; j++ {
	// 		if nums2[j] == nums1[i] {
	// 			start = j
	// 		} else if nums2[j] > nums1[i] && start >= 0 {
	// 			res = append(res, nums2[j])
	// 			hasFlag = true
	// 			break
	// 		}
	// 	}
	// 	if !hasFlag {
	// 		res = append(res, -1)
	// 	}
	// }
	// return res

	// 2. 基于单调栈解决next greater问题
	ans := make(map[int]int)
	var s Mystack
	length := len(nums2)
	for i := length - 1; i >= 0; i-- {
		for !s.isEmpty() && s.top() <= nums2[i] {
			s.pop()
		}
		var answer int
		if s.isEmpty() {
			answer = -1
		} else {
			answer = s.top()
		}

		ans[nums2[i]] = answer
		s.push(nums2[i])
	}

	res := []int{}
	for _, val := range nums1 {
		res = append(res, ans[val])
	}
	return res
}

// @lc code=end


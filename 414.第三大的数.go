package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 *
 * https://leetcode-cn.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (32.76%)
 * Likes:    74
 * Dislikes: 0
 * Total Accepted:    13.2K
 * Total Submissions: 40K
 * Testcase Example:  '[3,2,1]'
 *
 * 给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
 *
 * 示例 1:
 *
 *
 * 输入: [3, 2, 1]
 *
 * 输出: 1
 *
 * 解释: 第三大的数是 1.
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1, 2]
 *
 * 输出: 2
 *
 * 解释: 第三大的数不存在, 所以返回最大的数 2 .
 *
 *
 * 示例 3:
 *
 *
 * 输入: [2, 2, 3, 1]
 *
 * 输出: 1
 *
 * 解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
 * 存在两个值为2的数，它们都排第二。
 *
 *
 */

// @lc code=start
func thirdMax(nums []int) int {
	// 维护一个3个长度的数组，排序+查找插入(三个元素二分查找就没啥优势了，多了有优势)。特别注意题目要求数字不可重复。
	resArr := []int{}
	for _, val := range nums {
		if len(resArr) < 3 && arrIndex(resArr, val) == -1 {
			resArr = append(resArr, val)
			sort.Ints(resArr)
		} else {
			searchInsert(resArr, val)
		}
	}

	lastLen := len(resArr)
	if lastLen < 3 {
		return resArr[lastLen-1]
	}
	return resArr[0]
}

func arrIndex(arr []int, target int) int {
	for index, val := range arr {
		if val == target {
			return index
		}
	}
	return -1
}

func searchInsert(arr []int, target int) {
	if target < arr[0] {
		return
	}

	length := len(arr)
	// 这里使用遍历查找。进一步性能提升可用二分。
	for i := length - 1; i >= 0; i-- {
		// 数字不能重复
		if target == arr[i] {
			break
		} else if target > arr[i] {
			for j := 0; j < i; j++ {
				arr[j] = arr[j+1]
			}
			arr[i] = target
			break
		}
	}
}

// func main() {
// 	nums := []int{5, 2, 4, 1, 3, 6, 0}
// 	res := thirdMax(nums)
// 	fmt.Println(res)
// }

// @lc code=end

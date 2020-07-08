package main

import "sort"

/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 *
 * https://leetcode-cn.com/problems/wiggle-sort-ii/description/
 *
 * algorithms
 * Medium (33.56%)
 * Likes:    70
 * Dislikes: 0
 * Total Accepted:    6K
 * Total Submissions: 17.4K
 * Testcase Example:  '[1,5,1,1,6,4]'
 *
 * 给定一个无序的数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
 *
 * 示例 1:
 *
 * 输入: nums = [1, 5, 1, 1, 6, 4]
 * 输出: 一个可能的答案是 [1, 4, 1, 5, 1, 6]
 *
 * 示例 2:
 *
 * 输入: nums = [1, 3, 2, 2, 3, 1]
 * 输出: 一个可能的答案是 [2, 3, 1, 3, 1, 2]
 *
 * 说明:
 * 你可以假设所有输入都会得到有效的结果。
 *
 * 进阶:
 * 你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？
 *
 */

// @lc code=start
func wiggleSort(nums []int) {
	// 1. O(nlogn)方法：排序后反序，穿插插入
	// 排序后，左右反序，穿插插入
	sort.Ints(nums)
	length := len(nums)

	mid := (length + 1) / 2
	// 左右两边各反序
	l, r := 0, mid-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	l, r = mid, length-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	l = 0
	// 执行穿插
	for mid < length && l+1 < length {
		tmp := nums[mid]
		for j := mid; j > l+1; j-- {
			nums[j] = nums[j-1]
		}
		nums[l+1] = tmp
		mid++
		l += 2
	}

	// 2. O(n)算法：基于快速选择寻找中位数，然后3路partition执行分区,
	// 然后就是左右反序，穿插即可
	// length := len(nums)
	// // k的选择有讲究，要length + 1再除以2，为了奇数的正确
	// k := (length + 1) / 2
	// // 快速选择法获取第k大的元素(不去重)
	// midVal := getKth(nums, k)
	// // 执行3路partition分区
	// l, r, cur := 0, length-1, 0
	// for cur <= r {
	// 	if nums[cur] < midVal {
	// 		nums[l], nums[cur] = nums[cur], nums[l]
	// 		l++
	// 		cur++
	// 	} else if nums[cur] > midVal {
	// 		nums[cur], nums[r] = nums[r], nums[cur]
	// 		r--
	// 	} else {
	// 		cur++
	// 	}
	// }

	// // 排序完毕，左右反序, 左边的和右边相等或者始终比右边多1
	// l, r = 0, k-1
	// for l < r {
	// 	nums[l], nums[r] = nums[r], nums[l]
	// 	l++
	// 	r--
	// }
	// l, r = k, length-1
	// for l < r {
	// 	nums[l], nums[r] = nums[r], nums[l]
	// 	l++
	// 	r--
	// }

	// // 反序完毕，执行穿插
	// cur = 1
	// for i := k; i < length; i++ {
	// 	tmp := nums[i]
	// 	for j := i; j > cur; j-- {
	// 		nums[j] = nums[j-1]
	// 	}
	// 	nums[cur] = tmp
	// 	cur += 2
	// }
}

func getKth(arr []int, k int) int {
	length := len(arr)
	if length == 1 {
		return arr[0]
	}
	// partition算法分组
	midVal := arr[0]
	l, r := 0, length-1
	for l < r {
		for l < r && arr[r] > midVal {
			r--
		}
		if l < r {
			arr[l] = arr[r]
			l++
		}
		for l < r && arr[l] <= midVal {
			l++
		}
		if l < r {
			arr[r] = arr[l]
			r--
		}
	}
	arr[l] = midVal

	if l == k-1 {
		return arr[l]
	} else if l < k-1 {
		return getKth(arr[l+1:], k-1-l)
	} else {
		return getKth(arr[0:l], k)
	}
}

// func main() {
// 	// a := []int{1, 5, 1, 1, 6, 4}
// 	a := []int{5, 8, 1, 6, 3, 7, 2, 5, 5}
// 	// a := []int{1, 1, 2, 1, 2, 2, 1}
// 	wiggleSort(a)
// 	fmt.Println(a)
// }

// @lc code=end

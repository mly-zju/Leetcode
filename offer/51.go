package main

import "fmt"

var res int

func reversePairs(nums []int) int {
	// 基于归并排序来做
	res = 0
	nums = partition(nums)
	fmt.Println(nums)
	return res
}

func partition(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}

	mid := length / 2
	return merge(partition(nums[0:mid]), partition(nums[mid:]))
}

func merge(arr1, arr2 []int) []int {
	len1, len2 := len(arr1), len(arr2)
	i, j := 0, 0
	arr3 := []int{}
	// 从大到小排列
	for i < len1 && j < len2 {
		for i < len1 && j < len2 && arr1[i] > arr2[j] {
			res += len2 - j
			arr3 = append(arr3, arr1[i])
			i++
		}
		for j < len2 && i < len1 && arr2[j] >= arr1[i] {
			arr3 = append(arr3, arr2[j])
			j++
		}
	}
	if i < len1 {
		arr3 = append(arr3, arr1[i:]...)
	}
	if j < len2 {
		arr3 = append(arr3, arr2[j:]...)
	}
	return arr3
}

// func main() {
// 	// fmt.Println(reversePairs([]int{1, 3, 2, 3}))
// 	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
// }

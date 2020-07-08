func exchange(nums []int) []int {
	// 最简单当然是分别统计奇偶数组，然后顺序插入
	// 这里用首尾指针交换法
	length := len(nums)
	l, r := 0, length - 1
	for l < r {
		for l < r && nums[l] % 2 == 1 {
			l++
		}
		for l < r && nums[r] % 2 == 0 {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return nums
}
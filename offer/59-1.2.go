func maxSlidingWindow(nums []int, k int) []int {
	// 简单的滑动窗口法比较
	window := []int{}
	res := []int{}
	max := 0
	for i, length := 0, len(nums); i < length; i++ {
		if i < k {
			window = append(window, nums[i])
			if nums[i] > max {
				max = nums[i]
			}
			if i == k - 1 {
				res = append(res, max)
			}
		} else {
			outNum := window[0]
			window = window[1:]
			window = append(window, nums[i])
			
			// 更新最大值
			if outNum != max {
				// 如果划出去的不是最大值，只需要比较新的和最大值
				if nums[i] > max {
					max = nums[i]
				}
			} else {
				// 如果划出去的是最大值，需要重新寻找窗口内最大值
				max = window[0]
				for _, num := range window {
					if num > max {
						max = num
					}
				}
			}
			res = append(res, max)
		}
	}
	return res
}
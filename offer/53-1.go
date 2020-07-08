func search(nums []int, target int) int {
	// 上下界二分查找
	minIndex := binarySearch(nums, target, false)
	if minIndex == -1 {
		return 0
	}
	maxIndex := binarySearch(nums, target, true)
	return maxIndex - minIndex + 1
}

// 找不到返回-1
func binarySearch(nums []int, target int, isHigh bool) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	l, r := 0, length - 1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			if isHigh {
				if mid + 1 < length && nums[mid + 1] == target {
					l = mid + 1
				} else {
					return mid
				}
			} else {
				if mid - 1 >= 0 && nums[mid - 1] == target {
					r = mid - 1
				} else {
					return mid
				}
			}
		}
	}

	return -1
}

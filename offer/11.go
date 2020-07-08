func minArray(numbers []int) int {
	// 二分查找法, 查找旋转数组最小值, 二分时候，总有一个区间是有序的
	length := len(numbers)
	if length == 1 {
		return numbers[0]
	}

	l, r := 0, length - 1
	var mid int
	for l < r {
		mid = (l + r) / 2
		// 这里整体检测以右边区域为准，因为这样可以解决整个数组是递增的情况，将其缩减到index为0
		if numbers[mid] < numbers[r] {
			// 如果右边有序，说明旋转值在左边, 可能包含mid
			r = mid
		} else if numbers[mid] > numbers[r] {
			// 右边确定无序
			l = mid + 1
		} else if numbers[mid] == numbers[r] {
			// 无法判别右边有序还是无序，但至少知道左边可以缩小1个值
			r--
		}
	}

	// 结束后，进行检测
	return numbers[l]
}

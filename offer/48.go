func lengthOfLongestSubstring(s string) int {
	// 1. 滑动窗口
	// length := len(s)
	// if length == 0 {
	// 	return 0
	// }

	// seenMap := make(map[byte]bool)
	// tmpLen, maxLen := 0, 0
	// l, r := 0, 0
	// for l < length && r < length {
	// 	if !seenMap[s[r]] {
	// 		seenMap[s[r]] = true
	// 		r++
	// 		tmpLen++
	// 	} else {
	// 		// 如果遇到了相等的，循环增加l直到没有相等的为止
	// 		if tmpLen > maxLen {
	// 			maxLen = tmpLen
	// 		}
	// 		for seenMap[s[r]] {
	// 			seenMap[s[l]] = false
	// 			l++
	// 			tmpLen--
	// 		}
	// 	}
	// }
	// if tmpLen > maxLen {
	// 	maxLen = tmpLen
	// }

	// return maxLen

	// 2. 暴力求解
	length := len(s)
	if length < 2 {
		return length
	}

	// 从1开始扫描&try
	for i := 1; i <= length; i++ {
		flag := false
		for j := 0; j <= length - i; j++ {
			seenMap := make(map[byte]bool)
			k := j
			for ; k < j + i; k++ {
				if !seenMap[s[k]] {
					seenMap[s[k]] = true
				} else {
					break
				}
			}
			if k == j + i {
				flag = true
				break
			}
		}

		// 没有flag，说明当前大小不合适
		if !flag {
			return i - 1
		}
	}
	// 如果没有中途退出，说明是最大值
	return length
}
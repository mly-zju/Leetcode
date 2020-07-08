func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 1. 左下角算法
	// rNum := len(matrix)
	// if rNum == 0 {
	// 	return false
	// }
	// cNum := len(matrix[0])
	// if cNum == 0 {
	// 	return false
	// }

	// c, r := 0, rNum - 1
	// res := false
	// var tmp int

	// for c < cNum && r >= 0 {
	// 	tmp = matrix[r][c]	
	// 	if tmp > target {
	// 		r--
	// 	} else if tmp < target {
	// 		c++
	// 	} else {
	// 		res = true
	// 		break
	// 	}
	// }
	// return res

	// 2. 每行二分查找法
	rNum := len(matrix)
	if rNum == 0 {
		return false
	}
	cNum := len(matrix[0])
	if cNum == 0 {
		return false
	}

	res := false
	for i := 0; i < rNum; i++ {
		// 遍历每一行，每行使用二分查找，不过在此之前可以先判断最大和最小值
		l, r := 0, cNum - 1	
		if matrix[i][l] > target || matrix[i][r] < target {
			continue
		}

		var mid int
		for l <= r {
			mid = (l + r) / 2
			if matrix[i][mid] < target {
				l = mid + 1
			} else if matrix[i][mid] > target {
				r = mid - 1
			} else {
				res = true
				break
			}
		}
	}
	return res
}
func spiralOrder(matrix [][]int) []int {
	// 模拟题，一个状态变量记录当前方向，一个hash记录某个点是否访问过
	// 方向有L, R, U, D四种
	// direct := 'R'
	// res := []int{}
	// visitMap := make(map[string]bool)
	// rNum := len(matrix)
	// if rNum == 0 {
	// 	return res
	// }
	// cNum := len(matrix[0])
	// if cNum == 0 {
	// 	return res
	// }

	// // c,r代表当前坐标
	// c, r := 0, 0
	// // 在每个方向都可能终止
	// for {
	// 	// 当前点一定可取，这个由循环来保证
	// 	res = append(res, matrix[r][c])
	// 	visitMap[getKey(r, c)] = true

	// 	if direct == 'R' {
	// 		if c+1 < cNum && !visitMap[getKey(r, c+1)] {
	// 			c = c + 1
	// 		} else if r + 1 < rNum && !visitMap[getKey(r+1, c)] {
	// 			r = r + 1
	// 			direct = 'D'
	// 		} else {
	// 			break
	// 		}
	// 	} else if direct == 'D' {
	// 		if r + 1 < rNum && !visitMap[getKey(r+1, c)] {
	// 			r = r + 1
	// 		} else if c - 1 >= 0 && !visitMap[getKey(r,c-1)] {
	// 			c = c - 1
	// 			direct = 'L'
	// 		} else {
	// 			break
	// 		}
	// 	} else if direct == 'L' {
	// 		if c - 1 >= 0 && !visitMap[getKey(r,c-1)] {
	// 			c = c - 1
	// 		} else if r - 1 >= 0 && !visitMap[getKey(r-1, c)] {
	// 			r = r - 1
	// 			direct = 'U'
	// 		} else {
	// 			break
	// 		}
	// 	} else if direct == 'U' {
	// 		if r - 1 >= 0 && !visitMap[getKey(r-1, c)] {
	// 			r = r - 1
	// 		} else if c + 1 < cNum && !visitMap[getKey(r,c+1)] {
	// 			c = c + 1
	// 			direct = 'R'
	// 		} else {
	// 			break
	// 		}
	// 	}
	// }

	// return res

	// 2. 设定上下左右边界法
	res := []int{}
	direct := 'R'
	rNum := len(matrix)
	if rNum == 0 {
		return res
	}
	cNum := len(matrix[0])
	if cNum == 0 {
		return res
	}

	r, c := 0, 0
	// 设定上下左右边界
	Rlimit, Dlimit, Llimit, Ulimit := cNum - 1, rNum - 1, 0, 1
	for {
		// 开头推入，合法性由循环保障
		res = append(res, matrix[r][c])

		if direct == 'R' {
			if c + 1 <= Rlimit {
				c++
			} else if r + 1 <= Dlimit {
				r++
				direct = 'D'
				Rlimit--
			} else {
				break
			}
		} else if direct == 'D' {
			if r + 1 <= Dlimit {
				r++
			} else if c - 1 >= Llimit {
				c--
				direct = 'L'
				Dlimit--
			} else {
				break
			}
		} else if direct == 'L' {
			if c - 1 >= Llimit {
				c--
			} else if r - 1 >= Ulimit {
				r--
				direct = 'U'
				Llimit++
			} else {
				break
			}
		} else if direct == 'U' {
			if r - 1 >= Ulimit {
				r--
			} else if c + 1 <= Rlimit {
				c++
				direct = 'R'
				Ulimit++
			} else {
				break
			}
		}
	}
	return res
}

func getKey(i, j int) string {
	return strconv.Itoa(i) + "-" + strconv.Itoa(j)
}
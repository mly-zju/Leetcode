func exist(board [][]byte, word string) bool {
	rNum := len(board)
	if rNum == 0 {
		return false
	}
	cNum := len(board[0])
	if cNum == 0 {
		return false
	}
	
	for i := 0; i < rNum; i++ {
		for j := 0; j < cNum; j++ {
			// 起点必须匹配
			if board[i][j] == word[0] {
				statMap := make(map[string]bool)
				statMap[getKey(i, j)] = true
				if _exist(board, word, i, j, 1, statMap) {
					return true
				}
			}
		}
	}
	return false
}

func _exist(board [][]byte, word string, i, j, wIndex int, statMap map[string]bool) bool {
	if wIndex == len(word) {
		return true
	} else {
		// 上
		if i - 1 >= 0 && !statMap[getKey(i-1, j)] && board[i-1][j] == word[wIndex] {
			statMap[getKey(i-1, j)] = true
			tmpRes := _exist(board, word, i-1, j, wIndex + 1, statMap)
			if tmpRes {
				return true
			}
			statMap[getKey(i-1, j)] = false
		}
		// 下
		if i + 1 < len(board) && !statMap[getKey(i+1, j)] && board[i+1][j] == word[wIndex] {
			statMap[getKey(i+1, j)] = true
			tmpRes := _exist(board, word, i+1, j, wIndex + 1, statMap)
			if tmpRes {
				return true
			}
			statMap[getKey(i+1, j)] = false
		}
		// 左
		if j - 1 >= 0 && !statMap[getKey(i, j-1)] && board[i][j-1] == word[wIndex] {
			statMap[getKey(i, j-1)] = true
			tmpRes := _exist(board, word, i, j-1, wIndex + 1, statMap)
			if tmpRes {
				return true
			}
			statMap[getKey(i, j-1)] = false
		}
		// 右
		if j + 1 < len(board[0]) && !statMap[getKey(i, j+1)] && board[i][j+1] == word[wIndex] {
			statMap[getKey(i, j+1)] = true
			tmpRes := _exist(board, word, i, j+1, wIndex + 1, statMap)
			if tmpRes {
				return true
			}
			statMap[getKey(i, j+1)] = false
		}
		return false
	}
}

func getKey(i, j int) string {
	return strconv.Itoa(i) + "-" + strconv.Itoa(j)
}

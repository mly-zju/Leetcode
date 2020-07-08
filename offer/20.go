// 沙雕题目，不做了
func isNumber(s string) bool {
	// 先去除首尾空格
	s = strings.Trim(s, " ")
	arr := []byte(s)
	length := len(arr)
	dotExist := false
	eExist := false
	numExist := false

	for i, ch := range arr {
		if ch == '+' {
			// +只能出现在最开头
			if i != 0 {
				return false
			}
		} else if ch == '-' {
			// -只能出现在开头，或者e的下一位, 且不能是最后一位
			if i == 0 || (arr[i-1] == 'e' && i != length - 1) {
				continue
			} else {
				return false
			}
		} else if ch == '.' {
			// .只能出现一次, 且要么是第一, 如果非第一位，那么前一位必须是数字或者+/-
			if dotExist {
				return false
			}
			
			if i == 0 || ((arr[i-1] >= '0' && arr[i-1] <= '9') || arr[i-1] == '+' || arr[i-1] == '-') {
				dotExist = true
			} else {
				return false
			}
		} else if ch == 'e' {
			// e只能出现一次且不能在开头或者结尾, 并且e之前一定必须是数字
			if eExist || i == 0 || i == length - 1 || arr[i-1] < '0' || arr[i-1] > '9' {
				return false
			}
			eExist = true
		} else if ch < '0' || ch > '9' {
			// 其余的，不允许出现非数字
			return false
		} else {
			numExist = true
		}
	}
	return numExist
}
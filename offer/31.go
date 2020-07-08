// 模拟题
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	length := len(pushed)
	i, j := 0, 0
	for i < length && j < length {
		newLen := len(stack)
		if newLen > 0 && stack[newLen - 1] == popped[j] {
			// 如果栈顶有元素且等于当前值，出栈
			stack = stack[0:newLen-1]
			j++
		} else if i < length {
			// 否则入栈，直到遇到相等元素
			for i < length && pushed[i] != popped[j] {
				stack = append(stack, pushed[i])
				i++
			}

			if i == length {
				// 直到最后都没遇到，证明错误
				return false
			} else {
				// 遇到了相同的，省略入栈过程，二者都加
				i++
				j++
			}
		}
	}

	// 不可能遇到i没到j到了的情况，要么i和j同时到了，要么i到了j没到
	for j < length {
		newLen := len(stack)
		if stack[newLen - 1] != popped[j] {
			return false
		} else {
			stack = stack[0:newLen-1]
			j++
		}
	}
	return true
}
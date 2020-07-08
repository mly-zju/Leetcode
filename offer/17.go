func printNumbers(n int) []int {
	// 计算最大数据
	maxNum := 0
	for i := 0; i < n; i++ {
		maxNum = maxNum * 10 + 9
	}
	res := make([]int, maxNum)
	for i := 0; i < maxNum; i++ {
		res[i] = i + 1
	}
	return res
}
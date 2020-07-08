func findRepeatNumber(nums []int) int {
	statMap := make(map[int]bool)
	res := -1
	for _, num := range nums {
		if statMap[num] {
			res = num
			break
		}
		statMap[num] = true
	}
	return res
}
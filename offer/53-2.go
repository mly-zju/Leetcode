func missingNumber(nums []int) int {
	length := len(nums)
	sum := float64(1 + length) / 2 * float64(length)
	arrSum := 0
	for _, num := range nums {
		arrSum += num
	}
	return int(sum) - arrSum
}
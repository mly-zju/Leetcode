func singleNumber(nums []int) int {
	sort.Ints(nums)
	var res int
	for i, length := 0, len(nums); i < length; {
		if i + 2 < length && nums[i] == nums[i + 1] && nums[i] == nums[i + 2] {
			i = i + 3
		} else {
			res = nums[i]
			break
		}
	}
	return res
}
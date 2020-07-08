func twoSum(nums []int, target int) []int {
	friendMap := make(map[int]bool)
	for _, num := range nums {
		friend := target - num
		if friendMap[friend] {
			return []int{num, friend}
		}
		friendMap[num] = true
	}
	return []int{}
}
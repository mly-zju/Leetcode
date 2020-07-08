func majorityElement(nums []int) int {
	// 排序后求中间值, 因为不论如何，排序后最中心索引的一定是答案
	sort.Ints(nums)
	return nums[len(nums) / 2]
}
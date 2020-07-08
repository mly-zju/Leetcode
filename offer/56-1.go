func singleNumbers(nums []int) []int {
	// 先整体亦或，取某一位不为0的，然后执行分组，分别亦或即可
	comNum := 0
	for _, num := range nums {
		comNum ^= num
	}
	// 寻找某一位不为0的
	nonZero := 1
	for comNum > 0 {
		tmp := comNum % 2
		if tmp != 0 {
			break
		} else {
			nonZero = nonZero << 1
		}
		comNum = comNum / 2
	}

	num1, num2 := 0, 0
	for _, num := range nums {
		if nonZero & num == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}
	return []int{num1, num2}
}
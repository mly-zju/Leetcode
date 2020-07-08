package main

import (
	"math"
	"strconv"
)

func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}

	// 先计算所在数位的下界
	base := 1
	bitNum := 0
	maxIndex := 0
	minIndex := 0
	for maxIndex < n {
		base *= 10
		bitNum++
		minIndex = maxIndex
		maxIndex += base*bitNum - base/10*bitNum
	}

	// 在自己区域内，查看自己的索引
	n = n - minIndex - 1
	completeN := n / bitNum
	offsetN := n % bitNum

	realNum := int(math.Pow(float64(10), float64(bitNum-1))) + completeN
	realStr := strconv.Itoa(realNum)
	return int(realStr[offsetN] - '0')
}

// func main() {
// 	fmt.Println(findNthDigit(15))
// }

package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	// 1. 暴力法
	// maxIndex := (target + 1) / 2
	// res := [][]int{}
	// for i := 1; i <= maxIndex; i++ {
	// 	tmpSum := 0
	// 	tmpRes := []int{}
	// 	for j := i; j <= maxIndex; j++ {
	// 		tmpSum += j
	// 		tmpRes = append(tmpRes, j)
	// 		if tmpSum == target {
	// 			res = append(res, tmpRes)
	// 			break
	// 		} else if tmpSum > target {
	// 			break
	// 		}
	// 	}
	// }
	// return res

	// 2. 滑动窗口法
	l, r := 1, 1
	res := [][]int{}
	maxIndex := (target + 1) / 2
	curSum := 1
	for l <= maxIndex && r <= maxIndex {
		if curSum == target {
			tmpRes := []int{}
			for i := l; i <= r; i++ {
				tmpRes = append(tmpRes, i)
			}
			res = append(res, tmpRes)
			curSum -= l
			l++
			r++
			curSum += r
		} else if curSum < target {
			r++
			curSum += r
		} else if curSum > target {
			curSum -= l
			l++
		}
	}
	return res
}

func main() {
	fmt.Println(findContinuousSequence(9))
}

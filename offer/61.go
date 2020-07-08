package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isStraight([]int{11, 0, 9, 0, 0}))
}

func isStraight(nums []int) bool {
	joker := 0
	// 先排序
	sort.Ints(nums)
	// 统计数组中0的个数
	for _, num := range nums {
		if num == 0 {
			joker++
		}
	}

	nextNum := -1
	for _, num := range nums {
		fmt.Println(num, nextNum)
		if num == 0 {
			continue
		}

		if nextNum == num {
			nextNum++
		} else if nextNum != num {
			if nextNum == -1 {
				nextNum = num + 1
			} else if joker > 0 {
				for joker > 0 {
					joker--
					nextNum++
					if nextNum == num {
						break
					}
				}
				if nextNum == num {
					nextNum++
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}

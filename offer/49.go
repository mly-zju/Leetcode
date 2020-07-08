package main

import "fmt"

func nthUglyNumber(n int) int {
	// 三指针法：f2, f3, f5分别只和2、3、5相乘
	if n == 0 {
		return 0
	}
	res := []int{1}
	count := 1
	f2, f3, f5 := 0, 0, 0
	for count < n {
		n2 := res[f2] * 2
		n3 := res[f3] * 3
		n5 := res[f5] * 5
		minIndex := getMin(n2, n3, n5)
		if minIndex == 1 {
			if n2 != res[count-1] {
				res = append(res, n2)
				count++
			}
			f2++
		} else if minIndex == 2 {
			if n3 != res[count-1] {
				res = append(res, n3)
				count++
			}
			f3++
		} else if minIndex == 3 {
			if n5 != res[count-1] {
				res = append(res, n5)
				count++
			}
			f5++
		}
	}
	return res[count-1]
}

func getMin(a, b, c int) int {
	if a <= b && a <= c {
		return 1
	} else if b <= a && b <= c {
		return 2
	} else {
		return 3
	}
}

func main() {
	fmt.Println(nthUglyNumber(10))
}

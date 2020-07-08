package main

func minimumSwap(s1 string, s2 string) int {
	// 先查看词频是否相同
	len1, len2 := len(s1), len(s2)
	if len1 != len2 {
		return -1
	}
	// 一次遍历，统计各自的x/y词频，以及新的不同的单词(即忽略相等的位置)
	x1, y1, x2, y2 := 0, 0, 0, 0
	diffNum, xdiffNum := 0, 0
	for i := 0; i < len1; i++ {
		if s1[i] == 'x' {
			x1++
		} else {
			y1++
		}
		if s2[i] == 'x' {
			x2++
		} else {
			y2++
		}

		if s1[i] != s2[i] {
			diffNum++
			if s1[i] == 'x' {
				xdiffNum++
			}
		}
	}

	// x和y的总数都必须是偶数
	if (x1+x2)%2 != 0 || (y1+y2)%2 != 0 {
		return -1
	}

	ydiffNum := diffNum - xdiffNum
	evenNum := xdiffNum/2 + ydiffNum/2
	oddNum := xdiffNum%2 + ydiffNum%2

	return evenNum + oddNum
}

// func main() {
// 	// s1 := "xxyyxyxyxx"
// 	// s2 := "xyyxyxxxyx"
// 	s1 := "xxyx"
// 	s2 := "yyxy"
// 	fmt.Println(minimumSwap(s1, s2))
// }

package main

import "fmt"

func balancedString(s string) int {
	// 统计
	statMap := map[byte]int{
		'Q': 0, 'W': 0, 'E': 0, 'R': 0,
	}
	length := len(s)
	n := length / 4

	for i := 0; i < length; i++ {
		statMap[s[i]]++
	}

	// 寻找是否都相等，不等的话寻找需要找到的子串
	gapMap := map[byte]int{
		'Q': 0, 'W': 0, 'E': 0, 'R': 0,
	}
	gapNum := 0
	for key, val := range statMap {
		for val > n {
			gapMap[key]++
			val--
			gapNum++
		}
	}
	if gapNum == 0 {
		return 0
	}

	// 构建分段统计数组用于加速
	// statArr := [](map[byte]int){}
	// for i := 0; i < length; i++ {
	// 	tmpMap := map[byte]int{
	// 		'Q': 0, 'W': 0, 'E': 0, 'R': 0,
	// 	}
	// 	for j := i; j >= 0; j-- {
	// 		tmpMap[s[j]]++
	// 	}
	// 	statArr = append(statArr, tmpMap)
	// }

	// 接下来走窗口滑动
	minSize := length
	l, r := 0, 0
	statMap = map[byte]int{
		s[0]: 1,
	}
	for l <= r && r < length {
		if isContain(statMap, gapMap) {
			if r-l+1 < minSize {
				minSize = r - l + 1
			}
			statMap[s[l]]--
			l++
		} else {
			r++
			if r < length {
				statMap[s[r]]++
			}
		}
	}
	return minSize
}

func isContain(statMap, gapMap map[byte]int) bool {
	if statMap['Q'] >= gapMap['Q'] && statMap['W'] >= gapMap['W'] && statMap['E'] >= gapMap['E'] && statMap['R'] >= gapMap['R'] {
		return true
	}
	return false
}

func main() {
	// a := "WWEQERQWQWWRWWERQWEQ"
	// a := "QQER"
	// a := "WQWRQQQW"
	fmt.Println(balancedString(a))
}

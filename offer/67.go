package main

import (
	"fmt"
	"math"
)

func strToInt(str string) int {
	var res int64
	maxInt := int64(math.Pow(2, 31) - 1)
	numFlag := false
	minusFlag := false
	for i, length := 0, len(str); i < length; i++ {
		if !isNum(str[i]) {
			// 不是数字的，有几种情况：空格，正负号，字母
			if str[i] == ' ' {
				if !numFlag {
					continue
				} else {
					break
				}
			} else if str[i] == '+' {
				if !numFlag {
					numFlag = true
				} else {
					break
				}
			} else if str[i] == '-' {
				if !numFlag {
					numFlag = true
					minusFlag = true
				} else {
					break
				}
			} else {
				// 否则如果是小数点，跳出
				break
			}
		} else {
			res = res*10 + int64(str[i]-'0')
			numFlag = true
			if minusFlag && res >= maxInt+1 {
				res = maxInt + 1
				break
			} else if res >= maxInt {
				res = maxInt
				break
			}
		}
	}
	if minusFlag {
		return int(-res)
	}
	return int(res)
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {
	fmt.Println(strToInt("-0012a42"))
}

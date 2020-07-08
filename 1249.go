package main

import (
	"fmt"
)

func minRemoveToMakeValid(s string) string {
	stack := Mystack{}
	// 需要忽略的index
	igArr := []int{}

	res := ""
	length := len(s)
	for i := 0; i < length; i++ {
		if s[i] == '(' {
			stack.push(i)
		} else if s[i] == ')' {
			if !stack.isEmpty() {
				stack.pop()
			} else {
				igArr = append(igArr, i)
			}
		}
	}

	// 继续检测stack
	if stack.isEmpty() == false {
		igArr = append(igArr, stack...)
	}

	for i := 0; i < length; i++ {
		if len(igArr) > 0 && i == igArr[0] {
			igArr = igArr[1:]
		} else {
			res += string(s[i])
		}
	}

	return res
}

func main() {
	a := "lee(t(c)o)de)"
	fmt.Println(minRemoveToMakeValid(a))
}

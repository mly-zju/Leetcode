package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	// 先按照字符串排序
	sort.Slice(folder, func(i, j int) bool {
		return folder[i] < folder[j]
	})

	res := []string{}
	length := len(folder)
	for i := 0; i < length; i++ {
		rLen := len(res)
		if rLen == 0 {
			res = append(res, folder[i])
		} else {
			fLen := len(res[rLen-1])
			cLen := len(folder[i])
			if cLen > fLen && (strings.Compare(folder[i][0:fLen], res[rLen-1]) == 0) && folder[i][fLen] == '/' {
				continue
			} else {
				res = append(res, folder[i])
			}
		}
	}
	return res
}

func isChild(c, f string) bool {
	cArr := strings.Split(c, "/")
	fArr := strings.Split(f, "/")

	// 首先子文件夹层级必须大于父级
	if len(cArr) <= len(fArr) {
		return false
	}

	for index, val := range fArr {
		if val != cArr[index] {
			return false
		}
	}
	return true
}

func main() {
	a := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	fmt.Println(removeSubfolders(a))
}

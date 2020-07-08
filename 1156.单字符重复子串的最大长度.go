package main

import "fmt"

func maxRepOpt1(text string) int {
	type StrInfo struct {
		ch  byte
		num int
	}

	statArr := []StrInfo{}

	prevC := text[0]
	count := 0
	max := 0
	for i, length := 0, len(text); i < length; i++ {
		if text[i] == prevC {
			count++
		} else {
			statArr = append(statArr, StrInfo{
				ch:  prevC,
				num: count,
			})
			// 这个过程中先优先统计单独的最长子串
			if max < count {
				max = count
			}
			prevC = text[i]
			count = 1
		}
	}
	// 最后结束后还要检测一下
	statArr = append(statArr, StrInfo{
		ch:  prevC,
		num: count,
	})
	if max < count {
		max = count
	}

	// 然后循环检测statArr，看两段连接的情况下最大值
	length := len(statArr)
	for i := 0; i < length-2; i++ {
		for j := i + 2; j < length; j++ {
			if statArr[i].ch == statArr[j].ch {
				if j == i+2 && statArr[i+1].num == 1 {
					// 相隔一段的情况下如果满足条件可以全量连接
					var tmpMax int
					tmpMax = statArr[i].num + statArr[j].num
					// 还需要继续寻找有没有另一个不等于j并且ch相同的
					for k := 0; k < length; k++ {
						if k >= i && k <= j {
							continue
						}
						if statArr[k].ch == statArr[j].ch {
							tmpMax++
							break
						}
					}

					if max < tmpMax {
						max = tmpMax
					}
				} else {
					// 否则只能增加一
					var tmpMax int
					if statArr[i].num > statArr[j].num {
						tmpMax = statArr[i].num + 1
					} else {
						tmpMax = statArr[j].num + 1
					}
					if max < tmpMax {
						max = tmpMax
					}
				}
			}
		}
	}
	return max
}

func main() {
	a := "bbababaaaa"
	// 6 is right
	fmt.Println(maxRepOpt1(a))
}

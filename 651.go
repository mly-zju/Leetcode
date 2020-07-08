// 四键键盘问题，leetcode 651
// 这个不是简单的状态机，因为彼此之间会依赖
package main

import (
	"fmt"
	"strconv"
)

func getMaxA(N int) int {
	if N == 0 {
		return 0
	}

	// 方法一：基于复杂状态 + 记忆化递归
	// 定义状态：getDp(n, num, cacheNum), n代表当前是还剩几次操作，num代表当前屏幕字数，cacheNum代表缓冲字数
	// 定义备忘录
	// memo := map[string]int{}
	// return getDp(N, 0, 0, memo)

	// 方法二：基于以下一个事实
}

func getDp(n, num, cacheNum int, memo map[string]int) int {
	if n <= 0 {
		return num
	}

	key := getKey(n, num, cacheNum)
	if val, ok := memo[key]; ok {
		return val
	}

	// 1. 当前只做A操作
	val1 := getDp(n-1, num+1, cacheNum, memo)
	// 2. 当前做CV操作
	val2 := getDp(n-1, num+cacheNum, cacheNum, memo)
	// 3. 当前做了CA + CC操作, 那么n-2的时候，缓冲区有num
	val3 := getDp(n-2, num, num, memo)

	memo[key] = getMax(val1, getMax(val2, val3))
	return memo[key]
}

func getKey(n, num, cacheNum int) string {
	// 获取唯一标识
	return strconv.Itoa(n) + "-" + strconv.Itoa(num) + "-" + strconv.Itoa(cacheNum)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(getMaxA(7))
}

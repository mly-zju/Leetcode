package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 *
 * https://leetcode-cn.com/problems/ugly-number-ii/description/
 *
 * algorithms
 * Medium (47.55%)
 * Likes:    164
 * Dislikes: 0
 * Total Accepted:    11.8K
 * Total Submissions: 24.5K
 * Testcase Example:  '10'
 *
 * 编写一个程序，找出第 n 个丑数。
 *
 * 丑数就是只包含质因数 2, 3, 5 的正整数。
 *
 * 示例:
 *
 * 输入: n = 10
 * 输出: 12
 * 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
 *
 * 说明:
 *
 *
 * 1 是丑数。
 * n 不超过1690。
 *
 *
 */

// @lc code=start
func nthUglyNumber(n int) int {
	cache := []int{1}
	param := 1
	existMap := map[int]bool{}
	loopArr := []int{2, 3, 5}
	for i := 1; i < 2*n; {
		loopNum := 0
		for _, val := range loopArr {
			tmp := param * val
			if existMap[tmp] == false {
				cache = append(cache, tmp)
				existMap[tmp] = true
				loopNum++
			}
		}
		i += loopNum
		param++
	}
	sort.Ints(cache)
	fmt.Println(cache)
	return cache[n-1]
}

// @lc code=end

package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=464 lang=golang
 *
 * [464] 我能赢吗
 *
 * https://leetcode-cn.com/problems/can-i-win/description/
 *
 * algorithms
 * Medium (32.43%)
 * Likes:    67
 * Dislikes: 0
 * Total Accepted:    2.1K
 * Total Submissions: 6.4K
 * Testcase Example:  '10\n11'
 *
 * 在 "100 game" 这个游戏中，两名玩家轮流选择从 1 到 10 的任意整数，累计整数和，先使得累计整数和达到 100 的玩家，即为胜者。
 *
 * 如果我们将游戏规则改为 “玩家不能重复使用整数” 呢？
 *
 * 例如，两个玩家可以轮流从公共整数池中抽取从 1 到 15 的整数（不放回），直到累计整数和 >= 100。
 *
 * 给定一个整数 maxChoosableInteger （整数池中可选择的最大数）和另一个整数
 * desiredTotal（累计和），判断先出手的玩家是否能稳赢（假设两位玩家游戏时都表现最佳）？
 *
 * 你可以假设 maxChoosableInteger 不会大于 20， desiredTotal 不会大于 300。
 *
 * 示例：
 *
 * 输入：
 * maxChoosableInteger = 10
 * desiredTotal = 11
 *
 * 输出：
 * false
 *
 * 解释：
 * 无论第一个玩家选择哪个整数，他都会失败。
 * 第一个玩家可以选择从 1 到 10 的整数。
 * 如果第一个玩家选择 1，那么第二个玩家只能选择从 2 到 10 的整数。
 * 第二个玩家可以通过选择整数 10（那么累积和为 11 >= desiredTotal），从而取得胜利.
 * 同样地，第一个玩家选择任意其他整数，第二个玩家都会赢。
 *
 *
 */

// @lc code=start
var resultCache map[string]int

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal == 0 {
		return true
	}

	// 初始化全局缓存和全局可选数字
	resultCache = map[string]int{}
	// 初始化可选数字map
	availArr := []bool{}
	sum := 0
	for i := 0; i <= maxChoosableInteger; i++ {
		sum += i
		availArr = append(availArr, true)
	}
	if sum < desiredTotal {
		return false
	}
	return _canIWin(availArr, desiredTotal)
}

func _canIWin(availArr []bool, desiredTotal int) bool {
	if desiredTotal <= 0 {
		return true
	}

	key := getHashKey(availArr, desiredTotal)
	if resultCache[key] == -1 {
		return false
	} else if resultCache[key] == 1 {
		return true
	}

	// 如果直接已经比较大了，直接返回true
	ans := false
	for i, length := 1, len(availArr); i < length; i++ {
		if availArr[i] == true {
			if i >= desiredTotal {
				ans = true
				break
			} else {
				// 临时标记不可用
				availArr[i] = false
				if _canIWin(availArr, desiredTotal-i) == false {
					ans = true
				}
				availArr[i] = true
				if ans == true {
					break
				}
			}
		}
	}

	if ans == true {
		resultCache[key] = 1
	} else {
		resultCache[key] = -1
	}
	return ans
}

func getHashKey(availArr []bool, desiredTotal int) string {
	key := ""
	for i, length := 1, len(availArr); i < length; i++ {
		if availArr[i] == true {
			key = key + "-" + strconv.Itoa(i)
		}
	}
	return key + "_" + strconv.Itoa(desiredTotal)
}

// func main() {
// 	fmt.Println(canIWin(5, 50))
// }

// @lc code=end

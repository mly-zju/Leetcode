/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 *
 * https://leetcode-cn.com/problems/walking-robot-simulation/description/
 *
 * algorithms
 * Easy (30.09%)
 * Likes:    53
 * Dislikes: 0
 * Total Accepted:    3.8K
 * Total Submissions: 12.7K
 * Testcase Example:  '[4,-1,3]\n[]'
 *
 * 机器人在一个无限大小的网格上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令：
 *
 *
 * -2：向左转 90 度
 * -1：向右转 90 度
 * 1 <= x <= 9：向前移动 x 个单位长度
 *
 *
 * 在网格上有一些格子被视为障碍物。
 *
 * 第 i 个障碍物位于网格点  (obstacles[i][0], obstacles[i][1])
 *
 * 如果机器人试图走到障碍物上方，那么它将停留在障碍物的前一个网格方块上，但仍然可以继续该路线的其余部分。
 *
 * 返回从原点到机器人的最大欧式距离的平方。
 *
 *
 *
 * 示例 1：
 *
 * 输入: commands = [4,-1,3], obstacles = []
 * 输出: 25
 * 解释: 机器人将会到达 (3, 4)
 *
 *
 * 示例 2：
 *
 * 输入: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
 * 输出: 65
 * 解释: 机器人在左转走到 (1, 8) 之前将被困在 (1, 4) 处
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= commands.length <= 10000
 * 0 <= obstacles.length <= 10000
 * -30000 <= obstacle[i][0] <= 30000
 * -30000 <= obstacle[i][1] <= 30000
 * 答案保证小于 2 ^ 31
 *
 *
 */
package main

import (
	"strconv"
)

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	x, y := 0, 0
	// direct 0~3, 代表上、右、下、左
	direct := 0
	// unitGap代表不同direct时候的移动单位
	unitGap := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// 将障碍存入哈希表, 搜索时候快一点
	obsMap := make(map[string]bool)
	for _, pos := range obstacles {
		obsMap[getKey(pos[0], pos[1])] = true
	}
	ans := 0

	for _, instruct := range commands {
		if instruct < 0 {
			direct = getDirect(direct, instruct)
		} else {
			// 开始按照指示一步一步走(模拟机器人), 每走一步探测是否有障碍物
			for i := 0; i < instruct; i++ {
				nx := x + unitGap[direct][0]
				ny := y + unitGap[direct][1]
				// 查看是否是障碍物, 遇到了就返回
				if obsMap[getKey(nx, ny)] {
					break
				} else {
					x = nx
					y = ny
				}
			}
		}
		tmp := x*x + y*y
		if tmp > ans {
			ans = tmp
		}
	}

	return ans
}

func getDirect(direct, instruct int) int {
	if instruct == -1 {
		direct = (direct + 1) % 4
	} else if instruct == -2 {
		direct = (direct + 3) % 4
	}

	return direct
}

func getKey(x, y int) string {
	return strconv.Itoa(x) + ":" + strconv.Itoa(y)
}

// func main() {
// 	commands := []int{-2, 8, 3, 7, -1}
// 	obs := [][]int{{-4, -1}, {1, -1}, {1, 4}, {5, 0}, {4, 5}, {-2, -1}, {2, -5}, {5, 1}, {-3, -1}, {5, -3}}
// 	t := robotSim(commands, obs)
// 	fmt.Println(t)
// }

// @lc code=end

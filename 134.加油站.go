package main

/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 *
 * https://leetcode-cn.com/problems/gas-station/description/
 *
 * algorithms
 * Medium (48.88%)
 * Likes:    174
 * Dislikes: 0
 * Total Accepted:    16.4K
 * Total Submissions: 33K
 * Testcase Example:  '[1,2,3,4,5]\n[3,4,5,1,2]'
 *
 * 在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
 *
 * 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
 *
 * 如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
 *
 * 说明:
 *
 *
 * 如果题目有解，该答案即为唯一答案。
 * 输入数组均为非空数组，且长度相同。
 * 输入数组中的元素均为非负数。
 *
 *
 * 示例 1:
 *
 * 输入:
 * gas  = [1,2,3,4,5]
 * cost = [3,4,5,1,2]
 *
 * 输出: 3
 *
 * 解释:
 * 从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
 * 开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
 * 开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
 * 开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
 * 开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
 * 开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
 * 因此，3 可为起始索引。
 *
 * 示例 2:
 *
 * 输入:
 * gas  = [2,3,4]
 * cost = [3,4,3]
 *
 * 输出: -1
 *
 * 解释:
 * 你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
 * 我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
 * 开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
 * 开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
 * 你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
 * 因此，无论怎样，你都不可能绕环路行驶一周。
 *
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	// 1. 扫描一圈，获取可以的起点地址. 复杂度差不多是O(n^2)
	// length := len(gas)
	// for i := 0; i < length; i++ {
	// 	if gas[i] >= cost[i] {
	// 		if checkRound(gas, cost, i) {
	// 			return i
	// 		}
	// 	}
	// }
	// return -1

	// 2. O(n)的算法：原理就是油箱有剩余油的情况下都无法到达，那么这个也就没必要再作为起点
	i := 0
	for true {
		tmp, tmpIndex := checkRound(gas, cost, i)
		if tmp {
			return tmpIndex
		} else {
			// 如果已经绕了一圈，说明没有找到答案
			if tmpIndex < i {
				return -1
			}
			i = tmpIndex + 1
		}
	}
	return -1
}

// 改造checkRound，返回的是成功或者失败以及成功或者失败的index
func checkRound(gas []int, cost []int, start int) (bool, int) {
	oil := 0
	length := len(gas)
	res := true
	resIndex := start
	for i := start; i < length; i++ {
		oil += gas[i]
		if oil >= cost[i] {
			oil -= cost[i]
		} else {
			res = false
			resIndex = i
			break
		}
	}

	if res == false {
		return res, resIndex
	}

	// 接着从头开始
	for i := 0; i < start; i++ {
		oil += gas[i]
		if oil >= cost[i] {
			oil -= cost[i]
		} else {
			res = false
			resIndex = i
			break
		}
	}
	return res, resIndex
}

// func main() {
// 	gas := []int{1, 2, 3, 4, 5}
// 	cost := []int{3, 4, 5, 1, 2}
// 	fmt.Println(canCompleteCircuit(gas, cost))
// }

// @lc code=end

/*
 * @lc app=leetcode.cn id=1029 lang=golang
 *
 * [1029] 两地调度
 *
 * https://leetcode-cn.com/problems/two-city-scheduling/description/
 *
 * algorithms
 * Easy (54.86%)
 * Likes:    34
 * Dislikes: 0
 * Total Accepted:    3.8K
 * Total Submissions: 6.9K
 * Testcase Example:  '[[10,20],[30,200],[400,50],[30,20]]'
 *
 * 公司计划面试 2N 人。第 i 人飞往 A 市的费用为 costs[i][0]，飞往 B 市的费用为 costs[i][1]。
 *
 * 返回将每个人都飞到某座城市的最低费用，要求每个城市都有 N 人抵达。
 *
 *
 *
 * 示例：
 *
 * 输入：[[10,20],[30,200],[400,50],[30,20]]
 * 输出：110
 * 解释：
 * 第一个人去 A 市，费用为 10。
 * 第二个人去 A 市，费用为 30。
 * 第三个人去 B 市，费用为 50。
 * 第四个人去 B 市，费用为 20。
 *
 * 最低总费用为 10 + 30 + 50 + 20 = 110，每个城市都有一半的人在面试。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= costs.length <= 100
 * costs.length 为偶数
 * 1 <= costs[i][0], costs[i][1] <= 1000
 *
 *
 */
package main

import (
	"sort"
)

// @lc code=start
func twoCitySchedCost(costs [][]int) int {
	// 假设全部飞往A市，总价X。然后再选N个飞往B市，则增加的总价市priceB - priceA
	// 为了让增加的总价最少，选择N个priceB-priceA最小的值(可以是负数)
	// 也就是按照priceB-pricrA做排序，取前N个去B
	sort.Slice(costs, func(i, j int) bool {
		aSum := costs[i][1] - costs[i][0]
		bSum := costs[j][1] - costs[j][0]
		return aSum < bSum
	})
	N := len(costs) / 2
	sum := 0
	for i := 0; i < 2*N; i++ {
		if i < N {
			sum += costs[i][1]
		} else {
			sum += costs[i][0]
		}
	}
	return sum
}

// func main() {
// 	costs := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}
// 	sum := twoCitySchedCost(costs)
// 	fmt.Println(sum)
// }

// @lc code=end

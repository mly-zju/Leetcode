package main

import (
	"sort"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1169 lang=golang
 *
 * [1169] 查询无效交易
 *
 * https://leetcode-cn.com/problems/invalid-transactions/description/
 *
 * algorithms
 * Easy (24.58%)
 * Likes:    8
 * Dislikes: 0
 * Total Accepted:    2K
 * Total Submissions: 7.4K
 * Testcase Example:  '["alice,20,800,mtv","alice,50,100,beijing"]'
 *
 * 如果出现下述两种情况，交易 可能无效：
 *
 *
 * 交易金额超过 ¥1000
 * 或者，它和另一个城市中同名的另一笔交易相隔不超过 60 分钟（包含 60 分钟整）
 *
 *
 * 每个交易字符串 transactions[i] 由一些用逗号分隔的值组成，这些值分别表示交易的名称，时间（以分钟计），金额以及城市。
 *
 * 给你一份交易清单 transactions，返回可能无效的交易列表。你可以按任何顺序返回答案。
 *
 *
 *
 * 示例 1：
 *
 * 输入：transactions = ["alice,20,800,mtv","alice,50,100,beijing"]
 * 输出：["alice,20,800,mtv","alice,50,100,beijing"]
 * 解释：第一笔交易是无效的，因为第二笔交易和它间隔不超过 60 分钟、名称相同且发生在不同的城市。同样，第二笔交易也是无效的。
 *
 * 示例 2：
 *
 * 输入：transactions = ["alice,20,800,mtv","alice,50,1200,mtv"]
 * 输出：["alice,50,1200,mtv"]
 *
 *
 * 示例 3：
 *
 * 输入：transactions = ["alice,20,800,mtv","bob,50,1200,mtv"]
 * 输出：["bob,50,1200,mtv"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * transactions.length <= 1000
 * 每笔交易 transactions[i] 按 "{name},{time},{amount},{city}" 的格式进行记录
 * 每个交易名称 {name} 和城市 {city} 都由小写英文字母组成，长度在 1 到 10 之间
 * 每个交易时间 {time} 由一些数字组成，表示一个 0 到 1000 之间的整数
 * 每笔交易金额 {amount} 由一些数字组成，表示一个 0 到 2000 之间的整数
 *
 *
 */

// @lc code=start
type Order struct {
	name  string
	time  int
	money int
	city  string
}

func invalidTransactions(transactions []string) []string {
	orderList := []Order{}
	// 解析
	for _, val := range transactions {
		tmpArr := strings.Split(val, ",")
		time, _ := strconv.Atoi(tmpArr[1])
		money, _ := strconv.Atoi(tmpArr[2])
		orderList = append(orderList, Order{
			name:  tmpArr[0],
			time:  time,
			money: money,
			city:  tmpArr[3],
		})
	}
	// 按照time排序
	sort.Slice(orderList, func(i, j int) bool {
		return orderList[i].time < orderList[j].time
	})

	res := []string{}
	// 还要注意去重
	addMap := make(map[string]bool)
	for index, order := range orderList {
		// 如果不大于1000，从过往元素查询不同城市不超60分钟的, 由于排序，第一条如果不超，那么剩下的都不超
		for i := index - 1; i >= 0; i-- {
			if orderList[i].name == order.name && orderList[i].city != order.city {
				if order.time-orderList[i].time <= 60 {
					orderStr := getString(order)
					if addMap[orderStr] == false {
						res = append(res, orderStr)
						addMap[orderStr] = true
					}
					orderStr = getString(orderList[i])
					if addMap[orderStr] == false {
						res = append(res, orderStr)
						addMap[orderStr] = true
					}
				} else {
					break
				}
			}
		}
		if order.money > 1000 {
			orderStr := getString(order)
			if addMap[orderStr] == false {
				res = append(res, orderStr)
				addMap[orderStr] = true
			}
		}
	}

	return res
}

func getString(order Order) string {
	timeStr := strconv.Itoa(order.time)
	moneyStr := strconv.Itoa(order.money)
	return order.name + "," + timeStr + "," + moneyStr + "," + order.city
}

// func main() {
// 	a := []string{"bob,689,1910,barcelona", "alex,696,122,bangkok", "bob,832,1726,barcelona", "bob,820,596,bangkok", "chalicefy,217,669,barcelona", "bob,175,221,amsterdam"}
// 	fmt.Println(invalidTransactions(a))
// }

// @lc code=end

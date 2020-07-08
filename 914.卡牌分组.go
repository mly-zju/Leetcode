/*
 * @lc app=leetcode.cn id=914 lang=golang
 *
 * [914] 卡牌分组
 *
 * https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/description/
 *
 * algorithms
 * Easy (31.50%)
 * Likes:    61
 * Dislikes: 0
 * Total Accepted:    6.3K
 * Total Submissions: 19.8K
 * Testcase Example:  '[1,2,3,4,4,3,2,1]'
 *
 * 给定一副牌，每张牌上都写着一个整数。
 *
 * 此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：
 *
 *
 * 每组都有 X 张牌。
 * 组内所有的牌上都写着相同的整数。
 *
 *
 * 仅当你可选的 X >= 2 时返回 true。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[1,2,3,4,4,3,2,1]
 * 输出：true
 * 解释：可行的分组是 [1,1]，[2,2]，[3,3]，[4,4]
 *
 *
 * 示例 2：
 *
 * 输入：[1,1,1,2,2,2,3,3]
 * 输出：false
 * 解释：没有满足要求的分组。
 *
 *
 * 示例 3：
 *
 * 输入：[1]
 * 输出：false
 * 解释：没有满足要求的分组。
 *
 *
 * 示例 4：
 *
 * 输入：[1,1]
 * 输出：true
 * 解释：可行的分组是 [1,1]
 *
 *
 * 示例 5：
 *
 * 输入：[1,1,2,2,2,2]
 * 输出：true
 * 解释：可行的分组是 [1,1]，[2,2]，[2,2]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= deck.length <= 10000
 * 0 <= deck[i] < 10000
 *
 *
 *
 *
 */

// @lc code=start
func hasGroupsSizeX(deck []int) bool {
	// 统计不同数字频率，求频率数组的最大公约数
	arrMap := [10001]int{}
	for _, val := range deck {
		arrMap[val]++
	}
	statArr := []int{}
	for _, val := range arrMap {
		if val > 0 {
			statArr = append(statArr, val)
		}
	}
	// 求最大公约数
	gcdNum := statArr[0]
	for i, length := 1, len(statArr); i < length; i++ {
		gcdNum = gcd(gcdNum, statArr[i])
	}
	return gcdNum > 1
}

func gcd(a, b int) int {
	// 更相减损法
	if a == b {
		return a
	} else if a > b {
		return gcd(b, a-b)
	} else {
		return gcd(a, b-a)
	}
}

// @lc code=end


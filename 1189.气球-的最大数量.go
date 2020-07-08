/*
 * @lc app=leetcode.cn id=1189 lang=golang
 *
 * [1189] “气球” 的最大数量
 *
 * https://leetcode-cn.com/problems/maximum-number-of-balloons/description/
 *
 * algorithms
 * Easy (61.83%)
 * Likes:    10
 * Dislikes: 0
 * Total Accepted:    3.8K
 * Total Submissions: 6.2K
 * Testcase Example:  '"nlaebolko"'
 *
 * 给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。
 *
 * 字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：text = "nlaebolko"
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：text = "loonbalxballpoon"
 * 输出：2
 *
 *
 * 示例 3：
 *
 * 输入：text = "leetcode"
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= text.length <= 10^4
 * text 全部由小写英文字母组成
 *
 *
 */

// @lc code=start
func maxNumberOfBalloons(text string) int {
	// balloon需要1个b，1个a，2个l，2个o，1个n。统计text中这几个字母的频次，找出最小倍数
	mapNeed := map[rune]int{
		'b': 1, 'a': 1, 'l': 2, 'o': 2, 'n': 1,
	}
	mapReal := make(map[rune]int)

	// 执行统计
	for _, ch := range text {
		mapReal[ch]++
	}
	minNum := len(text)
	for key, val := range mapNeed {
		tmp := mapReal[key] / val
		if tmp < minNum {
			minNum = tmp
		}
	}
	return minNum
}

// @lc code=end


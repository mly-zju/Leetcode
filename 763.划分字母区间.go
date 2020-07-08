package main

/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 *
 * https://leetcode-cn.com/problems/partition-labels/description/
 *
 * algorithms
 * Medium (66.94%)
 * Likes:    81
 * Dislikes: 0
 * Total Accepted:    5.9K
 * Total Submissions: 8.5K
 * Testcase Example:  '"ababcbacadefegdehijhklij"'
 *
 * 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一个字母只会出现在其中的一个片段。返回一个表示每个字符串片段的长度的列表。
 *
 * 示例 1:
 *
 * 输入: S = "ababcbacadefegdehijhklij"
 * 输出: [9,7,8]
 * 解释:
 * 划分结果为 "ababcbaca", "defegde", "hijhklij"。
 * 每个字母最多出现在一个片段中。
 * 像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
 *
 *
 * 注意:
 *
 *
 * S的长度在[1, 500]之间。
 * S只包含小写字母'a'到'z'。
 *
 *
 */

// @lc code=start
func partitionLabels(S string) []int {
	// 不是简单的区间调度：区间调度只需要找到重叠的区间最小个数；
	// 而此题，则是要找出彼此完全不重叠的区间边界划分分组
	// 思路：每次只有遇到区间内所有元素在往后区间都没有再出现过的情况，才可以独立分组,
	// 因此可以先执行一次统计，统计每个元素最后一次出现的位置。然后第二次扫描时候来检测。

	// 最后一次出现位置的统计
	statArr := [26]int{}
	for index, ch := range S {
		statArr[ch-'a'] = index
	}

	// 开始执行第二次扫描
	// startIndex代表本次分组的开始位置
	startIndex := 0
	res := []int{}
	for i, length := 0, len(S); i < length; i++ {
		// 是否可以独立分组的标识
		flag := true
		for j := startIndex; j <= i; j++ {
			// 查看当前这个字母的最后元素出现位置是否在本区间
			if statArr[S[j]-'a'] > i {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, i-startIndex+1)
			startIndex = i + 1
		}
	}
	return res
}

// func main() {
// 	s := "ababcbacadefegdehijhklij"
// 	// s := "eaaaabaaec"
// 	// s := "dccccbaabe"
// 	fmt.Println(partitionLabels(s))
// }

// @lc code=end

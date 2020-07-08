/*
 * @lc app=leetcode.cn id=925 lang=golang
 *
 * [925] 长按键入
 *
 * https://leetcode-cn.com/problems/long-pressed-name/description/
 *
 * algorithms
 * Easy (43.22%)
 * Likes:    40
 * Dislikes: 0
 * Total Accepted:    4.7K
 * Total Submissions: 10.7K
 * Testcase Example:  '"alex"\n"aaleex"'
 *
 * 你的朋友正在使用键盘输入他的名字 name。偶尔，在键入字符 c 时，按键可能会被长按，而字符可能被输入 1 次或多次。
 *
 * 你将会检查键盘输入的字符 typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回 True。
 *
 *
 *
 * 示例 1：
 *
 * 输入：name = "alex", typed = "aaleex"
 * 输出：true
 * 解释：'alex' 中的 'a' 和 'e' 被长按。
 *
 *
 * 示例 2：
 *
 * 输入：name = "saeed", typed = "ssaaedd"
 * 输出：false
 * 解释：'e' 一定需要被键入两次，但在 typed 的输出中不是这样。
 *
 *
 * 示例 3：
 *
 * 输入：name = "leelee", typed = "lleeelee"
 * 输出：true
 *
 *
 * 示例 4：
 *
 * 输入：name = "laiden", typed = "laiden"
 * 输出：true
 * 解释：长按名字中的字符并不是必要的。
 *
 *
 *
 *
 * 提示：
 *
 *
 * name.length <= 1000
 * typed.length <= 1000
 * name 和 typed 的字符都是小写字母。
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
func isLongPressedName(name string, typed string) bool {
	// 双指针+状态记录来做是可以的
	// 词频分组统计数组更好一点
	type statInfo struct {
		name byte
		count int
	}
	len1, len2 := len(name), len(typed)
	if len2 < len1 {
		return false
	}

	var statArr1, statArr2 []statInfo

	// 构建分组统计
	prevI := 0
	for i := 0; i <= len1; i++ {
		if i == len1 || name[i] != name[prevI] {
			statArr1 = append(statArr1, statInfo{
				name: name[prevI],
				count: i - prevI,
			})
			prevI = i
		}
	}

	prevI = 0
	for i := 0; i <= len2; i++ {
		if i == len2 || typed[i] != typed[prevI] {
			statArr2 = append(statArr2, statInfo{
				name: typed[prevI],
				count: i - prevI,
			})
			prevI = i
		}
	}

	lena, lenb := len(statArr1), len(statArr2)
	if lena != lenb {
		return false
	}
	for i := 0; i < lena; i++ {
		if statArr1[i].name != statArr2[i].name || statArr1[i].count > statArr2[i].count {
			return false
		}
	}
	return true
}

// @lc code=end


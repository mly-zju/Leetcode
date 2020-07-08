/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 *
 * https://leetcode-cn.com/problems/monotone-increasing-digits/description/
 *
 * algorithms
 * Medium (40.18%)
 * Likes:    34
 * Dislikes: 0
 * Total Accepted:    2.7K
 * Total Submissions: 6.7K
 * Testcase Example:  '10'
 *
 * 给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
 *
 * （当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）
 *
 * 示例 1:
 *
 * 输入: N = 10
 * 输出: 9
 *
 *
 * 示例 2:
 *
 * 输入: N = 1234
 * 输出: 1234
 *
 *
 * 示例 3:
 *
 * 输入: N = 332
 * 输出: 299
 *
 *
 * 说明: N 是在 [0, 10^9] 范围内的一个整数。
 *
 */

// @lc code=start
func monotoneIncreasingDigits(N int) int {
	// 注意不需要严格单调递增
	// 则对于任意dxxxx(xxxx不单调递增), 至少我们可以得到(d-1)9999, 这是损耗最小的。从大到小走这个过程直到单调递增
	tmpArr := getArr(N)

	length := len(tmpArr)
	for true {
		// 是否可以停止循环的标识
		tmpFlag := true

		for i := 0; i < length-1; i++ {
			if tmpArr[i] > tmpArr[i+1] {
				tmpFlag = false
				tmpArr[i]--
				for j := i + 1; j < length; j++ {
					tmpArr[j] = 9
				}
				break
			}
		}

		if tmpFlag {
			break
		}
	}

	res := 0
	for _, val := range tmpArr {
		res = res*10 + val
	}
	return res
}

func getArr(N int) []int {
	res := []int{}
	for N > 0 {
		tmp := N % 10
		res = append(res, tmp)
		N = N / 10
	}

	// 反转
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}

// @lc code=end


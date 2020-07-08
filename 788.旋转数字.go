/*
 * @lc app=leetcode.cn id=788 lang=golang
 *
 * [788] 旋转数字
 *
 * https://leetcode-cn.com/problems/rotated-digits/description/
 *
 * algorithms
 * Easy (56.98%)
 * Likes:    49
 * Dislikes: 0
 * Total Accepted:    6.5K
 * Total Submissions: 11.3K
 * Testcase Example:  '10'
 *
 * 我们称一个数 X 为好数, 如果它的每位数字逐个地被旋转 180 度后，我们仍可以得到一个有效的，且和 X 不同的数。要求每位数字都要被旋转。
 *
 * 如果一个数的每位数字被旋转以后仍然还是一个数字， 则这个数是有效的。0, 1, 和 8 被旋转后仍然是它们自己；2 和 5 可以互相旋转成对方；6 和
 * 9 同理，除了这些以外其他的数字旋转以后都不再是有效的数字。
 *
 * 现在我们有一个正整数 N, 计算从 1 到 N 中有多少个数 X 是好数？
 *
 *
 * 示例:
 * 输入: 10
 * 输出: 4
 * 解释:
 * 在[1, 10]中有四个好数： 2, 5, 6, 9。
 * 注意 1 和 10 不是好数, 因为他们在旋转之后不变。
 *
 *
 * 注意:
 *
 *
 * N 的取值范围是 [1, 10000]。
 *
 *
 */

// @lc code=start
func rotatedDigits(N int) int {
	count := 0
	for i := 0; i <= N; i++ {
		if isGood(i) {
			count++
		}
	}
	return count
}

func isGood(n int) bool {
	// 旋转后有效且等于自己的为0，不等于自己的为1，否则为-1
	goodArr := [10]int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}
	// 对每位数字颠倒
	otherFlag := false
	var number int
	for n > 0 {
		number = n % 10
		if goodArr[number] == -1 {
			return false
		} else if goodArr[number] == 1 {
			otherFlag = true
		}
		n = n / 10
	}
	return otherFlag
}

// @lc code=end


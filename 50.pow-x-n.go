/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode-cn.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (33.24%)
 * Likes:    214
 * Dislikes: 0
 * Total Accepted:    39.9K
 * Total Submissions: 119.2K
 * Testcase Example:  '2.00000\n10'
 *
 * 实现 pow(x, n) ，即计算 x 的 n 次幂函数。
 *
 * 示例 1:
 *
 * 输入: 2.00000, 10
 * 输出: 1024.00000
 *
 *
 * 示例 2:
 *
 * 输入: 2.10000, 3
 * 输出: 9.26100
 *
 *
 * 示例 3:
 *
 * 输入: 2.00000, -2
 * 输出: 0.25000
 * 解释: 2^-2 = 1/2^2 = 1/4 = 0.25
 *
 * 说明:
 *
 *
 * -100.0 < x < 100.0
 * n 是 32 位有符号整数，其数值范围是 [−2^31, 2^31 − 1] 。
 *
 *
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	// 1. 直观方法，相乘n次, 当然要区分正负数
	// // 运行结果：超时了
	// if x == 0 {
	// 	return x
	// }

	// if n < 0 {
	// 	x = 1 / x
	// 	n = -n
	// }
	// var res float64 = 1
	// for i := 0; i < n; i++ {
	// 	res *= x
	// }
	// return res

	// 2. 另一种方案：x^2n = x^n * x^n, 不断二分下去

	// 先处理边界
	if x == 0 {
		return x
	} else if n == 0 {
		return 1
	}

	// 都转换成n为正数
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return _myPow(x, n)
}

func _myPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}

	if n%2 == 0 {
		tmp := _myPow(x, n/2)
		return tmp * tmp
	}

	tmp := _myPow(x, (n-1)/2)
	return tmp * tmp * x
}

// @lc code=end


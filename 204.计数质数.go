/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (32.86%)
 * Likes:    320
 * Dislikes: 0
 * Total Accepted:    50.1K
 * Total Submissions: 151.9K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 * 
 * 示例:
 * 
 * 输入: 10
 * 输出: 4
 * 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 * 
 * 
 */

// @lc code=start
func countPrimes(n int) int {
	// 1. 常规解法, 可能超时
	// count := 0
	// for i := 2; i < n; i++ {
	// 	if isPrime(i) {
	// 		count++
	// 	}
	// }
	// return count

	// 2. 排除法
	isPrime := make([]bool, n)
	// 都初始化为true
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			// 如果i是素数，那么i的倍数都不是素数
			tmp := i
			for j := 2; j < n; j++ {
				tmp = i * j
				if tmp >= n {
					break
				} else {
					isPrime[tmp] = false
				}
			}
		}
	}
	
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}

func isPrime(x int) bool {
	sqrtx := int(math.Sqrt(float64(x)))
	for i := 2; i <= sqrtx; i++ {
		if x % i == 0 {
			return false
		}
	}
	return true
}
// @lc code=end


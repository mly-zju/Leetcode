/*
 * @lc app=leetcode.cn id=1013 lang=golang
 *
 * [1013] 将数组分成和相等的三个部分
 *
 * https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/description/
 *
 * algorithms
 * Easy (40.86%)
 * Likes:    94
 * Dislikes: 0
 * Total Accepted:    31.1K
 * Total Submissions: 76.4K
 * Testcase Example:  '[0,2,1,-6,6,-7,9,1,2,0,1]'
 *
 * 给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。
 *
 * 形式上，如果可以找出索引 i+1 < j 且满足 (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ...
 * + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[0,2,1,-6,6,-7,9,1,2,0,1]
 * 输出：true
 * 解释：0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
 *
 *
 * 示例 2：
 *
 * 输入：[0,2,1,-6,6,7,9,-1,2,0,1]
 * 输出：false
 *
 *
 * 示例 3：
 *
 * 输入：[3,3,6,5,-2,2,5,1,-9,4]
 * 输出：true
 * 解释：3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= A.length <= 50000
 * -10^4 <= A[i] <= 10^4
 *
 *
 */
package main

// @lc code=start
func canThreePartsEqualSum(A []int) bool {
	// 先求sum
	sum := 0
	for _, val := range A {
		sum += val
	}
	if sum%3 != 0 {
		return false
	}
	sum = sum / 3

	sumCount := 0
	gapSum := 0
	thirdStart := 0
	// 第三个区间要考虑最后有0的情况，所以单独处理
	for index, val := range A {
		gapSum += val
		if gapSum == sum {
			gapSum = 0
			sumCount++
			if sumCount == 2 {
				thirdStart = index + 1
				break
			}
		}
	}

	length := len(A)
	if thirdStart == 0 || thirdStart > length-1 {
		return false
	}
	gapSum = 0
	for i := thirdStart; i < length; i++ {
		gapSum += A[i]
	}
	if gapSum == sum {
		return true
	}
	return false
}

// func main() {
// 	fmt.Println(canThreePartsEqualSum([]int{12, -4, 16, -5, 9, -3, 3, 8, 0}))
// }

// @lc code=end

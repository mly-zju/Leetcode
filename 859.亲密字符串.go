/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 *
 * https://leetcode-cn.com/problems/buddy-strings/description/
 *
 * algorithms
 * Easy (26.83%)
 * Likes:    58
 * Dislikes: 0
 * Total Accepted:    6.6K
 * Total Submissions: 24.3K
 * Testcase Example:  '"ab"\n"ba"'
 *
 * 给定两个由小写字母构成的字符串 A 和 B ，只要我们可以通过交换 A 中的两个字母得到与 B 相等的结果，就返回 true ；否则返回 false
 * 。
 *
 *
 *
 * 示例 1：
 *
 * 输入： A = "ab", B = "ba"
 * 输出： true
 *
 *
 * 示例 2：
 *
 * 输入： A = "ab", B = "ab"
 * 输出： false
 *
 *
 * 示例 3:
 *
 * 输入： A = "aa", B = "aa"
 * 输出： true
 *
 *
 * 示例 4：
 *
 * 输入： A = "aaaaaaabc", B = "aaaaaaacb"
 * 输出： true
 *
 *
 * 示例 5：
 *
 * 输入： A = "", B = "aa"
 * 输出： false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= A.length <= 20000
 * 0 <= B.length <= 20000
 * A 和 B 仅由小写字母构成。
 *
 *
 */

// @lc code=start
func buddyStrings(A string, B string) bool {
	// 分相等和不相等两种情况求解
	if A == B {
		// 如果相等，则只需要找是否在不同位置有俩相同字符即可
		return buddySame(A)
	} else {
		return buddyDiff(A, B)
	}
}

func buddySame(A string) bool {
	mapArr := [26]int{}
	for _, ch := range A {
		mapArr[ch-'a']++
		if mapArr[ch-'a'] > 1 {
			return true
		}
	}
	return false
}

func buddyDiff(A, B string) bool {
	lena, lenb := len(A), len(B)
	if lena != lenb {
		return false
	}

	var a, b byte
	flag := 0
	for i := 0; i < lena; i++ {
		if A[i] != B[i] {
			flag++
			if flag == 1 {
				// 如果是第一次，记录下次希望的值
				a = B[i]
				b = A[i]
			} else if A[i] != a || B[i] != b || flag > 2 {
				return false
			}
		}
	}
	return true
}

// @lc code=end


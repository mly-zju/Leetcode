package main

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (27.14%)
 * Likes:    1505
 * Dislikes: 0
 * Total Accepted:    146.2K
 * Total Submissions: 528.2K
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	// 先用朴素思路求解：两次扫描，每次以index为中点
	// length := len(s)
	// if length == 0 {
	// 	return ""
	// }
	// start, end := 0, 0
	// for i := 0; i < length; i++ {
	// 	tmpS, tmpE := getPalind(s, i)
	// 	if tmpE-tmpS > end-start {
	// 		start = tmpS
	// 		end = tmpE
	// 	}
	// }
	// return s[start : end+1]

	// 再来动态规划: dp(i,j) = (dp(i+1,j-1) and si == sj, 如果)
	dp := [][]bool{}
	length := len(s)
	if length == 0 {
		return ""
	}

	// 先初始化一下dp
	for i := 0; i < length; i++ {
		dp = append(dp, []bool{})
	}

	start, end := 0, 0
	for i := length - 1; i >= 0; i-- {
		for j := 0; j < length; j++ {
			if j < i {
				dp[i] = append(dp[i], false)
			} else if j == i {
				dp[i] = append(dp[i], true)
			} else if j == i+1 {
				isEqual := s[i] == s[j]
				dp[i] = append(dp[i], isEqual)
				if isEqual && j-i > end-start {
					start = i
					end = j
				}
			} else {
				ii := i + 1
				jj := j - 1
				if ii < length && jj >= 0 {
					isEqual := (dp[ii][jj] && s[i] == s[j])
					dp[i] = append(dp[i], isEqual)
					if isEqual && j-i > end-start {
						start = i
						end = j
					}
				}
			}
		}
	}

	return s[start : end+1]
}

func getPalind(s string, mid int) (start, end int) {
	length := len(s)
	gap := 1
	maxLen := 0
	// 先按照奇数的求解
	for mid-gap >= 0 && mid+gap < length {
		if s[mid-gap] == s[mid+gap] {
			maxLen = 2*gap + 1
			start = mid - gap
			end = mid + gap
			gap++
		} else {
			break
		}
	}

	// 再按照偶数的方法求解，mid作为偶数中较小的一个
	gap = 0
	for mid-gap >= 0 && mid+1+gap < length {
		if s[mid-gap] == s[mid+1+gap] {
			if 2*gap+2 > maxLen {
				maxLen = 2*gap + 2
				start = mid - gap
				end = mid + 1 + gap
			}
			gap++
		} else {
			break
		}
	}
	return
}

// func main() {
// 	s := "cbbd"
// 	fmt.Println(longestPalindrome(s))
// }

// @lc code=end

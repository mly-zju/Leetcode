package main

import "fmt"

// 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	if len1 == 0 || len2 == 0 {
		return 0
	}

	// dp[i][j]代表text1[0~i]与text2[0~j]的最长公共子串长度
	// 为了简单起见，dp[i][j]代表text1[0~i-1]和text2[0~j-1]
	dp := [][]int{}
	// 初始化
	for i := 0; i <= len1; i++ {
		dp = append(dp, make([]int, len2+1))
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = getMax(dp[i-1][j-1], getMax(dp[i-1][j], dp[i][j-1]))
			}
		}
	}

	return dp[len1][len2]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := "bsbininm"
	b := "jmjkbkjkv"
	// 1 is right
	fmt.Println(longestCommonSubsequence(a, b))
}

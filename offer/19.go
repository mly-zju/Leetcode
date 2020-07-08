package main

import "fmt"

func isMatch(s string, p string) bool {
	// 动态规划求解
	sLen, pLen := len(s), len(p)
	// dp[i][j]代表0-i的s是否匹配0-j的p
	dp := [][]bool{}
	// 初始化
	for i := 0; i <= sLen; i++ {
		dp = append(dp, make([]bool, pLen+1))
	}
	// 几个初始化
	dp[0][0] = true
	for j := 1; j <= pLen; j++ {
		// 查看是否都可以删除
		if p[j-1] == '*' {
			if j-2 >= 0 {
				dp[0][j] = dp[0][j-2]
			}
		}
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if j-2 >= 0 {
					if p[j-2] == s[i-1] || p[j-2] == '.' {
						// 1. 星号取1个字符的情况
						// 2. 星号取多个字符的情况
						dp[i][j] = dp[i-1][j] || dp[i][j-1]
					}
					// 3. 星号取0个字符的情况, 一定是前进2位
					dp[i][j] = dp[i][j] || dp[i][j-2]
				}
			}
		}
	}

	return dp[sLen][pLen]
}

func main() {
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}

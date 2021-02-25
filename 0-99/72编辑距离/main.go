package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	// 自下向上
	n := len(word1)
	m := len(word2)
	if n == 0 || m == 0 {
		return n + m
	}
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	// 边界状态初始化
	for i := 0; i < n+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < m+1; j++ {
		dp[0][j] = j
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			// 主要是看最后一个字符是否相等相等则 dp[i][j]=dp[i-1][j-1]
			if word1[i-1] == word2[j-1] {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]-1)
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])

			}

		}
	}
	return dp[n][m]

}

func min(x, y, z int) int {
	if x < y {
		if z < x {
			return z
		}
		return x
	}
	if z < y {
		return z
	}
	return y

}

func minDistance1(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	if n == 0 || m == 0 {
		return n + m
	}
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			dp[i][j] = -1
		}
	}
	return helper(word1, word2, n, m, dp)
}

func helper(word1 string, word2 string, i, j int, dp [][]int) int {
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	if i == 0 {
		dp[i][j] = j
		return dp[i][j]
	}
	if j == 0 {
		dp[i][j] = i
		return dp[i][j]
	}
	if word1[i-1] == word2[j-1] {
		dp[i][j] = 1 + min(helper(word1, word2, i-1, j, dp), helper(word1, word2, i, j-1, dp), helper(word1, word2, i-1, j-1, dp)-1)
	} else {
		dp[i][j] = 1 + min(helper(word1, word2, i-1, j, dp), helper(word1, word2, i, j-1, dp), helper(word1, word2, i-1, j-1, dp))
	}
	return dp[i][j]

}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
	fmt.Println(minDistance1(word1, word2))

}

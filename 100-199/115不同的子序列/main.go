package main

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int, m+1)
	// dp[i][j]表示s[i:]中出现t[j:]个子序列的个数
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1 // 当j=n时 空字符串是任意字符串的子序列。为1
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]

}

func main() {

}

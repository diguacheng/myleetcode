package main

import "fmt"

func longestPalindrome(s string) string {
	// 自底向上 动态规划
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}

func longestPalindrome3(s string) string {
	// 自底向上 动态规划
	ans := ""
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		for i := j; i >= 0; i-- {
			if i == j {
				dp[i][j] = 1
			} else if j-i == 1 {
				if s[i] == s[j] {
					dp[i][j] = 2
				}
			} else if i+1 < j {
				if s[i] == s[j] && dp[i+1][j-1] > 0 {
					dp[i][j] = dp[i+1][j-1] + 2
				}
			}
			if dp[i][j] > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

func longestPalindrome1(s string) string {
	// 中心扩展法
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

func longestPalindrome2(s string) string {
	//超时 动态规划
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	var help func(i, j int) int
	help = func(i, j int) int {
		if dp[i][j] != 0 {
			return dp[i][j]
		}
		if j-i == 0 {
			dp[i][j] = 1
			return dp[i][j]
		}
		if j-i == 1 && s[i] == s[j] {
			dp[i][j] = 1
			return dp[i][j]
		}
		if s[i] == s[j] {
			return help(i+1, j-1)
		}
		dp[i][j] = -1
		return dp[i][j]
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if help(i, j) == 1 && j+1-i > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestPalindrome3("babad"))

}

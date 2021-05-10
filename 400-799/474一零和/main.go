package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	//  超时 要 转成dp
	var getnums func(s string) (zeros, ones int)
	getnums = func(s string) (zeros, ones int) {
		zeros, ones = 0, 0
		n := len(s)
		for i := 0; i < n; i++ {
			if s[i] == '0' {
				zeros++
			} else {
				ones++
			}
		}
		return
	}

	max := 0
	l := len(strs)
	var backTrack func(i int, counts int, m, n int)
	backTrack = func(i int, counts int, m, n int) {
		if counts > max {
			max = counts
		}
		if i == l {

			return
		}
		zeros, ones := getnums(strs[i])
		if m >= zeros && n >= ones {
			backTrack(i+1, counts+1, m-zeros, n-ones)
		}
		backTrack(i+1, counts, m, n)
	}
	backTrack(0, 0, m, n)
	return max
}

func findMaxForm1(strs []string, m int, n int) int {
	//  超时 要 转成dp
	var getnums func(s string) (zeros, ones int)
	getnums = func(s string) (zeros, ones int) {
		zeros, ones = 0, 0
		n := len(s)
		for i := 0; i < n; i++ {
			if s[i] == '0' {
				zeros++
			} else {
				ones++
			}
		}
		return
	}

	l := len(strs)
	dp := make([][][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	var help func(i int, m, n int) int
	help = func(i int, m, n int) int {
		if i == -1 {
			return 0
		}
		if dp[i][m][n] > 0 {
			return dp[i][m][n]
		}
		zeros, ones := getnums(strs[i])
		var v1, v2 int
		if m >= zeros && n >= ones {
			v1 = help(i-1, m-zeros, n-ones) + 1
		}
		v2 = help(i-1, m, n)
		dp[i][m][n] = max(v1, v2)
		return dp[i][m][n]

	}
	return help(l-1, m, n)
}

func findMaxForm2(strs []string, m int, n int) int {
	// 更慢了
	l := len(strs)
	getnums := map[string][]int{}
	var zeros, ones int
	var sl int
	for i := 0; i < l; i++ {
		sl = len(strs[i])
		zeros, ones = 0, 0
		for k := 0; k < sl; k++ {
			if strs[i][k] == '0' {
				zeros++
			} else {
				ones++
			}
		}
		getnums[strs[i]] = []int{zeros, ones}
	}

	dp := make([][][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	var help func(i int, m, n int) int
	help = func(i int, m, n int) int {
		if i == -1 {
			return 0
		}
		if dp[i][m][n] > 0 {
			return dp[i][m][n]
		}
		zeros, ones := getnums[strs[i]][0], getnums[strs[i]][1]
		var v1, v2 int
		if m >= zeros && n >= ones {
			v1 = help(i-1, m-zeros, n-ones) + 1
		}
		v2 = help(i-1, m, n)
		dp[i][m][n] = max(v1, v2)
		return dp[i][m][n]

	}
	return help(l-1, m, n)
}

func findMaxForm3(strs []string, m int, n int) int {

	// 最好

	var getnums func(s string) (zeros, ones int)
	getnums = func(s string) (zeros, ones int) {
		zeros, ones = 0, 0
		n := len(s)
		for i := 0; i < n; i++ {
			if s[i] == '0' {
				zeros++
			} else {
				ones++
			}
		}
		return
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	l := len(strs)
	for i := 0; i < l; i++ {
		a, b := getnums(strs[i])
		for zero := m; zero >= a; zero-- {
			for one := n; one >= b; one-- {
				dp[zero][one] = max(1+dp[zero-a][one-b], dp[zero][one])
			}
		}
	}
	return dp[m][n]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	Array := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	fmt.Println(findMaxForm1(Array, m, n))
}

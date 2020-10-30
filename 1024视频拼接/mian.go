package main

import "fmt"

func videoStitching(clips [][]int, T int) int {
	dp := make([]int, T+1)
	dp[0] = 0
	for i := 1; i <= T; i++ {
		dp[i] = 1 << 31
	}
	for i := 1; i <= T; i++ {
		for _, c := range clips {
			l, r := c[0], c[1]
			if l < i && i <= r && dp[l]+1 < dp[i] {
				dp[i] = dp[l] + 1
			}

		}
	}
	if dp[T] == 1<<31 {
		return -1
	}
	return dp[T]

}

func videoStitching1(clips [][]int, T int) int {
	last, pre := 0, 0
	maxn := make([]int, T)
	for _, c := range clips {
		l, r := c[0], c[1]
		if l < T && r > maxn[l] {
			maxn[l] = r
		}
	}
	ans := 0
	for i, v := range maxn {
		if v > last {
			last = v
		}
		if i == last {
			return -1
		}
		if i == pre {
			ans++
			pre = last
		}
	}
	return ans

}

func main() {
	clips := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	fmt.Println(videoStitching(clips, 9))

}

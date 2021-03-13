package main

import "fmt"

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var res int
	var DFS func(temp, i int)

	DFS = func(temp, i int) {
		if i == n {
			if temp == S {
				res++
			}
			return
		}
		DFS(temp+nums[i], i+1)
		DFS(temp-nums[i], i+1)

	}
	DFS(nums[0], 1)
	DFS(-nums[0], 1)
	return res
}

func findTargetSumWays1(nums []int, S int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2001)
	}
	dp[0][nums[0]+1000] = 1
	dp[0][-nums[0]+1000]++
	for i := 1; i < n; i++ {
		for sum := -1000; sum <= 1000; sum++ {
			if dp[i-1][sum+1000] > 0 {
				dp[i][sum+nums[i]+1000] += dp[i-1][sum+1000]
				dp[i][sum-nums[i]+1000] += dp[i-1][sum+1000]
			}
		}

	}
	if S > 1000 {
		return 0
	}
	return dp[n-1][S+1000]
}

func findTargetSumWays2(nums []int, S int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, 2001)
	dp[nums[0]+1000] = 1
	dp[-nums[0]+1000]++ // 若nums[0]=0 则加减 都是0
	next := make([]int, 2001)
	for i := 1; i < n; i++ {
		next = make([]int, 2001)
		for sum := -1000; sum <= 1000; sum++ {
			if dp[sum+1000] > 0 {
				next[sum+nums[i]+1000] += dp[sum+1000]
				next[sum-nums[i]+1000] += dp[sum+1000]
			}
		}
		dp = next

	}
	if S > 1000 {
		return 0
	}
	return dp[S+1000]

}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	S := 3
	fmt.Println(findTargetSumWays(nums, S))

}

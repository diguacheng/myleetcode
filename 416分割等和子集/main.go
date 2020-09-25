package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	// 本质是dp
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}

	isVisited := make([][]int, n)
	for i := 0; i < n; i++ {
		isVisited[i] = make([]int, sum+1)
	}
	sum = sum / 2
	var DFS func(i, temp int) bool

	DFS = func(i, temp int) bool {

		if temp == sum {
			return true
		}
		if i >= n {
			return false
		}
		if isVisited[i][temp] == 1 {
			return true
		}
		if isVisited[i][temp] == -1 {
			return false
		}
		flag := DFS(i+1, temp+nums[i]) || DFS(i+1, temp)
		if flag == true {
			isVisited[i][temp] = 1
		} else {
			isVisited[i][temp] = -1
		}
		return flag
	}
	return DFS(1, nums[0]) || DFS(1, 0)

}

func canPartition1(nums []int) bool {
	// 本质是dp
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum+1)
	}
	if nums[0] <= sum {
		dp[0][nums[0]] = true
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= sum; j++ {
			if nums[i] == j {
				dp[i][j] = true
				continue
			}
			if nums[i] < j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[n-1][sum]

}

func canPartition2(nums []int) bool {
	// 本质是dp
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum+1)
	}
	dp[0][0] = true
	if nums[0] <= sum {
		dp[0][nums[0]] = true
	}
	for i := 1; i < n; i++ {

		for j := 0; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			if nums[i] < j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
		if dp[i][sum] {
			return true
		}

	}
	return dp[n-1][sum]

}

func canPartition3(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	dp := make([]bool, sum+1)

	dp[0] = true
	if nums[0] <= sum {
		dp[nums[0]] = true
	}
	for i := 1; i < n; i++ {
		for j := sum; nums[i] <= j; j-- {
			if dp[sum] {
				return true
			}
			dp[j] = dp[j] || dp[j-nums[i]]
		}

	}
	return dp[sum]
}

func main() {
	S := []int{1, 5, 11, 5}
	fmt.Println(canPartition2(S))

}

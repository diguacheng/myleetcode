package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, 1)
	dp[0] = 1
	maxans := 1
	for i := 1; i < len(nums); i++ {
		maxval := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if maxval < dp[j] {
					maxval = dp[j]

				}
			}
		}
		dp=append(dp,maxval + 1)
		if maxans < dp[i] {
			maxans = dp[i]

		}
	}
	return maxans

}

func main() {
	nums := []int{4, 10, 4, 3, 8, 9}
	fmt.Println(lengthOfLIS(nums))

}

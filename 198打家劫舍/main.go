package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n==2{
		return max(nums[0], nums[1])
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	dp[2] = max(nums[0], nums[1])
	for i := 3; i < n+1; i++ {
		dp[i] = max(nums[i-1]+dp[i-2], dp[i-1])
	}
	fmt.Println(dp)
	return dp[n]

}

func rob1(nums []int) int {
	// 优化 使用滚动数字 

	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	first:=nums[0]
	second:=max(nums[0], nums[1])
	for i:=2;i<n;i++{
		first,second=second,max(nums[i]+first,second)
	}

	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := []int{1,10,2,1,1,7,1,3,1}
	fmt.Println(rob(s))
	s = []int{10,2,1,1,7,1,3}
	fmt.Println(rob(s))
	s = []int{10,2,1,1,7,1,3,1}
	fmt.Println(rob(s))

}
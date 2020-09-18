package main

import "fmt"

func rob(nums []int)int{
	n:=len(nums)
	if n==1{
		return nums[0]
	}
	
	return max(help(nums[1:]), help(nums[:n-1]))
}

func help(nums []int) int {
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
	s := []int{2,1,1,7,1,1,1}
	fmt.Println(rob(s))

}
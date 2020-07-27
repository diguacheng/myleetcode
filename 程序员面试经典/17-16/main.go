package main

import (
	"fmt"
)

func massage1(nums []int) int {
	// 递归算法
	// 超时
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a := nums[0] + massage(nums[2:])
	b := massage(nums[1:])
	if a > b {
		return a
	}
	return b

}

func massage(nums []int) int {
	// 递归算法
	// 用列表存储结果
	helper := make([]int, len(nums))
	for i:=0;i<len(nums);i++ {
		helper[i]=-1
	}
	return dpp(nums, helper, 0)

}

func dpp(nums, helper []int, i int) int {
	if i == len(nums) {
		return 0
	}
	if i == len(nums)-1 {
		return nums[i]
	}
	if helper[i] < 0 {
		a := nums[i] + dpp(nums, helper, i+2)
		b := dpp(nums, helper, i+1)
		if a > b {
			helper[i] = a
		} else {
			helper[i] = b
		}
	}
	return helper[i]

}

func massage3(nums []int)int{
	l:=len(nums)
	if l==0{
		return 0
	}
	dp0:=0
	dp1:=nums[0]
	for i:=1;i<l;i++{
		tdp0:=max(dp0,dp1)
		tdp1:=dp0+nums[i]
		dp0=tdp0// 前i个预约，第i个预约不接的最长用时
		dp1=tdp1// 前i个预约，第i个预约接的最长用时
	}
	return max(dp0,dp1)
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func main() {
	nums := []int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}
	fmt.Println(massage3(nums))

}

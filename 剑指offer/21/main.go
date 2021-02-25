package main

import (
	"fmt"
)

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	start := 0
	end := len(nums) - 1
	for start < end {
		for nums[start]%2 == 1 && start < len(nums)-1 {
			start++
		}
		for nums[end]%2 == 0 && end > 0 {
			end--
		}
		if start < end {
			nums[start], nums[end] = nums[end], nums[start]
		}
	}
	return nums

}

func main() {
	nums := []int{2, 4, 6}
	fmt.Println(exchange(nums))

}

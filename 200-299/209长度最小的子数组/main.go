package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	left, right := 0, 0
	sum := 0
	minl := len(nums) + 1
	n := len(nums)
	for right < n {
		sum += nums[right]
		right++
		for sum >= s {

			if right-left < minl {
				minl = right - left
			}

			sum -= nums[left]
			left++
		}
	}
	if minl == len(nums)+1 {
		return 0
	}
	return minl

}

func main() {
	s := 11
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(minSubArrayLen(s, nums))

}

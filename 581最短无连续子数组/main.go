package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	max := -1 << 31
	min := 1 << 31

	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			if nums[i] < min {
				min = nums[i]
			}
			if nums[i-1] > max {
				max = nums[i-1]
			}
		}
	}
	if min == 1<<31 {
		return 0
	}
	start, end := 0, n-1
	for start < n && nums[start] <= min {
		start++
	}

	for end >= 0 && nums[end] >= max {
		end--
	}
	return end - start + 1

}

func main() {
	s := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(s))

}

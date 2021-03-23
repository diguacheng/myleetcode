package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	start := 0
	max := 0
	var i int
	for i = 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			if max < i-start {
				max = i - start
			}
			start = i
		}
	}
	if i-start > max {
		max = i - start
	}
	return max
}

func main() {
	s := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(s))
}

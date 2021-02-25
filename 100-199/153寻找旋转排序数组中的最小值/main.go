package main

import "fmt"

func findMin(nums []int) int {
	n := len(nums)
	if n == 1 || nums[0] < nums[n-1] {
		return nums[0]
	}
	start, end := 0, n-1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] > nums[n-1] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]

}

func main() {
	s := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(s))

}

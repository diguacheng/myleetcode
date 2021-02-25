package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	start, end := 0, n-1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] > mid {
			if mid-1 >= 0 && nums[mid-1] == mid-1 {
				return mid
			}
			end = mid - 1
		}
		if nums[mid] == mid {
			if mid+1 < n && nums[mid+1] > mid+1 {
				return mid + 1
			}
			start = mid + 1
		}
	}
	if nums[n-1] == n {
		return 0
	}
	return n

}

func main() {
	nums := []int{0, 1}
	fmt.Println(missingNumber(nums))

}

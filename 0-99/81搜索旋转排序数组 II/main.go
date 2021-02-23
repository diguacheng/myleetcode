package main

import "fmt"

func search(nums []int, target int) bool {
	n := len(nums)

	start, end := 0, n-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid]==nums[start]{
			start++
			continue
		}
		if nums[mid] > nums[0] {
			if nums[mid] > target && nums[0] <= target {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				start = mid + 1
			} else {
				end = mid - 1

			}
		}
	}
	return false
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 3
	fmt.Println(search(nums, target))
}
package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	start, end := 0, n-1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[mid] > target && nums[0] <= target {
				end = mid - 1

			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < target && nums[n-1] >= target {
				start = mid + 1
			} else {
				end = mid - 1
			}

		}

	}
	return -1

}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))

}

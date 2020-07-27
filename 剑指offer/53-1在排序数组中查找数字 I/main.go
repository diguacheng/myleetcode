package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := 0

	start, end := 0, n-1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			res++
			for i := mid - 1; i >= 0; i-- {
				if nums[i] == target {
					res++
				}
			}
			for i := mid + 1; i < n; i++ {
				if nums[i] == target {
					res++
				}
			}
			break
		}
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		}
	}
	return res

}
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(search(nums, 6))

}

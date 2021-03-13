package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)

	n := len(nums)
	var idx int
	for i := 0; i < n; i++ {
		idx = nums[i]
		if idx < 0 {
			idx = -idx
		}
		if nums[idx-1] > 0 {
			nums[idx-1] = -nums[idx-1]
		}

	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res

}

func main() {
	s := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(s))

}

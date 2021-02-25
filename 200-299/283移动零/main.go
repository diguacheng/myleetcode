package main

import "fmt"

func moveZeroes(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	count := 0
	start, end := 0, 0
	for end < n {
		for end < n && nums[end] == 0 {
			count++
			end++
		}
		if end < n {
			nums[start] = nums[end]
			start++
			end++
		}

	}

	for start < n {
		nums[start] = 0
		start++
	}

}

func main() {
	s := []int{0, 0}
	moveZeroes(s)
	fmt.Println(s)

}

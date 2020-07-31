package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < n; second++ {

			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			value := target - (nums[first] + nums[second])
			third := second + 1
			fourth := n - 1
			for third < fourth {
				if nums[third]+nums[fourth] == value {
					res = append(res, []int{nums[first], nums[second], nums[third], nums[fourth]})
				}
				if nums[third]+nums[fourth] > value {
					fourth--
					for fourth > third && nums[fourth] == nums[fourth+1] {
						fourth--

					}
				} else {
					third++
					for fourth > third && nums[third] == nums[third-1] {
						third++

					}
				}

			}
		}
	}
	return res

}

func main() {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	target := 0
	fmt.Println(fourSum(nums, target))

}

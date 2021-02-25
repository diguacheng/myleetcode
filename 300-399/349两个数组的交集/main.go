package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	res := make([]int, 0)
	for _, v := range nums1 {
		set[v] = true
	}
	for _, v := range nums2 {
		if flag, ok := set[v]; flag && ok {
			res = append(res, v)
			set[v] = false
		}
	}
	return res

}

func main() {
	var nums1 = []int{1, 2, 2, 1}
	var nums2 = []int{2, 2}
	fmt.Println(intersection(nums1, nums2))

}

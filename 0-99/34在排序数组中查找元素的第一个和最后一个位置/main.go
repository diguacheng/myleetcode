package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	res := []int{-1, -1}
	start, end := 0, n-1
	var mid, i int
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			res = []int{mid, mid}
			i = mid - 1
			for i >= 0 && nums[i] == target {
				i--
			}
			res[0] = i + 1
			i := mid + 1
			for i < n && nums[i] == target {
				i++
			}
			res[1] = i - 1
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

func searchRange1(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	res := []int{-1, -1}
	start, end := 0, n-1
	var mid int
	for start < end {
		mid = (start + end) / 2
		if nums[mid] >= target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if nums[start] == target {
		i := start
		for i+1 < n && nums[i+1] == target {
			i++
		}
		res = []int{start, i}
	}
	return res
}

func searchRange2(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	start := binarySerach(nums, target)
	if nums[start] != target || start == n {
		return []int{-1, -1}
	}
	end := binarySerach(nums, target+1) - 1
	return []int{start, end}

}

func binarySerach(arr []int, target int) int {
	n := len(arr)
	l, r := 0, n
	for l <= r {
		mid := (l + r) >> 1
		if arr[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
func main() {
	nums := []int{1}
	fmt.Println(searchRange1(nums, 1))

}

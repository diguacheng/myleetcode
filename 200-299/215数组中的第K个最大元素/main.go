package main

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	idx := povit(nums)
	if idx+1 == k {
		return nums[idx]
	} else if idx+1 < k {
		return findKthLargest(nums[idx+1:], k-idx-1)
	}
	return findKthLargest(nums[:idx], k)
}

func quickSort(nums []int) {
	if len(nums) > 0 {
		idx := povit(nums)
		quickSort(nums[:idx])
		quickSort(nums[idx+1:])
	}

}

func povit(nums []int) int {
	n := len(nums)
	l := 0
	r := n - 1
	p := nums[l]
	idx := l
	for l < r {
		for l < r && nums[r] <= p {
			r--
		}
		if l < r {
			nums[l] = nums[r]
			idx = r
		}
		for l < r && nums[l] >= p {
			l++
		}
		if l < r {
			nums[r] = nums[l]
			idx = l
		}
	}
	nums[idx] = p
	return idx
}

func heapSort(nums []int) {
	l := len(nums) - 1
	f := len(nums) / 2
	for l >= f {

	}
}

func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))
}

package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	s, e := 0, n-1
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
	s, e = 0, k-1
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
	s, e = k, n-1
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}

}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(s, 3)
	fmt.Println(s)

}

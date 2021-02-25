package main

import (
	"fmt"
)

// 超时2
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	table := make(map[int]int)
	for num, i := range nums {
		for item, idx := range table {
			if abs(idx-i) <= k && abs(num-item) <= t {
				return true
			}
		}
		table[num] = i
	}
	return false

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{}))

}

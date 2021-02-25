package main

import (
	"fmt"
)

func findRepeatNumber(nums []int) int {
	/*
		哈希表

		45ms左右
	*/
	d := make(map[int]bool)
	for _, num := range nums {
		if d[num] != true {
			d[num] = true
		} else {
			return num
		}
	}
	return 0
}

func findRepeatNumber1(nums []int) int {
	/*
		100ms左右
	*/
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			} else {
				temp := nums[nums[i]]
				nums[nums[i]] = nums[i]
				nums[i] = temp
			}
		}
	}
	return 0

}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	//nums1:=[]int{1,1,1}
	fmt.Println(findRepeatNumber1(nums))
}

package main

import "sort"

// func findKthLargest(nums []int, k int) int {
// 	sort.Ints(nums)
// 	start:=len(nums)-1
// 	if k==1{
// 		return nums[start]
// 	}
// 	for k>1{
// 		if nums[start-1]!=nums[start]{
// 			k--
// 		}
// 		start--
// 	}
// 	return nums[start]



// }

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
	



}

func main() {

}
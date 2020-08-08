package main

import "fmt"

func search(nums []int, target int) int {

	start, end := 0, len(nums)-1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}

	}
	return -1
}

func main() {
	x:=[]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(search(x,1))
	fmt.Println(search(x,7))

}



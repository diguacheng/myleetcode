package main

import ("fmt"
	"sort"
)

func sortArray(nums []int) []int {
	ssort(nums)
	return nums

}

func ssort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	flag := nums[0]
	index := 0
	start, end := 0, len(nums)-1
	for start < end {
		for start < end && nums[end] >= flag {
			end--
		}
		if start < end {
			nums[start] = nums[end]
			index = end

		}
		for start < end && nums[start] <= flag {
			start++
		}
		if start < end {
			nums[end] = nums[start]
			index = start
		}
	}
	nums[index] = flag
	ssort(nums[:index])
	ssort(nums[index+1:])

}

func sortArray1(nums []int) []int {
	sort.Ints(nums)
	return nums


}

func main() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(sortArray(nums))

}

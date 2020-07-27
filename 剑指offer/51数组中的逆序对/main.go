package main

import "fmt"

func reversePairs(nums []int) int {

	return help(nums, 0, len(nums)-1)

}

func help(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	cnt := help(nums, start, mid) + help(nums, mid+1, end)
	temp := make([]int, 0)
	i, j := start, mid+1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			temp = append(temp, nums[i])
			cnt += j - (mid + 1)
			i++
		} else {
			temp = append(temp, nums[j])

			j++
		}
	}
	for ; i <= mid; i++ {
		temp = append(temp, nums[i])
		cnt += end - (mid + 1) + 1

	}
	for ; j <= end; j++ {
		temp = append(temp, nums[j])

	}
	for i := start; i <= end; i++ {
		nums[i] = temp[i-start]
	}
	return cnt

}

func reversePairs1(nums []int) int {

	return help(nums, 0, len(nums)-1)

}

func main() {
	nums := []int{7, 5, 6, 4}
	fmt.Println(reversePairs(nums))

}

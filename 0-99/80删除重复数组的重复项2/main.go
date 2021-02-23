package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}
	i := 0
	k := 0
	var count int
	for j := 1; j < n; j++ {
		if nums[j] != nums[i] {
			count = 0
			for i < j && count < 2 {
				nums[k] = nums[i]
				count++
				k++
				i++
			}
			i = j
		}
	}
	count = 0
	for i< n && count < 2 {
		nums[k] = nums[i]
		count++
		k++
		i++
	}
	return k
}

func main() {
	nums:= []int{1,1,1,2,2,3}
	fmt.Println(removeDuplicates(nums))

}
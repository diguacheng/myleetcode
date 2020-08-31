package main

import (
	"fmt"	
	"sort"
)


func permuteUnique(nums []int) [][]int {

	res := make([][]int, 0)
	temp := make([]int, 0)
	sort.Ints(nums)
	help(nums, temp, &res)
	return res

}

func help(nums, temp []int, res *[][]int) {
	n := len(nums)
	if n == 0 {
		t := make([]int, len(temp))
		copy(t, temp)

		*res = append(*res, t)
	}
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		tempnums := append([]int{}, nums[:i]...)
		tempnums = append(tempnums, nums[i+1:]...)
		help(tempnums, temp, res)
		temp = temp[:len(temp)-1]

	}

}
func main() {
	nums := []int{1, 2, 3, 3}
	fmt.Println( permuteUnique(nums))

}
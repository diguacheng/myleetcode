package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	sort.Ints(nums)
	backTrack(nums, 0, list, &res)
	return res
}

func backTrack(nums []int, pos int, list []int, res *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*res = append(*res, ans)
	for i := pos; i < len(nums); i++ {
		if i != pos && nums[i] == nums[i-1] {
			continue
		}
		list = append(list, nums[i])
		backTrack(nums, i+1, list, res)
		list = list[:len(list)-1]
	}

}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))

}

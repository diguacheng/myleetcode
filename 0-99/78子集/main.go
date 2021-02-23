package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{}
	list := make([]int, 0)
	backTrack(nums, 0, list, &res)
	return res
}

func backTrack(nums []int, pos int, list []int, res *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*res = append(*res, ans)
	for i := pos; i < len(nums); i++ {
		list := append(list, nums[i])
		backTrack(nums, i+1, list, res)
		list = list[:len(list)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))

}
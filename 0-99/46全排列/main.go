package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	temp := make([]int, 0)
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

		temp = append(temp, nums[i])
		tempnums := append([]int{}, nums[:i]...)
		tempnums = append(tempnums, nums[i+1:]...)
		help(tempnums, temp, res)
		temp = temp[:len(temp)-1]

	}

}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(permute(nums))

}

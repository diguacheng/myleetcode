package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	s := make([]int, 0) // 存索引
	i:=0
	round:=0
	for round<2{
		for len(s) > 0 && nums[s[len(s)-1]] < nums[i] {
			res[s[len(s)-1]] = nums[i]
			s = s[:len(s)-1]
		}
		
		if round==0{
			s = append(s, i)
		}
		i=i+1
		if i==len(nums){
			round++
			i=0
		}
	}
	for _, i := range s {
		res[i] = -1
	}
	return res
}

func main() {
	s := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(s))
}
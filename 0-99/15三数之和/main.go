package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	sort.Ints(nums)
	fmt.Println(nums)
	
	
	for first := 0; first <n; first++ {
		if first>0&&nums[first]==nums[first-1]{
			continue
		}
		third := n - 1
		target:=-1*nums[first]
		for second:=first+1; second<n; second++ {
			if second>first+1&&nums[second]==nums[second-1]{
				continue
			}
			for second<third&&nums[second]+nums[third]>target{
				third--
			}
			if second==third{
				break
			}
			if nums[first]+nums[second]+nums[third]==0{
				res= append(res,[]int{nums[first],nums[second],nums[third]})
			}

		}

	}
	return res
	
}


func main() {
	in:=[]int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(in))
}
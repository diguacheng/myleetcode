package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	m := 1 << 32
	res := 0
	for first := 0; first < n; first++ {
		second:=first+1
		third:=n-1
		for second<third{
			temp:=nums[first]+nums[second]+nums[third]
			if temp==target{
				return target
			}
			if temp>target{
				third--
			}else{
				second++
			}
			x:=abs(temp- target)
			if x<m{
				m=x
				res=temp
			}
		}
	}
	return res

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums := []int{0, 2, 1, -3}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

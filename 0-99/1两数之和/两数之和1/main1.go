package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int 
	n := len(nums)
	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			if nums[i]+nums[j] == target{
				res = []int{i, j}
				
			}  
		}
	}
	return res
	
}

func main() {
	nums:=[]int{3,2,4}
	target:=6
	fmt.Println(twoSum(nums,target))
	
}
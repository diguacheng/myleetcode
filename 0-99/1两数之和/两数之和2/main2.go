package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	var res []int
	n := len(nums)
	for i:=0;i<n;i++{
		hashmap[nums[i]]=i
	}
	for i:=0;i<n;i++{
		a:=target-nums[i]
		if index,ok:=hashmap[a];index!=i &&ok==true{
			res=[]int{i,hashmap[a]}
			break
		}
	}
	return res
	
}

func main() {
	nums:=[]int{3,2,4}
	target:=6
	fmt.Println(twoSum(nums,target))
	
}
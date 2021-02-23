package main

import ("fmt")

func searchInsert(nums []int, target int) int {
	if nums[0]>target{
		return 0
	}
	if nums[len(nums)-1]<target{
		return len(nums)
	}

	for i,k := range nums {
		if k == target {
			return i
		}
		if k>target{
			return i
		}
	}
	return 0




    
}

func main() {
	nums:=[]int{1,3,5,6}
	target:=4
	fmt.Println(searchInsert(nums,target))

}
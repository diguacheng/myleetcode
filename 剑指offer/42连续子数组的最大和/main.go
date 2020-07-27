package main

import "fmt"

func maxSubArray(nums []int) int {
	res:=make([]int,0)
	curr:=nums[0]
	res=append(res,curr)
	max:=0
	for i:=1;i<len(nums);i++{
		if curr>=0{
			curr=curr+nums[i]
		}else{
			curr=nums[i]
		}
		if max<curr{
			max=curr
		}
		
	}
	return max



}


func main() {
	nums :=[]int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))

}
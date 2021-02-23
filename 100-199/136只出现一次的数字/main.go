package main

import "fmt"

func singleNumber(nums []int) int {
	n:=len(nums)
	res:=0
	for i:=0;i<n;i++ {
		
		res=res^nums[i]
	}
	return res

}

func main(){
	num:=[]int{2,2,1}
	fmt.Println(singleNumber(num))

}
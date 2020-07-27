package main

import (
	"fmt"
	"math"
)

func singleNumber(nums []int) int {
	m:=make(map[int]int)
	for _,num := range nums {
		m[num]++
	}
	for k,v := range m{
		if v==1{
			return k
		}
	}
	return 0

}
func singleNumber1(nums []int) int {
	if len(nums)==0{
		return 0
	}
	bitSum:=[32]int{}
	for i:=0;i<len(nums);i++{
		bitMask :=1
		for j:=0;j<32;j++{
			bit:=nums[i]&bitMask
			if bit!=0{
				bitSum[j]++
			}
			bitMask=bitMask<<1

		}
	}
	result:=0
	for i:=0;i<32;i++{
		if bitSum[i]%3!=0{
			result=result+int(math.Pow(float64(2),float64(i)))
		}
	}
	return result


}


func main() {
	nums:=[]int{3,4,3,3}
	fmt.Println(singleNumber1(nums))

}
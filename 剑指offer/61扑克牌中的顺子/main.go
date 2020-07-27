package main

import (
	"fmt"
	"sort"
)

func isStraight(nums []int) bool {
	sort.Ints(nums)
	i:=0
	zero:=0
	for nums[i]==0{
		i++
		zero++
	}
	for i=zero+1;i<5;i++{
		x := nums[i] - nums[i-1]-1
		if x>0{
			if x<=zero{
				zero=zero-x
			}else{
				return false
			}
		}else if x<0{
			return false
		}

	}
	return true

	
}

func main() {
	nums:=[]int{0,0,1,2,5}
	fmt.Println(isStraight(nums))

}

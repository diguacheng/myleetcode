package main

import "fmt"

func canJump(nums []int) bool {

	i := len(nums) - 2
	for i >= 0 {
		if nums[i] == 0 {
			i--
			j := 1
			for i >= 0 {
				if nums[i] > j {
					break
				} else {
					i--
					j++
				}
			}
			if i < 0 {
				return false
			}
			i--
		} else {
			i--
		}
	}
	return true

}

func canJump1(nums []int) bool {
	//贪心策略
	n:=len(nums)-1
	max:=0
	for i := 0; i <=n; i++{
		if i<=max{
			if max>i+nums[i]{
				max=i+nums[i]
			}
			if max>=n{
				return true
			}
		}
	}
	return false




}

func main() {
	nums := []int{0,1}
	fmt.Println(canJump(nums))

}

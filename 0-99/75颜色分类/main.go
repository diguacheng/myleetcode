package main

import "fmt"

func sortColors(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	start, end := 0, n-1
	cur := 0
	for cur <= end {
		fmt.Println(nums,cur)
		if nums[cur] == 0 {
			nums[cur], nums[start] = nums[start], nums[cur]
			// 因为nums[cur]如果是2的话，会被移动到后面，此时cur的值是不变的，当它为0时,会和前面的元素换，只能时1，cur可以向右边移动
			
			start++
			cur++
		} else if nums[cur] == 2{
			nums[cur], nums[end] = nums[end], nums[cur]
			end--
		}else{
			cur++
		}
		
	}

}

func main() {
	nums := []int{1,1,1,1,2,0,1,2}
	sortColors(nums)
	fmt.Println(nums)
}
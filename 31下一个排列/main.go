package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		nums[i],nums[j] = nums[j],nums[i]

	}
	l:=i+1+(n-i-1)/2
	for j:=i+1;j<l; j++ {
		nums[j],nums[n-j+i]=nums[n-j+i],nums[j]
	}
}

func main() {
	nums := []int{2,3,1,3,3}
	nextPermutation(nums)
	fmt.Println(nums)

}
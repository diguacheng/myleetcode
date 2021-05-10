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
		nums[i], nums[j] = nums[j], nums[i]

	}
	l := i + 1 + (n-i-1)/2
	for j := i + 1; j < l; j++ {
		nums[j], nums[n-j+i] = nums[n-j+i], nums[j]
	}
}

func nextPermutation2(nums []int)  {
    n:=len(nums)
    if n<=1{
        return 
    }
    i:=n-2
    for i>=0&&nums[i]>=nums[i+1]{
        i--
    }
    if i<0{
        reverse(nums)
        return 
    }
    left:=i 
    for i=left+1;i<n;i++{
        if nums[i]<=nums[left]{
            break
        }
    }
    right:=i-1
    nums[left],nums[right]=nums[right],nums[left]
    reverse(nums[left+1:])
    return 
}

func reverse(arr []int){
    n:=len(arr)
    for i:=0;i<n/2;i++{
        arr[i],arr[n-1-i]=arr[n-1-i],arr[i]
    }
}

func main() {
	nums := []int{5,1,1}
	nextPermutation2(nums)
	fmt.Println(nums)

}

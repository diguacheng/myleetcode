package main

import "sort"

func findDuplicate1(nums []int) int {
	sort.Ints(nums)
	for i:=1;i<len(nums);i++ {
		if nums[i]==nums[i-1] {
			return nums[i]
		}
	}
	return 0
}


func findDuplicate2(nums []int) int {
	sort.Ints(nums)
	for i:=len(nums)-1;i>=0;i--{
		if nums[i]>i{
			return nums[i]
		}
	}
	return 0
}

func findDuplicate3(nums []int) int {
	// 二分法 先排序 再通过二分法找重复元素
	sort.Ints(nums)
	l,r:=0,len(nums)-1
	var mid int
	for l<r{
		mid=(l+r)/2
		if nums[mid]<=mid{
			r=mid-1
		}else{
			l=mid+1
		}
	}
	return nums[l]

}

func findDuplicate4(nums []int) int {
	sort.Ints(nums)
	for i:=len(nums)-1;i>=0;i--{
		if nums[i]>i{
			return nums[i]
		}
	}
	return 0
}

func findDuplicate5(nums []int) int {
	//快慢指针 重复出现 说明有环 floyd 算法 想找相遇点 再找入口点
	slow,fast:=nums[0],nums[nums[0]]
	for slow!=fast {
		slow,fast=nums[slow],nums[nums[fast]]
	}
	slow=0
	for slow!=fast {
		slow=nums[slow]
		fast=nums[fast]
	}
	return slow
}

func findDuplicate(nums []int) int {
	//快慢指针 重复出现 说明有环 floyd 算法 想找相遇点 再找入口点
	slow,fast:=nums[0],nums[nums[0]]
	for slow!=fast {
		slow,fast=nums[slow],nums[nums[fast]]
	}
	slow=0
	for slow!=fast {
		slow=nums[slow]
		fast=nums[fast]
	}
	return slow
}



func main(){

}
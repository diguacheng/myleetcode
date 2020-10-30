package main

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	n:=len(nums)
	for i := 1; i < n; i++ {
		if nums[i]==nums[i-1] {
			return true
		}
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	t:=map[int]int{}
	for _,v:=range nums{
		t[v]++
		if t[v]>1{
			return false
		}
	}
	return true

}

func main() {

}
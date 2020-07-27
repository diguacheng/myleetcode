package main

import "fmt"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
        return []int{-1,-1}
	}
    res := []int{-1,-1}
	start, end := 0, n-1
	var mid ,i int
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			res=[]int{mid,mid}
			i=mid-1
			for i>=0&&nums[i]== target {
				i--
			}
			res[0]=i+1
			i:=mid+1
			for  i<n && nums[i]== target{
				i++
			}
			res[1]=i-1
			break
		}
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		}
	}
	return res
}
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))

}
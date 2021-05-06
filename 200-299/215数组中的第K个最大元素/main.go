package main

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	idx := povit(nums)
	if idx+1 == k {
		return nums[idx]
	} else if idx+1 < k {
		return findKthLargest(nums[idx+1:], k-idx-1)
	}
	return findKthLargest(nums[:idx], k)
}

func quickSort(nums []int) {
	if len(nums) > 0 {
		idx := povit(nums)
		quickSort(nums[:idx])
		quickSort(nums[idx+1:])
	}

}

func povit(nums []int) int {
	n := len(nums)
	l := 0
	r := n - 1
	p := nums[l]
	idx := l
	for l < r {
		for l < r && nums[r] <= p {
			r--
		}
		if l < r {
			nums[l] = nums[r]
			idx = r
		}
		for l < r && nums[l] >= p {
			l++
		}
		if l < r {
			nums[r] = nums[l]
			idx = l
		}
	}
	nums[idx] = p
	return idx
}


func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func findKthLargest3(nums []int, k int) int {
    Qselect(nums,0,len(nums)-1,k)
    return nums[k-1]
}

func Qselect(arr []int,l,r int,k int){
    p:=Povit(arr,l,r)
    if p+1==k{
        return 
    }
    if p+1>k{
        Qselect(arr,l,p-1,k)
    }else{
        Qselect(arr,p+1,r,k-p)
    }
}

func Povit(arr []int,l,r int)int{
    c:=r
    for l<r{
        for l<r&&arr[l]>=arr[r]{
            l++
        }
        if l<r{
            arr[l],arr[r]=arr[r],arr[l]
            c=l
        }
        for l<r&&arr[r]<=arr[l]{
            r--
        }
        if l<r{
            arr[l],arr[r]=arr[r],arr[l]
            c=r
        }
    }
    return c 
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest3(nums, 4))
}

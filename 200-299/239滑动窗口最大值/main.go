package main

import (
	"container/heap"
	"sort"
)

// func maxSlidingWindow(nums []int, k int) []int {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	res:=make([]int,0)
// 	if len(nums) >= k && k >= 1 {
// 		list := make([]int, 0)
// 		for i := 0; i < k; i++ {
// 			for len(list) != 0 && nums[i] >= nums[list[len(list)-1]] {
// 				list = list[:len(list)-1]
// 			}
// 			list = append(list,i)
// 		}
// 		for i := k; i < len(nums); i++ {
// 			res=append(res,nums[list[0]])
// 			for len(list) != 0 && nums[i] >= nums[list[len(list)-1]] {
// 				list = list[:len(list)-1]
// 			}
// 			if len(list) != 0 && list[0] <= i-k {
//                 list = list[1:]
// 			}
// 			list = append(list,i)
// 		}
// 		res=append(res,nums[list[0]])

// 	}
// 	return res
// }

// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	list := []int{}
	push := func(i int) {
		for len(list) > 0 && nums[i] >= nums[list[len(list)-1]] {
			list = list[:len(list)-1]
		}
		list = append(list, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[list[0]]
	for i := k; i < n; i++ {
		push(k)
		for list[0] <= i-k {
			list = list[1:]
		}
		ans = append(ans, list[0])
	}
	return ans

}

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow1(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

func main() {

}

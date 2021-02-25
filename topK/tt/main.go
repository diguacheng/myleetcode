package main

import (
	"container/heap"
	"fmt"
)

type Intheap []int

func (h Intheap) Len() int           { return len(h) }
func (h Intheap) Less(i, j int) bool { return h[i] < h[j] }
func (h Intheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Intheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Intheap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func main() {

	nums := &Intheap{2, 5, 3, 9, 4, 6, 8}
	heap.Init(nums)
	fmt.Println(nums) // [2 4 3 9 5 6 8]
	heap.Push(nums, 1)
	fmt.Println(nums) //[1 2 3 4 5 6 8 9]
	heap.Push(nums, 3)
	fmt.Println(nums) // [1 2 3 3 5 6 8 9 4]

	heap.Remove(nums, 4)
	fmt.Println(nums) //[1 2 3 3 4 6 8 9]

	(*nums)[1] = 10
	fmt.Println(nums) //[1 10 3 3 4 6 8 9]
	heap.Fix(nums, 1)
	fmt.Println(nums) //[1 3 3 9 4 6 8 10]

	x := heap.Pop(nums) //1
	fmt.Println(x)

}

package main

import (
	"container/heap"	
	"sort"
)


type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key,value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int,k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

func topKFrequent1(nums []int, k int) []int {
    m := make(map[int]int)
    s := make([]int,0)
    for _,v := range nums {
        i,ok := m[v]
        if ok {
            m[v] = i+1
        }else{
            m[v] = 1
            s = append(s, v)
        }
    }

    sort.Slice(s, func(i, j int) bool {
		return m[s[i]] > m[s[j]]
	})

    return s[:k]
}


func main() {

}
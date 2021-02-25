package main

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() interface{} {
	n := len(*m)
	x := (*m)[n-1]
	*m = (*m)[:n-1]
	return x
}

func heapSort(arr []int) []int {
	m := minHeap(arr)
	heap.Init(&m)
	res := []int{}
	for len(m) != 0 {
		res = append(res, heap.Pop(&m).(int))
	}
	return res
}

func main() {
	arr := []int{2, 9, 5, 1, 7, 6, 3}
	fmt.Println(heapSort(arr))

}

func heapSort1(arr []int) []int {
	n := len(arr)
	heapify := func(arr []int, k int) {
		for {
			left := 2*k + 1  // 左子树下标
			right := 2*k + 2 // 右子树下标
			idx := k         // 根 左 右 最大值的索引
			if left < n && arr[left] > arr[idx] {
				idx = left
			}
			if right < n && arr[right] > arr[idx] {
				idx = right
			}
			if idx == k {
				// 若根节点较大，不用将节点下沉
				break
			}
			arr[k], arr[idx] = arr[idx], arr[k]
			// 交换过后，idx保存的不是最大值的索引，是从根节点换下的索引
			k = idx
			// 继续节点下沉
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, i)
	}
	for i := n - 1; i >= 1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		n-- // 注意这里的长度要减去一
		heapify(arr, 0)
	}
	return arr
}

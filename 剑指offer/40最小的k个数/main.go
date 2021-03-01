package main

import (
	"container/heap"
	"fmt"
)

func getLeastNumbers(arr []int, k int) []int {
	//744ms 5.38% 6.2mb 100%
	j := 0
	for j < k {
		for i := len(arr) - 1; i >= j+1; i-- {
			if arr[i] < arr[i-1] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
		j++
	}
	return arr[:k]

}

func getLeastNumbers1(arr []int, k int) []int {
	// 320ms 5.91% 6.3mb 100%
	j := 0
	for j < k {
		min := arr[j]
		mindex := j
		for i := len(arr) - 1; i >= j; i-- {
			if min > arr[i] {
				min = arr[i]
				mindex = i
			}
		}
		if mindex != j {
			arr[mindex], arr[j] = arr[j], arr[mindex]
		}
		j++
	}
	return arr[:k]

}

func getLeastNumbers2(arr []int, k int) []int {
	// 快速选择 没看懂
	if k >= len(arr) {
		return arr
	}

	var recur func(arr []int, k int)
	recur = func(arr []int, k int) {
		s, e := 0, len(arr)-1
		c := arr[s]
		for s < e {
			for s < e {
				if arr[s] > c {
					break
				}
				s++
			}

			for s < e {
				if arr[e] <= c {
					break
				}
				e--
			}
			if s < e {
				arr[s], arr[e] = arr[e], arr[s]
			}
		}
		if arr[e] < arr[0] {
			arr[0], arr[e] = arr[e], arr[0]
		}
		if e > k {
			recur(arr[0:e], k)
		} else if e < k {
			recur(arr[e:], k-e)
		} else {
			return
		}
	}
	recur(arr, k)
	return arr[0:k]

}

func getLeastNumbers3(arr []int, k int) []int {
	if k >= len(arr) {
		return arr
	}
	qselect(arr, k)
	return arr[0:k]

}

func qselect(arr []int, k int) {
	start, end := 0, len(arr)-1
	curr := arr[0]
	i := 0
	for start < end {
		for start < end && arr[end] >= curr {
			end--
		}
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
			i = end
		}
		for start < end && arr[start] < curr {
			start++
		}
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
			i = start
		}

	}
	arr[i] = curr
	if i < k-1 {
		qselect(arr[i+1:], k-i-1)
	} else if i > k {
		qselect(arr[:i], k)
	} else {
		return
	}
}


type MHeap []int 

func (h MHeap)Len()int {return len(h)}

func (h MHeap)Less(i,j int)bool {return h[i]>h[j]}

func (h MHeap)Swap(i,j int){
	h[i],h[j]=h[j],h[i]
}

func (h *MHeap)Push(x interface{}){
	*h=append(*h, x.(int))

}

func (h *MHeap)Pop()interface{}{
	old:=*h 
	n:=len(old)
	x:=old[n-1]
	*h=old[:n-1]
	return x
}



func main() {
	arr := []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}
	k := 8

	fmt.Println(getLeastNumbers3(arr, k))
	mh:=&MHeap{}
	heap.Push()

}

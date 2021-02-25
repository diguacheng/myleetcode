package main

import (
	"fmt"
	"math"
	"sort"
)

// type maxheap [][]int

// func (m maxheap) Len() int { return len(m) }

// func (m maxheap) Less(i,j int) bool {
// 	return sqrt(m[i])>sqrt(m[j])
// }

// func (m maxheap)Swap(i,j int){
// 	m[i],m[j]=m[j],m[i]
// }

// func (m *maxheap)Push(x interface{}){
// 	*m=append(*m, x.([]int))
// }

// func (m *maxheap)Pop()interface{}{
// 	n:=len(*m)
// 	x:=(*m)[n-1]
// 	*m=(*m)[:n-1]
// 	return x
// }

// func sqrt(point []int) float64 {
// 	x, y :=point[0],point[1]
// 	return math.Sqrt(float64(x*x+y*y))
// }

// func kClosest(points [][]int, K int) [][]int {
// 	maxh:=&maxheap{}
// 	heap.Init(maxh)
// 	n:=len(points)
// 	for i:=0;i<n;i++{
// 		l:=maxh.Len()
// 		if l<K{
// 			heap.Push(maxh,points[i])
// 		}else if l==K{
// 			if sqrt(points[i])<sqrt((*maxh)[0]){
// 				heap.Pop(maxh)
// 				heap.Push(maxh,points[i])
// 			}
// 		}
// 	}
// 	return *maxh

// }

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return sqrt(points[i]) < sqrt(points[j])
	})
	return points[:K]

}

func sqrt(point []int) float64 {
	x, y := point[0], point[1]
	return math.Sqrt(float64(x*x + y*y))
}

func main() {
	points := [][]int{{1, 3}, {2, -2}, {-2, 2}}
	k := 2
	fmt.Println(kClosest(points, k))

}

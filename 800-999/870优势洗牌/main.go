package main

import (
	"fmt"
	"sort"
)

func advantageCount(A []int, B []int) []int {
	tmpB := make([]int, len(B))
	copy(tmpB, B)
	sort.Ints(A)
	sort.Ints(tmpB)
	assigned := make(map[int][]int)
	remaining := make([]int, 0)
	j := 0
	for _, a := range A {
		if a > tmpB[j] {
			assigned[tmpB[j]] = append(assigned[tmpB[j]], a)
			j++
		} else {
			remaining = append(remaining, a)
		}
	}
	ans := make([]int, 4)
	var item int
	for i, b := range B {
		if len(assigned[b]) > 0 {
			item = assigned[b][len(assigned[b])-1]
			assigned[b] = assigned[b][:len(assigned[b])-1]
			ans[i] = item
		} else {
			item = remaining[len(remaining)-1]
			remaining = remaining[:len(remaining)-1]
			ans[i] = item
		}
	}
	return ans

}

func main() {
	A := []int{2, 0, 4, 1, 2}
	B := []int{1, 3, 0, 0, 2}
	fmt.Println(advantageCount(A, B))

}

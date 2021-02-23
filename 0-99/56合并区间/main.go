package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	res := make([][]int, 0)
	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if temp[1] >= intervals[i][0] {
			temp = []int{temp[0], max(temp[1], intervals[i][1])}
		} else {
			res = append(res, temp)
			temp = intervals[i]
		}
	}
	res = append(res, temp)
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//intervals := [][]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}}
	intervals := [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals))

}

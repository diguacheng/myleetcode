package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	res := make([][]int, 0)
	start, end := newInterval[0], newInterval[1]

	i := 0
	for i < n && start > intervals[i][0] {
		res = append(res, intervals[i])
		i++
	}
	if len(res) == 0 || res[len(res)-1][1] < start {
		res = append(res, newInterval)
	} else {
		res[len(res)-1][1] = max(res[len(res)-1][1], end)
	}
	for i < n {
		s, e := intervals[i][0], intervals[i][1]
		if res[len(res)-1][1] < s {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], e)
		}
		i++
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newinterval := []int{2, 5}
	fmt.Println(insert(intervals, newinterval))

}

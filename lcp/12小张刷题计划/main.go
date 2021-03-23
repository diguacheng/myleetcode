package main

import "fmt"

func minTime(time []int, m int) int {

	l, r := 0, sum(time)
	for l < r {
		mid := (l + r) >> 1
		if check(mid, time, m) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func check(limit int, cost []int, day int) bool {
	useDay, totalTime, maxTime := 1, 0, cost[0]
	for _, i := range cost[1:] {
		if totalTime+min(maxTime, i) <= limit {
			totalTime = totalTime + min(maxTime, i)
			maxTime = max(maxTime, i)
		} else {
			useDay += 1
			totalTime = 0
			maxTime = i
		}

	}
	return useDay <= day

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

func main() {
	t := []int{1, 2, 3, 3}
	m := 2
	fmt.Println(minTime(t, m))
}

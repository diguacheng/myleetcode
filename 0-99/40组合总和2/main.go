package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	temp := make([]int, 0)
	sort.Ints(candidates)
	backTrack(candidates, target, temp, &res)
	return res

}

func backTrack(candidates []int, target int, temp []int, res *[][]int) {
	if target == 0 {
		ans := make([]int, len(temp))
		copy(ans, temp)
		*res = append(*res, ans)
	}
	for i := 0; i < len(candidates); i++ {
		if i != 0 && candidates[i] == candidates[i-1] {
			continue
		}

		if target-candidates[i] >= 0 {
			temp = append(temp, candidates[i])
			backTrack(candidates[i+1:], target-candidates[i], temp, res)
			temp = temp[:len(temp)-1]
		}
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}

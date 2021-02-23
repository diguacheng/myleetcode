package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	temp := make([]int, 0)
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
		if target-candidates[i] >= 0 {
			temp = append(temp, candidates[i])
			backTrack(candidates[i:], target-candidates[i], temp, res)
			temp = temp[:len(temp)-1]
		}
	}
}

func main() {
	candidates:=[]int{2,3,6,7}
	target :=7
	fmt.Println(combinationSum(candidates, target) )

}
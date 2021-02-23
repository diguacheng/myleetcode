package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	temp := make([]int, 0)
	res := make([][]int, 0)
	backTrack(candidates, k, n, temp, &res)
	return res

}

func backTrack(candidates []int, k, n int, temp []int, res *[][]int) {
	if k == 0 && n == 0 {
		ans := make([]int, len(temp))
		copy(ans, temp)
		*res = append(*res, ans)
	}
	if k > 0 && n > 0 {
		l := len(candidates)
		for i := 0; i < l; i++ {
			if n < candidates[i] {
				break
			}
			temp = append(temp, candidates[i])
			backTrack(candidates[i+1:], k-1, n-candidates[i], temp, res)
			temp = temp[:len(temp)-1]
		}
	}
}

func main() {
	fmt.Println(combinationSum3(3,9))

}
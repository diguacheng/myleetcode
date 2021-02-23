package main

import "github.com/spf13/cobra/cobra/cmd"

func solveNQueens(n int) [][]string {
	temp := []string{}
	res := make([][]string, 0)
	isVisited := make([][]bool, n)
	for i := 0; i < n; i++ {
		isVisited[i] = make([]bool, n)
	}
	backTrack(isVisited, 0, n, temp, &res)

}
func backTrack(isVisited [][]bool, i, n int, temp []string, res *[][]string) {
	if i == n {
		ans := make([]string, len(temp))
		copy(ans, temp)
		*res = append(*res, ans)
	}
	for k := i; k < n; k++ {
		for j := 0; j < n; j++ {
			if isVisited[k][j] == false {
				cmd
			}
		}
	}
}

func main() {

}
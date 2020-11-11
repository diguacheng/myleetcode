package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	l := len(word)
	if m*n < l {
		return false
	}
	NewIsvisiTted := func(m, n int) [][]bool {
		res := make([][]bool, m)
		for i := 0; i < m; i++ {
			res[i] = make([]bool, n)
		}
		return res
	}

	var DFS func(isVisited [][]bool, i, j int, word string) bool

	DFS = func(isVisited [][]bool, i, j int, word string) bool {
		if len(word) == 0 {
			return true
		}
		dx := []int{1, 0, -1, 0}
		dy := []int{0, -1, 0, 1}
		res := false
		for d := 0; d < 4; d++ {
			newi, newj := i+dx[d], j+dy[d]
			if newi >= 0 && newi < m && newj >= 0 && newj < n && isVisited[newi][newj] == false && board[newi][newj] == word[0] {
				isVisited[newi][newj] = true
				res = res || DFS(isVisited, newi, newj, word[1:])
				isVisited[newi][newj] = false
			}

		}
		return res
	}
	isVisited := NewIsvisiTted(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				isVisited[i][j] = true
				if DFS(isVisited, i, j, word[1:]) == true {
					return true
				}
				isVisited[i][j] = false
			}

		}
	}
	return false
}

func exist1(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	l := len(word)
	isvisted := make([][]bool, m)
	for i := 0; i < m; i++ {
		isvisted[i] = make([]bool, n)
	}

	var DFS func(i, j, k int) bool

	DFS = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == l-1 {
			return true
		}

		isvisted[i][j] = true
		dx := []int{1, 0, -1, 0}
		dy := []int{0, -1, 0, 1}
		for d := 0; d < 4; d++ {
			newi, newj := i+dx[d], j+dy[d]
			if newi >= 0 && newi < m && newj >= 0 && newj < n && isvisted[newi][newj] == false {
				if DFS(newi, newj, k+1) == true {
					return true
				}
			}

		}
		isvisted[i][j] = false
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if DFS(i, j, 0) == true {
				return true
			}

		}
	}
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	// board := [][]byte{
	// 	{'a', 'a', 'a', 'a'},
	// 	{'a', 'a', 'a', 'a'},
	// 	{'a', 'a', 'a', 'a'},
	// 	{'a', 'a', 'a', 'a'},
	// 	{'a', 'a', 'a', 'b'},
	// }
	// fmt.Println(exist1(board,"aaaaaaaaaaaaaaaaaaaa"))
	fmt.Println(exist1(board, "ABCCED"))
	fmt.Println(exist1(board, "ABCB"))

}

package main

import "fmt"

func gameOfLife(board [][]int) {
	neighbors := [3]int{0, 1, -1}
	rows := len(board)
	cols := len(board[0])
	copyboard := make([][]int, rows)
	copy(copyboard, board)
	for i := 0; i < rows; i++ {
		copyboard[i] = make([]int, cols)
		copy(copyboard[i], board[i])
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			liveNeignbors := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if !(neighbors[i] == 0 && neighbors[j] == 0) {
						r := row + neighbors[i]
						c := col + neighbors[j]

						if r < rows && r >= 0 && c < cols && c >= 0 && copyboard[r][c] == 1 {
							liveNeignbors++

						}

					}
				}
			}
			if copyboard[row][col] == 1 && (liveNeignbors < 2 || liveNeignbors > 3) {
				board[row][col] = 0
			}
			if copyboard[row][col] == 0 && liveNeignbors == 3 {
				board[row][col] = 1

			}
		}

	}

}
func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0}}
	gameOfLife(board)
	fmt.Println(board)

}

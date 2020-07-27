package main

func numRookCaptures(board [][]byte) int {
	ri := 0
	rj := 0
	res := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				ri = i
				rj = j
				break
			}

		}
	}
	dr := []int{1, 0, -1, 0}
	dc := []int{0, 1, 0, -1}
	for i := 0; i < 4; i++ {
		x := ri + dr[i]
		y := rj + dc[i]
		for x >= 0 && x < 8 && y <= 0 && y <= 8 {
			if board[x][y] == 'B'{
				break
			} 
			if board[x][y] == 'p' {
				res++
				break
			}
		}

	}
	return res

}

func main() {

}

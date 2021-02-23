package main



func solve(board [][]byte) {
	n := len(board)
	if n <= 2 {
		return
	}
	m:=len(board[0])
	if m<=2{
		return 
	}
	isVisited := make([][]bool, n)
	for i := 0; i < n; i++ {
		isVisited[i] = make([]bool, m)
	}
	var help func(x, y int)
	help = func(x, y int) {
		if isVisited[x][y] {
			return
		}
		isVisited[x][y] = true
		if x+1 < n {
			if board[x+1][y] == 'O' {
				help(x+1, y)
			}
		}
		if x-1 >= 0 {
			if board[x-1][y] == 'O' {
				help(x-1, y)
			}
		}
		if y+1 < m {
			if board[x][y+1] == 'O' {
				help(x, y+1)
			}
		}
		if y-1 >= 0 {
			if board[x][y-1] == 'O' {
				help(x, y-1)
			}
		}

	}

	for i := 0; i < m; i++ {
		if board[0][i] == 'O' {
			help(0, i)
		}
		if board[n-1][i] == 'O' {
			help(n-1, i)
		}
	}
	for i:=1;i<n-1;i++{
		if board[i][0] == 'O' {
			help(i,0)
		}
		if board[i][m-1] == 'O' {
			help(i,m-1)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' && isVisited[i][j] == false {
				board[i][j] = 'X'
			}
		}
	}

}

func main() {
	s := [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'O', 'X'},
	}
	solve(s)

}

package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	n, m := len(grid), len(grid[0])
	isvisted := make([][]int, n)

	for i := 0; i < n; i++ {
		isvisted[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' && isvisted[i][j] == 0 {
				isvisted[i][j] = 1
				helper(grid, isvisted, i, j, n, m)
				fmt.Println(i, j)
				count++
			}
		}
	}
	return count

}

func helper(grid [][]byte, isvisted [][]int, i, j, n, m int) {
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if 0 <= x && x < n && 0 <= y && y < m {
			if grid[x][y] == '1' && isvisted[x][y] == 0 {
				isvisted[x][y] = 1
				helper(grid, isvisted, x, y, n, m)
			}

		}

	}
}

func numIslands1(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				helper1(grid, i, j, n, m)
				count++
			}
		}
	}
	return count

}

func helper1(grid [][]byte, i, j, n, m int) {
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if 0 <= x && x < n && 0 <= y && y < m {
			if grid[x][y] == '1' {
				grid[x][y] = '0'
				helper1(grid, x, y, n, m)
			}

		}

	}
}

func numIslands4(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	isvisteed := make([][]int, m)
	for i := 0; i < m; i++ {
		isvisteed[i] = make([]int, n)
	}
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && isvisteed[i][j] == 0 {
				isvisteed[i][j] = 1
				DFS(grid, isvisteed, i, j)
				fmt.Println(i, j)
				cnt++
			}
		}
	}
	return cnt
}

func DFS(grid [][]byte, isvisteed [][]int, i, j int) {
	m := len(grid)
	n := len(grid[0])
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	for i := 0; i < 4; i++ {
		ni := i + dx[i]
		nj := i + dy[i]
		if ni >= 0 && ni < m && nj >= 0 && nj < n && isvisteed[ni][nj] == 0 && grid[ni][nj] == '1' {
			isvisteed[ni][nj] = 1
			DFS(grid, isvisteed, ni, nj)
		}
	}
}
func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands4(grid))
	fmt.Println(numIslands1(grid))

}

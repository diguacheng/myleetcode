package main

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	table := make([][]int, m)
	for i := 0; i < m; i++ {
		table[i] = make([]int, n)
	}
	var dp func(x, y int) int
	dp = func(x, y int) int {
		if x >= m || y >= n {
			return 1 << 31
		}
		if x==m-1&&y==n-1{
			table[x][y] = grid[x][y]
			return table[x][y]
		}
		r := table[x][y]
		if r == 0 {
			temp := min(dp(x+1, y), dp(x, y+1)) + grid[x][y]
			table[x][y] = temp
		}
		return table[x][y]
	}
	return dp(0, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	table := make([][]int, m)
	for i := 0; i < m; i++ {
		table[i] = make([]int, n)
	}
	table[0][0] =grid[0][0]
	for i:=1;i<m;i++ {
		table[i][0] = table[i-1][0]+grid[i][0]
	}
	for i:=1;i<n;i++{
		table[0][i]=table[0][i-1]+grid[0][i]
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			table[i][j]=min(table[i-1][j],table[i][j-1])+grid[i][j]
		}
	}
	return table[m-1][n-1]
}


func main() {
	s := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum1(s))
}
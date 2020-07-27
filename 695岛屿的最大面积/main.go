package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	isvisited := make([][]int, len(grid))
	for i, v := range grid {
		isvisited[i] = make([]int, len(v))
	}
	count := 0
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 && isvisited[i][j] == 0 {
				visitGrid(grid, isvisited, i, j, &count)
				if max < count {
					max = count
				}
			}
			count = 0
		}
	}
	return max

}
func visitGrid(grid [][]int, isvisited [][]int, x, y int, count *int) {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
		if isvisited[x][y] == 0 && grid[x][y] == 1 {
			isvisited[x][y] = 1
			*count++
			visitGrid(grid, isvisited, x-1, y, count)
			visitGrid(grid, isvisited, x+1, y, count)
			visitGrid(grid, isvisited, x, y-1, count)
			visitGrid(grid, isvisited, x, y+1, count)
		}

	}

}
func main() {

	/* 	grid:=[][]int{{0,0,1,0,0,0,0,1,0,0,0,0,0},
	{0,0,0,0,0,0,0,1,1,1,0,0,0},
	{0,1,1,0,1,0,0,0,0,0,0,0,0},
	{0,1,0,0,1,1,0,0,1,0,1,0,0},
	{0,1,0,0,1,1,0,0,1,1,1,0,0},
	{0,0,0,0,0,0,0,0,0,0,1,0,0},
	{0,0,0,0,0,0,0,1,1,1,0,0,0},
	{0,0,0,0,0,0,0,1,1,0,0,0,0}} */
	grid := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
	fmt.Println(maxAreaOfIsland(grid))

}

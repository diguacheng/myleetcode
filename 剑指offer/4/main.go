package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])
	if cols == 0 {
		return false
	}
	x := 0
	y := cols - 1
	for x < rows && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func main() {
	/*
		matrix := [][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}
	*/

	fmt.Println(findNumberIn2DArray([][]int{{-5}}, -5))

}

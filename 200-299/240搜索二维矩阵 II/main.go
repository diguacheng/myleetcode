package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	start, end := 0, m-1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if end == -1 {
		return false
	}
	i := end
	for i >= 0 {
		start, end = 0, n-1
		for start <= end {
			mid := (start + end) / 2
			if matrix[i][mid] == target {
				return true
			}
			if matrix[i][mid] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}

		}
		i--
	}
	return false
}

func main() {
	m := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(m, 5))
}

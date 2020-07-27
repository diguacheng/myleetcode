package main

import "fmt"

func updateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	isvised := make([][]int, n)
	for i := 0; i < n; i++ {
		isvised[i] = make([]int, m)
	}
	dir := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	list := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				list = append(list, []int{i, j})
				isvised[i][j] = 1

			}
		}
	}
	for len(list) != 0 {
		x, y := list[0][0], list[0][1]
		list = list[1:]
		for d := 0; d < 4; d++ {
			nx := x + dir[d][0]
			ny := y + dir[d][1]
			if 0 <= nx && nx < n && 0 <= ny && ny < m && isvised[nx][ny] == 0 {
				res[nx][ny] = res[x][y] + 1
				list = append(list, []int{nx, ny})
				isvised[nx][ny] = 1

			}
		}
	}
	return res

}

func updateMatrix1(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] != 0 {
				res[i][j] = 500
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i-1 >= 0 {
				res[i][j] = min(res[i][j], res[i-1][j]+1)
			}
			if j-1 >= 0 {
				res[i][j] = min(res[i][j], res[i][j-1]+1)
			}
		}
	}
	

	for i := n-1; i >=0; i-- {
		for j := m-1; j >=0; j-- {
			if i+1<n {
				res[i][j] = min(res[i][j], res[i+1][j]+1)
			}
			if j+1<m {
				res[i][j] = min(res[i][j], res[i][j+1]+1)
			}
		}
	}
	return res


}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix := [][]int{
		{0, 0, 0}, {0, 1, 0}, {1, 1, 1},
	}
	fmt.Println(updateMatrix1(matrix))

}

package main

import "fmt"

func rotate2(matrix [][]int) {

	if len(matrix) == 0 {
		return
	}
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {

			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}
func rotate(matrix [][]int) {
	// 保存临时变量

	if len(matrix) == 0 {
		return
	}
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = temp
		}
	}

}

func rotate1(matrix [][]int) {
	// 不满足要求 原地旋转
	if len(matrix) == 0 {
		return
	}
	n := len(matrix)
	temp := make([][]int, n)
	for i := 0; i < n; i++ {
		temp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			temp[j][n-i-1] = matrix[i][j]
		}
	}
	copy(matrix, temp)

}

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate2(matrix)
	fmt.Println(matrix)

}

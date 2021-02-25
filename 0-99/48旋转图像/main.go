package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}
	flag := n % 2
	l := n / 2
	var tmp int
	for i := 0; i < l+flag; i++ {
		for j := 0; j < l; j++ {
			tmp = matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = tmp
		}
	}
	return
}

func main() {
	// matrix := [][]int{
	// 	{5, 1, 9, 11},
	// 	{2, 4, 8, 10},
	// 	{13, 3, 6, 7},
	// 	{15, 14, 12, 16},
	// }
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)

}

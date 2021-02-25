package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	// 暴力法
	if len(matrix) == 0 {
		return 0
	}
	n := len(matrix)
	if len(matrix[0]) == 0 {
		return 0
	}
	m := len(matrix[0])
	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				temp := isSquare(matrix, i, j, n, m)
				if temp == 4 {
					fmt.Println(i, j)

				}
				if max < temp {
					max = temp

				}

			}

		}

	}
	return max

}

func isSquare(matrix [][]byte, i, j, n, m int) int {
	edg := 1
	var a, b int
	for i+edg < n && j+edg < m {
		for a, b = i+edg, j; b <= j+edg; b++ {
			if matrix[a][b] == '0' {
				return edg * edg
			}
		}
		for a, b = i, j+edg; a < i+edg; a++ {
			if matrix[a][b] == '0' {
				return edg * edg
			}
		}
		edg++
	}

	return edg * edg

}

func maximalSquare1(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxEdg := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxEdg = 1
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxEdg {
					maxEdg = dp[i][j]
				}
			}
		}
	}
	return maxEdg * maxEdg

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	matrix := [][]byte{
		{'0', '0', '0', '1', '0', '1', '1', '1'},
		{'0', '1', '1', '0', '0', '1', '0', '1'},
		{'1', '0', '1', '1', '1', '1', '0', '1'},
		{'0', '0', '0', '1', '0', '0', '0', '0'},
		{'0', '0', '1', '0', '0', '0', '1', '0'},
		{'1', '1', '1', '0', '0', '1', '1', '1'},
		{'1', '0', '0', '1', '1', '0', '0', '1'},
		{'0', '1', '0', '0', '1', '1', '0', '0'},
		{'1', '0', '0', '1', '0', '0', '0', '0'},
	}
	fmt.Println(maximalSquare1(matrix))

	// matrix := [][]byte{
	// 	{'1', '0', '1', '0', '0'},
	// 	{'1', '0', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'1', '0', '0', '1', '0'},
	// }
}

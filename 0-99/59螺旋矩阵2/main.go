package main

import "fmt"

func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	l := n * n
	i := 1
	x := 0
	y := 0
	circle := 0
	for i <= l {
		for i <= l && y < n-circle {
			res[x][y] = i
			i++
			y++
		}
		y--
		x++
		for i <= l && x < n-circle {
			res[x][y] = i
			i++
			x++
		}
		x--
		y--
		for i <= l && y >= circle {
			res[x][y] = i
			i++
			y--
		}
		y++
		x--
		for i <= l && x >= circle+1 {
			res[x][y] = i
			i++
			x--
		}
		x++
		y++
		circle++
	}
	return res

}

func main() {
	n := 6
	x := generateMatrix(6)
	for i := 0; i < n; i++ {
		fmt.Println(x[i])
	}

}

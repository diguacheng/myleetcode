package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	if len(matrix[0]) == 0 {
		return nil
	}
	directionx := [4]int{0, 1, 0, -1}
	directiony := [4]int{1, 0, -1, 0}
	res := make([]int, 0)
	xstart, xend := 0, len(matrix)-1
	ystart, yend := 0, len(matrix[0])-1
	x := 0
	y := 0
	curr := 0 // 当前方向
	for x >= xstart && x <= xend && y >= ystart && y <= yend {
		res = append(res, matrix[x][y])
		x = x + directionx[curr]
		y = y + directiony[curr]
		if !(x >= xstart && x <= xend && y >= ystart && y <= yend) {
			x = x - directionx[curr]
			y = y - directiony[curr]
			curr = (curr + 1) % 4
			x = x + directionx[curr]
			y = y + directiony[curr]
			if directionx[curr] == 0 && directiony[curr] == -1 {
				yend--
			}
			if directionx[curr] == -1 && directiony[curr] == 0 {
				xend--
			}
			if directionx[curr] == 0 && directiony[curr] == 1 {
				ystart++
			}
			if directionx[curr] == 1 && directiony[curr] == 0 {
				xstart++
			}
		}
	}
	return res
}

func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	sum := len(matrix) * len(matrix[0])
	res := make([]int, 0)
	xstart, xend := 0, len(matrix)-1
	ystart, yend := 0, len(matrix[0])-1
	step := 0
	for step < sum {
		for i := ystart; i <= yend && step < sum; i++ {
			res = append(res, matrix[xstart][i])
			step++
		}
		xstart++
		for i := xstart; i <= xend && step < sum; i++ {
			res = append(res, matrix[i][yend])
			step++
		}
		yend--
		for i := yend; i >= ystart && step < sum; i-- {
			res = append(res, matrix[xend][i])
			step++
		}
		xend--
		for i := xend; i >= xstart && step < sum; i-- {
			res = append(res, matrix[i][ystart])
			step++
		}
		ystart++

	}

	return res
}

func main() {

	matrix := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}}
	fmt.Println(spiralOrder1(matrix))
}

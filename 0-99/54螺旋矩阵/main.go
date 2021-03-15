package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	sum := m * n
	res := make([]int, sum)
	step := 0
	xstart, xend := 0, m-1
	ystart, yend := 0, n-1
	for step < sum {
		fmt.Println(res)
		for i := ystart; i <= yend && step < sum; i++ {
			res[step] = matrix[xstart][i]
			step++
		}
		xstart++
		for i := xstart; i <= xend && step < sum; i++ {
			res[step] = matrix[i][yend]
			step++
		}
		yend--
		for i := yend; i >= ystart && step < sum; i-- {
			res[step] = matrix[xend][i]
			step++
		}
		xend--
		for i := xend; i >= xstart && step < sum; i-- {
			res[step] = matrix[i][ystart]
			step++
		}
		ystart++
	}
	return res
}

func main() {
	s := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(s))

}
package main

import "fmt"

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	table := make([][]int, R)
	for i := 0; i < R; i++ {
		table[i] = make([]int, C)
	}
	list := [][]int{}
	res := [][]int{}
	list = append(list, []int{r0, c0})
	table[r0][c0] = 1
	res = append(res, []int{r0, c0})
	c := 1
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, 1, -1}
	for len(list) != 0 {
		newlist := [][]int{}
		for _, p := range list {
			x, y := p[0], p[1]
			for i := 0; i < 4; i++ {
				if x+dx[i] >= 0 && x+dx[i] < R && y+dy[i] >= 0 && y+dy[i] < C && table[x+dx[i]][y+dy[i]] == 0 && distence(r0, c0, x+dx[i], y+dy[i]) == c {
					table[x+dx[i]][y+dy[i]] = 1
					res = append(res, []int{x + dx[i], y + dy[i]})
					newlist = append(newlist, []int{x + dx[i], y + dy[i]})
				}
			}
		}
		c++
		list = newlist
	}
	return res

}

func distence(x1, y1, x2, y2 int) int {
	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	return abs(x1, x2) + abs(y1, y2)

}

func main() {
	fmt.Println(allCellsDistOrder(2, 3, 1, 2))

}

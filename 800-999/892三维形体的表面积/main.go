package main

func surfaceArea(grid [][]int) int {
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}
	l := len(grid)
	res := 0
	for i := 0; i < l; i++ {
		for j := 0; i < l; j++ {
			if grid[i][j] > 0 {
				res += 2
				for k := 0; k < 4; k++ {
					temp := 0
					if 0 <= i+dr[k] && i+dr[k] < l && 0 <= j+dc[k] && j+dc[k] < l {
						temp = grid[i+dr[k]][j+dc[k]]
					}
					res += max(grid[i][j]-temp, 0)
				}
			}
		}
	}
	return res

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {

}

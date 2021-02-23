package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	
	depth := len(triangle) - 1
	table:=make([][]int,len(triangle))
	for i := 0; i < len(triangle); i++ {
		table[i]=make([]int, len(triangle[i]))
	}
	var dp func(x, y int) int
	dp = func(x, y int) int {
		if x > depth {
			return 0
		}
		r:=table[x][y]
		if r==0{
			temp:=min(dp(x+1,y),dp(x+1,y+1)) + triangle[x][y]
			table[x][y]=temp
		}
		return table[x][y]
	}
	return dp(0, 0)

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	a := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(a))

}

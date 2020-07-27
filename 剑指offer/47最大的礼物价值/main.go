package main

import "fmt"

func maxValue(grid [][]int) int {
	if len(grid)==0{
		return 0
	}
	if len(grid[0])==0{
		return 0
	}
	dpp:=make([][]int,len(grid))
	for i:=0;i<len(grid);i++{
		dpp[i]=make([]int,len(grid[i]))

	}
	n:=len(grid)-1
	m:=len(grid[0])-1
	var dp func(int,int)int
	dp=func(x,y int) int {
		if x<0||y<0{
			return 0
		}
		if dpp[x][y]>0{
			return dpp[x][y]
		}
		dpp[x][y]=max(dp(x-1,y),dp(x,y-1))+grid[x][y]
		return dpp[x][y]

	}
	return dp(n,m)

}

func max(x,y int) int {
	if x>y{
		return x
	}
	return y
}

func maxValue1(grid [][]int) int {
	if len(grid)==0{
		return 0
	}
	if len(grid[0])==0{
		return 0
	}
	
	n:=len(grid)
	
	m:=len(grid[0])
	for i:=1;i<n;i++{
		grid[i][0]+=grid[i-1][0]
	}
	for i:=1;i<m;i++{
		grid[0][i]+=grid[0][i-1]
	}
	i,j:=1,1
	for i!=n&&j!=m{
		for k:=i;k<n;k++{
			grid[k][j]+=max(grid[k-1][j],grid[k][j-1])
		}
		for k:=j+1;k<m;k++{
			grid[i][k]+=max(grid[i-1][k],grid[i][k-1])
		}
		i++
		j++
	}
	return grid[n-1][m-1]

}


func main() {
	grid:=[][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	fmt.Println(maxValue1(grid))
}
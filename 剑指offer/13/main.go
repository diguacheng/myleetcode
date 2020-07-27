package main

import (
	"fmt"
)

func movingCount(m int, n int, k int) int {
	isVisted:=make([][]int ,m)
	for i:=range isVisted{
		isVisted[i]=make([]int,n)
	}
	res:=count(0,0,m,n,isVisted,k)
	return res


}

func count(x,y,m,n int,isVisted [][]int,k int)int{
	if x>=0&&x<m && y>=0&&y<n&&isVisted[x][y]==0{
		if digitSum(x)+digitSum(y)<=k{
			isVisted[x][y]=1
			sum:=count(x-1,y,m,n,isVisted,k)+count(x,y-1,m,n,isVisted,k)+count(x+1,y,m,n,isVisted,k)+count(x,y+1,m,n,isVisted,k)
			return 1+sum
		}
	}
	return 0
}

func digitSum(x int)int{
	res:=0
	for x!=0{
		res=res+x%10
		x/=10
	}
	return res
}
func main() {
	m := 2
	n:= 3 
	k:= 1
	fmt.Println(movingCount(m,n,k))

}

